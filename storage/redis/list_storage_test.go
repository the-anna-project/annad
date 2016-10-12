package redis

import (
	"testing"

	"github.com/rafaeljusto/redigomock"
)

func Test_ListStorage_PopFromList(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("BRPOP", "prefix:test-key").Expect("test-element")

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
	c.Command("BRPOP", "prefix:test-key").ExpectError(queryExecutionFailedError)

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
