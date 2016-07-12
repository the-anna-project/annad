# network
The network can be seen as neural network. It brings all components together.
Signals coming in over the network's [gateway](gateway.md). They are
transformed to [impulses](impulse.md). The incoming impulse and its
corresponding actual [input](input.md) are prepared as input request. This
input request is provided to the Input [CLG](clg.md). From there the way the
impulse is going is completely dynamic. Which way the impulse goes relies on
the current neural [connections](connection.md) given by the networks
[state](state.md). At some point in time the Output CLG will be triggered. Here
the network hooks in again. As soon as the network receives a message from the
output CLG, it is returned. In case the given expectation, if any given, does
not match the responded output, the involved connections of the failed
iteration will be "punished", otherwise "rewarded".

The following picture illustrates the only hard coded Input and Output CLGs
using the red circles. They represent the entrance and exit of the neural
network. The black arrows represent information and behaviour flow into certain
directions from peers to peers. Here each peer decides on its own how to
proceed, until the walk through the neural network is finished and some final
output is generated.

![network](image/network.png)
