# What is this layer?
This layer acts as a glue to connect the outer data, coming from the outer circles, to the inner circles, handling API requests for example. Please take care of not mixing business logic here, only concerns about technological logic.
This layer is composed by 3 directories:
- `Controllers` - In charge of handling API requests that come from the outer layer (convert any data coming from outer circle to a format that the business can understand)
- `Presenters` - In charge of the Output Port (convert the database's data format to send it to the outcircle)
- `Repository` - Contains the implementation of [repository](usecase/repository) abstraction on the usecase folder