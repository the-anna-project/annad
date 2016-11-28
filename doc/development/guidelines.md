# guidelines

### code
This section documents some code guidelines that should be considered when
working in the codebase.

###### CRUD operations
When a service implements [CRUD
operations](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete) the
following naming should be considered.
```
create
search
update
delete
```

### trace
Having insides into complex systems is key. Events going through neural
networks need to be highly comprehensible in detail. A lot of different things
will go on during operational daily business. We need to be able to separate
single actions to make them visible and to trace them down with all their
implications. Data needs to be collected. All information we can gather need to
be visualized somehow. [Logging](/doc/concept/logging.md) and
[instrumentation](/doc/concept/instrumentation.md) are major concepts to help
in this direction.

### plugin
A plugable architecture having decent interfaces provides a strong faundation
for complex projects like this here. Think about software like having lego
blocks you can compose, plug together and replace. For instance the concept of
[CLGs](/doc/concept/clg.md) is fully extensible by simply adding functions
implementing behaviour. The only thing to do there is to add some method and
let Anna explore how to use it.

### test
Functionality needs to be guaranteed by testing actual software
implementations. Automated. Painless. Fast. This means unit tests in the first
place. It is hard to come up with decent integration tests. Here we need to
find ways to create integration tests that do not hurt, that are beautiful and
in best case even fast. The last point is in most cases contradictionary.
Anyway [make gotest](makefile.md#gotest) is your friend. Further code coverage
is tracked at https://codecov.io/gh/the-anna-project/annad.

### data format
Data formats depend on the use case. E.g. for simplicity of some network API,
JSON should be good enough in the first place. Where required other data
formats like protocol-buffers for gRPC API's can also be used. Config files may
or may not want to make use of a JSON encoded format.

### interface
There are two types of API's we want to care about here. Library interfaces and
network API's.

Library interfaces should always be well defined. Software packages of this
project need to have a real purpose on their own, so they can stand alone, or
be used by something else. Interfaces are gold for abstraction and testability
reasons.

Network API's should always be well defined. They should simply represent
business logic implementation wrapped by some network protocol middleware.

# state
AI lives from its intelligence. This can only be achieved by information that
are more or less structured. Business logic and state must be fully decoupled.
That way Anna is able to completely backup and restore her whole state. Her
inner workings aim to be that flexible and dynamic that intelligence arises
from the combination of sufficient [connections](/doc/concept/connection.md)
between [inputs](/doc/concept/input.md), [outputs](/doc/concept/output.md) and
[connections](/doc/concept/connection.md) provided by the sorounding
environment. The concept of having such a state model is important for a lot of
reasons. State can be backed up and restored. Once Anna is shut down, she can
backup her state. Once she is coming back to life, she can restore her state
like nothing happened. Imagine to test a specific part or structure of some
self awareness. Think about analysing, debugging and manipulating Anna's mind.
Future dreams go along the matrix way.  Upload the state of a doctor's mind.
Inject the knowledge of a helicopter pilot. You name it.
