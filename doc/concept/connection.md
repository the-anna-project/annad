# connection
The connection model is the most important concept of the Anna project. A
connection can be seen as a neural connection. Such connections represent
relationships between certain information and behaviors, represented in a
multi dimensional space. [Inputs](input.md), [outputs](output.md) and
[CLGs](clg.md) are wired together that way. Many of these connections joined
together form Anna's neural [network](network.md).

### space
The connection space can be seen as a multi dimensional vector space. In theory
it's dimensions can represent everything: space, time or even conceptional
weights representing something like emotions. Information and behaviors are
mapped onto the multi dimensional connection space. Here similarities between
information, and similarities between behaviors can be calculated, because they
are represented by connection paths. These connection paths consist of
coordinates, which identify the connection path's location within the
connection space. Coordinates can be used to calculate distances to other
coordinates in the surrounding area. Close connection paths can indicate some
kind of similarity.

In the connection space information are represented by the coordinates of input
sequences. Each vector, that is represented by coordinates, reflects a specific
character of a specific input sequence. You can think of the joined coordinates
of an input sequence as a drawn information path in the multi dimensional
space. Information paths in the surrounding area are evidence of some kind of
similarity input wise. Over time the information paths are transformed when
each dimension is pulling it's coordinates into their own direction. That way a
balanced alignment is achieved that makes each information path unique and
comparable against other information paths in the surrounding area. The same
concept that applies to information paths also applies to behavior paths, but
only on a different problem domain and in more complex structure. This is
because information paths are represented by linear input sequences, and
behavior paths are represented by CLG trees. Here behavior is mapped to an
executable CLG tree, which coordinates are mapped onto the multi dimensional
connection space. The relationship between an input sequence and a CLG tree
that solved a problem for this input sequence creates the link between
information and behavior.

The following picture illustrates the multi dimensional connection space. For
simplicity it only shows two dimensions, `x` and `y`. We see two different
paths within the coordinate system like space. It doesn't matter if the paths
are information or behavior paths. The same principle applies to both of these
concepts. We see some peers which are pretty near to each other. We assume that
a smaller distance is an indicator for common connection patterns which are
aligned over time.

![connection space](image/connection_space.png)

### creation
When creating new connections it is important not to create weak connections.
Each connection that exists only exists because it brought some kind of value
in the past. The process of creating connections is a continuous task, that is
fully dynamic and learned by experience and can be described as follows.

When information is provided in the form of input, and no connection path for
the current input sequence exists yet, it is mapped onto a multi dimensional
space. Each character of the input sequence is represented by their own
coordinates. That way an input sequence draws a connection path within the
multi dimensional connection space. This is then the connection path of
information.

The creation of behavior connections takes place on the CLG level. Connections
of behaviors are only persisted in case they belong to a successful behavior
connection path that in turn helped to solve some problem. So, the connection
path of behaviors is made out of connections that belong to behaviors. Here a
connection is represented as single key-value pair where it's key being the
string representation of coordinates links to another string representation of
coordinates. A CLG is identified using its unique coordinates within the multi
dimensional connection space in combination with its very unique connection
path, which in turn is a represented by a CLG tree ID. A reference from
coordinates to an actual CLG ID is also stored for mapping of coordinates to
actual behavior later on. This is then the connection path of behaviors.

At some point, each single CLG needs to decide where to forward its own signal
to. Once forwarded, the coordinates of the CLG receiving the forwarded impulse
are randomly created with some offset. The offset of the CLG's coordinates is
orientated to the coordinates of the CLG actually being forwarding the signal.
All newly created connections are persisted within a trial scope in the first
place. Such trial scope is basically realized by some specific storage prefix
that simply identifies certain key-value pairs as being part of such a trial
scope. This is done to label the current creation process of behavior
connections to something that is volatile. If the neural network succeeds to
solve a problem with some newly created connection path, the behavior
connections stored and marked within a trial scope are persisted as regular
connections. In case the created connection path did not lead to some
successful operation, all behaior connections marked within a trial scope are
simply removed again. Anyway, there needs a decision to be made to forward to
some CLG. These strategies are considered when it comes to draw new connections
within a multi dimensional connection space.

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

Following is a picture describing the process of connection creation within the
neural network. `[0]` marks the point at which it is tried to map the given
input sequence to a CLG tree ID in the available storage. This information is
attached to the current impulse. `[1]` marks the points at which it is tried to
lookup connections of the current CLG tree. This is only possible in case there
is a CLG tree ID known. If there is no CLG tree ID available, connections are
created as described above.

![connection creation](image/connection_creation.png)

### lookup
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
storage prefix simply used to prefix data structures to a certain scope.

---

###### map input sequence to information ID
When having an input sequence given it needs to be mapped to an information ID.
An input sequence represents some externally provided information. The
following key maps an input sequence to an information ID.

```
<prefix>:input-sequence:information-id:<input-sequence>    <information-id>
```

---

###### map information ID to input object
When having an information ID given it needs to be mapped to an input object.
That way accociated meta data of an input sequences can be tracked by one
reference object. The following key maps an information ID to an input object.

```
<prefix>:information-id:input-object:<information-id>    {input-sequence: <input-sequence>, information-coordinates: <information-coordinates>}
```

---

###### map information ID to CLG tree IDs
When having an information ID given it needs to be mapped to CLG tree IDs.
Here we make use of sets being able to create intersections with other sets, or
being able to weight the set members for other sophisticated lookups. Note that
this key-value pair represents a link between information and behavior due to
the reference from an information ID to CLG tree IDs. The following key maps an
information ID to CLG tree IDs.

```
<prefix>:information-id:clg-tree-id:<information-id>    <clg-tree-id>,<clg-tree-id>,...
```

---

###### map information coordinates to information ID
When having information coordinates given they need to be mapped to their
information ID. Having information coordinates indexed as keys enables fast
scans when it needs to be found out which information are near to the
surrounding area of a given information within the multi dimensional connection
space. That way information can be mapped and aligned to other matching
information. The following key maps information coordinates to it's
information ID.

```
<prefix>:information-coordinates:information-id:<information-coordinates>    <information-id>
```

---

###### map behavior coordinates to behavior coordinates
When having behavior coordinates given they need to be mapped to behavior
coordinates. This mapping represents a single connection of behaviors within
the multi dimensional connection space. CLGs can lookup connections supposed to
be used to forward signals to their peers using their own coordinates. The
following key maps behavior coordinates to behavior coordinates.

```
<prefix>:behavior-coordinates:behavior-coordinates:<behavior-coordinates>    <behavior-coordinates>,<behavior-coordinates>,...
```

---

###### map behavior coordinates to CLG ID
When having behavior coordinates given they need to be mapped to their CLG ID.
Having behavior coordinates indexed as keys enables fast scans when it needs to
be found out which behaviors are near to the surrounding area of a given
behavior within the multi dimensional connection space. That way behavior can
be mapped and aligned to matching behaviors. Further it needs to be known where
to forward impulses to when walking through connection paths. The following key
maps behavior coordinates to it's CLG ID.

```
<prefix>:behavior-coordinates:behavior-id:<behavior-coordinates>    <CLG-id>
```
