# character tree

### challenge

###### simple input
This input is send in one request.
```
Split the string "ab" in the middle.
```

###### divided input
This two inputs are send in different requests.
```
Here is the string "ab".
Split it in the middle.
```

input variations
```
Cut the string "ab" in the middle.
Split the string "ab" into half.
Make the string "ab" split into half.
```

output
```
a b
```

### false positive

input
```
The string "ab" in the middle.
```

outupt
```

```

### data structures

###### impulse-input relationship
Relationships between impulse and input can be stored as a hash map. `key` is a
impulse ID holding basic information about itself. Information that should be
stored are this.

- given input
- given expected output
- count of how often the impulse's input was triggered to generate sufficient output
- count of how often the impulse's input helped to create a successful response
- date of occurrence

```
key
  hashkey: hashvalue
  hashkey: hashvalue
  ...      ...
```

---

###### impulse-session relationship
Relationships between impulse and session can be stored as a key-value pair.
`key` is an impulse ID. `value` is a session ID. This is a 1:1 relationship.
One impulse is associated with exactly one session.

```
key: value
```

---

###### session-impulse relationship
Relationships between session and impulse can be stored as a list. `key` is a
session ID. `listitem` is an impulse ID. This is a 1:n relationship. One
session is associated with multiple impulses.
```
key
  listitem, listitem, ...
```

---

###### basic character information
Basic character information is stored as a hash map. `key` is a character
sequence holding basic information about itself. How many characters should
form the sequence should be dynamically found out using different strategies.
Default might be 3. Information that should be stored are this.

- count of how often this sequence was ever seen
- count of how often this sequence helped to create a successful response
- date of first occurrence
- date of last occurrence

```
key
  hashkey: hashvalue
  hashkey: hashvalue
  ...      ...
```

---

###### basic word information
Basic word information is stored as a hash map. `key` is a word sequence
holding basic information about itself. How many word should form the sequence
should be dynamically found out using different strategies. Default might be 3.
Information that should be stored are this.

- count of how often this sequence was ever seen
- count of how often this sequence helped to create a successful response
- date of first occurrence
- date of last occurrence
- impulse ID
- session ID


```
key
  hashkey: hashvalue
  hashkey: hashvalue
  ...      ...
```

---

# ???

impulse-generated-output relationship
```
impulseID
  split, mid dle, ...
```

### semantic relationship
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
