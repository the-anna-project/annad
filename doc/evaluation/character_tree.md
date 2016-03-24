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

how to structure and measure semantic relationships?

---

# ???

impulse-generated-output relationship
```
impulseID
  split, mid dle, ...
```
