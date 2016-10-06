# starting a conversation by greeting
**Note that the following is not fully thought through and is currently only
there to help to form more concrete ideas around solutions.**

What we want to achieve is something like this.

```
human: Hello Anna!
anna: Hello! Who is there?
...
```

This looks pretty straight forward and simple. It is not. Here a lot of really
important skills take place. Lets have a look at what happens and lets break
the situation down to analyse each step.

### identify letters
Anna analyses incoming input. Here she alrwady needs to show how smart she
really is. What she sees is actually only some mystery bytes. To turn these
mystery bytes into something understandable we need to [guide
her](https://github.com/xh3b4sd/anna/blob/f0d28faa0f5d9ecba407b6678e3ebd767142e58c/doc/evaluation/guidance.md)
beforehand that there are digits and letters.

### create character tree
As soon as she knows what letters and words are, she can create a character
tree. The word `hello` would result in the following character tree: `h - e - l - l - o`.
The character tree should be "publicly" available across the whole basic core,
so it can be extended and used across multiple and parallel contexts. Each
node of the tree holds information about the string it is part of up to the
nodes position. That is, `e` holds `he`. Each node holds and adds further
information about itself and the string it is part of during its lifetime. The
node `o` would hold the string `hello` and the information that the nodes
associated string is actually some kind of greeting.

Questions
- What relation should `hello` - `greeting` have?
- Simply meta string information within a character neuron?
- A link to a new neuron? What neuron type?

### lookup historical events
Once the character tree is created and information of the input's associated
neurons are added to the current context, Anna can go on to further fill the
context with historical information about similar events she is now dealing
with.

Questions
- Where are historical events are stored?
- In which form are historical events stored?
- What needs to be stored besides request/response relationship?

### create response
When we have filled the context with all necessary information Anna can create
a response.

Creativity, inspriration, spontaneity and other human like fresh skills need to
influence behaviour as well. To fill a conversation with life is basically the
simple result of using already known information and choosing this one that
fits the current situation the most. What that actually means depends on the
experiences an individual made so far. After clustering context related
information one can be picked pseudo randomly and combined to actually create a
response for the current request.

### store historical data
The created response needs to be persisted ina structured way. So Anna can
comprehend later on what requests were there and what responses she gave and
how she came up with these responses.
