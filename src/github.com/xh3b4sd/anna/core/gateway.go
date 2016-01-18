package core

// gateway

type Gateway interface {
	Close() error

	GetTextGateway() TextGateway
}

func NewGateway() Gateway {
	return gateway{
		TextGateway: NewTextGateway(),
	}
}

type gateway struct {
	TextGateway TextGateway
}

func (g gateway) Close() error {
	return nil
}

func (g gateway) GetTextGateway() TextGateway {
	return g.TextGateway
}

// text gateway

func NewTextGateway() TextGateway {
	return textGateway{
		StringChannel: make(chan string, 1000),
	}
}

type TextGateway interface {
	GetStringChannel() chan string
}

type textGateway struct {
	StringChannel chan string
}

func (g textGateway) GetStringChannel() chan string {
	return g.StringChannel
}
