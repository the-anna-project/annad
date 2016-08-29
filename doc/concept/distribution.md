# distribution
A distribution is a calculation object that describes the distribution of
vectors within space. This vectors can be of arbitrary dimensions. That way
characteristics of features and their location in space can be represented and
alanysed. See https://godoc.org/github.com/xh3b4sd/anna/spec#Distribution for
implementation details.

### balance system
The distribution can be used as balance system. That describes the amount of
signals and how harmonic or extreme the distribution of these currently is
within a system. That way the state of an organism can be represented. Extremes
of signal occurrence cause imbalance within the organism. Extremes can be seen
as low or high.

No signals at all represent a low extreme. This is the bar chart representation
of no signals.
```
x  x  x

1  2  3
```

To many signals at the same time represent a high extreme. This is the bar
chart representation of too many signals.
```
      x
      x
      x
      x
x  x  x

1  2  3
```

The organism's motivation is the balance of the signal distribution. This is
the bar chart representation of a more balanced distribution of signals.
```
      x
x     x
x  x  x
x  x  x
x  x  x

1  2  3
```

As we see there are different channels obtaining their own signal population.
This can indicate different evaluations of whatever is going on within the
organism. Each channel has an separate input and an separate output. Pushing an
signal to the input of one channel causes the signal distribution for the given
channel to increase. The whole signal population within a balance system is
capped. In case the overall signal population is already saturated, an adaption
of the other channels happens automatically. That means, that at the same time
some input is received on one channel, some output is received on the other
channels. The population of one channel increases and the population of the
other channels decreases symmetrically. Having three saturated channels and one
channel pushes an signal, the value of it is divided by two and population
worth the half of the pushed signal is decreased from the other two channels,
causing a balance or imbalance.

# weighted analysis
Weighted analysis can represent vector population within space. The more
channels the distribution has the more differentiations can be made between
vectors being located within channels. E.g. analysing sequence features
requires an abstracted way of representing patterns. Using a distribution
feature locations within sequences can be visualized and compared. Thinking
about the sequence `.`, humans instantly recognize the period and associate the
location being the end of a sentence. Since all this thinking is obvious for
humans, it is not for machines. The distribution helps to find out that the
sequence `.` is in fact a feature, because it can be detected as recurring
pattern. The distribution for the feature `.` would look something like this,
because periods are almost always located at the end of sentences.
```
                                                                             x
                                                                             x
                                                                             x
                                                                        x    x
                                                                    x   x    x

5  10  15  20  25  30  35  40  45  50  55  60  65  70  75  80  85  90  95  100
```
