# annactl
The client implementation of Anna's network API is reprsented by the command
line tool `annactl`.

### build
Compile the client and check the help usage for more information.

```
make client
.workspace/bin/annactl -h
```

### text interface
The [text interface](interface.md#text) is used to feed Anna with text input.

```
.workspace/bin/annactl interface text read plain some text input
```

### log control

The [log control](control.md#log) is used to configure Anna's logging behavior
at runtime.

```
# only log messages emitted with log level error or fatal
.workspace/bin/annactl control log set levels E F
```

```
# log messages emitted with log object core or network
.workspace/bin/annactl control log set objects core network
```

```
# log messages emitted with log verbosity 3 or lower
.workspace/bin/annactl control log set verbosity 3
```

```
# log messages emitted with default log levels
.workspace/bin/annactl control log reset levels
```

```
# log messages emitted with default log objects
.workspace/bin/annactl control log reset objects
```

```
# log messages emitted with default log verbosities
.workspace/bin/annactl control log reset verbosity
```
