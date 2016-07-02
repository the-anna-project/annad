# strategy
A strategy is basically an ordered chain of [CLGs](clg.md). In the simplest
form a strategy can be a wrapper around one single CLG. The complexity of a
strategy is infinite in theory. In praxis it will not be sufficient to have
overly complicated strategies. Due to their nature this would lead to way to
static behaviour. Instead smaller strategies should be layered across multiple
[stages](stage.md). That way the [output](output.md) of a previously executed
strategy can be reviewed and maybe even reset.

The process of creating a strategy can be described as follows.
[Connections](connection.md) point into a certain contextual direction of
strategies. Context is provided by some [input](input.md). Strategies
satisfying the current context through their associated connections are
selected and grouped for further processing. Strategies are not only connected
with contextual information, but also with other strategies. This helps
creating more sophisticated strategies, because it is possible to draw a path
through the neural connections into a certain direction. Once strategies are
selected, they are compared against their weight. Over all the highest weight
always wins, but how the weight is being calculated is fully dynamic, and
further dependening on the current context.
