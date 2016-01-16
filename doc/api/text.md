# text
Here we cover Annas text interface. This is used to read text over network.
Text can either be pulled or pushed.

### fetch url

###### request url
```
/interface/text/action/fetchurl
```

###### request body
```
{
  "url": "<url>"
}
```

---

### pull file

###### read file
```
/interface/text/action/readfile
```

###### request body
```
{
  "file": "<file>"
}
```

---

### read stream

###### request url
```
/interface/text/action/readstream
```

###### request body
Note that this is does not work and needs thinking.
```
{
  "stream": "<stream>"
}
```

---

### read plain

###### request url
```
/interface/text/action/readplain
```

###### request body
```
{
  "plain": "<plain>"
}
```
