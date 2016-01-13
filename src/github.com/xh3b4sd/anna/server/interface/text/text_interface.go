package textinterface

import ()

type TextInterface interface {
	FetchURL(url string) error
	ReadFile(file string) error
	ReadStream(stream string) error
	ReadPlain(plain string) error
}

type NewTextInterfaceConfig struct {
	StringGateway chan string
}

func NewTextInterface(config NewTextInterfaceConfig) TextInterface {
	return textInterface{
		NewTextInterfaceConfig: config,
	}
}

type textInterface struct {
	NewTextInterfaceConfig
}

// TODO this should actually fetch a url from the web
func (ti textInterface) FetchURL(url string) error {
	return nil
}

// TODO this should actually read a file from file system
func (ti textInterface) ReadFile(file string) error {
	return nil
}

// TODO this should actually be streamed
func (ti textInterface) ReadStream(stream string) error {
	return nil
}

func (ti textInterface) ReadPlain(plain string) error {
	ti.StringGateway <- plain
	return nil
}
