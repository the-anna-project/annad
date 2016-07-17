# connection
The connection model is the most important concept of the Anna project. A
connection can be seen as a neural connection. Such connections represent
relationships between certain information and bevahiors, represented in a
multi dimensional space. [Inputs](input.md), [outputs](output.md) and
[CLGs](clg.md) are wired together that way. Many of these connections joined
together form Anna's neural [network](network.md).

### creation

When creating new connections it is important not to create weak connections.
Each connection that exists only exists because it brought some kind of value
in the past. The process of creating connections is fully dynamic and learned
by experience and can be described as follows. These strategies are considered
when it comes to draw new connections within a multi dimensional space.

1. Bias is some manually provided hint, intended to guide some connection path
   into a certain direction. Read more on on this in this issue:
   https://github.com/xh3b4sd/anna/issues/44.

2. Intuition is some sort of vague feeling that points into a certain
   direction. Drawing distantly related connections across multiple levels can
   gather information and generate new relations between peers.

3. Copy connections from other branches looks up possible connection structures
   from different problem domains. Connections that have been useful in one
   problem domain might be useful as well in another.

4. Random connections can be drawn if none of the preceding options are
   available.

### lookup

The process of looking up relevant information and bevahiors looks as follows.
When [input](input.md) is provided, it is mapped onto a multi dimensional
space. The given input draws an information path. This information path is used
to lookup [CLG](clg.md) paths in the sorounding area, which represent behavior
paths. Over time the connection paths are formed while each dimension is
pulling on connection peers into their own direction. That way a balanced
alignment is achieved that makes each connection unique in terms of information
and behavior. The following picture illustrates the multi dimensional
connection space. For simplicity it only shows three dimensions. In theory this
dimensions can represent anything: space, time or conceptional weights
representing even something like emotions. Here we see two different connection
paths. Some peers are pretty near to each other. This small distance is an
indicator for common connection patterns that are aligned over time in case
such common connection cause challenges to be accomplished.

![connection](image/connection.png)

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
<prefix>:information:sequence:<input-sequence>    <information-id>
```

---

###### map information ID to input sequence
When having an information ID given it needs to be mapped to an input sequence.
That way a mapping to the original information can be achieved. The following
key maps an information ID to an input sequence.

```
<prefix>:information-id:input-sequence:<information-id>   <input-sequence>
```

---

###### map information ID to input tree ID
When having an information ID given it needs to be mapped to an input tree ID.
That way a mapping to an optimization structure can be achieved. The following
key maps an information ID to an input tree ID.

```
<prefix>:information-id:input-tree-id:<information-id>   <input-tree-id>
```

---

###### map input tree ID to input tree
When having an input tree ID given it needs to be mapped to an input tree. An
input tree represents an organizational structure that holds ordered IDs of
input sequences broken down into their features. The following key maps an
input tree ID to an input tree.

```
<prefix>:input-tree-id:input-tree:<input-tree-id>    [[<sequence-id>,<sequence-id>,...],[<sequence-id>,<sequence-id>,...],...]
```

---

###### map input tree ID to input tree coordinates
When having an input tree ID given it's position within the connection space
needs to be looked up. Such lookups are necessary when conceptionaly related
connections between input trees are required during operations on information
level. The following key maps an input tree ID to input tree coordinates within
the connection space.

```
<prefix>:input-tree-id:input-tree-coordinates:<input-tree-id>    [<x>,<y>,...],[<x>,<y>,...],...
```

---

###### map input tree coordinates to input tree ID
When having input tree coordinates given they need to be mapped to their input
tree ID. Having input tree coordinates indexed as keys enables fast scans when
it needs to be found out which input trees are near to the surrounding area of a
given input tree within the connection space. That way information can be
mapped and aligned to matching input trees. The following key maps input tree
coordinates to it's input tree ID.

```
<prefix>:input-tree-coordinates:input-tree-id:[<x>,<y>,...],[<x>,<y>,...],...    <input-tree-id>
```

---

###### map input tree ID to CLG tree ID
When having an input tree ID given it needs to be mapped to a CLG tree ID. This
is the key that maps information to behavior. The following key maps input tree
coordinates to input tree it's ID.

```
<prefix>:input-tree-id:clg-tree-id:<input-tree-id>    <clg-tree-id>
```

---

###### map CLG tree ID to CLG tree
When having a CLG tree ID given it needs to be mapped to a CLG tree. A CLG tree
represents an organizational structure that holds ordered IDs of CLGs forming
an executable behavior network. Note that a CLG tree can only be valid in case
it starts with the ID of the Input CLG, and ends with the ID of the Output CLG
in any branch. The following key maps an CLG tree ID to a CLG tree.

```
<prefix>:clg-tree-id:clg-tree:<clg-tree-id>    {<clg-id>: {<clg-id>: {...}, <clg-id>: {...}, ...}}
```

---

###### map CLG tree ID to CLG tree coordinates
When having a CLG tree ID given it's position within the connection space needs
to be looked up. The following key maps a CLG tree ID to CLG tree coordinates
within the connection space.

```
<prefix>:clg-tree-id:clg-tree-coordinates:<clg-tree-id>    [<x>,<y>,...],[<x>,<y>,...],...
```

---

###### map CLG tree coordinates to CLG tree ID
When having CLG tree coordinates given they need to be mapped to their CLG tree
ID. Having CLG tree coordinates indexed as keys enables fast scans when it
needs to be found out which CLG trees are near to the surrounding area of a
given CLG tree within the connection space. That way behavior can be mapped and
aligned to matching CLG trees. The following key maps CLG tree coordinates to
it's CLG tree ID.

```
<prefix>:clg-tree-coordinates:clg-tree-id:[<x>,<y>,...],[<x>,<y>,...],...    <clg-tree-id>
```
