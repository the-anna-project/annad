package core

// gateway

type Gateway interface {
	Close() error

	Text() TextGateway
}

func NewGateway() Gateway {
	return gateway{
		Text: NewTextGateway(),
	}
}

func (g gateway) Text() TextGateway {
	return g.StringGateway
}

// text gateway

func NewTextGateway() TextGateway {
	return textGateway{
		String: make(chan string, 1000),
	}
}

type textGateway interface {
	String() chan string
}

type textGateway struct {
	String chan string
}

func (g textGateway) String() chan string {
	return g.StringGateway
}
