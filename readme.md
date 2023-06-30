## Payment Gateway Dev Test

## Dependencies

* Docker
* Docker compose

## Technologies used

[Golang](https://go.dev), [Typescript](https://www.typescriptlang.org), [gRPC](https://grpc.io/), [MariaDB](https://mariadb.org), [Redis](https://redis.io).

## Execution

* Clone this repository
* Execute `docker compose up -d`

After execution, the following services will be available for access.

Transactions microservice = http://127.0.0.1:9007

Authorization microservice = tcp://127.0.0.1:50051

MariaDB = root:13152c0e-ce51-4e03-bebd-8e2c4fb359bf@127.0.0.1:13306

Redis = 127.0.0.1:16379

PHPMyAdmin = http://127.0.0.1:13307

Redis Commander = http://127.0.0.1:16380

To view running services, just execute the command `docker compose ps`.

To view logs of running microservices, just execute the command `docker compose logs -f ctransaction cauthorization cnotification`.

### About the project

The project was made using some concepts about Domain Driven Design, like layered application, bounded context and repository pattern.

The domain layer have more than 95% of corevage, all unit tests, just to cover all cases success and mapped errors as a possible. This was possible by the implementation of DIP "Dependency Inversion Principle", represented by letter D in SOLID.

The domain layer is agnostic for the source data, was possible to create a command for API listen HTTP, but, the domain layer can receive data from any source, like event driven for example.

Another interesting point is the entities of domain isn't related to a database modeling and was proposital, as a way to show this separation, that allows improvements, fixes or new implementations, in the domain layer side or in other layer side, like in this case that was in infra.

The infra layer contains all necessary code to sustain the domain layer, as the code that interact the external database, api controllers, integration with other microservices by gprc or event dispatch with pubsub concepts. Some design patterns were used to be more dynamic the operation, like abstract factory for creation of domain entities from the request data or the opposite, creation of response entities from the domain entities.

