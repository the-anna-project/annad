# connection
The connection model is the most important concept of the Anna project. A
connection can be seen as a neural connection. Such connections represent
relationships between certain information and behaviours, represented in a multi
dimensional space. [Inputs](input.md), [outputs](output.md) and [CLGs](clg.md)
are wired together that way. Many of these connections interacting together
dynamically, represent Anna's neural [network](network.md).

### space
The connection space can be seen as a multi dimensional vector space. In theory
it's dimensions can represent everything: space, time or even conceptual weights
representing something like emotions. Information and behaviours are mapped onto
the connection space by having coordinates applied. A coordinate represents a
location within the connection space, which is calculated by interacting with
[the balance system](distribution.md#balance-system). The balance system is used
to make dimensions create tensions on coordinates within the connection space.
Each dimension is pulling into its own direction. The important idea behind the
tensions created here is that they are balanced. One dimension pulling
coordinates spontaneously into a certain direction makes another dimension move
proportionally into another direction. If one dimension goes up, another
dimension goes down. Thus a balance is created. This creates a simulation of
psychological interactions on different levels. Like disgust and attraction
cannot go together. Aim of such a balance system is to provide built in
mechanisms of leveraging event interactions and calculating connection locations
within the connection space.

Spontaneous modificatins of connections lead to variances within applied
behaviours. Leads some behaviour to successful results, slight adjustments of a
CLG tree's coordinates can be stored, which causes the CLG tree to be updated.
This will improve CLG trees over time, because behaviour was adjusted to the
given environmental feedback that Anna has received. That means the balance
system influences connections within the connection space, and thus behaviour of
the whole neural network. Thus coordinates are used to calculate dynamic and
individual connection locations within the connection space. They are also used
to calculate certain similarities between information, and certain similarities
between behaviours.

The following picture illustrates the connection space. For simplicity it only
shows two dimensions, `x` and `y`. We see two different CLG trees within the
connection space. Here some peers are pretty near to each other. We can
calculate a similarity based on the distances between peers. An indicator for
common connection patterns can also be derived from the shape of a connection
path.

![connection space](image/connection_space.png)

### creation
When creating new connections it is important to not create weak connections.
Each connection that exists only exists because it brought some kind of value in
the past. The process of creating connections is a continuous task, that is
fully dynamic and learned by experience.

When information is provided, it is stored within the underlying storage. The
creation of behaviour connections takes place on the CLG level. A connection is
represented as single key-value pair. The key of such a key-value pair simply
consists of a behaviour ID. The value of such a key-value pair is a list of
behaviour IDs. All newly created connections are persisted within a trial scope
in the first place. The purpose of such a trial scope is to label all stored
data within the current creation process as volatile. If the neural network
succeeds to solve a problem with some newly created connections, the connections
stored within a trial scope are persisted as regular connections and thus not
considered volatile anymore. In case the created connections did not lead to
some successful operation, all volatile connections stored within a trial scope
are simply removed.

A connection can also have a TTL. In this case the connection itself is somehow
volatile on purpose, because it is supposed to connect peers only for a certain
amount of time. Why this is done and how long such a connection exists is up to
the neural network itself.

Anyway, there needs a decision to be made to forward signals to some CLGs. The
following strategies are considered when it comes to create new connections
within the connection space.

1. *Bias* is some manually provided hint, intended to guide some CLG tree into a
   certain direction. Read more on this in this issue:
   https://github.com/xh3b4sd/anna/issues/44.

2. *Intuition* is some sort of vague feeling that points into a certain
   direction. Drawing distantly related connections across multiple levels can
   gather information and generate new relations between peers.

3. *Copy* connections from other branches looks up possible connection
   structures from different problem domains. Connections that have been useful
   in one problem domain might be useful as well in another.

4. *Random* connections can be drawn if none of the preceding options are
   available. This is the most weak way to create new connections, because it
   does not consider any additional information.

### lookup
The process of looking up connections is triggered on demand and thus must be
optimized for fast execution.

The lookup of connections happens within each CLG's execution scope to fetch all
the peers that needs to be known to forward signals to. Connections will be
looked up as they are stored. Spontaneous connections are made by CLGs who think
this is sufficient.

### data structure
Designing a data structure is quite important. Smart systems need to store
information efficiently. The wrong data structures will cause huge amounts of
data or cause high latency for business logic tasks. The following data
structure design aims to be efficient and fast while meeting the requirements of
Anna's business logic. Note that the following documentation does not aim to
track a complete list of all data structures used for the Anna project. We only
try to document the most important key-value pairs to transport ideas. It might
happen that some information are missing. In this case the missing information
should be added below. Anyway, we use key-value pairs to store data and describe
relations between objects where possible because of simplicity and speed. The
notation of the described data structures reads as follows.

- Everything within angle brackets (`<>`) reads as variable.
- On the left is the key, on the right is the value of the key-value pairs.
- When talking about a `<prefix>`, we talk about some internal storage prefix,
  which is simply used to prefix data structures to a certain scope. This
  prefix might also indicate a relation to some `<trial-scope>`.
- When talking about a `<clg-tree-id>`, we talk about an identifier for
  combined, executable behaviour.
- When talking about a `<behaviour-coordinate>`, we talk about a single point
  within the connection space, which represents one single CLG associated with a
  very specific CLG tree. A CLG's coordinate represents its location within the
  connection space. This representation of a CLG's location is calculated
  continuously by interacting with the balance system. Thus coordinates might be
  recalculated to align them to successful CLG trees.
- When talking about a `<behaviour-ID>`, we talk about a very unique identifier
  of a CLG, which is generated for each CLG execution that is not related to an
  already known CLG tree ID. That way a single CLG can be represented as unique
  CLG with its own very unique connections and behaviours.
- When talking about an `<information-coordinate>`, we talk about a single point
  within the connection space, which represents one single information sequence.
  An information sequence's coordinate represents its location within the
  connection space. This representation of an information sequence's location is
  calculated continuously by interacting with the balance system. Thus
  coordinates might be recalculated to align them to successful connection
  paths.
- When talking about an `<information-sequence>`, we talk about a single piece
  of information, which is either provided from the outside, or generated
  internally.

---

###### map information sequence to information ID
When having an information sequence given it needs to be mapped to its
information ID. This mapping resolves a single information ID from its
information sequence. The following key maps an information sequence to its
information ID.

```
<prefix>:information-sequence:<information-sequence>:information-id    <information-id>
```

---

###### map information ID to information sequence
When having an information ID given it needs to be mapped to its information
sequence. This mapping resolves a single information sequence from its
information ID. The following key maps an information ID to its information
sequence.

```
<prefix>:information-id:<information-id>:information-sequence    <information-sequence>
```

---

###### map information ID to information coordinate
When having an information ID given it needs to be mapped to its information
coordinate. This mapping resolves a single information coordinate from its
information ID. A coordinate is used to calculate its location within the
connection space. The following key maps an information ID to its information
coordinate.

```
<prefix>:information-id:<information-id>:information-coordinate    <information-coordinate>
```

---

###### map information ID to CLG tree IDs
When having an information ID given it needs to be mapped to CLG tree IDs. This
mapping resolves CLG tree IDs from an information ID. The following key maps an
information ID to its information coordinate.

```
<prefix>:information-id:<information-id>:clg-tree-ids    <clg-tree-id>,<clg-tree-id>,...
```

---

###### map information coordinate to information ID
When having an information coordinate given it needs to be mapped to its
information ID. This mapping resolves a single information ID from its
information coordinate. A coordinate is used to calculate its location within
the connection space. The keyscan is very slow and only supposed to be done in
the background as persistent neural activity. The following key maps an
information coordinate to its information ID.

```
<prefix>:information-coordinate:<information-coordinate>:information-id    <information-id>
```

---

###### map behaviour coordinate to behaviour ID
When having a behaviour coordinate given it needs to be mapped to its behaviour
ID. This mapping resolves a single behaviour ID from its own coordinate within
the connection space. Having the behaviour coordinate indexed as key enables
lookups based on similarities when scanning the key space within the underlying
storage. A coordinate is used to calculate its location within the connection
space. The keyscan is very slow and only supposed to be done in the background
as persistent neural activity. The following key maps a behaviour coordinate to
its behaviour ID.

```
<prefix>:behaviour-coordinate:<behaviour-coordinate>:behaviour-id    <behaviour-id>
```

---

###### map behaviour ID to behaviour coordinate
When having a behaviour ID given it needs to be mapped to its behaviour
coordinate. This mapping resolves a single behaviour coordinate from its own ID
within the connection space. A coordinate is used to calculate its location
within the connection space. The following key maps a behaviour ID to its
behaviour coordinate.

```
<prefix>:behaviour-id:<behaviour-id>:behaviour-coordinate    <behaviour-coordinate>
```

---

###### map CLG tree ID to behaviour ID
When having a CLG tree ID given it needs to be mapped to the very first CLG
within a specific CLG tree. This very first CLG is represented by a behaviour ID.
This mapping resolves all behaviours of a whole CLG tree. The following key maps
a CLG tree ID to the very first behaviour within this specific CLG tree.

```
<prefix>:clg-tree-id:<clg-tree-id>:first-behaviour-id    <behaviour-id>
```

---

###### map behaviour ID to behaviour IDs set
When having a single behaviour ID given it needs to be mapped to multiple
behaviour IDs. This mapping represents connections linking from one behaviour to
other behaviours within the connection space. That way dynamic CLG trees can be
referenced. CLGs can lookup connections using their own behaviour IDs to know
where to forward signals to. In case an exact expectation is known, the behaviour
should change slightly over time to find more optimized ways of solving
problems. In case a problem was able to be solved using some modifications, the
behaviour can be updated completely. Note that the order of the stored behaviour
IDs is not important, because we only need to make sure that signals are
dispatched to each connected CLG. The following key maps a single behaviour ID to
multiple behaviour IDs.

```
<prefix>:behaviour-id:<behaviour-id>:forward-behaviour-ids    <behaviour-id>,<behaviour-id>,...
```

---

###### map behaviour ID to behaviour IDs list
When having a single behaviour ID given it needs to be mapped to multiple
behaviour IDs. This mapping represents connections linking from one behaviour to
other behaviours within the connection space. That way dynamic CLG trees can be
referenced. CLGs can lookup connections using their own behaviour IDs to know
which payloads to accept when being activated. In case an exact expectation is
known, the behaviour should change slightly over time to find more optimized ways
of solving problems. In case a problem was able to be solved using some
modifications, the behaviour can be updated completely. Note that the order of
the stored behaviour IDs is important, because it reflects the input interface of
the requested CLG. The following key maps a single behaviour ID to multiple
behaviour IDs.

```
<prefix>:behaviour-id:<behaviour-id>:activate-behaviour-ids    <behaviour-id>,<behaviour-id>,...
```

---

###### map behaviour ID to CLG name
When having a single behaviour ID given it needs to be mapped to its unique CLG
name. That way behaviour can be resolved from its very unique ID to some actual
functionality. This works even across reboots, because CLG IDs change where
their names don't. The following key maps a single behaviour ID to its unique CLG
name.

```
<prefix>:behaviour-id:<behaviour-id>:behaviour-name    <CLG-name>
```

---

###### map behaviour name to weighted behaviour names
When having a single behaviour name given it needs to be mapped to multiple
weighted bahavior names. That way patterns of behaviour connections can be
tracked. The following key maps a single behaviour name to multiple behaviour
weighted names.

```
<prefix>:behaviour-name:<behaviour-name>:weighted-behaviour-names    <behaviour-name>:<weight>,<behaviour-name>:<weight>,...
```

### abstraction
The described model above can be abstracted and used for other different things
as well. E.g. a mapping of conceptional ideas can be achieved by implementing
certain CLGs. For instance one CLG which adds a mapping between two given
strings, where these strings represent some concepts like `tree` or `house`.
Another CLG reads such mappings and acts according to some pattern it decided
for. Having enough entropy and time given this behaviour will lead to
surprisingly results. The important thing for on this level of abstraction is
simplicity. We only use very simple key-pairs to represent connections, which
gives the neural network enough room to figure itself out how to act best.
