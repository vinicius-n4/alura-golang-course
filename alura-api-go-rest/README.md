# alura-api-go-rest

1. Subir o banco de dados PostgreSQL:
```shell
$ docker-compose up
```

2. Rodar o servidor localmente:
```sh
$ go run main.go
```

Vamos integrar uma aplicação front-end desenvolvida em React para consumir a API.

Verifique se Nodejs está instalado em sua máquina, executando o seguinte comando no terminal:
```sh
node --version
```

Se sua versão aparecer na tela, pode passar para o próximo passo. Caso a versão do Nodejs não apareça na tela, realize
por gentileza o download do Nodejs em sua máquina, de acordo com seu sistema operacional.

Agora abra um terminal na pasta do projeto e digite o seguinte comando para instalar as dependências do React:
```sh
npm install
```

Atualize as dependências do npm com o seguinte comando:
```sh
npm update
```

Para finalizar, ainda no terminal digite o comando abaixo para subir o servidor do React e aguarde:
```sh
npm start
```

Após o carregamento, será exibida a página do projeto.
