# connection
The connection model is the most important concept of the Anna project. A
connection can be seen as a neural connection. Such connections represent
relationships between certain information and behaviors, represented in a multi
dimensional space. [Inputs](input.md), [outputs](output.md) and [CLGs](clg.md)
are wired together that way. Many of these connections interacting together in
a dynamic fashion form Anna's neural [network](network.md).

### space
The connection space can be seen as a multi dimensional vector space. In theory
it's dimensions can represent everything: space, time or even conceptional
weights representing something like emotions. Information and behaviors are
mapped onto the connection space. Here similarities between information, and
similarities between behaviors can be calculated.

The following picture illustrates the connection space. For simplicity it only
shows two dimensions, `x` and `y`. We see two different connection paths within
the connection space. Here some peers are pretty near to each other. We assume
that a smaller distance is an indicator for common connection patterns which
are aligned over time under the effect of pulling dimensions.

![connection space](image/connection_space.png)

### balance
[The balance system](distribution.md#balance-system) is used to make dimensions
create tensions on coordinates within the connection space. Each dimension is
pulling into its own direction based on learned patterns. The important idea
behind the tensions created here is that they are balanced. One dimension
pulling coordinates spontaneously into a certain direction makes another
dimension move into the same direction. This creates a simulation of
psychological interactions on different levels. Like disgust and attraction
cannot go together. Aim of such a balance system is to provide buildin
mechanisms of leveraging event interactions.

Spontaneous modificatins of connections lead to variances within applied
behaviors. Leads some behavior to successful results, slight adjustments of a
connection path's coordinates can be stored, which cause the connection path to
be updated. This will lead connection paths to be improved over time, because
behavior was adjusted to the given environmental feedback Anna is receiving.

### creation TODO
When creating new connections it is important to not create weak connections.
Each connection that exists only exists because it brought some kind of value
in the past. The process of creating connections is a continuous task, that is
fully dynamic and learned by experience and can be described as follows.

When information is provided in the form of input, and no connection path for
the current input sequence exists yet, it is mapped onto the connection space.
Each character of the input sequence is represented by their own coordinates.
That way an input sequence draws a connection path within the connection space.
This is then the connection path of information.

The creation of behavior connections takes place on the CLG level. Connections
of behaviors are only persisted in case they belong to a successful behavior
connection path that in turn helped to solve some problem. So, the connection
path of behaviors is made out of connections that belong to behaviors. Here a
connection is represented as single key-value pair where it's key being the
string representation of coordinates links to another string representation of
coordinates. A CLG is identified using its unique coordinates within the
connection space in combination with its very unique connection path, which in
turn is represented by a CLG tree ID. New coordinates are randomly created with
some offset to the coordinates of the requesting CLG. A reference from
coordinates to an actual CLG ID is also stored for mapping of coordinates to
actual behavior later on. This is then the connection path of behaviors.

All newly created connections are persisted within a trial scope in the first
place. Such trial scope is basically realized by some specific storage prefix
that simply identifies certain key-value pairs as being part of such a trial
scope. This is done to label the current creation process of behavior
connections to something that is volatile. If the neural network succeeds to
solve a problem with some newly created connection path, the behavior
connections stored and marked within a trial scope are persisted as regular
connections. In case the created connection path did not lead to some
successful operation, all behavior connections marked within a trial scope are
simply removed again. Anyway, there needs a decision to be made to forward to
some CLG. These strategies are considered when it comes to draw new connections
within the connection space.

1. *Bias* is some manually provided hint, intended to guide some connection path
   into a certain direction. Read more on this in this issue:
   https://github.com/xh3b4sd/anna/issues/44.

2. *Intuition* is some sort of vague feeling that points into a certain
   direction. Drawing distantly related connections across multiple levels can
   gather information and generate new relations between peers.

3. *Copy* connections from other branches looks up possible connection structures
   from different problem domains. Connections that have been useful in one
   problem domain might be useful as well in another.

4. *Random* connections can be drawn if none of the preceding options are
   available. This is the most weak way to create new connections, because it
   does not consider any additional information.

### lookup TODO
The process of looking up connections is triggered on demand and thus must be
optimized for fast execution.

When information is provided in the form of input, and there does a connection
path for the current input sequence exist, the information ID of the input
sequence is fetched. This information ID links to some meta data associated
with this input sequence, which also contains CLG tree IDs. Using such CLG tree
IDs it is possible to lookup the connection path of behaviors. Within each
CLG's scope a lookup happens to fetch all the peers that needs to be known to
forward signals to.

### data structure
Designing a data structure is quite important. Smart systems need to store
information efficiently. The wrong data structures will cause huge amounts of
data or cause high latency for business logic tasks. The following data
structure design aims to be efficient and fast while meeting the requirements
of Anna's business logic. We use key-value pairs to store data and describe
relations between objects where possible because of simplicity and speed.

The notation of the described data structures reads as follows. Everything
within angle brackets (`<>`) reads as variable. On the left is the key, on the
right is the value of the key-value pairs. `<prefix>` represents some internal
storage prefix simply used to prefix data structures to a certain scope. When
talking about a `behavior-coordinate`, we talk about a single point within the
connection space, which represents one single CLG associated with a very
specific connection path.

---

###### map behavior coordinate to behavior coordinates
When having a single behavior coordinate given it needs to be mapped to
multiple behavior coordinates. This mapping represents connections linking to N
behaviors within the connection space. That way dynamic CLG trees can be
formed. CLGs can lookup connections using their own coordinates. The found
connections are supposed to be used to forward signals to. The following key
maps a single behavior coordinate to multiple behavior coordinates.

```
<prefix>:behavior-coordinate:behavior-coordinates:<behavior-coordinate>    <behavior-coordinate>,<behavior-coordinate>,...
```

###### map behavior coordinate and information sequence to behavior coordinates
When having a single behavior coordinate and an information sequence given,
this combination needs to be mapped to multiple behavior coordinates. This
mapping represents connections linking to N behaviors within the connection
space. That way dynamic CLG trees can be formed. CLGs can lookup connections
using their own coordinates and provided information sequences. The found
connections are supposed to be used to forward signals to. The following key
maps a single behavior coordinate an information sequence to multiple behavior
coordinates.

```
<prefix>:behavior-coordinate:information-sequence:behavior-coordinates:<behavior-coordinate>:<information-sequence>    <behavior-coordinate>,<behavior-coordinate>,...
```

---

###### map behavior coordinate to CLG name
When having a single behavior coordinate given it needs to be mapped to its
unique CLG name. That way behavior can be resolved from its very unique
coordinate to some actual functionality. This works even across reboots,
because CLG IDs change where their names don't. The following key maps a single
behavior coordinate to its CLG name.

```
<prefix>:behavior-coordinate:behavior-name:<behavior-coordinate>    <CLG-name>
```

### abstraction
The described model above can be abstracted and used for other different things
as well. E.g. a mapping of conceptional ideas can be achieved by implementing
certain CLGs. For instance one CLG which adds a mapping between two given
strings, where these strings represent some concepts like `tree` or `house`.
Another CLG reads such mappings and acts according to some pattern it decided
for. Having enough entropy and time given this behavior will lead to
surprisingly results. The important thing for on this level of abstraction is
simplicity. We only use very simple key-pairs to represent connections, which
gives the neural network enough room to figure itself out how to act best.
