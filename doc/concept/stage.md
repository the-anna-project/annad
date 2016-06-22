# stage
A stage can be seen as a layer of a neural network. In every stage a
[strategy](strategy.md) is created and executed. The strategy creation and its
specific behavior is defined by its applied [connections](connection.md).

The stage design helps to fill the gap between information and behavior. Any
problem can be solved using the the self contained stage design, regardless the
problem's complexity.

![stage](image/stage.png)
