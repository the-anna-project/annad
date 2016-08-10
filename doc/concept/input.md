# input
Input should be self explaining. From a ten thousand feet view it represents
the input that can be provided through [interfaces](interface.md) to request
calculations from the [neural network](network.md). This can be any kind of
input. In the first place only string input can be provided, because the first
goals are going into the direction of natural language understanding. Later on
byte streams can be provided to learn image recognition or something like that.
The result of such calculations is some [output](output.md), which is streamed
back to the client.
