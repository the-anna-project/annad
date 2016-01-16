package server

import (
	"log"
	"net/http"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/language"
	"github.com/xh3b4sd/anna/server/interface/text"
)

func Listen() {
	ctx := context.Background()

	// language network
	ln := language.NewLanguageNetwork()

	// text interface
	config := textinterface.NewTextInterfaceConfig{
		StringGateway: ln.Gateway().String(),
	}
	ti := textinterface.NewTextInterface(config)
	handlers := textinterface.NewHandlers(ctx, ti)

	// http
	for url, handler := range handlers {
		http.Handle(url, handler)
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}
