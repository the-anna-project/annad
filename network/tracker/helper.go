package tracker

import (
	"github.com/xh3b4sd/anna/spec"
)

func networkPayloadToConnectionPaths(networkPayload spec.NetworkPayload) ([]string, error) {
	var connectionPaths []string

	d := string(networkPayload.GetDestination())
	for _, s := range networkPayload.GetSources() {
		cp := string(s) + "," + d
		connectionPaths = append(connectionPaths, cp)
	}

	return connectionPaths, nil
}
