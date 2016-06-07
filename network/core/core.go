// Package core implements spec.Network to provide CLG execution. Gateways send
// signals to the core network to request calculations. The core network
// translates a signal into an impulse. So the core network is the starting
// point for all impulses. Once an impulse finished its walk through the core
// network, the impulse's output is translated back to the requesting signal
// and the signal is send back through the gateway to its requestor.
package core
