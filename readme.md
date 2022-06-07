# Project Layout
For purpose to implement a clean code architecture, I will use the following directory structure:

```
├── README.md
├── cmd // 1. 
│   ├── api
│   │   ├── main.go // 2. 
│  ...
├── pkg // 3.
│   ├── postgres // 4.
│   ├── redis // 5. 
├   ├── utils // 6. 
├  ...
├── internal // 7.
│   ├── model // 8.
│   ├── infra // 9.
│   │   ├── repositories // 10.
├   ├   ├── http // 11.
│   ├── usecase // 12.
├
├── migrations // 13.


```

Legend:

1. cmd package contains all executable entry point.
1. contains the http service entry point
1. contains all packages that can be used by other aplication. 
1. contains of connection to db
1. contains of connection and function of redis
1. contains of helper
1. any package defined under internal can not be imported by other package outside this repo. 
1. as an example model package will contain all repository models contract.
1. contains the implementation detail for models and other dependencies to external system
1. repositories package will contains the implementation detail for all modules under model that implemented as an SQL repository, or PostgreSQL to be precise.
1. any package relate to how deliver from user to usecase. Kind of request, response, and handler.
1. usecase package will contain all usecase for this service, for now we used model directly from handler, but idealy handler must depend only on usecases instead of model. the dependency flow must follow this pattern: input layer must only depend on usecase, and usecase must only depend on the contract instead of the implementation detail.
1. migration contains of migration to db.


## Documentation
https://documenter.getpostman.com/view/13002065/Uz5JJGEk