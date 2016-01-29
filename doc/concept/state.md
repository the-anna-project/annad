# state
AI lives from its intelligence. This can only be achieved by information that
are more or less structured. The information that form some sort of
intelligence are called state. State is implemented as such that all objects
([core](core.md), [network](network.md), neuron, etc.) only contain their pure
state. All information an object holds is stored within its state. Business
logic and state is fully decoupled. That way Anna is able to completely backup
and restore her whole state.

The concept of having such a state model is important for a lot of reasons.
State can be backed up and restored. Once Anna is shut down, she can backup her
state. Once she is comming back to life, she can restore her state like nothing
happened. Imagine to test a specific part or structure of some self awareness.
Think about analysing, debugging and manipulating Anna's mind.
