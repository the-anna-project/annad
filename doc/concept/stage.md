# stage
A stage can be seen as a layer of a neural network. The number of stages to be
processed is not hardcoded, but fully dynamic and context dependent. In every
stage a [strategy](strategy.md) is created and executed. The strategy creation
and its specific behavior is defined by its applied
[connections](connection.md).

The stage design helps to fill the gap between information and behavior. Any
problem can be solved using the the self contained stage design, regardless the
problem's complexity.

![stage](image/stage.png)
