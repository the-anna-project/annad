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
mapped onto the connection space. Here similarities between information, and
similarities between behaviors can be calculated. In the connection space
information are represented by the coordinates of input trees. Each vector,
that is a list of coordinates, reflects a specific reusable feature of an input
sequence. The distance calculation is done with respect to the input tree's
coordinates. You can think of the joined coordinates of the input tree as a
drawn information path in multi dimensional space. Information paths in the
surrounding area are evidence of some kind of similarity. Over time the
information paths are formed while each dimension is pulling on the peers of
the information path into their own direction. That way a balanced alignment is
achieved that makes each information path unique. The same concept that applies
to information paths also applies to behavior paths, but only on a different
problem domain. Here behavior is mapped to an executable CLG tree, which
coordinates are mapped onto the multi dimensional connection space.

The following picture illustrates the multi dimensional connection space. For
simplicity it only shows three dimensions. Here we see two different paths. It
doesn't matter if they are information or behavior paths. The same principle
applies to both of them. We see some peers are pretty near to each other. A
smaller distance is an indicator for common connection patterns which are
aligned over time.

![connection](image/connection.png)

### creation
When creating new connections it is important not to create weak connections.
Each connection that exists only exists because it brought some kind of value
in the past. The process of creating connections is a continuous task, that is
fully dynamic and learned by experience and can be described as follows. The
creation of connections takes place on the CLG level if there is no known
connection path yet. That means that connections are only persisted in case
they belong to a successful connection path that in turn helped to solve some
problem. So, the connection path is made out of connections. A connection is
represented as single key-value pair where its key being the string
representation of coordinates links to another string representation of
coordinates. Therefore coordinates act as some kind of IDs. A CLG is identified
using its unique coordinates within the multi dimensional connection space in
combination with its very unique connection path. At some point, each single
CLG needs to decide where to forward its own signal to. Once forwarded, the
coordinates of the CLG receiving the forwarded impulse are randomly created
with some offset. The offset of the CLG's coordinates is orientated to the
coordinates of the CLG actually being forwarding the impulse. All newly created
connections are persisted within a trial scope in the first place. This is done
to label the current creation process to something that is volatile. If the
neural network succeeds to solve a problem with some newly created connection
path, the connections stored and marked within a trial scope are persisted as
regular connections. In case the created connection path did not lead to some
successful operation, all connections marked within a trial scope are simply
removed again. Anyway, there needs a decision to be made to forward to some
CLG. These strategies are considered when it comes to draw new connections
within a multi dimensional connection space.

1. Bias is some manually provided hint, intended to guide some connection path
   into a certain direction. Read more on this in this issue:
   https://github.com/xh3b4sd/anna/issues/44.

2. Intuition is some sort of vague feeling that points into a certain
   direction. Drawing distantly related connections across multiple levels can
   gather information and generate new relations between peers.

3. Copy connections from other branches looks up possible connection structures
   from different problem domains. Connections that have been useful in one
   problem domain might be useful as well in another.

4. Random connections can be drawn if none of the preceding options are
   available. This is the most weak way to create new connections, because it
   does not consider any additional information.

### lookup

TODO

Lookup happens within each CLG's scope to fetch all the peers supposed to be forwarded to.

The process of looking up connections is triggered on demand and thus must be
optimized for fast execution. When information is provided in the form of
[input](input.md), it is mapped onto a multi dimensional space to enable the
lookup of some behavior. The given input consists of many different reusable
features. These features are resolved to their information IDs, which translate
to many of so called input tree IDs. Here we can lookup information IDs of all
features and create an intersection of the created set. The resulting input
tree ID is then identified to represent the full original input sequence.
However this input tree ID maps to an input tree, which itself maps to
information IDs. Their coordinates identify the input tree's location within
the connection space. The input tree's coordinates can be used to calculate
distances to other input trees in the surrounding area. Close connection paths
can indicate helpful information or some kind of similarity. Due to the already
created connection we have a mapping between the input tree and the
[CLG](clg.md) tree, which we actually want to lookup. This relationship creates
the link between information and behavior, because the connected CLG tree maps
directly to executable behavior.

### data structure
Designing a data structure is quite important. Smart systems need to store
information efficiently. The wrong data structures will cause even more huge
amounts of data or cause high latency for business logic tasks. The following
data structure design aims to be efficient and fast while meeting the
requirements of Anna's business logic. We use key-value pairs to store data and
describe relations between objects because of simplicity and speed.

The notation of the described data structures reads as follows. Everything
within angle brackets (`<>`) reads as variable. On the left is the key, on the
right is the value of the key-value pairs. `<prefix>` represents some internal
storage prefix simply used to prefix data structures to a certain scope.

---

###### map input sequence to information ID
When having an input sequence given it needs to be registered. An input
sequence represents some externally provided information as it is in the first
place. Over time input sequences are broken down into their separate features
based on their usage. The following key maps an input sequence to an
information ID.

```
<prefix>:input-sequence:information-id:<input-sequence>    <information-id>
```

---

###### map information ID to input object
When having an information ID given it needs to be mapped to an input object.
That way all accociated meta data of an input sequences can be tracked by one
reference object. The following key maps an information ID to an input object.

```
<prefix>:information-id:input-object:<information-id>    {input-sequence: <input-sequence>, information-coordinates: <information-coordinates>, clg-tree-ids: [<clg-tree-id>,<clg-tree-id>,...]}
```

---

###### map information coordinates to information ID
When having information coordinates given they need to be mapped to their
information ID. Having information coordinates indexed as keys enables fast
scans when it needs to be found out which informations are near to the
surrounding area of a given information within the connection space. That way
information can be mapped and aligned to matching information. The following
key maps information coordinates to it's information ID.

```
<prefix>:information-coordinates:information-id:[<x>,<y>,...],[<x>,<y>,...],...    <information-id>
```

---

TODO
- How to scope persisted single key-value pair connections? Which role does the CLG tree ID have?
- Is there a way to make connections "longer" so we store coordinates of more peers than two? This would make 10000 feet view lookups faster.

###### map CLG tree ID to CLG tree
When having a CLG tree ID given it needs to be mapped to a CLG tree. A CLG tree
represents an organizational structure that holds ordered behavior IDs forming
an executable behavior network. A behavior ID is effectively a CLG ID. Note
that a CLG tree can only be valid in case it starts with the ID of the Input
CLG, and ends with the ID of the Output CLG in any branch. The following key
maps an CLG tree ID to a CLG tree.

```
<prefix>:clg-tree-id:clg-tree:<clg-tree-id>    {<behavior-id>: {<behavior-id>: {...}, <behavior-id>: {...}, ...}}
```

---

###### map behavior ID to behavior object
When having an behavior ID given it needs to be mapped to an behavior object.
That way all accociated meta data of an behavior can be tracked by one
reference object. The following key maps an behavior ID to an behavior object.

```
<prefix>:behavior-id:behavior-coordinates:<behavior-id>    {behavior-coordinates: <behavior-coordinates>, clg-tree-id: <clg-tree-id>}
```

---

###### map behavior coordinates to behavior ID
When having behavior coordinates given they need to be mapped to their behavior
ID. Having behavior coordinates indexed as keys enables fast scans when it
needs to be found out which behaviors are near to the surrounding area of a
given behavior within the connection space. That way behavior can be mapped and
aligned to matching behaviors. The following key maps behavior coordinates to
it's behavior ID.

```
<prefix>:behavior-coordinates:behavior-id:[<x>,<y>,...],[<x>,<y>,...],...    <behavior-id>
```
