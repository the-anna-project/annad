package textinterface

import ()

type TextInterface interface {
	FetchURL(url string) error
	ReadFile(file string) error
	ReadStream(stream string) error
	ReadPlain(plain string) error
}

type TextInterfaceConfig struct {
	StringChannel chan string
}

func DefaultTextInterfaceConfig() TextInterfaceConfig {
	return TextInterfaceConfig{
		StringChannel: make(chan string),
	}
}

func NewTextInterface(config TextInterfaceConfig) TextInterface {
	return textInterface{
		TextInterfaceConfig: config,
	}
}

type textInterface struct {
	TextInterfaceConfig
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
	ti.StringChannel <- plain
	return nil
}
