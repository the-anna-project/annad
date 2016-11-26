package spec

// EndpointCollection represents a collection of endpoint services. This scopes
// different endpoint service implementations in a simple container, which can
// easily be passed around.
type EndpointCollection interface {
	Boot()
	Metric() EndpointService
	SetMetricService(metricService EndpointService)
	SetTextService(textService EndpointService)
	Shutdown()
	Text() EndpointService
}
