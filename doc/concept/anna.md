# anna
This is the ten thousand feet view of Anna. To understand how she looks like
from the very top we consider the following 4 layers.

1. The `i/o` layer describes a set of network protocols Anna understands. Data
   can be written to and retreived from her over network. I/O is flowing to and
   coming from the server.

2. The `server` layer describes the actual server listening for traffic of
   implemented network protocols. It provides so called `interfaces` that are
   used to differentiate between different types of inputs that serve different
   types of purposes. Interfaces dispatch information to and from gateways.

3. The `gateway` layer describes a gateway where data is exchanged. The concept
   of a separate gateway is important architectural wise to fully decouple the
   server and the core.

4. The `core` layer describes the implementation of Anna's most inner workings.
   It bundles everything around data processing and intelligence. The core
   itself contains multiple `networks`. Signals provided by the gateway are are
   translated to impulses that, at first, pass the scheduler network. From
   there impulses find their way through the networks to finally trigger
   actions in form of a response to incoming signals.

This is how it basically looks like. Note that the white pale boxes represent
ideas that are not yet implemented. The strong grey boxes in fact represent
components that are, at least partly, implemented.

![anna](image/anna.png)
