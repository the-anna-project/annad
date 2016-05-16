package redisstorage

import (
	"reflect"
	"testing"

	"github.com/cenk/backoff"
	"github.com/garyburd/redigo/redis"
	"github.com/rafaeljusto/redigomock"

	"github.com/xh3b4sd/anna/spec"
)

func testMaybeNewStorageWithConn(t *testing.T, c redis.Conn) spec.Storage {
	newConfig := DefaultConfigWithConn(c)
	newStorage, err := NewRedisStorage(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newStorage
}

func Test_RedisStorage_GetID(t *testing.T) {
	firstStorage, err := NewRedisStorage(DefaultConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	secondStorage, err := NewRedisStorage(DefaultConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if firstStorage.GetID() == secondStorage.GetID() {
		t.Fatal("expected", "different IDs", "got", "equal IDs")
	}
}

// Get

func Test_RedisStorage_Get_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("GET", "foo").Expect("bar")

	newStorage := testMaybeNewStorageWithConn(t, c)

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
	c.Command("GET", "foo").ExpectError(queryExecutionFailedError)

	newStorage := testMaybeNewStorageWithConn(t, c)

	_, err := newStorage.Get("foo")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

// GetElementsByScore

func Test_RedisStorage_GetElementsByScore_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZREVRANGEBYSCORE", "foo", 0.8, 0.8, "LIMIT", 0, 3).Expect([]interface{}{"bar"})

	newStorage := testMaybeNewStorageWithConn(t, c)

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

func Test_RedisStorage_GetElementsByScore_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZREVRANGEBYSCORE", "foo", 0.8, 0.8, "LIMIT", 0, 3).ExpectError(queryExecutionFailedError)

	newStorage := testMaybeNewStorageWithConn(t, c)

	_, err := newStorage.GetElementsByScore("foo", 0.8, 3)
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

// GetStringMap

func Test_RedisStorage_GetStringMap_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("HGETALL", "foo").Expect([]interface{}{[]byte("k1"), []byte("v1")})

	newStorage := testMaybeNewStorageWithConn(t, c)

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
	c.Command("HGETALL", "foo").ExpectError(queryExecutionFailedError)

	newStorage := testMaybeNewStorageWithConn(t, c)

	_, err := newStorage.GetStringMap("foo")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

// GetHighestScoredElements

func Test_RedisStorage_GetHighestScoredElements_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZREVRANGE", "foo", 0, 2, "WITHSCORES").Expect([]interface{}{"one", "0.8", "two", "0.5"})

	newStorage := testMaybeNewStorageWithConn(t, c)

	values, err := newStorage.GetHighestScoredElements("foo", 3)
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

func Test_RedisStorage_GetHighestScoredElements_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZREVRANGE", "foo", 0, 2, "WITHSCORES").ExpectError(queryExecutionFailedError)

	newStorage := testMaybeNewStorageWithConn(t, c)

	_, err := newStorage.GetHighestScoredElements("foo", 3)
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

// Set

func Test_RedisStorage_Set_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SET", "foo", "bar").Expect("OK")

	newStorage := testMaybeNewStorageWithConn(t, c)

	err := newStorage.Set("foo", "bar")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_RedisStorage_Set_NoSuccess(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SET", "foo", "bar").Expect("invalid")

	newStorage := testMaybeNewStorageWithConn(t, c)

	err := newStorage.Set("foo", "bar")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_RedisStorage_Set_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SET", "foo", "bar").ExpectError(queryExecutionFailedError)

	newStorage := testMaybeNewStorageWithConn(t, c)

	err := newStorage.Set("foo", "bar")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

// SetElementByScore

func Test_RedisStorage_SetElementByScore_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZADD", "key", 0.8, "element").Expect(int64(1))

	newStorage := testMaybeNewStorageWithConn(t, c)

	err := newStorage.SetElementByScore("key", "element", 0.8)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_RedisStorage_SetElementByScore_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZADD", "key", 0.8, "element").ExpectError(queryExecutionFailedError)

	newStorage := testMaybeNewStorageWithConn(t, c)

	err := newStorage.SetElementByScore("key", "element", 0.8)
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

// PushToSet

func Test_RedisStorage_PushToSet(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SADD", "test-key", "test-element").Expect(int64(1))

	newStorage := testMaybeNewStorageWithConn(t, c)

	err := newStorage.PushToSet("test-key", "test-element")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_RedisStorage_PushToSet_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SADD", "test-key", "test-element").ExpectError(queryExecutionFailedError)

	newStorage := testMaybeNewStorageWithConn(t, c)

	err := newStorage.PushToSet("test-key", "test-element")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

// RemoveFromSet

func Test_RedisStorage_RemoveFromSet(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SREM", "test-key", "test-element").Expect(int64(1))

	newStorage := testMaybeNewStorageWithConn(t, c)

	err := newStorage.RemoveFromSet("test-key", "test-element")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_RedisStorage_RemoveFromSet_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SREM", "test-key", "test-element").ExpectError(queryExecutionFailedError)

	newStorage := testMaybeNewStorageWithConn(t, c)

	err := newStorage.RemoveFromSet("test-key", "test-element")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

// RemoveScoredElement

func Test_RedisStorage_RemoveScoredElement(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZREM", "test-key", "test-element").Expect(int64(1))

	newStorage := testMaybeNewStorageWithConn(t, c)

	err := newStorage.RemoveScoredElement("test-key", "test-element")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_RedisStorage_RemoveScoredElement_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZREM", "test-key", "test-element").ExpectError(queryExecutionFailedError)

	newStorage := testMaybeNewStorageWithConn(t, c)

	err := newStorage.RemoveScoredElement("test-key", "test-element")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

// SetStringMap

func Test_RedisStorage_SetStringMap_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("HMSET", "foo", "k1", "v1").Expect("OK")

	newStorage := testMaybeNewStorageWithConn(t, c)

	err := newStorage.SetStringMap("foo", map[string]string{"k1": "v1"})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_RedisStorage_SetStringMap_NotOK(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("HMSET", "foo", "k1", "v1").Expect("Not OK")

	newStorage := testMaybeNewStorageWithConn(t, c)

	err := newStorage.SetStringMap("foo", map[string]string{"k1": "v1"})
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_RedisStorage_SetStringMap_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("HMSET", "foo", "k1", "v1").ExpectError(queryExecutionFailedError)

	newStorage := testMaybeNewStorageWithConn(t, c)

	err := newStorage.SetStringMap("foo", map[string]string{"k1": "v1"})
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

// WalkScoredElements

func Test_RedisStorage_WalkScoredElements(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZSCAN", "test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		int64(0),
		[]string{"test-value-1", "0.8", "test-value-2", "0.8"},
	})

	newStorage := testMaybeNewStorageWithConn(t, c)

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
	c.Command("ZSCAN", "test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		int64(0),
		[]string{"test-value-1", "0.8", "test-value-2", "0.8"},
	})

	newStorage := testMaybeNewStorageWithConn(t, c)

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
	c.Command("ZSCAN", "test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		int64(0),
		[]string{"test-value-1", "0.8", "test-value-2", "0.8"},
	})

	newStorage := testMaybeNewStorageWithConn(t, c)

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

	newStorage := testMaybeNewStorageWithConn(t, c)

	err := newStorage.WalkScoredElements("test-key", nil, nil)
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_RedisStorage_WalkScoredElements_CallbackError(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZSCAN", "test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		int64(0),
		[]string{"test-value-1", "0.8", "test-value-2", "0.8"},
	})

	newStorage := testMaybeNewStorageWithConn(t, c)

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
	c.Command("SSCAN", "test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		int64(0),
		[]string{"test-value"},
	})

	newStorage := testMaybeNewStorageWithConn(t, c)

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
	c.Command("SSCAN", "test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		int64(0),
		[]string{"test-value-1", "test-value-2"},
	})

	newStorage := testMaybeNewStorageWithConn(t, c)

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
	c.Command("SSCAN", "test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		int64(0),
		[]string{"test-value-1", "test-value-2"},
	})

	newStorage := testMaybeNewStorageWithConn(t, c)

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

	newStorage := testMaybeNewStorageWithConn(t, c)

	err := newStorage.WalkSet("test-key", nil, nil)
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_RedisStorage_WalkSet_CallbackError(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SSCAN", "test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		int64(0),
		[]string{"test-value-1", "test-value-2"},
	})

	newStorage := testMaybeNewStorageWithConn(t, c)

	err := newStorage.WalkSet("test-key", nil, func(element string) error {
		return maskAny(queryExecutionFailedError)
	})
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}
