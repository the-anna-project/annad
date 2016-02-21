package log

import (
	"strings"
	"testing"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/spec"
)

// rootLogger implements spec.RootLogger and is used to capture log messages.

type rootLogger struct {
	Args []interface{}
}

func (rl *rootLogger) ArgsToString() string {
	args := ""
	for _, v := range rl.Args {
		if arg, ok := v.(string); ok {
			args += arg
		}
	}
	return args
}

func (rl *rootLogger) Println(v ...interface{}) {
	rl.Args = v
}

func (rl *rootLogger) ResetArgs() {
	rl.Args = []interface{}{}
}

// core implements spec.Object and is used to provide an object for the log tags.

type core struct{}

func (c core) GetID() spec.ObjectID {
	return spec.ObjectID("")
}

func (c core) GetType() spec.ObjectType {
	return common.ObjectType.Core
}

func (c core) MarshalJSON() ([]byte, error) {
	return nil, nil
}
func (c core) UnmarshalJSON(bytes []byte) error {
	return nil
}

// tracer implementes spec.Tracer and is used to provide a tracer for the
// log tags.

type tracer struct{}

func (t tracer) GetTraceID() spec.TraceID {
	return spec.TraceID("test-trace-id")
}

// Test_Log_001 checks that different combinations of logger configuration and
// log tags result in logged messages as expected.
func Test_Log_001(t *testing.T) {
	testCases := []struct {
		Tags         spec.Tags
		F            string
		V            []interface{}
		Expected     string
		LogObjects   []spec.ObjectType
		LogLevels    []string
		LogTraceID   spec.TraceID
		LogVerbosity int
	}{
		// Logs should not be logged when no tags are given.
		{
			Tags:         spec.Tags{},
			F:            "test message",
			V:            []interface{}{},
			Expected:     "",
			LogObjects:   []spec.ObjectType{},
			LogLevels:    []string{},
			LogVerbosity: 10,
		},

		// Logs should not be logged using invalid log level.
		{
			Tags:         spec.Tags{L: "weird", V: 10},
			F:            "test message",
			V:            []interface{}{},
			Expected:     "invalid log level: weird",
			LogObjects:   []spec.ObjectType{},
			LogLevels:    []string{},
			LogVerbosity: 10,
		},

		// Logs should be logged with proper formats and operands.
		{
			Tags:         spec.Tags{L: "I", V: 9},
			F:            "test %s %s message",
			V:            []interface{}{"message", "test"},
			Expected:     "test message test message",
			LogObjects:   []spec.ObjectType{},
			LogLevels:    []string{},
			LogVerbosity: 10,
		},

		// Object core logs should be logged when all object logs are allowed by default.
		{
			Tags:         spec.Tags{O: core{}, L: "I", V: 9},
			F:            "test message",
			V:            []interface{}{},
			Expected:     "[O: core ",
			LogObjects:   []spec.ObjectType{},
			LogLevels:    []string{},
			LogVerbosity: 10,
		},

		// Object core logs should not be logged when only object network logs are allowed.
		{
			Tags:         spec.Tags{O: core{}, L: "I", V: 9},
			F:            "test message",
			V:            []interface{}{},
			Expected:     "",
			LogObjects:   []spec.ObjectType{common.ObjectType.StrategyNetwork},
			LogLevels:    []string{},
			LogVerbosity: 10,
		},

		// Verbosity 9 logs should not be logged when only verbosity 8 logs and lower are allowed.
		{
			Tags:         spec.Tags{O: core{}, L: "I", V: 9},
			F:            "test message",
			V:            []interface{}{},
			Expected:     "",
			LogObjects:   []spec.ObjectType{},
			LogLevels:    []string{},
			LogVerbosity: 8,
		},

		// Debug logs should not be logged when only info logs are allowed.
		{
			Tags:         spec.Tags{O: core{}, L: "D", V: 9},
			F:            "test message",
			V:            []interface{}{},
			Expected:     "",
			LogObjects:   []spec.ObjectType{},
			LogLevels:    []string{"I"},
			LogVerbosity: 10,
		},

		// Debug logs should be logged when only debug logs are allowed.
		{
			Tags:         spec.Tags{O: core{}, L: "D", V: 9},
			F:            "test message",
			V:            []interface{}{},
			Expected:     "test message",
			LogObjects:   []spec.ObjectType{},
			LogLevels:    []string{"D"},
			LogVerbosity: 10,
		},

		// Log message using trace ID should be logged when tracer is given in log
		// tags and log config's trace ID is empty by default.
		{
			Tags:         spec.Tags{O: core{}, T: tracer{}, L: "D", V: 9},
			F:            "test message",
			V:            []interface{}{},
			Expected:     "test-trace-id",
			LogObjects:   []spec.ObjectType{},
			LogLevels:    []string{},
			LogTraceID:   spec.TraceID(""),
			LogVerbosity: 10,
		},

		// Log message using trace ID should NOT be logged when tracer is given in
		// log tags and log config's trace ID is different.
		{
			Tags:         spec.Tags{O: core{}, T: tracer{}, L: "D", V: 9},
			F:            "test message",
			V:            []interface{}{},
			Expected:     "",
			LogObjects:   []spec.ObjectType{},
			LogLevels:    []string{},
			LogTraceID:   spec.TraceID("other"),
			LogVerbosity: 10,
		},
	}

	for i, testCase := range testCases {
		newRootLogger := &rootLogger{Args: []interface{}{}}

		newLogConfig := DefaultConfig()
		newLogConfig.Objects = testCase.LogObjects
		newLogConfig.Levels = testCase.LogLevels
		newLogConfig.TraceID = testCase.LogTraceID
		newLogConfig.RootLogger = newRootLogger
		newLogConfig.Verbosity = testCase.LogVerbosity
		newLog := NewLog(newLogConfig)

		newLog.WithTags(testCase.Tags, testCase.F, testCase.V...)

		result := newRootLogger.ArgsToString()

		if (testCase.Expected == "" && result != "") || (!strings.Contains(result, testCase.Expected)) {
			t.Fatalf("%d. test case failed: logged message '%s' does not match expected result '%s'", i+1, result, testCase.Expected)
		}
	}
}

// Test_Log_002 checks that setting and resetting levels works as expected.
func Test_Log_002(t *testing.T) {
	newRootLogger := &rootLogger{Args: []interface{}{}}

	newLogConfig := DefaultConfig()
	newLogConfig.RootLogger = newRootLogger
	newLog := NewLog(newLogConfig)

	// Logging a normal log message should work.
	newLog.WithTags(spec.Tags{L: "W", O: core{}, V: newLogConfig.Verbosity}, "test message")
	result := newRootLogger.ArgsToString()
	if !strings.Contains(result, "test message") {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "test message")
	}

	// Reset the test logger to clean the log message of the previous logging.
	newRootLogger.ResetArgs()
	// Set levels so log messages should not be logged.
	err := newLog.SetLevels("I,D")
	if err != nil {
		t.Fatalf("Log.SetLevels returned error: %#v", err)
	}

	// Because the level we use is not allowed, the message should not be logged.
	newLog.WithTags(spec.Tags{L: "W", O: core{}, V: newLogConfig.Verbosity}, "test message")
	result = newRootLogger.ArgsToString()
	if result != "" {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "")
	}

	// Resetting the levels should log the same log message.
	newRootLogger.ResetArgs()
	err = newLog.ResetLevels()
	if err != nil {
		t.Fatalf("Log.ResetLevels returned error: %#v", err)
	}
	newLog.WithTags(spec.Tags{L: "W", O: core{}, V: newLogConfig.Verbosity}, "test message")
	result = newRootLogger.ArgsToString()
	if !strings.Contains(result, "test message") {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "test message")
	}
}

// Test_Log_003 checks that setting invalid levels throws an error.
func Test_Log_003(t *testing.T) {
	newRootLogger := &rootLogger{Args: []interface{}{}}

	newLogConfig := DefaultConfig()
	newLogConfig.RootLogger = newRootLogger
	newLog := NewLog(newLogConfig)

	// Logging a normal log message should work.
	newLog.WithTags(spec.Tags{L: "W", O: core{}, V: newLogConfig.Verbosity}, "test message")
	result := newRootLogger.ArgsToString()
	if !strings.Contains(result, "test message") {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "test message")
	}

	// Reset the test logger to clean the log message of the previous logging.
	newRootLogger.ResetArgs()
	// Set levels so log messages should not be logged.
	err := newLog.SetLevels("foo")
	if !IsInvalidLogLevel(err) {
		t.Fatalf("Log.SetLevels NOT returned proper error")
	}
}

// Test_Log_004 checks that setting levels with empty string does not have any
// effect.
func Test_Log_004(t *testing.T) {
	newRootLogger := &rootLogger{Args: []interface{}{}}

	newLogConfig := DefaultConfig()
	newLogConfig.RootLogger = newRootLogger
	newLog := NewLog(newLogConfig)

	// Logging a normal log message should work.
	newLog.WithTags(spec.Tags{L: "W", O: core{}, V: newLogConfig.Verbosity}, "test message")
	result := newRootLogger.ArgsToString()
	if !strings.Contains(result, "test message") {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "test message")
	}

	// Reset the test logger to clean the log message of the previous logging.
	newRootLogger.ResetArgs()
	// Try to set levels to empty string. This should have not effect.
	err := newLog.SetLevels("")
	if err != nil {
		t.Fatalf("Log.SetLevels returned error: %#v", err)
	}

	// Because nothing should have changed, the same log still should be logged.
	newLog.WithTags(spec.Tags{L: "W", O: core{}, V: newLogConfig.Verbosity}, "test message")
	result = newRootLogger.ArgsToString()
	if !strings.Contains(result, "test message") {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "test message")
	}
}

// Test_Log_005 checks that setting and resetting objects works as expected.
func Test_Log_005(t *testing.T) {
	newRootLogger := &rootLogger{Args: []interface{}{}}

	newLogConfig := DefaultConfig()
	newLogConfig.RootLogger = newRootLogger
	newLog := NewLog(newLogConfig)

	// Logging a normal log message should work.
	newLog.WithTags(spec.Tags{L: "W", O: core{}, V: newLogConfig.Verbosity}, "test message")
	result := newRootLogger.ArgsToString()
	if !strings.Contains(result, "test message") {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "test message")
	}

	// Reset the test logger to clean the log message of the previous logging.
	newRootLogger.ResetArgs()
	// Set objects so log messages should not be logged.
	err := newLog.SetObjects("strategy-network,impulse")
	if err != nil {
		t.Fatalf("Log.SetObjects returned error: %#v", err)
	}

	// Because the object we use is not allowed, the message should not be logged.
	newLog.WithTags(spec.Tags{L: "W", O: core{}, V: newLogConfig.Verbosity}, "test message")
	result = newRootLogger.ArgsToString()
	if result != "" {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "")
	}

	// Resetting the objects should log the same log message.
	newRootLogger.ResetArgs()
	err = newLog.ResetObjects()
	if err != nil {
		t.Fatalf("Log.ResetObjects returned error: %#v", err)
	}
	newLog.WithTags(spec.Tags{L: "W", O: core{}, V: newLogConfig.Verbosity}, "test message")
	result = newRootLogger.ArgsToString()
	if !strings.Contains(result, "test message") {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "test message")
	}
}

// Test_Log_006 checks that setting invalid objects throws an error.
func Test_Log_006(t *testing.T) {
	newRootLogger := &rootLogger{Args: []interface{}{}}

	newLogConfig := DefaultConfig()
	newLogConfig.RootLogger = newRootLogger
	newLog := NewLog(newLogConfig)

	// Logging a normal log message should work.
	newLog.WithTags(spec.Tags{L: "W", O: core{}, V: newLogConfig.Verbosity}, "test message")
	result := newRootLogger.ArgsToString()
	if !strings.Contains(result, "test message") {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "test message")
	}

	// Reset the test logger to clean the log message of the previous logging.
	newRootLogger.ResetArgs()
	// Set objects so log messages should not be logged.
	err := newLog.SetObjects("foo")
	if !IsInvalidLogObject(err) {
		t.Fatalf("Log.SetObjects NOT returned proper error")
	}
}

// Test_Log_007 checks that setting objects with empty string does not have any
// effect.
func Test_Log_007(t *testing.T) {
	newRootLogger := &rootLogger{Args: []interface{}{}}

	newLogConfig := DefaultConfig()
	newLogConfig.RootLogger = newRootLogger
	newLog := NewLog(newLogConfig)

	// Logging a normal log message should work.
	newLog.WithTags(spec.Tags{L: "W", O: core{}, V: newLogConfig.Verbosity}, "test message")
	result := newRootLogger.ArgsToString()
	if !strings.Contains(result, "test message") {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "test message")
	}

	// Reset the test logger to clean the log message of the previous logging.
	newRootLogger.ResetArgs()
	// Try to set objects to empty string. This should have not effect.
	err := newLog.SetObjects("")
	if err != nil {
		t.Fatalf("Log.SetObjects returned error: %#v", err)
	}

	// Because nothing should have changed, the same log still should be logged.
	newLog.WithTags(spec.Tags{L: "W", O: core{}, V: newLogConfig.Verbosity}, "test message")
	result = newRootLogger.ArgsToString()
	if !strings.Contains(result, "test message") {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "test message")
	}
}

// Test_Log_008 checks that setting and resetting verbosity works as expected.
func Test_Log_008(t *testing.T) {
	newRootLogger := &rootLogger{Args: []interface{}{}}

	newLogConfig := DefaultConfig()
	newLogConfig.RootLogger = newRootLogger
	newLog := NewLog(newLogConfig)

	// Logging a normal log message should work.
	newLog.WithTags(spec.Tags{L: "I", O: core{}, V: newLogConfig.Verbosity}, "test message")
	result := newRootLogger.ArgsToString()
	if !strings.Contains(result, "test message") {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "test message")
	}

	// Reset the test logger to clean the log message of the previous logging.
	newRootLogger.ResetArgs()
	// Set verbosity lower than what we are going to use next.
	err := newLog.SetVerbosity(newLogConfig.Verbosity - 1)
	if err != nil {
		t.Fatalf("Log.SetVerbosity returned error: %#v", err)
	}

	// Because the verbosity is lower than what we use, the message should not be
	// logged.
	newLog.WithTags(spec.Tags{L: "I", O: core{}, V: newLogConfig.Verbosity}, "test message")
	result = newRootLogger.ArgsToString()
	if result != "" {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "")
	}

	// Resetting the verbosity should log the same log message.
	newRootLogger.ResetArgs()
	err = newLog.ResetVerbosity()
	if err != nil {
		t.Fatalf("Log.ResetVerbosity returned error: %#v", err)
	}
	newLog.WithTags(spec.Tags{L: "I", O: core{}, V: newLogConfig.Verbosity}, "test message")
	result = newRootLogger.ArgsToString()
	if !strings.Contains(result, "test message") {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "test message")
	}
}
