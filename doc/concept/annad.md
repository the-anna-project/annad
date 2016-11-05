# annad
This is the ten thousand feet view of the anna daemon. To understand how she
looks like from the very top we consider the following layers.

- The `i/o` layer describes a set of network protocols Anna understands. Data
  can be written to and retrieved from her over network. I/O is flowing to and
  coming from the server.

- The `server` layer describes the actual server listening for traffic of
  implemented network protocols. It provides so called
  [interfaces](interface.md) that are used to differentiate between different
  types of inputs that serve different types of purposes.

- The `network` layer describes the implementation of Anna's most inner
  workings. This can be seen as the neural network. It bundles everything
  around data processing and intelligence. The network itself contains
  [CLGs](clg.md) acting as some sort of artificial neurons. They connect to
  each other and form the network.

- The `storage` layer describes the data storage responsible for storing any
  kind of information like the neural [connections](connection.md). This can be
  seen as Anna's memory.

This is how it basically looks like. Note that the pale boxes represent ideas
that are not yet implemented.

![anna](image/anna.png)
