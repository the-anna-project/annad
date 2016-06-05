# core network
The core network receives signals from the gateway. These are translated to
impulses that walk through the core network to create some output with respect
to the given input. The CoreNet executes one [CLG](clg.md) after another and
forms dynamic [strategies](strategy.md) that way. How the strategy then looks
like is determined by the [knowledge network](knowledge_network.md). Is there
no information available in the knowledge network, a strategy is cerated using
an implementation of a [permutation factory](permutation_factory).
