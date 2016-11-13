package connection

import (
	"fmt"
	"time"

	objectspec "github.com/xh3b4sd/anna/object/spec"
)

func (s *service) sortPeers(a, b objectspec.Peer) (objectspec.Peer, objectspec.Peer) {
	if a.Value() > b.Value() {
		return b, a
	}

	return a, b
}

func (s *service) newUnixSeconds() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}
