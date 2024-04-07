# Como implementar design orientado a domínio (DDD) em Golang

## O que é DDD
Domain-Driven Design é uma forma de estruturar e modelar o software de acordo com o domínio ao qual pertence. O que isso significa é que primeiro deve ser considerado um domínio para o software que é escrito. O domínio é o tema ou problema que o software pretende trabalhar. O software deve ser escrito para refletir o domínio.

Etapas

* Entidades
  * item
  * person 
* objetos de Valor
  * transaction
* Agregando domínio
  * customer
    * test
* Repositórios em memória
  * repository
  * memory
    * test
* Serviços
* Cas

## Entidades
Uma entidade é uma struct que possui um Identificador e que pode mudar de estado, por mudar de estado queremos dizer que os valores da entidade podem mudar.

Criaremos duas entidades, para começar, Person e Item . Gosto de manter minhas entidades em um pacote separado para que possam ser usadas por todos os outros domínios.

## Objetos de Valor
Pode haver ocorrências onde temos structs que são imutáveis ​​e não precisam de um identificador único, essas structs são chamadas de Value Objects . Portanto, estruturas sem identificador e valores persistentes após a criação. Objetos de valor são frequentemente encontrados dentro de domínios e usados ​​para descrever certos aspectos desse domínio. Estaremos criando um objeto de valor por enquanto que é Transaction , uma vez que uma transação é executada, ela não pode mudar de estado.

## Agregando domínio
## Repositórios em memória

## Serviços
Um serviço vinculará todos os repositórios fracamente acoplados a uma lógica de negócios que atenda às necessidades de um determinado domínio. No caso da taverna, poderíamos ter um serviço de Pedido , responsável por encadear repositórios para realizar um pedido. Portanto, o serviço terá acesso a um CustomerRepository e a um ProductRepository

## Cas

