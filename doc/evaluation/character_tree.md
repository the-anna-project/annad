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

# relative character relevance
Relevance of character sequences can be efficiently stored in a diagram like
matrix. The matrix provides `n` rows dependening on the required distributional
accuracy. The more rows the matrix has the more differentiations can be made.
The value of each row represents the sequence having highest score for this
distribution. A sequence's relevant score is compared against each row of the
matrix. Is the score bigger than the row value, the next row is compared. Is
the score lower then the next row's value, the former row'S value will be set
to the score of the current sequence.

That way a relative relevance can be detected by comparing all `n` rows of the
matrix and calculating medians, percentiles or other metrics. Using this method
metric spikes can be ignored or used. This approach further benefits from its
limited amount of information required to make the metrics detection work. Note
that the sequence's score can be anything. A counter for occurrence, success, or
failure, or simply a timestamp.

For several reasons the methods metrics can be capped. To prevent technical
overflows or subsets fulfilling certain requirements. Lets say there is a
maximum value given scores must not exceed. This cap then must be implemented
by preserving the relative distances between the matrix's rows. To prevent
overflowing all rows values can be cut down by some percentage, e.g. 50%. This
leads to the problem that the manipulating sequences still hold their 100%
scores. The matrix defines the cap that must be applied to each sequence, if
not yet done. For this there must be the information of the cap value, and if
the cap should be applied before operating against the matrix. Here that means
any operation. Reads and writes.

This is how such a matrix looks like.
```
 1    -
 2    -
 3    -
 4    --
 5    --
 6    --
 7    ---
 8    ---
 9    ----
10    ----------
```
