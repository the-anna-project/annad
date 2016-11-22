# data flow
To understand how Anna works it is mandatory to understand the data flow within
her various components. We need clean interfaces between these components to
ease the development of their specific business logic.

The following picture helps to understand the data flowing within Anna. Here
the solid arrows show in which direction data types are transformed, while the
dashed arrows indicate in which direction data types are transported. On the
very left is described in which layer the shown data flow actually happens.

- Entrypoint here is [annactl](annactl.md), the command line client.
  [input](input.md) is received from the user and transformed into a
  [TextRequest](https://godoc.org/github.com/the-anna-project/annad/object/networkresponse#TextRequest).
  This request object is provided to the [client](client.md).
- The client is asked to execute some business logic with respect to the given
  TextRequest, which is transformed into a
  [StreamTextRequest](https://godoc.org/github.com/the-anna-project/annad/object/networkresponse#StreamTextRequest)
  and send to the [server](server.md).
- The server receives the StreamTextRequest and transforms it back into a
  TextRequest. This is then forwarded to the [network](network.md).
- The network takes the TextRequest, transforms it to a
  [NetworkPayload](https://godoc.org/github.com/the-anna-project/annad/spec#NetworkPayload).
  [CLGs](clg.md) take over with by processing and dispatching further network
  payloads through the neural network. At some point the neural network decides
  to stop processing and returns the output as shown below until data is flown
  back to the user.

![data flow](image/data_flow.png)
