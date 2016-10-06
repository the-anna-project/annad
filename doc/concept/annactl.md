# annactl
The command line client implementation of Anna's network API is reprsented by
the `annactl` binary. For convenience this can be compiled using the
[makefile](/doc/development/makefile.md). `annactl` is written using the
[cobra](https://github.com/spf13/cobra) CLI framework. The binary can be
executed to interact with Anna's remote API. In this communication the
[client](client.md) implementation forwards any request to the
[server](server.md). For a better understanding of this communication see the
[data flow](data_flow.md) documentation.

### autocompletion
For convenience there are [autocompletion scripts](autocompletion.md).

### build
Compile the client and check the help usage for more information.

```yaml
make annactl
.workspace/bin/annactl -h
```

### log control

The [log control](control.md#log) is used to configure Anna's logging behaviour
at runtime.

```yaml
# only log messages emitted with log level error or fatal
.workspace/bin/annactl control log set levels E F
```

```yaml
# log messages emitted with log object core or network
.workspace/bin/annactl control log set objects core network
```

```yaml
# log messages emitted with log verbosity 3 or lower
.workspace/bin/annactl control log set verbosity 3
```

```yaml
# log messages emitted with default log levels
.workspace/bin/annactl control log reset levels
```

```yaml
# log messages emitted with default log objects
.workspace/bin/annactl control log reset objects
```

```yaml
# log messages emitted with default log verbosities
.workspace/bin/annactl control log reset verbosity
```

### text interface
The [text interface](interface.md#text) is used to feed Anna with text input.

```yaml
# feed Anna with the text "some text input"
.workspace/bin/annactl interface text read plain some text input --expected="output expected"
```
