# concept
Here are concepts documented that drive the development of this project.
Concepts are always open to be discussed, wether they are implemented or not.

### tracability
Having insides into complex systems is key. Events going through neural
networks need to be highly comprehensible in detail. Data needs to be
collected. All information we can gather need to be visualized somehow.

### plugability
A plugable architecture having decent interfaces ...

### testability
Functionality needs to be guaranteed by testing actual software
implementations. Automated. Painless. Fast.

### data formats
For simplicity JSON should be good enough for now.

### API's
There are two forms of API's we want to care about here. Library interfaces and
network API's.

Library interfaces should always be well defined. Software packages of this
project need to have a real purpose on their own, so they can stand alone, or
be used by something else.

Network API's should always be well defined. They should simply represent
business logic implementation wrapped by some network protocol middleware.
