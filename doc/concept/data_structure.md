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

### stage
Storing the [stage](stage.md) related data structures is done using the
following key-value pairs. Note the values described here are weighted lists.

The following keys represent
- the incrementable number of a stage in which a strategy is created and executed
- the data type of some input
- the data value of some input

The following values represent lists of weighted strategy IDs, in which the
numerical strength of the connection between a the key and a strategy is
described. Purpose of this key-value pairs is to store the importance and the
relation between a stages, input types, input values and strategy. This
information is required for strategy creation.

```
<prefix>:<stage>                   <strategy-id>:<weight>, <strategy-id>:<weight>, ...
<prefix>:<strategy-input-type>     <strategy-id>:<weight>, <strategy-id>:<weight>, ...
<prefix>:<strategy-input-value>    <strategy-id>:<weight>, <strategy-id>:<weight>, ...
```

### strategy
Storing the raw strategy objects is done using the following key-value pair,
where the key represents the ID of a [strategy](strategy.md), and the value
holds the `<strategy-object>`, that represents the storable data of [the
strategy object](https://godoc.org/github.com/xh3b4sd/anna/spec#Strategy).

```
<prefix>:<strategy-id>    <strategy-object>
```

Storing the strategy path related data structures is done using the following
key-value pair. Note the value described here is a weighted list. The following
key represents the ID of some strategy. The following value represents a list
of weighted strategy IDs, in which the numerical strength of the connection
between a the key and a strategy is described. Purpose of this key-value pair
is to store the importance and the relation between strategies. This
information is required for strategy creation.

```
<prefix>:<strategy-id>    <strategy-id>:<weight>, <strategy-id>:<weight>, ...
```
