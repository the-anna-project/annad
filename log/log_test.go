package log

import (
	"strings"
	"testing"
	"time"

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

func testMustNewRootLogger(t *testing.T) spec.RootLogger {
	return &rootLogger{Args: []interface{}{}}
}

// object implements spec.Object and is used to provide an object for the log
// tags.
type object struct{}

func (o *object) GetID() spec.ObjectID {
	return spec.ObjectID("object-id")
}

func (o *object) GetType() spec.ObjectType {
	return spec.ObjectType("object-type")
}

func testMustNewObject(t *testing.T) spec.Object {
	return &object{}
}

// tracer implements spec.Tracer and is used to provide a tracer for the
// log tags.

type context struct {
	CLGTreeID string
	ID        string
	SessionID string
}

func (c *context) Deadline() (time.Time, bool) {
	return time.Time{}, false
}

func (c *context) Done() <-chan struct{} {
	return nil
}

func (c *context) Err() error {
	return nil
}

func (c *context) GetCLGTreeID() string {
	return c.CLGTreeID
}

func (c *context) GetID() string {
	return c.ID
}

func (c *context) GetSessionID() string {
	return c.SessionID
}

func (c *context) SetCLGTreeID(clgTreeID string) {
	c.CLGTreeID = clgTreeID
}

func (c *context) SetSessionID(sessionID string) {
	c.SessionID = sessionID
}

func (c *context) Value(key interface{}) interface{} {
	return nil
}

func testMustNewContext(t *testing.T) spec.Context {
	newContext := &context{
		CLGTreeID: "clg-tree-id",
		ID:        "context-id",
		SessionID: "session-id",
	}

	return newContext
}

// Test_Log_001 checks that different combinations of logger configuration and
// log tags result in logged messages as expected.
func Test_Log_001(t *testing.T) {
	testCases := []struct {
		Tags         spec.Tags
		F            string
		V            []interface{}
		Expected     string
		LogContextID string
		LogObjects   []spec.ObjectType
		LogLevels    []string
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
			Tags:         spec.Tags{C: nil, L: "weird", O: nil, V: 9},
			F:            "test message",
			V:            []interface{}{},
			Expected:     "invalid log level: weird",
			LogObjects:   []spec.ObjectType{},
			LogLevels:    []string{},
			LogVerbosity: 10,
		},

		// Logs should be logged with proper formats and operands.
		{
			Tags:         spec.Tags{C: nil, L: "I", O: nil, V: 9},
			F:            "test %s %s message",
			V:            []interface{}{"message", "test"},
			Expected:     "test message test message",
			LogObjects:   []spec.ObjectType{},
			LogLevels:    []string{},
			LogVerbosity: 10,
		},

		// Object object logs should be logged when all object logs are allowed by default.
		{
			Tags:         spec.Tags{C: nil, L: "I", O: testMustNewObject(t), V: 9},
			F:            "test message",
			V:            []interface{}{},
			Expected:     "[L: I] [O: object-type / object-id] [V:  9] test message",
			LogObjects:   []spec.ObjectType{},
			LogLevels:    []string{},
			LogVerbosity: 10,
		},

		// Object core logs should not be logged when only object impulse logs are allowed.
		{
			Tags:         spec.Tags{C: nil, L: "I", O: testMustNewObject(t), V: 9},
			F:            "test message",
			V:            []interface{}{},
			Expected:     "",
			LogObjects:   []spec.ObjectType{spec.ObjectType("impulse")},
			LogLevels:    []string{},
			LogVerbosity: 10,
		},

		// Verbosity 9 logs should not be logged when only verbosity 8 logs and lower are allowed.
		{
			Tags:         spec.Tags{C: nil, L: "I", O: testMustNewObject(t), V: 9},
			F:            "test message",
			V:            []interface{}{},
			Expected:     "",
			LogObjects:   []spec.ObjectType{},
			LogLevels:    []string{},
			LogVerbosity: 8,
		},

		// Debug logs should not be logged when only info logs are allowed.
		{
			Tags:         spec.Tags{C: nil, L: "D", O: testMustNewObject(t), V: 9},
			F:            "test message",
			V:            []interface{}{},
			Expected:     "",
			LogObjects:   []spec.ObjectType{},
			LogLevels:    []string{"I"},
			LogVerbosity: 10,
		},

		// Debug logs should be logged when only debug logs are allowed.
		{
			Tags:         spec.Tags{C: nil, L: "D", O: testMustNewObject(t), V: 9},
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
			Tags:         spec.Tags{C: testMustNewContext(t), L: "D", O: testMustNewObject(t), V: 9},
			F:            "test message",
			V:            []interface{}{},
			Expected:     "[C: context-id] [L: D] [O: object-type / object-id] [V:  9] test message",
			LogContextID: "",
			LogObjects:   []spec.ObjectType{},
			LogLevels:    []string{},
			LogVerbosity: 10,
		},

		// Log message using trace ID should NOT be logged when tracer is given in
		// log tags and log config's trace ID is different.
		{
			Tags:         spec.Tags{C: testMustNewContext(t), L: "D", O: testMustNewObject(t), V: 9},
			F:            "test message",
			V:            []interface{}{},
			Expected:     "",
			LogContextID: "other",
			LogObjects:   []spec.ObjectType{},
			LogLevels:    []string{},
			LogVerbosity: 10,
		},
	}

	for i, testCase := range testCases {
		newRootLogger := testMustNewRootLogger(t)

		newLogConfig := DefaultConfig()
		newLogConfig.ContextID = testCase.LogContextID
		newLogConfig.Objects = testCase.LogObjects
		newLogConfig.Levels = testCase.LogLevels
		newLogConfig.RootLogger = newRootLogger
		newLogConfig.Verbosity = testCase.LogVerbosity
		newLog := NewLog(newLogConfig)

		newLog.WithTags(testCase.Tags, testCase.F, testCase.V...)

		result := newRootLogger.(*rootLogger).ArgsToString()

		if (testCase.Expected == "" && result != "") || (!strings.Contains(result, testCase.Expected)) {
			t.Fatalf("%d. test case failed: logged message '%s' does not match expected result '%s'", i+1, result, testCase.Expected)
		}
	}
}

// Test_Log_002 checks that setting and resetting levels works as expected.
func Test_Log_002(t *testing.T) {
	newRootLogger := testMustNewRootLogger(t)

	newLogConfig := DefaultConfig()
	newLogConfig.RootLogger = newRootLogger
	newLog := NewLog(newLogConfig)

	// Logging a normal log message should work.
	newLog.WithTags(spec.Tags{C: nil, L: "W", O: testMustNewObject(t), V: newLogConfig.Verbosity}, "test message")
	result := newRootLogger.(*rootLogger).ArgsToString()
	if !strings.Contains(result, "test message") {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "test message")
	}

	// Reset the test logger to clean the log message of the previous logging.
	newRootLogger.(*rootLogger).ResetArgs()
	// Set levels so log messages should not be logged.
	err := newLog.SetLevels("I,D")
	if err != nil {
		t.Fatalf("Log.SetLevels returned error: %#v", err)
	}

	// Because the level we use is not allowed, the message should not be logged.
	newLog.WithTags(spec.Tags{C: nil, L: "W", O: testMustNewObject(t), V: newLogConfig.Verbosity}, "test message")
	result = newRootLogger.(*rootLogger).ArgsToString()
	if result != "" {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "")
	}

	// Resetting the levels should log the same log message.
	newRootLogger.(*rootLogger).ResetArgs()
	err = newLog.ResetLevels()
	if err != nil {
		t.Fatalf("Log.ResetLevels returned error: %#v", err)
	}
	newLog.WithTags(spec.Tags{C: nil, L: "W", O: testMustNewObject(t), V: newLogConfig.Verbosity}, "test message")
	result = newRootLogger.(*rootLogger).ArgsToString()
	if !strings.Contains(result, "test message") {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "test message")
	}
}

// Test_Log_003 checks that setting invalid levels throws an error.
func Test_Log_003(t *testing.T) {
	newRootLogger := testMustNewRootLogger(t)

	newLogConfig := DefaultConfig()
	newLogConfig.RootLogger = newRootLogger
	newLog := NewLog(newLogConfig)

	// Logging a normal log message should work.
	newLog.WithTags(spec.Tags{C: nil, L: "W", O: testMustNewObject(t), V: newLogConfig.Verbosity}, "test message")
	result := newRootLogger.(*rootLogger).ArgsToString()
	if !strings.Contains(result, "test message") {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "test message")
	}

	// Reset the test logger to clean the log message of the previous logging.
	newRootLogger.(*rootLogger).ResetArgs()
	// Set levels so log messages should not be logged.
	err := newLog.SetLevels("foo")
	if !IsInvalidLogLevel(err) {
		t.Fatalf("Log.SetLevels NOT returned proper error")
	}
}

// Test_Log_004 checks that setting levels with empty string does not have any
// effect.
func Test_Log_004(t *testing.T) {
	newRootLogger := testMustNewRootLogger(t)

	newLogConfig := DefaultConfig()
	newLogConfig.RootLogger = newRootLogger
	newLog := NewLog(newLogConfig)

	// Logging a normal log message should work.
	newLog.WithTags(spec.Tags{C: nil, L: "W", O: testMustNewObject(t), V: newLogConfig.Verbosity}, "test message")
	result := newRootLogger.(*rootLogger).ArgsToString()
	if !strings.Contains(result, "test message") {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "test message")
	}

	// Reset the test logger to clean the log message of the previous logging.
	newRootLogger.(*rootLogger).ResetArgs()
	// Try to set levels to empty string. This should have not effect.
	err := newLog.SetLevels("")
	if err != nil {
		t.Fatalf("Log.SetLevels returned error: %#v", err)
	}

	// Because nothing should have changed, the same log still should be logged.
	newLog.WithTags(spec.Tags{C: nil, L: "W", O: testMustNewObject(t), V: newLogConfig.Verbosity}, "test message")
	result = newRootLogger.(*rootLogger).ArgsToString()
	if !strings.Contains(result, "test message") {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "test message")
	}
}

// Test_Log_005 checks that setting and resetting objects works as expected.
func Test_Log_005(t *testing.T) {
	newRootLogger := testMustNewRootLogger(t)

	newLogConfig := DefaultConfig()
	newLogConfig.RootLogger = newRootLogger
	newLog := NewLog(newLogConfig)

	// Logging a normal log message should work.
	newLog.WithTags(spec.Tags{C: nil, L: "W", O: testMustNewObject(t), V: newLogConfig.Verbosity}, "test message")
	result := newRootLogger.(*rootLogger).ArgsToString()
	if !strings.Contains(result, "test message") {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "test message")
	}

	newLog.Register(spec.ObjectType("strategy-network"))
	newLog.Register(spec.ObjectType("impulse"))

	// Reset the test logger to clean the log message of the previous logging.
	newRootLogger.(*rootLogger).ResetArgs()
	// Set objects so log messages should not be logged.
	err := newLog.SetObjects("strategy-network,impulse")
	if err != nil {
		t.Fatalf("Log.SetObjects returned error: %#v", err)
	}

	// Because the object we use is not allowed, the message should not be logged.
	newLog.WithTags(spec.Tags{C: nil, L: "W", O: testMustNewObject(t), V: newLogConfig.Verbosity}, "test message")
	result = newRootLogger.(*rootLogger).ArgsToString()
	if result != "" {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "")
	}

	// Resetting the objects should log the same log message.
	newRootLogger.(*rootLogger).ResetArgs()
	err = newLog.ResetObjects()
	if err != nil {
		t.Fatalf("Log.ResetObjects returned error: %#v", err)
	}
	newLog.WithTags(spec.Tags{C: nil, L: "W", O: testMustNewObject(t), V: newLogConfig.Verbosity}, "test message")
	result = newRootLogger.(*rootLogger).ArgsToString()
	if !strings.Contains(result, "test message") {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "test message")
	}
}

// Test_Log_006 checks that setting invalid objects throws an error.
func Test_Log_006(t *testing.T) {
	newRootLogger := testMustNewRootLogger(t)

	newLogConfig := DefaultConfig()
	newLogConfig.RootLogger = newRootLogger
	newLog := NewLog(newLogConfig)

	// Logging a normal log message should work.
	newLog.WithTags(spec.Tags{C: nil, L: "W", O: testMustNewObject(t), V: newLogConfig.Verbosity}, "test message")
	result := newRootLogger.(*rootLogger).ArgsToString()
	if !strings.Contains(result, "test message") {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "test message")
	}

	// Reset the test logger to clean the log message of the previous logging.
	newRootLogger.(*rootLogger).ResetArgs()
	// Set objects so log messages should not be logged.
	err := newLog.SetObjects("foo")
	if !IsInvalidLogObject(err) {
		t.Fatalf("Log.SetObjects NOT returned proper error")
	}
}

// Test_Log_007 checks that setting objects with empty string does not have any
// effect.
func Test_Log_007(t *testing.T) {
	newRootLogger := testMustNewRootLogger(t)

	newLogConfig := DefaultConfig()
	newLogConfig.RootLogger = newRootLogger
	newLog := NewLog(newLogConfig)

	// Logging a normal log message should work.
	newLog.WithTags(spec.Tags{C: nil, L: "W", O: testMustNewObject(t), V: newLogConfig.Verbosity}, "test message")
	result := newRootLogger.(*rootLogger).ArgsToString()
	if !strings.Contains(result, "test message") {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "test message")
	}

	// Reset the test logger to clean the log message of the previous logging.
	newRootLogger.(*rootLogger).ResetArgs()
	// Try to set objects to empty string. This should have not effect.
	err := newLog.SetObjects("")
	if err != nil {
		t.Fatalf("Log.SetObjects returned error: %#v", err)
	}

	// Because nothing should have changed, the same log still should be logged.
	newLog.WithTags(spec.Tags{C: nil, L: "W", O: testMustNewObject(t), V: newLogConfig.Verbosity}, "test message")
	result = newRootLogger.(*rootLogger).ArgsToString()
	if !strings.Contains(result, "test message") {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "test message")
	}
}

// Test_Log_008 checks that setting and resetting verbosity works as expected.
func Test_Log_008(t *testing.T) {
	newRootLogger := testMustNewRootLogger(t)

	newLogConfig := DefaultConfig()
	newLogConfig.RootLogger = newRootLogger
	newLog := NewLog(newLogConfig)

	// Logging a normal log message should work.
	newLog.WithTags(spec.Tags{C: nil, L: "I", O: testMustNewObject(t), V: newLogConfig.Verbosity}, "test message")
	result := newRootLogger.(*rootLogger).ArgsToString()
	if !strings.Contains(result, "test message") {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "test message")
	}

	// Reset the test logger to clean the log message of the previous logging.
	newRootLogger.(*rootLogger).ResetArgs()
	// Set verbosity lower than what we are going to use next.
	err := newLog.SetVerbosity(newLogConfig.Verbosity - 1)
	if err != nil {
		t.Fatalf("Log.SetVerbosity returned error: %#v", err)
	}

	// Because the verbosity is lower than what we use, the message should not be
	// logged.
	newLog.WithTags(spec.Tags{C: nil, L: "I", O: testMustNewObject(t), V: newLogConfig.Verbosity}, "test message")
	result = newRootLogger.(*rootLogger).ArgsToString()
	if result != "" {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "")
	}

	// Resetting the verbosity should log the same log message.
	newRootLogger.(*rootLogger).ResetArgs()
	err = newLog.ResetVerbosity()
	if err != nil {
		t.Fatalf("Log.ResetVerbosity returned error: %#v", err)
	}
	newLog.WithTags(spec.Tags{C: nil, L: "I", O: testMustNewObject(t), V: newLogConfig.Verbosity}, "test message")
	result = newRootLogger.(*rootLogger).ArgsToString()
	if !strings.Contains(result, "test message") {
		t.Fatalf("logged message '%s' does not match expected result '%s'", result, "test message")
	}
}
