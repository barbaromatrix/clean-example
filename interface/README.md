# What is this layer?
This layer acts as a glue to connect the outer data, coming from the outer circles, to the inner circles, handling API requests and talking to the database.
This layer is composed by 3 directories:
- `Controllers` - In charge of handling API requests that come from the outer layer.
- `Presenters` - In charge of the Output Port (what goes to outside)
- `Repository` - Contains the implementation of [repository](usecase/repository) abstraction on the usecase folder