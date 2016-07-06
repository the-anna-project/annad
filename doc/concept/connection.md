# connection
The connection model is the most important concept of the Anna project. It can
be seen as a neural connection. Such a connections represent relationships
between certain information and behaviours. [Inputs](input.md),
[outputs](output.md) and [CLGs](clg.md) are wrired together that way. Many of
these connections joined together form Anna's neural [network](network.md).

The process of creating connections can be described as follows. Context is
provided by some [input](input.md). Connections satisfying the current context
through their associated information or behaviours are selected and grouped for
further processing. This helps creating more sophisticated connectins, because
it is possible to draw a path through the neural network into a certain
direction. Once connections are selected, they are compared against their
weight. Overall the highest weight always wins, but how the weight is being
calculated is fully dynamic, and further dependening on the current context.

The following picture illustrates the idea of the neural network. White circles
represent peers that are not (yet) connected. Red circles represent information
being connected into a certain direction. The direction is visualized by black
arrows. The dotted arrow shows a new connection being made to extend an already
existing strategy path. Creating connections is a process that itself is driven
by specific CLG. That way the neural network improves itself.

![connection](image/connection.png)
