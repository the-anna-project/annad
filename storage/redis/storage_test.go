package redis

import (
	"reflect"
	"testing"

	"github.com/garyburd/redigo/redis"
	"github.com/rafaeljusto/redigomock"

	"github.com/xh3b4sd/anna/spec"
)

func testMustNewStorageWithConn(t *testing.T, c redis.Conn) spec.Storage {
	newStorage, err := NewStorage(DefaultStorageConfigWithConn(c))
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newStorage
}

// NewStorage

func Test_Storage_NewRedisStorage_Error_BackOffFactory(t *testing.T) {
	newStorageConfig := DefaultStorageConfig()
	newStorageConfig.BackOffFactory = nil
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
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_Storage_NewRedisStorage_Error_Pool(t *testing.T) {
	newStorageConfig := DefaultStorageConfig()
	newStorageConfig.Log = nil
	_, err := NewStorage(newStorageConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
	newStorageConfig.Pool = nil
	_, err = NewStorage(newStorageConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

// Get

func Test_RedisStorage_Get_Success(t *testing.T) {
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

func Test_RedisStorage_Get_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("GET", "prefix:foo").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	_, err := newStorage.Get("foo")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

// GetElementsByScore

//func Test_RedisStorage_GetElementsByScore_Success(t *testing.T) {
//	c := redigomock.NewConn()
//	c.Command("ZREVRANGEBYSCORE", "prefix:foo", 0.8, 0.8, "LIMIT", 0, 3).Expect([]interface{}{"bar"})
//
//	newStorage := testMustNewStorageWithConn(t, c)
//
//	values, err := newStorage.GetElementsByScore("foo", 0.8, 3)
//	if err != nil {
//		t.Fatal("expected", nil, "got", err)
//	}
//	if len(values) != 1 {
//		t.Fatal("expected", 1, "got", len(values))
//	}
//	if values[0] != "bar" {
//		t.Fatal("expected", "bar", "got", values[0])
//	}
//}

func Test_RedisStorage_GetElementsByScore_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZREVRANGEBYSCORE", "prefix:foo", 0.8, 0.8, "LIMIT", 0, 3).ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	_, err := newStorage.GetElementsByScore("foo", 0.8, 3)
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

// GetHighestScoredElements

//func Test_RedisStorage_GetHighestScoredElements_Success(t *testing.T) {
//	c := redigomock.NewConn()
//	c.Command("ZREVRANGE", "prefix:foo", 0, 2, "WITHSCORES").Expect([]interface{}{"one", "0.8", "two", "0.5"})
//
//	newStorage := testMustNewStorageWithConn(t, c)
//
//	values, err := newStorage.GetHighestScoredElements("foo", 3)
//	if err != nil {
//		t.Fatal("expected", nil, "got", err)
//	}
//	if len(values) != 4 {
//		t.Fatal("expected", 1, "got", len(values))
//	}
//	if values[0] != "one" {
//		t.Fatal("expected", "one", "got", values[0])
//	}
//	if values[1] != "0.8" {
//		t.Fatal("expected", "0.8", "got", values[1])
//	}
//	if values[2] != "two" {
//		t.Fatal("expected", "two", "got", values[2])
//	}
//	if values[3] != "0.5" {
//		t.Fatal("expected", "0.5", "got", values[3])
//	}
//}

func Test_RedisStorage_GetHighestScoredElements_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZREVRANGE", "prefix:foo", 0, 2, "WITHSCORES").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	_, err := newStorage.GetHighestScoredElements("foo", 3)
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

// GetRandomKey

func Test_RedisStorage_GetRandomKey_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("RANDOMKEY").Expect("key1")

	newStorage := testMustNewStorageWithConn(t, c)

	randomKey, err := newStorage.GetRandomKey()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if randomKey != "key1" {
		t.Fatal("expected", "key1", "got", randomKey)
	}
}

func Test_RedisStorage_GetRandomKey_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("RANDOMKEY").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	_, err := newStorage.GetRandomKey()
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

// GetStringMap

func Test_RedisStorage_GetStringMap_Success(t *testing.T) {
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

func Test_RedisStorage_GetStringMap_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("HGETALL", "prefix:foo").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	_, err := newStorage.GetStringMap("foo")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

// Set

func Test_RedisStorage_Set_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SET", "prefix:foo", "bar").Expect("OK")

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.Set("foo", "bar")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_RedisStorage_Set_NoSuccess(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SET", "prefix:foo", "bar").Expect("invalid")

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.Set("foo", "bar")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_RedisStorage_Set_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SET", "prefix:foo", "bar").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.Set("foo", "bar")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

// SetElementByScore

func Test_RedisStorage_SetElementByScore_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZADD", "prefix:key", 0.8, "element").Expect(int64(1))

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.SetElementByScore("key", "element", 0.8)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_RedisStorage_SetElementByScore_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZADD", "prefix:key", 0.8, "element").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.SetElementByScore("key", "element", 0.8)
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

// PushToSet

func Test_RedisStorage_PushToSet(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SADD", "prefix:test-key", "test-element").Expect(int64(1))

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.PushToSet("test-key", "test-element")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_RedisStorage_PushToSet_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SADD", "prefix:test-key", "test-element").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.PushToSet("test-key", "test-element")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

// RemoveFromSet

func Test_RedisStorage_RemoveFromSet(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SREM", "prefix:test-key", "test-element").Expect(int64(1))

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.RemoveFromSet("test-key", "test-element")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_RedisStorage_RemoveFromSet_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SREM", "prefix:test-key", "test-element").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.RemoveFromSet("test-key", "test-element")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

// RemoveScoredElement

func Test_RedisStorage_RemoveScoredElement(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZREM", "prefix:test-key", "test-element").Expect(int64(1))

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.RemoveScoredElement("test-key", "test-element")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_RedisStorage_RemoveScoredElement_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZREM", "prefix:test-key", "test-element").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.RemoveScoredElement("test-key", "test-element")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

// SetStringMap

func Test_RedisStorage_SetStringMap_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("HMSET", "prefix:foo", "k1", "v1").Expect("OK")

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.SetStringMap("foo", map[string]string{"k1": "v1"})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_RedisStorage_SetStringMap_NotOK(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("HMSET", "prefix:foo", "k1", "v1").Expect("Not OK")

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.SetStringMap("foo", map[string]string{"k1": "v1"})
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_RedisStorage_SetStringMap_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("HMSET", "prefix:foo", "k1", "v1").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.SetStringMap("foo", map[string]string{"k1": "v1"})
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

// WalkScoredElements

func Test_RedisStorage_WalkScoredElements(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZSCAN", "prefix:test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-value-1"), []uint8("0.8"), []uint8("test-value-2"), []uint8("0.8")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	var values []interface{}
	err := newStorage.WalkScoredElements("test-key", nil, func(element string, score float64) error {
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

func Test_RedisStorage_WalkScoredElements_CloseDirectly(t *testing.T) {
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
	err := newStorage.WalkScoredElements("test-key", closer, func(element string, score float64) error {
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

func Test_RedisStorage_WalkScoredElements_CloseAfterCallback(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZSCAN", "prefix:test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-value-1"), []uint8("0.8"), []uint8("test-value-2"), []uint8("0.8")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	var count int
	closer := make(chan struct{}, 1)

	err := newStorage.WalkScoredElements("test-key", closer, func(element string, score float64) error {
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

func Test_RedisStorage_WalkScoredElements_QueryError(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZSCAN").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.WalkScoredElements("test-key", nil, nil)
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_RedisStorage_WalkScoredElements_CallbackError(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZSCAN", "prefix:test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-value-1"), []uint8("0.8"), []uint8("test-value-2"), []uint8("0.8")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.WalkScoredElements("test-key", nil, func(element string, score float64) error {
		return maskAny(queryExecutionFailedError)
	})
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

// WalkSet

func Test_RedisStorage_WalkSet(t *testing.T) {
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

func Test_RedisStorage_WalkSet_CloseDirectly(t *testing.T) {
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

func Test_RedisStorage_WalkSet_CloseAfterCallback(t *testing.T) {
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

func Test_RedisStorage_WalkSet_QueryError(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SSCAN").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.WalkSet("test-key", nil, nil)
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_RedisStorage_WalkSet_CallbackError(t *testing.T) {
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
