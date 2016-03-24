# mind

**TODO differentiate between mind, consciousness and imagination**

I think what is key for Anna to be intelligent is to be self aware. She needs
to have an idea about herself. About the difference between herself and the
sorounding environment.

One idea how such a mind could be implemented would be this. When thinking
about the mind and how imagination works, this has a lot to do with
visualization. Certain brain movies are created and more or less recognized. In
any way such pictures, ideas and reflections influence the human life a lot.
All light and dark sides of ourselves evolve out of this. So necessary woud be
some content generation that is then more or less recognized by Anna. The
content Anna produces here would be partially influenced by the current
environment and partially random. This content can be in form of pictures or
text. Depending on how much of certain brain movies would be recognized, Anna's
behaviour could be influenced. That way a simulation of the real mind would go
a step further. The surrounding environment influences created thoughts.
Created thoughts are recognized. Recognized thoughts influence behaviour.
Behaviour influences the surrounding environment. Note that this idea
inevitably includes free will and represents a serious risk.

To detect meaning, we need to see knowledge as a network, no matter if we can
make information visible or not. We need to find the nodes with the most
connections. These information can be used to feed the ballance system
dynamically. Anyway I think we really need visualization to truly understand
the meaning of things, to grasp contextual information over time and space
where information flow in. Further imagine the following problem. Having an
arbitrary sentence given, there are words that are more and there are words
that are less descriptive.  E.g. a noun like `tree` has more semantic meaning
and is more descriptive than the word `have`. We can try to find out what
sequences have a dedicated meaning by looking up connections between nodes
within the knowledge network. Inspecting the lookups outcome we will see that
nodes connected with the `tree` node are more tightly coupled and
interconnected. The picture of the connections of the `have` node should be
more mixed across all kinds of different semantics, causing contradictionary
information.

Imagine inputs like that.
```
Trees have leaves.
Trees have roots.
Children have lollipos.
Children have shoes.
```

We see that `leaves` and `roots` are more tightly coupled to `trees`. This
connections form a dense cluster where participating nodes have first class
connections.
```
trees
  - leaves
  - roots
```

We see that the nouns have more broad connections across different topics where
nodes semantically don't really have tight connections to each other.
```
have
  - lollipos
  - children
  - shoes
  - trees
  - leaves
```

**One idea could be to detect this pattern using k-means clustering and using
the occupied space of a clusters vectors to weight semantic relationships.**

See
- https://www.youtube.com/watch?v=Q-B_ONJIEcE
- https://www.youtube.com/watch?v=hBpetDxIEMU
- https://en.wikipedia.org/wiki/Damasio%27s_theory_of_consciousness.
