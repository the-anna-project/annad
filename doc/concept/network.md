# network
The network can be seen as neural network. It brings all components together.
Impulses are coming in over the network's channels. They represent the actual
[input](input.md), which is provided to the input [CLG](clg.md). From there the
way the impulses are going is completely dynamic. Which way the impulses are
going relies on the current neural [connections](connection.md) given by the
networks [state](state.md). At some point in time the output CLG will be
triggered. Here the network hooks in again. As soon as the network receives a
message from the output CLG, it is streamed back to the client. In case the
given expectation, if any given, does not match the responded output, the
involved connections of the failed iteration will be "punished" by being ask to
calculate again, otherwise "rewarded", what means that the successful
connection path is persisted.

The following picture illustrates the only hard coded Input and Output CLGs
using the red circles. They represent the entrance and exit of the neural
network. The black arrows represent information and behaviour flow into certain
directions from peers to peers. Here each peer decides on its own how to
proceed, until the walk through the neural network is finished and some final
output is generated.

![network](image/network.png)
