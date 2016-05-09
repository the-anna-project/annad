package profile

// TODO
func (g *generator) getCLGRightSideNeighbours(collection spec.CLGCollection, clgName string, methodValue reflect.Value, canceler <-chan struct{}) ([]string, error) {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call getCLGRightSideNeighbours")

	// TODO
	// Fill a queue with all CLG names.
	clgNameQueue := g.getCLGNameQueue(g.CLGNames)

	// Initialize the profile creation.
	for {
		select {
		case <-canceler:
			return nil, maskAny(workerCanceledError)
		case clgName := <-clgNameQueue:
		}
	}

	//     find right side neighbours for given clg name
	//         if no profile for checked neighbour
	//             push neighbour name back to channel

	return nil, nil
}

// TODO
func (g *generator) isRightSideCLGNeighbour(collection spec.CLGCollection, left, right spec.CLGProfile) (bool, error) {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call isRightSideNeighbour")

	// run clg chain
	// if error
	//     return false

	return false, nil
}
