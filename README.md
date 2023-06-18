## Descrição
Esse projeto visa receber um arquivo txt/csv tratar seus dados e persistir suas informações em um banco de dados.
O banco de dados escolhido foi o `postgres` ele é bem estavel oferece uma grande variedade de recursos alem de ser bastante utilizado.

## Padronização
Foi decido utilizar os padroes descritos no seguinte artigo `https://blog.devgenius.io/golang-naming-conventions-72bbaf84e959`

# Tabelas
O projeto dispoe das seguinte tabela `purchase`.

## purchase

Essa tabela é responsavel por armazenar os dados da compra.
Ela dispoe dos seguintes campos.
### id SERIAL PRIMARY KEY
Reponsavel pela chave primaria da tabela em questão

### cpf_cnpj VARCHAR(18) NOT NULL
Resposavel por armazenar os dados do CPF/CNPJ esses dados não pode ser nulos, os dados serão persistidos com as devidas maskaras, foi decidido assim já que as mesmas não muda com frequencia.

### private BOOLEAN
Responsavel por armazenar o valor booleano.

###  incomplete BOOLEAN
Responsavel por armazenar o valor booleano.

### date_last_compra timestamp NULL
Responsavel por armazenar a data em formato de timestamp, esse formato foi selecionado por permitir uma maior precisão sendo possivel armazenar informações sobre ano, mês, dia , hora, minutos e segundos, pode ser formato de varias formas dentre outras

### average_ticket FLOAT NULL
Responsavel por armazenar o valor flutuante.

### last_purchase_ticket FLOAT NULL
Responsavel por armazenar o valor flutuante.
### store_more_frequent VARCHAR(18)
Resposavel por armazenar os dados do CNPJ esses dados podem ser nulos, os dados serão persistidos com as devidas maskaras.

### store_last_purchase VARCHAR(18) NULL
Resposavel por armazenar os dados do CNPJ esses dados podem ser nulos, os dados serão persistidos com as devidas maskaras.


## Requisitos:
É necessario ter instalado o `Docker` e ter acesso a `internet`.
### Primeiros passos:

1. Rodar o comando `docker compose up -d` para criar os containers do postgres e do golang.
2. Irá aparecer a mensagem `Server is running on port :8080`

### Testes

Para executar os testes é necessário rodar o seguinte comando:

`$ go test ./...`

## Dependencias utilizadas

1. Viper `https://github.com/spf13/viper`
2. CPFCNPJ `https://github.com/klassmann/cpfcnpj`
3. Mockery  `https://github.com/vektra/mockery`

## Viper

O viper oferece um gerenciamento de configurações flexiveis, alem de suportar uma boa quantidade de formatos como JSON, YAML dentre outros.
Ele também suporta a configuração de varios ambientes.

## CPFCNPJ

o CPFCNPJ foi escolhido por ser facil de utilizar e por oferecer funções que é necessaria para a validação dos dados e os regex necessarios para gerar as maskaras.

## Mockery

Ele é de facil utilização com sintaxe simples, ele acelera o desenvolvimento dos testes por gerar os mocks substituindo objetos reais.

## Postgres

Foi decidido utilizar o driver padrão do GO, alem de não ser necessario a instalação de libs de terceiros, ele atende as necessedades do projeto em tarefas não muito complexas.
