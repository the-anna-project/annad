package redisstorage

import (
	"testing"

	"github.com/rafaeljusto/redigomock"
)

func Test_RedisStorage_GetID(t *testing.T) {
	firstStorage := NewRedisStorage(DefaultConfig())
	secondStorage := NewRedisStorage(DefaultConfig())

	if firstStorage.GetID() == secondStorage.GetID() {
		t.Fatal("expected", "different IDs", "got", "equal IDs")
	}
}

// Get

func Test_RedisStorage_Get_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("GET", "foo").Expect("bar")

	newConfig := defaultConfigWithConn(c)
	newStorage := NewRedisStorage(newConfig)

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

	newConfig := defaultConfigWithConn(c)
	newStorage := NewRedisStorage(newConfig)

	_, err := newStorage.Get("foo")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

// GetElementsByScore

func Test_RedisStorage_GetElementsByScore_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZREVRANGEBYSCORE", "foo", 0.8, 0.8, "LIMIT", 0, 3).Expect([]interface{}{"bar"})

	newConfig := defaultConfigWithConn(c)
	newStorage := NewRedisStorage(newConfig)

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

	newConfig := defaultConfigWithConn(c)
	newStorage := NewRedisStorage(newConfig)

	_, err := newStorage.GetElementsByScore("foo", 0.8, 3)
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

// GetHighestScoredElements

func Test_RedisStorage_GetHighestScoredElements_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZREVRANGE", "foo", 0, 2, "WITHSCORES").Expect([]interface{}{"one", "0.8", "two", "0.5"})

	newConfig := defaultConfigWithConn(c)
	newStorage := NewRedisStorage(newConfig)

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

	newConfig := defaultConfigWithConn(c)
	newStorage := NewRedisStorage(newConfig)

	_, err := newStorage.GetHighestScoredElements("foo", 3)
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

// Set

func Test_RedisStorage_Set_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SET", "foo", "bar").Expect("OK")

	newConfig := defaultConfigWithConn(c)
	newStorage := NewRedisStorage(newConfig)

	err := newStorage.Set("foo", "bar")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_RedisStorage_Set_NoSuccess(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SET", "foo", "bar").Expect("invalid")

	newConfig := defaultConfigWithConn(c)
	newStorage := NewRedisStorage(newConfig)

	err := newStorage.Set("foo", "bar")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_RedisStorage_Set_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SET", "foo", "bar").ExpectError(queryExecutionFailedError)

	newConfig := defaultConfigWithConn(c)
	newStorage := NewRedisStorage(newConfig)

	err := newStorage.Set("foo", "bar")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

// SetElementByScore

func Test_RedisStorage_SetElementByScore_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZADD", "key", 0.8, "element").Expect(int64(1))

	newConfig := defaultConfigWithConn(c)
	newStorage := NewRedisStorage(newConfig)

	err := newStorage.SetElementByScore("key", "element", 0.8)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_RedisStorage_SetElementByScore_NoSuccess(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZADD", "key", 0.8, "element").Expect(int64(3))

	newConfig := defaultConfigWithConn(c)
	newStorage := NewRedisStorage(newConfig)

	err := newStorage.SetElementByScore("key", "element", 0.8)
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_RedisStorage_SetElementByScore_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZADD", "key", 0.8, "element").ExpectError(queryExecutionFailedError)

	newConfig := defaultConfigWithConn(c)
	newStorage := NewRedisStorage(newConfig)

	err := newStorage.SetElementByScore("key", "element", 0.8)
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}
