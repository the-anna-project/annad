package redis

import (
	"strings"
	"testing"

	"github.com/the-anna-project/collection"
	"github.com/the-anna-project/log"
	servicespec "github.com/the-anna-project/spec/service"
)

// rootLogger implements spec.RootLogger and is used to capture log messages.
type rootLogger struct {
	Args []interface{}
}

func (rl *rootLogger) ArgsToString() string {
	args := ""
	for _, v := range rl.Args {
		if arg, ok := v.(error); ok {
			args += " " + arg.Error()
		}
		if arg, ok := v.(string); ok {
			args += " " + arg
		}
	}

	return args[1:]
}

func (rl *rootLogger) Log(v ...interface{}) error {
	rl.Args = v
	return nil
}

func (rl *rootLogger) ResetArgs() {
	rl.Args = []interface{}{}
}

func testMustNewRootLogger(t *testing.T) servicespec.RootLogger {
	return &rootLogger{Args: []interface{}{}}
}

func Test_RedisStorage_retryErrorLogger(t *testing.T) {
	storageService := New()
	storageService.SetPrefix("test-prefix")

	logService := log.New()
	rootLoggerService := testMustNewRootLogger(t)
	logService.SetRootLogger(rootLoggerService)

	serviceCollection := collection.New()
	serviceCollection.SetLogService(logService)

	logService.SetServiceCollection(serviceCollection)
	storageService.SetServiceCollection(serviceCollection)

	storageService.(*service).retryErrorLogger(invalidConfigError, 0)
	result := rootLoggerService.(*rootLogger).ArgsToString()

	if !strings.Contains(result, invalidConfigError.Error()) {
		t.Fatal("expected", invalidConfigError.Error(), "got", result)
	}
}

func Test_RedisStorage_withPrefix(t *testing.T) {
	storageService := New()
	storageService.SetPrefix("test-prefix")

	expected := "test-prefix:my:test:key"
	newKey := storageService.(*service).withPrefix("my", "test", "key")
	if newKey != expected {
		t.Fatal("expected", expected, "got", newKey)
	}
}

func Test_RedisStorage_withPrefix_Empty(t *testing.T) {
	storageService := New()
	storageService.SetPrefix("test-prefix")

	newKey := storageService.(*service).withPrefix()
	if newKey != "test-prefix" {
		t.Fatal("expected", "test-prefix", "got", newKey)
	}
}
