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

The notation of the described data structures reads as follows. On the left is
the key, on the right is the value described. `<prefix>` represents some
internal storage prefix.

---

###### map information to ID
When having input given it needs to be registered, to be remembered. The
following key maps a character sequence to it's space representation. Unknown
character sequences are created and placed within the connection space using
randomly chosen coordinates.

```
<prefix>:information:sequence:<character><character>...    <information-id>
```

---

###### map information ID to coordinates
When having an information ID given it's position within the connection space
needs to be looked up. The following key maps an information ID to coordinates
within the connection space.

```
<prefix>:information-id:coordinates:<information-id>    [<x>,<y>,...],[<x>,<y>,...],...
```

---

###### map information coordinates to behavior IDs
When having information coordinates given they need to be mapped to behaviors.
The following key maps information coordinates to behavior IDs. Having
coordinates indexed as keys enables fast scans when it needs to be found out
which information are near to the sorounding area of a given behavior within
the connection space. That way information can be mapped and aligned to
matching behaviors simply by scanning for connections within a certain padding.

```
<prefix>:information-coordinates:behavior-ids:[<x>,<y>,...],[<x>,<y>,...],...    <behavior-id>,<behavior-id>,...
```

---

###### map behavior ID to coordinates
When having a behavior ID given it's position within the connection space needs
to be looked up. The following key maps a behavior ID to coordinates within the
connection space.

```
<prefix>:behavior-id:coordinates:<behavior-id>    [<x>,<y>,...],[<x>,<y>,...],...
```

---

###### map behavior coordinates to information IDs
When having behavior coordinates given they need to be mapped to information.
The following key maps behavior coordinates to information IDs. Having
coordinates indexed as keys enables fast scans when it needs to be found out
which behaviors are near to the sorounding area of given information within the
connection space. That way behaviors can be mapped and aligned to matching
information simply by scanning for connections within a certain padding.

```
<prefix>:behavior-coordinates:information-ids:[<x>,<y>,...],[<x>,<y>,...],...    <information-id>,<information-id>,...
```
