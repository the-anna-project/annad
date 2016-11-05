package input

import (
	"reflect"
	"testing"

	redigo "github.com/garyburd/redigo/redis"
	"github.com/rafaeljusto/redigomock"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/context"
	"github.com/xh3b4sd/anna/key"
	servicespec "github.com/xh3b4sd/anna/service/spec"
	systemspec "github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage"
	"github.com/xh3b4sd/anna/storage/redis"
)

type testErrorIDService struct{}

// New is only a test implementation of servicespec.ID to do nothing but
// returning some error we can check against.
func (f *testErrorIDService) New() (string, error) {
	return "", maskAny(invalidConfigError)
}

func (f *testErrorIDService) WithType(idType servicespec.IDType) (string, error) {
	return "", nil
}

type testIDService struct{}

// New is only a test implementation of servicespec.ID to do nothing but
// returning some error we can check against.
func (f *testIDService) New() (string, error) {
	return "new-ID", nil
}

func (f *testIDService) WithType(idType servicespec.IDType) (string, error) {
	return "", nil
}

type testServiceCollection struct {
	IDService servicespec.ID
}

func (c *testServiceCollection) FS() servicespec.FS {
	return nil
}

func (c *testServiceCollection) ID() servicespec.ID {
	return c.IDService
}

func (c *testServiceCollection) Permutation() servicespec.Permutation {
	return nil
}

func (c *testServiceCollection) Random() servicespec.Random {
	return nil
}

func (c *testServiceCollection) Shutdown() {
}

func (c *testServiceCollection) TextOutput() servicespec.TextOutput {
	return nil
}

func testMustNewErrorServiceCollection(t *testing.T) servicespec.Collection {
	return &testServiceCollection{
		IDService: &testErrorIDService{},
	}
}

func testMustNewServiceCollection(t *testing.T) servicespec.Collection {
	return &testServiceCollection{
		IDService: &testIDService{},
	}
}

func testMustNewStorageCollection(t *testing.T) systemspec.StorageCollection {
	newCollection, err := storage.NewCollection(storage.DefaultCollectionConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newCollection
}

func testMustNewStorageCollectionWithConn(t *testing.T, c redigo.Conn) systemspec.StorageCollection {
	newFeatureStorage, err := redis.NewStorage(redis.DefaultStorageConfigWithConn(c))
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	newGeneralStorage, err := redis.NewStorage(redis.DefaultStorageConfigWithConn(c))
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	newStorageCollectionConfig := storage.DefaultCollectionConfig()
	newStorageCollectionConfig.FeatureStorage = newFeatureStorage
	newStorageCollectionConfig.GeneralStorage = newGeneralStorage
	newStorageCollection, err := storage.NewCollection(newStorageCollectionConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newStorageCollection
}

func testMustNewNetworkPayload(t *testing.T, ctx systemspec.Context, input string) systemspec.NetworkPayload {
	newNetworkPayloadConfig := api.DefaultNetworkPayloadConfig()
	newNetworkPayloadConfig.Args = []reflect.Value{reflect.ValueOf(input)}
	newNetworkPayloadConfig.Context = ctx
	newNetworkPayloadConfig.Destination = "destination"
	newNetworkPayloadConfig.Sources = []systemspec.ObjectID{"source"}
	newNetworkPayload, err := api.NewNetworkPayload(newNetworkPayloadConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newNetworkPayload
}

func Test_CLG_Input_KnownInputSequence(t *testing.T) {
	newCLG := MustNew()
	newCtx := context.MustNew()
	newStorageCollection := testMustNewStorageCollection(t)

	// Create record for the test input.
	informationID := "123"
	newInput := "test input"
	informationIDKey := key.NewNetworkKey("information-sequence:%s:information-id", newInput)
	err := newStorageCollection.General().Set(informationIDKey, informationID)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Set prepared storage to CLG we want to test.
	newCLG.(*clg).StorageCollection = newStorageCollection

	// Execute CLG.
	err = newCLG.(*clg).calculate(newCtx, newInput)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Check if the information ID was set to the context.
	injectedInformationID, _ := newCtx.GetInformationID()
	if informationID != injectedInformationID {
		t.Fatal("expected", informationID, "got", injectedInformationID)
	}
}

func Test_CLG_Input_UnknownInputSequence(t *testing.T) {
	newCLG := MustNew()
	newCtx := context.MustNew()
	newServiceCollection := testMustNewServiceCollection(t)
	newStorageCollection := testMustNewStorageCollection(t)

	// Note we do not create a record for the test input. This test is about an
	// unknown input sequence.
	newInput := "test input"

	// Set prepared storage to CLG we want to test.
	newCLG.(*clg).ServiceCollection = newServiceCollection
	newCLG.(*clg).StorageCollection = newStorageCollection

	// Execute CLG.
	err := newCLG.(*clg).calculate(newCtx, newInput)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Check if the information ID was set to the context.
	injectedInformationID, _ := newCtx.GetInformationID()
	if injectedInformationID != "new-ID" {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_CLG_Input_DataProperlyStored(t *testing.T) {
	newCLG := MustNew()
	newCtx := context.MustNew()
	newStorageCollection := testMustNewStorageCollection(t)
	newServiceCollection := testMustNewServiceCollection(t)

	// Note we do not create a record for the test input. This test is about an
	// unknown input sequence.
	newInput := "test input"
	// Our test ID factory always returns the same ID. That way we are able to
	// check for the ID being used during the test.
	newID, err := newServiceCollection.ID().New()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Set prepared storage to CLG we want to test.
	newCLG.(*clg).ServiceCollection = newServiceCollection
	newCLG.(*clg).StorageCollection = newStorageCollection

	// Execute CLG.
	err = newCLG.(*clg).calculate(newCtx, newInput)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	informationIDKey := key.NewNetworkKey("information-sequence:%s:information-id", newInput)
	storedID, err := newStorageCollection.General().Get(informationIDKey)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if storedID != string(newID) {
		t.Fatal("expected", newID, "got", storedID)
	}

	informationSequenceKey := key.NewNetworkKey("information-id:%s:information-sequence", newID)
	storedInput, err := newStorageCollection.General().Get(informationSequenceKey)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if storedInput != newInput {
		t.Fatal("expected", newInput, "got", storedInput)
	}
}

func Test_CLG_Input_IDServiceError(t *testing.T) {
	newCLG := MustNew()
	newCtx := context.MustNew()
	newStorageCollection := testMustNewStorageCollection(t)
	newServiceCollection := testMustNewErrorServiceCollection(t)

	// Note we do not create a record for the test input. This test is about an
	// unknown input sequence.
	newInput := "test input"

	// Set prepared storage to CLG we want to test.
	newCLG.(*clg).ServiceCollection = newServiceCollection
	newCLG.(*clg).StorageCollection = newStorageCollection

	// Execute CLG.
	err := newCLG.(*clg).calculate(newCtx, newInput)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_CLG_Input_SetInformationIDError(t *testing.T) {
	newCLG := MustNew()
	newCtx := context.MustNew()
	newServiceCollection := testMustNewServiceCollection(t)

	// Prepare the storage connection to fake a returned error.
	newInput := "test input"
	informationIDKey := key.NewNetworkKey("information-sequence:%s:information-id", newInput)
	// Our test ID factory always returns the same ID. That way we are able to
	// check for the ID being used during the test.
	newID, err := newServiceCollection.ID().New()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	c := redigomock.NewConn()
	c.Command("GET", "prefix:"+informationIDKey).ExpectError(redigo.ErrNil)
	c.Command("SET", "prefix:"+informationIDKey, string(newID)).ExpectError(invalidConfigError)
	newStorageCollection := testMustNewStorageCollectionWithConn(t, c)

	// Set prepared storage to CLG we want to test.
	newCLG.(*clg).StorageCollection = newStorageCollection
	newCLG.(*clg).ServiceCollection = newServiceCollection

	// Execute CLG.
	err = newCLG.(*clg).calculate(newCtx, newInput)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_CLG_Input_SetInformationSequenceError(t *testing.T) {
	newCLG := MustNew()
	newCtx := context.MustNew()
	newServiceCollection := testMustNewServiceCollection(t)

	// Prepare the storage connection to fake a returned error.
	newInput := "test input"
	informationIDKey := key.NewNetworkKey("information-sequence:%s:information-id", newInput)
	// Our test ID factory always returns the same ID. That way we are able to
	// check for the ID being used during the test.
	newID, err := newServiceCollection.ID().New()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	informationSequenceKey := key.NewNetworkKey("information-id:%s:information-sequence", newID)

	c := redigomock.NewConn()
	c.Command("GET", "prefix:"+informationIDKey).ExpectError(redigo.ErrNil)
	c.Command("SET", "prefix:"+informationIDKey, string(newID)).Expect("OK")
	c.Command("SET", "prefix:"+informationSequenceKey, newInput).ExpectError(invalidConfigError)
	newStorageCollection := testMustNewStorageCollectionWithConn(t, c)

	// Set prepared storage to CLG we want to test.
	newCLG.(*clg).StorageCollection = newStorageCollection
	newCLG.(*clg).ServiceCollection = newServiceCollection

	// Execute CLG.
	err = newCLG.(*clg).calculate(newCtx, newInput)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_CLG_Input_GetInformationIDError(t *testing.T) {
	newCLG := MustNew()
	newCtx := context.MustNew()

	newInput := "test input"
	informationIDKey := key.NewNetworkKey("information-sequence:%s:information-id", newInput)

	// Prepare the storage connection to fake a returned error.
	c := redigomock.NewConn()
	c.Command("GET", "prefix:"+informationIDKey).ExpectError(invalidConfigError)
	newStorageCollection := testMustNewStorageCollectionWithConn(t, c)

	// Set prepared storage to CLG we want to test.
	newCLG.SetStorageCollection(newStorageCollection)

	// Execute CLG.
	err := newCLG.(*clg).calculate(newCtx, newInput)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}
