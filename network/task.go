package network

func (n *network) Boot() {
	// Create new worker pool.
	listenPool.Execute(n.listenFunc)
	// Create new worker pool.
	networkPool.Execute(n.networkFunc)
}

func (n *network) listenFunc() {
	for {
		select {
		case <-n.Closer:
			return
		case textRequest := <-n.TextInput:
			// TODO
		}
	}
}

func (n *network) networkFunc() {
	for ID, CLG := range n.CLGs {
		go func(ID spec.ObjectID, CLG spec.CLG) {
			for {
				select {
				case <-n.Closer:
					return
				case payload := <-CLG.GetInputChannel():
					// TODO
				}
			}
		}(ID, CLG)
	}
}
