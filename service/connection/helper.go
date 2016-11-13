package connection

import (
	"fmt"
	"time"
)

func (s *service) sortPeers(a, b string) (string, string) {
	if a > b {
		return b, a
	}

	return a, b
}

func (s *service) newUnixSeconds() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}
