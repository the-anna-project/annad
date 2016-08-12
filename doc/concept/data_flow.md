# data flow
To understand how Anna works it is mandatory to understand the data flow within
her various components. We need clean interfaces between these components to
ease the development of their specific business logic. The following picture
helps to understand the data flowing within Anna. Here the solid arrows show in
which direction data types are transformed, while the dashed arrows indicate in
which direction data types are transported. On the very left is described in
which layer the shown data actually happens.

![data flow](image/data_flow.png)
