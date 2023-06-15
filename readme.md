## Credsystem Dev Test

## Dependências

* Docker
* Docker compose

## Execução

* Clonar este repositório
* Executar `docker compose up -d`

Após execução, os seguintes serviços estarão disponíveis para acesso.

Microserviço de Transactions = http://127.0.0.1:9007

Microserviço de Autorização = tcp://127.0.0.1:50051

MariaDB = root:13152c0e-ce51-4e03-bebd-8e2c4fb359bf@127.0.0.1:13306

Redis = 127.0.0.1:16379

PHPMyAdmin = http://127.0.0.1:13307

Redis Commander = http://127.0.0.1:16380

Para visualizar os serviços em execução, utilizar o comando `docker compose ps`.

Para visualizar os logs dos microserviços em execução, utilizar o comando `docker compose logs -f ctransaction cauthorization cnotification`.

### Sobre o projeto

O projeto foi feito usando alguns conceitos de Domain-Driven Design, como a aplicação de camadas, contexto delimitado (bounded context) e repository pattern.

A camada de domínio tem 100% de cobertura de testes, sendo todos eles unitários, justamente para cobrir todos os casos de sucesso e erros possíveis. O coverage pode ser consultado no comando `make coverage-docker` quando o projeto estiver em execução, essa possiblidade foi por conta da implementação do DIP "Dependency Inversion Principle", representado pela letra D no SOLID.

Não vi necessidade de implementação de testes na camada de infra, pois neste ponto é melhor a implementação de testes de integração em alguns casos e isto estaria muito fora do escopo.

Como a camada de domínio é agnóstica a origem de entrada de dados, foi possível fazer um comando para o listen de api HTTP, porém, a camada de domínio pode receber dados de qualquer fonte, como um evento por exemplo.

Um outro ponto interessante é ver que as entidades do domínio não estão presos a modelagem de banco de dados, sendo que o schema implementado foi proposital como forma de mostrar esta separação, o que permite melhorias, correções ou novas implementações, tanto do lado do domínio, quanto do lado da infra, nesse caso, a estrutura de dados.

Já a camada de infra contém todo o código necessário para a sustentação da camada de domínio, como o código que de fato faz a interação com o banco de dados, controllers da api, integração com outros microserviços via grpc ou disparo de eventos com conceitos de pubsub. Alguns design patterns foram utilizados para tornar mais dinâmica a operação, como abstract factory para criação das entidades de domínio a partir dos dados de request ou o contrário, criação das entidades de response a partir das entidades de domínio.

