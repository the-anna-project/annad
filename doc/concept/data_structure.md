# data structure
Designing a data structure is quite important. Smart systems need to store
information efficiently. The wrong data structures will cause even more huge
amounts of data or cause high latency for business logic tasks. The following
data structure design aims to be efficient and fast while meeting the
requirements of Anna's business logic. We use key-value pairs to store data and
describe relations between objects because of simplicity and speed.

The notation of the described data structures reads as follows. On the left is
the key, on the right is the value described. `<prefix>` represents some
internal storage prefix.

The following values represent lists of weighted data, in which the numerical
strength of the connection between a key and its value is described. Purpose of
this key-value pairs is to store the importance and the relation between
behaviour and information. These relationships are required to draw context
related paths through the neural network.

```
<prefix>:clg:clg:<clg-id>       <clg-id>:<weight>, <clg-id>:<weight>, ...
<prefix>:clg:info:<clg-id>      <info-id>:<weight>, <info-id>:<weight>, ...
<prefix>:info:clg:<info-id>     <clg-id>:<weight>, <clg-id>:<weight>, ...
<prefix>:info:info:<info-id>    <info-id>:<weight>, <info-id>:<weight>, ...
```
