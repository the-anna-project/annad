package network

// TODO
func (n *network) Boot() {
	// Create new worker pool.
	inputPool.Execute(n.InputListener)
	// Create new worker pool.
	eventPool.Execute(n.EventListener)
}

func (n *network) InputListener() error {
	CLG, err := n.clgByName("input")
	if err != nil {
		return maskAny(err)
	}

	for {
		select {
		case <-n.Closer:
			return nil
		case textRequest := <-n.TextInput:
			ctx := context.MustNew()
			err := n.InputHandler(ctx, CLG, textRequest)
			if err != nil {
				n.Log.WithTags(spec.Tags{C: nil, L: "E", O: n, V: 4}, "%#v", maskAny(err))
			}
		}
	}
}

func (n *network) InputHandler(ctx spec.Context, CLG spec.CLG, textRequest spec.TextRequest) error {
	// This should only be used for testing to bypass the neural network
	// and directly respond with the received input.
	if textRequest.GetEcho() {
		// TODO set CLG to output CLG
	}

	// TODO create network payload from text request
}

func (n *network) EventListener() {
	for {
		select {
		case <-n.Closer:
			return
		default:
			// TODO
			// TODO create context according to event
			// TODO lookup CLG according to event
			// TODO read network payload from Storage.PopFromList(key)
			err := n.EventHandler(ctx, CLG, networkPayload)
			if err != nil {
				n.Log.WithTags(spec.Tags{C: nil, L: "E", O: n, V: 4}, "%#v", maskAny(err))
			}
		}
	}
}

func (n *network) EventHandler(ctx spec.Context, CLG spec.CLG, networkPayload spec.NetworkPayload) {
}
