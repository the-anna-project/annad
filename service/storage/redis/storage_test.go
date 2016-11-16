package redis

import (
	"reflect"
	"testing"

	"github.com/garyburd/redigo/redis"

	"github.com/the-anna-project/spec/service"
)

func testMustNewStorageWithConn(t *testing.T, c redis.Conn) spec.Storage {
	newPoolConfig := DefaultPoolConfig()
	newMockDialConfig := defaultMockDialConfig()
	newMockDialConfig.RedisConn = redisConn
	newPoolConfig.Dial = newMockDial(newMockDialConfig)
	newPool := NewPool(newPoolConfig)

	// storage
	newStorageConfig := DefaultStorageConfig()
	newStorageConfig.pool = newPool
	newStorage, err := NewStorage(newStorageConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newStorage
}

func Test_ListStorage_PopFromList(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("BRPOP", "prefix:test-key", 0).Expect([]interface{}{
		[]uint8("test-key"),
		[]uint8("test-element"),
	})

	newStorage := testMustNewStorageWithConn(t, c)

	element, err := newStorage.PopFromList("test-key")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if element != "test-element" {
		t.Fatal("expected", "test-element", "got", element)
	}
}

func Test_ListStorage_PopFromList_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("BRPOP", "prefix:test-key", 0).ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	_, err := newStorage.PopFromList("test-key")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_ListStorage_PopFromList_Error_OneReturnValue(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("BRPOP", "prefix:test-key", 0).Expect([]interface{}{
		[]uint8("test-key"),
	})

	newStorage := testMustNewStorageWithConn(t, c)

	_, err := newStorage.PopFromList("test-key")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_ListStorage_PushToList(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("LPUSH", "prefix:test-key", "test-element").Expect(int64(1))

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.PushToList("test-key", "test-element")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_ListStorage_PushToList_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("LPUSH", "prefix:test-key", "test-element").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.PushToList("test-key", "test-element")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_ScoredSetStorage_GetElementsByScore_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZREVRANGEBYSCORE", "prefix:foo", 0.8, 0.8, "LIMIT", 0, 3).Expect([]interface{}{[]uint8("bar")})

	newStorage := testMustNewStorageWithConn(t, c)

	values, err := newStorage.GetElementsByScore("foo", 0.8, 3)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if len(values) != 1 {
		t.Fatal("expected", 1, "got", len(values))
	}
	if values[0] != "bar" {
		t.Fatal("expected", "bar", "got", values[0])
	}
}

func Test_ScoredSetStorage_GetElementsByScore_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZREVRANGEBYSCORE", "prefix:foo", 0.8, 0.8, "LIMIT", 0, 3).ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	_, err := newStorage.GetElementsByScore("foo", 0.8, 3)
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_ScoredSetStorage_GetHighestScoredElements_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZREVRANGE", "prefix:foo", 0, 2, "WITHSCORES").Expect([]interface{}{
		[]uint8("one"), []uint8("0.8"), []uint8("two"), []uint8("0.5"),
	})

	newStorage := testMustNewStorageWithConn(t, c)

	values, err := newStorage.GetHighestScoredElements("foo", 2)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if len(values) != 4 {
		t.Fatal("expected", 1, "got", len(values))
	}
	if values[0] != "one" {
		t.Fatal("expected", "one", "got", values[0])
	}
	if values[1] != "0.8" {
		t.Fatal("expected", "0.8", "got", values[1])
	}
	if values[2] != "two" {
		t.Fatal("expected", "two", "got", values[2])
	}
	if values[3] != "0.5" {
		t.Fatal("expected", "0.5", "got", values[3])
	}
}

func Test_ScoredSetStorage_GetHighestScoredElements_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZREVRANGE", "prefix:foo", 0, 2, "WITHSCORES").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	_, err := newStorage.GetHighestScoredElements("foo", 2)
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_ScoredSetStorage_SetElementByScore_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZADD", "prefix:key", 0.8, "element").Expect(int64(1))

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.SetElementByScore("key", "element", 0.8)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_ScoredSetStorage_SetElementByScore_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZADD", "prefix:key", 0.8, "element").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.SetElementByScore("key", "element", 0.8)
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_ScoredSetStorage_RemoveScoredElement(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZREM", "prefix:test-key", "test-element").Expect(int64(1))

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.RemoveScoredElement("test-key", "test-element")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_ScoredSetStorage_RemoveScoredElement_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZREM", "prefix:test-key", "test-element").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.RemoveScoredElement("test-key", "test-element")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_ScoredSetStorage_WalkScoredSet(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZSCAN", "prefix:test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-value-1"), []uint8("0.8"), []uint8("test-value-2"), []uint8("0.8")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	var values []interface{}
	err := newStorage.WalkScoredSet("test-key", nil, func(element string, score float64) error {
		values = append(values, element, score)
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if !reflect.DeepEqual(values, []interface{}{"test-value-1", 0.8, "test-value-2", 0.8}) {
		t.Fatal("expected", []interface{}{"test-value-1", 0.8, "test-value-2", 0.8}, "got", values)
	}
}

func Test_ScoredSetStorage_WalkScoredSet_CloseDirectly(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZSCAN", "prefix:test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-value-1"), []uint8("0.8"), []uint8("test-value-2"), []uint8("0.8")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	// Directly close and end walking.
	closer := make(chan struct{}, 1)
	closer <- struct{}{}

	var values []interface{}
	err := newStorage.WalkScoredSet("test-key", closer, func(element string, score float64) error {
		values = append(values, element, score)
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if values != nil {
		t.Fatal("expected", nil, "got", values)
	}
}

func Test_ScoredSetStorage_WalkScoredSet_CloseAfterCallback(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZSCAN", "prefix:test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-value-1"), []uint8("0.8"), []uint8("test-value-2"), []uint8("0.8")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	var count int
	closer := make(chan struct{}, 1)

	err := newStorage.WalkScoredSet("test-key", closer, func(element string, score float64) error {
		count++

		// Close and end walking.
		closer <- struct{}{}

		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if count != 1 {
		t.Fatal("expected", 1, "got", count)
	}
}

func Test_ScoredSetStorage_WalkScoredSet_QueryError(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZSCAN").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.WalkScoredSet("test-key", nil, nil)
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_ScoredSetStorage_WalkScoredSet_CallbackError(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZSCAN", "prefix:test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-value-1"), []uint8("0.8"), []uint8("test-value-2"), []uint8("0.8")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.WalkScoredSet("test-key", nil, func(element string, score float64) error {
		return maskAny(queryExecutionFailedError)
	})
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_SetStorage_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SET", "prefix:foo", "bar").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.Set("foo", "bar")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_SetStorage_GetAllFromSet_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SMEMBERS", "prefix:foo").Expect([]interface{}{
		[]uint8("one"), []uint8("two"),
	})

	newStorage := testMustNewStorageWithConn(t, c)

	values, err := newStorage.GetAllFromSet("foo")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if len(values) != 2 {
		t.Fatal("expected", 1, "got", len(values))
	}
	if values[0] != "one" {
		t.Fatal("expected", "one", "got", values[0])
	}
	if values[1] != "two" {
		t.Fatal("expected", "two", "got", values[2])
	}
}

func Test_SetStorage_GetAllFromSet_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SMEMBERS", "prefix:foo").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	_, err := newStorage.GetAllFromSet("foo")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_SetStorage_PushToSet(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SADD", "prefix:test-key", "test-element").Expect(int64(1))

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.PushToSet("test-key", "test-element")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_SetStorage_PushToSet_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SADD", "prefix:test-key", "test-element").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.PushToSet("test-key", "test-element")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_SetStorage_RemoveFromSet(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SREM", "prefix:test-key", "test-element").Expect(int64(1))

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.RemoveFromSet("test-key", "test-element")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_SetStorage_RemoveFromSet_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SREM", "prefix:test-key", "test-element").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.RemoveFromSet("test-key", "test-element")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_SetStorage_WalkSet(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SSCAN", "prefix:test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-value")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	var element1 string
	err := newStorage.WalkSet("test-key", nil, func(element string) error {
		element1 = element
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if element1 != "test-value" {
		t.Fatal("expected", "test-value", "got", element1)
	}
}

func Test_SetStorage_WalkSet_CloseDirectly(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SSCAN", "prefix:test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-value-1"), []uint8("test-value-2")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	// Directly close and end walking.
	closer := make(chan struct{}, 1)
	closer <- struct{}{}

	var element1 string
	err := newStorage.WalkSet("test-key", closer, func(element string) error {
		element1 = element
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if element1 != "" {
		t.Fatal("expected", "", "got", element1)
	}
}

func Test_SetStorage_WalkSet_CloseAfterCallback(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SSCAN", "prefix:test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-value-1"), []uint8("test-value-2")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	var count int
	closer := make(chan struct{}, 1)

	err := newStorage.WalkSet("test-key", closer, func(element string) error {
		count++

		// Close and end walking.
		closer <- struct{}{}

		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if count != 1 {
		t.Fatal("expected", 1, "got", count)
	}
}

func Test_SetStorage_WalkSet_QueryError(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SSCAN").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.WalkSet("test-key", nil, nil)
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_SetStorage_WalkSet_CallbackError(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SSCAN", "prefix:test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-value-1"), []uint8("test-value-2")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.WalkSet("test-key", nil, func(element string) error {
		return maskAny(queryExecutionFailedError)
	})
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Storage_DefaultStorageConfigWithAddr(t *testing.T) {
	newStorageConfig := DefaultStorageConfigWithAddr("foo")
	_, err := newStorageConfig.Pool.Dial()
	if err.Error() != "dial tcp: missing port in address foo" {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_Storage_NewRedisStorage_Error_BackoffFactory(t *testing.T) {
	newStorageConfig := DefaultStorageConfig()
	newStorageConfig.BackoffFactory = nil
	_, err := NewStorage(newStorageConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_Storage_NewRedisStorage_Error_Log(t *testing.T) {
	newStorageConfig := DefaultStorageConfig()
	newStorageConfig.Log = nil
	_, err := NewStorage(newStorageConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Storage_NewRedisStorage_Error_Pool(t *testing.T) {
	newStorageConfig := DefaultStorageConfig()
	newStorageConfig.Pool = nil
	_, err := NewStorage(newStorageConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Storage_NewRedisStorage_Error_Prefix(t *testing.T) {
	newStorageConfig := DefaultStorageConfig()
	newStorageConfig.Prefix = ""
	_, err := NewStorage(newStorageConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Storage_Shutdown(t *testing.T) {
	newStorageConfig := DefaultStorageConfig()
	newStorage, err := NewStorage(newStorageConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newStorage.Shutdown()
	newStorage.Shutdown()
	newStorage.Shutdown()
}

func Test_StringMapStorage_GetStringMap_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("HGETALL", "prefix:foo").Expect([]interface{}{[]byte("k1"), []byte("v1")})

	newStorage := testMustNewStorageWithConn(t, c)

	value, err := newStorage.GetStringMap("foo")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if !reflect.DeepEqual(value, map[string]string{"k1": "v1"}) {
		t.Fatal("expected", map[string]string{"k1": "v1"}, "got", value)
	}
}

func Test_StringMapStorage_GetStringMap_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("HGETALL", "prefix:foo").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	_, err := newStorage.GetStringMap("foo")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_StringMapStorage_SetStringMap_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("HMSET", "prefix:foo", "k1", "v1").Expect("OK")

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.SetStringMap("foo", map[string]string{"k1": "v1"})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_StringMapStorage_SetStringMap_NotOK(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("HMSET", "prefix:foo", "k1", "v1").Expect("Not OK")

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.SetStringMap("foo", map[string]string{"k1": "v1"})
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_StringMapStorage_SetStringMap_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("HMSET", "prefix:foo", "k1", "v1").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.SetStringMap("foo", map[string]string{"k1": "v1"})
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_StringStorage_Get_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("GET", "prefix:foo").Expect("bar")

	newStorage := testMustNewStorageWithConn(t, c)

	value, err := newStorage.Get("foo")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if value != "bar" {
		t.Fatal("expected", "bar", "got", value)
	}
}

func Test_StringStorage_Get_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("GET", "prefix:foo").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	_, err := newStorage.Get("foo")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_StringStorage_Get_Error_NotFound(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("GET", "prefix:foo").ExpectError(redis.ErrNil)

	newStorage := testMustNewStorageWithConn(t, c)

	_, err := newStorage.Get("foo")
	if !IsNotFound(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_StringStorage_GetRandom_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("RANDOMKEY").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	_, err := newStorage.GetRandom()
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_StringStorage_GetRandom_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("RANDOMKEY").Expect("key1")

	newStorage := testMustNewStorageWithConn(t, c)

	randomKey, err := newStorage.GetRandom()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if randomKey != "key1" {
		t.Fatal("expected", "key1", "got", randomKey)
	}
}

// Test_StringStorage_Remove_Error ensures that Remove does not throw any not
// found error.
func Test_StringStorage_Remove_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("DEL", "prefix:foo").Expect(int64(0))

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.Remove("foo")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_StringStorage_Remove_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("DEL", "prefix:foo").Expect(int64(1))

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.Remove("foo")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_StringStorage_Set_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SET", "prefix:foo", "bar").Expect("OK")

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.Set("foo", "bar")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_StringStorage_Set_NoSuccess(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SET", "prefix:foo", "bar").Expect("invalid")

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.Set("foo", "bar")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_StringStorage_WalkKeys(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SCAN", int64(0), "MATCH", "*", "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-key")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	var count int
	var element1 string

	err := newStorage.WalkKeys("*", nil, func(key string) error {
		count++
		element1 = key
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if count != 1 {
		t.Fatal("expected", 1, "got", count)
	}
	if element1 != "test-key" {
		t.Fatal("expected", "test-key", "got", element1)
	}
}

func Test_StringStorage_WalkKeys_CloseDirectly(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SCAN", int64(0), "MATCH", "*", "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-key")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	var count int
	// Directly close and end walking.
	closer := make(chan struct{}, 1)
	closer <- struct{}{}

	err := newStorage.WalkKeys("*", closer, func(key string) error {
		count++
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if count != 0 {
		t.Fatal("expected", 0, "got", count)
	}
}

func Test_StringStorage_WalkKeys_CloseAfterCallback(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SCAN", int64(0), "MATCH", "*", "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-key")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	var count int
	var element1 string
	closer := make(chan struct{}, 1)

	err := newStorage.WalkKeys("*", closer, func(key string) error {
		count++
		element1 = key

		// Close and end walking.
		closer <- struct{}{}

		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if count != 1 {
		t.Fatal("expected", 1, "got", count)
	}
	if element1 != "test-key" {
		t.Fatal("expected", "test-key", "got", element1)
	}
}

func Test_StringStorage_WalkKeys_QueryError(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SCAN").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.WalkKeys("*", nil, nil)
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_StringStorage_WalkKeys_CallbackError(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SCAN", int64(0), "MATCH", "*", "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-key")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.WalkKeys("*", nil, func(key string) error {
		return maskAny(queryExecutionFailedError)
	})
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}
