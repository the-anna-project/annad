# guidelines

### tracability
Having insides into complex systems is key. Events going through neural
networks need to be highly comprehensible in detail. A lot of different things
will go on during operational daily business. We need to be able to separate
single actions to make them visible and to trace them down with all their
implications. Data needs to be collected. All information we can gather need to
be visualized somehow.

### plugability
A plugable architecture having decent interfaces provides a strong faundation
for complex projects like this here. Think about software like having lego
blocks you can compose, plug together and replace.

### testability
Functionality needs to be guaranteed by testing actual software
implementations. Automated. Painless. Fast. This means unit tests in the first
place. It is hard to come up with decent integration tests. Here we need to
find ways to create integration tests that do not hurt, that are beautiful and
in best case even fast. The last point is in most cases contradictionary.

### data formats
This depends on the use case. E.g. for simplicity of some network API, JSON
should be good enough for now. Config files may want to make use of some
ini-style format.

### API's
There are two forms of API's we want to care about here. Library interfaces and
network API's.

Library interfaces should always be well defined. Software packages of this
project need to have a real purpose on their own, so they can stand alone, or
be used by something else. Interfaces are gold for abstraction and testability
reasons.

Network API's should always be well defined. They should simply represent
business logic implementation wrapped by some network protocol middleware.
