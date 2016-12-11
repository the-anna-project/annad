package service

// LayerCollection represents a collection of layer services. This scopes
// different layer service implementations in a simple container, which can
// easily be passed around.
type LayerCollection interface {
	Boot()
	Behaviour() LayerService
	Information() LayerService
	Position() LayerService
	SetBehaviourService(behaviourService LayerService)
	SetInformationService(informationService LayerService)
	SetPositionService(positionService LayerService)
}
