package network

type State interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
}

func NewState() State {
	return state{}
}

type state struct{}

func (s state) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (s state) UnmarshalJSON([]byte) error {
	return nil
}
