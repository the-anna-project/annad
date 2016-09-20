# CLG, complex logic gate
CLG, Complex logic gate, implements BB, Basic Behavior, when receiving input in
form of BI, Basic Inormation. Basic behavior is created when some sort of
complex logic operation is processed with respect to some fundamental input. In
computer science there are logic gates implementing logic operations causing
some action. The concept of CLGs presumes the need of more sophisticated logic
gates to mimic human behavior. The theory is that as soon as there are enough
sufficient CLGs, intellectual explosion will happen. CLGs are supposed to be
the minimal and fundamental inborn logic units that can be used out of the box.
This builtin functionality should be the only hard coded logic. All other
behavior must result out of the combination of learned data, CLGs and
[connections](connection.md) between these two.

A CLG creates behavior when having information provided during execution. This
is how information is translated into behavior. Using some function
implementing some functionality.

Random fact: currently Anna is capable of more than 130 CLG implementations.
This number is expected to be raised even further. For now manually, later on
automatically. See https://godoc.org/github.com/xh3b4sd/anna/spec#CLGIndex.

### BI, basic information
BI, Basic Information, represent fundamental information that must be known to
reach a higher level of intelligence. The most fundamental BIs like numbers
from the decimal number system or letters from the latin alphabet simply need
to be provided. A human and a machine respectively cannot think of something
totally unknown out of nothing. There needs to be information in the first
place, so some kind of intelligence can collect and connect these to generate
new knowledge. Looking at the process of how humans learn, we have teachers and
other humans who explain fundamental concepts that help to explore the world by
our own. Having that said there is a need for some guidance for a machine to
learn. Goal should be to make a machine able to gather information by itself
using available sensors, interfaces and resources it is able to access.

### BB, basic behavior
BB, Basic Behavior, can probably be seen as instinct, intuition or talent.
There are presets within humans. These enable us to do things without actively
thinking about them. These things are might be in our DNA, subconscious mind or
as football trainers would point at it, in our blood. BB lets us breathe from
the first second of our life, make us fear dangerous things and let us explore
our environment using our inborn senses. So some sort of BB a machine needs to
have to learn in a guided fashion and later autonomously.

### the intelligence gap
We can write the following formular to express how intelligence is created: `I
= ibⁿ`. Intelligence equals information times exponential behavior. Simply put,
one piece of information, combined with multiple pieces of behaviour, creates
some sort of intelligence.

Anyway, there is a technical issue that needs to be solved to fix the
intelligence gap, to make artificial intelligence work. The intelligence gap
describes the gap between information and behavior. This gap needs to be filled
to make the intelligence equation work. This technical issue is solved by the
concept of CLGs.

### implementation
The following describes implementation details necessary to consider to make
the neural network work. The described details intend to suggest one possible
approach. There are maybe others which have not ben found yet.

###### input (4) done
- receive user provided input
- write user provided input into storage using information ID
- add information ID to current context

###### read input done
- lookup information ID from context
- read user provided input from storage by information ID
- return user provided input

###### feature size (1) wontfix
- check when feature size was last updated
- if last updated is long enough ago, increment or decrement feature size by 1
- store feature size into storage
- write last updated into storage

###### split features (2) done
- receive information
- read feature size from storage
- split information into features
- write features into storage

###### pair syntactic (5) done
- read features from storage
- lookup syntactic relations of features
- combine features that belong together
- write pair into storage

###### pair semantic (3) wontfix
- read features from storage
- lookup conceptual relations of features
- combine features that belong together
- write pair into storage

###### evaluate certenty (7)
- read pairs from storage
- lookup relations of pair and user input
- calculate certenty
- remove pair from current key
- store pair with key having certenty applied

###### derive behavior (8)
- read information from certenty level
- lookup relations of information and behavior
- write spontaneous connection into storage

###### maximum certenty
- read maximum certenty from information pyramid
- write maximum certenty into storage, if maximum certenty is bigger than the current one

###### output (6)
- read information from certenty level
- read certenty range from storage
- write minimum certenty into storage, if expectation matches and minimum certenty is smaller than current one
- if information is within certenty range, return information, otherwise return error to keep neural activity up
