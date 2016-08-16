// Package clg provides CLGs that implement basic behaviour leveraged in the
// neural network.
//
// Note that this package defines a go generate statement to generate fully
// functional source code for all CLGs.
//
//go:generate ${GOPATH}/bin/clggen generate --clg-dir=. --template-dir=../template
//
package clg
