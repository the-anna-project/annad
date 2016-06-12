# instrumentation
Instrumentation is important to get insights into the runtime internals. In the
Anna project we oriented on the [prometheus project](https://prometheus.io).
Within the project's code we make use of the [prometheus client
library](https://godoc.org/github.com/prometheus/client_golang/prometheus).
This client is abstracted by the [instrumentation
package](https://godoc.org/github.com/xh3b4sd/anna/instrumentation), which can
be used to emit metrics to your likes. Note there is a memory implementation
that does basically nothing. To actually emit metrics there need to be the
configured prometheus implementation used. For graph visualization
[grafana](http://grafana.org) can be used.
