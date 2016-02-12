# pattern recognition
**Note that the following is not fully thought through and is currently only
there to help to form more concrete ideas around solutions.**

Each implementation of pattern recognition is a
[BB](/doc/concept/clg.md#bb-basic-behavior). The goal is to make Anna find
patterns herself and even make her find patterns for pattern finding. Thus she
needs to be able to discover new patterns herself.

Currently we only concentrate on text patterns, because at the moment there is
only the text interface implemented. Analysing binary data e.g. for image
analysis will be ignored for now.

### text patterns
Lets assume we have the following text inputs to analyse.

```
What is the name of the capital of england?
What is the size of the capital of france?
Who is the major of the capital of germany?
The name of the capital of portugal is lisbon.
Lisbon is the name of the capital of portugal.
```

A human is able to recognize some patterns here.
```
What   is the name  of the capital of england?
What   is the size  of the capital of france?
Who    is the major of the capital of germany?
          The name  of the capital of portugal  is lisbon.
Lisbon is the name  of the capital of portugal.
```

So Anna must be able to recognize patterns herself and even discover new
patterns to apply them to increase success on responses. Probably the list of
the following patterns are not complete.

### order
The word order makes Anna understand how likely it is what word follows the
other. The analysed word order can look something like this.
```
count    1st        2nd        3rd

2        What       is         the
2        is         the        name
3        the        name       of
3        name       of         the
5        of         the        capital
5        the        capital    of
1        capital    of         england
1        of         england    ?
```

### distance
The relative distance between words makes Anna understand what patterns exist
in human language. Using this knowledge Anna can apply these patterns in
return. It is also possible to find out how to fill the gaps between distant
words. The analysed word distance can look something like this.
```
count    distance    start    end

3        3           name     capital
2        2           capital  portugal
```

### position
The word position makes Anna understand where words can be put. The most
interesting information captured here are probably the very first and the very
last words or characters of the inout. E.g. the very first words will probably
be upper cased, the very last characters will probably imply the end of
sentences. Applying these information is another challenge. The analysed word
position can look something like this.
```
count    pos    word
1        0      Who         # How to learn that this is the first one and needs to be capitalized?
2        0      What
1        1      name
2        3      name
5        6      capital
2        8      portugal
3        9      ?           # How to learn that this is the last one and marks questions?
```
