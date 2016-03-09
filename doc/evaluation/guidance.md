# guidance
**Note that the following is not fully thought through and is currently only
there to help to form more concrete ideas around solutions.**

Guided learning is essential to gather fundamental information Anna cannot know
out of nothing. This can improve and speed up Anna's learning process. It is
administrative task that needs to be hard coded though. Since this provides
mechanisms to manipulate and help Anna structuring information, this option
should only be considered as very last option. The most important thing is to
make Anna learn herself, to make her dynamic and smart. Guiding her through all
problems will not push her into the right direction, but keep her dumb and
static. So it always must be very consciously thought about when, why and how
Anna should be taught actively.

Finally the strategic results of guidance should be recognized by Anna and lead
to a dynamically learned behavior and cause consciously connected neurons within
Anna's neural networks.

### provide BI, Basic Information
[BIs](/doc/concept/clg.md#bi-basic-information) need to be provided to reach
the next level of knowledge. This is only permitted where no other option is
sufficient. E.g. this BIs probably need to be provided from the very beginning
- all digits from 0 to 9
- all letters from a to z

### connect neurons
Anna needs to connect neurons to gather knowledge. It may happen that there
cannot be made all important connections by herself. Then there could be the
option to tell her what neurons to connect and why. Here it is important that
this behavior is implemented in a generic way so we don't create exceptions
only to support this special case. This approach must be part of Anna's
[BBs](/doc/concept/clg.md#bb-basic-behavior). The interface of such generic
behavior model could be seen as `when x then y` action.
