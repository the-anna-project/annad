package prometheus

import (
	"fmt"
)

func helpFor(metricType, key string) string {
	return fmt.Sprintf("This %s is represented by the metric name %s.", metricType, key)
}
