package network

type Gateway interface {
	String() chan string
}

func NewGateway() Gateway {
	return gateway{
		StringGateway: make(chan string, 1000),
	}
}

type gateway struct {
	StringGateway chan string
}

func (g gateway) String() chan string {
	return g.StringGateway
}
