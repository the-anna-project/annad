# connection
A connection is the most important concept of the Anna project. It can be seen
as a neural connection. Many connections joined to [strategies](strategy.md)
form Anna's neural network.

### data structure
Designing a data structure is quiet important. Smart systems need to store
information efficiently. The wrong data structures will cause even more huge
amounts of data or causing high latency for business logic tasks. The following
data structure design aims to be efficient and fast while meeting the
requirements of Anna's business logic. We use key-value pairs to store data and
describe relations between objects because of simplicity and speed.

The notation of the described data structures reads as follows. On the left is
the key, on the right is the value described. `<prefix>` represents some
internal storage prefix.

###### strategy
Storing the raw strategy objects is done using the following key-value pair.
`<strategy-id>` represents the ID of a [strategy](strategy.md). It holds the
`<strategy-object>`, that represents the storable data of [the strategy
object](https://godoc.org/github.com/xh3b4sd/anna/spec#Strategy).

```
<prefix>:<strategy-id>   <strategy-object>
```

###### stage
Storing the [stage](stage.md) related data structures is done using the
following key-value pairs. Note the values described here are weighted lists.
`<stage>` represents the incrementable number of a stage in which a
[strategy](strategy.md) is created and executed. `<strategy-id>` represents
the ID of a [strategy](strategy.md). `<weight>` represents the numerical
strength of the connection between a stage and a strategy.

```
<prefix>:<stage>                   <strategy-id>:<weight>, <strategy-id>:<weight>, ...
<prefix>:<strategy-input-type>     <strategy-id>:<weight>, <strategy-id>:<weight>, ...
<prefix>:<strategy-input-value>    <strategy-id>:<weight>, <strategy-id>:<weight>, ...
```

###### path
```
<prefix>:<strategy-id>             <strategy-id>:<weight>, <strategy-id>:<weight>, ...
```



![anna](image/anna.png)
