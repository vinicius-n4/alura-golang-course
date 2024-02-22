# alura-api-go-rest

## Run Back-end Locally

1. Start the PostgreSQL database:
```sh
$ docker-compose up
```

2. Run the server locally:
```sh
$ go run main.go
```

## Run Front-end Locally

We will integrate a front-end application developed in React to consume the API.

Check if Node.js is installed on your machine by running the following command in the terminal:
```sh
$ node --version
```

If your version appears on the screen, you can move on to the next step. If the Node.js version does not appear on the screen, please download Node.js on your machine, according to your operating system.

Now open a terminal in the project folder and type the following command to install the React dependencies:
```sh
$ npm install
```

Update the npm dependencies with the following command:
```sh
$ npm update
```

To finish, still in the terminal type the command below to start the React server and wait:
```sh
$ npm start
```

After loading, the project page will be displayed.

## API Documentation

This API documentation provides information on the endpoints available for this service.

All URLs referenced in the documentation have the following base: `https://localhost:8000`

## Endpoints

### `GET /api/personalidades`

Retrieves a list containing all the personalities recorded on database.

#### Parameters

None.

#### Example Request

```sh
$ curl --location 'localhost:8000/api/personalidades'
```

#### Example Response

```json
{
    "id": 2,
    "nome": "Carmela Dutra",
    "historia": "Carmela Teles Leite Dutra foi a primeira-dama do Brasil, de 31 de janeiro de 1946 até a sua morte, tendo sido a esposa de Eurico Gaspar Dutra, 16.º Presidente do Brasil. Era, carinhosamente, chamada de Dona Santinha, pela sua forte religiosidade, fazendo seu marido abrir uma capelinha no Palácio Guanabara."
}
```

---
### `GET /api/personalidades/:id`

Retrieves a personality recorded on database by its ID.

#### Parameters

`id`: **Required**. Id of the personality to fetch.

#### Example Request

```sh
$ curl --location 'localhost:8000/api/personalidades/1'
```

#### Example Response

```json
{
    "id": 1,
    "nome": "Deodato Petit Wertheimer",
    "historia": "Deodato Petit Wertheimer foi um médico e político brasileiro, seus primeiros anos de vida foram em São Paulo, mas logo mudou para Nova Friburgo no Estado do Rio de Janeiro e com 11 anos de idade ingressou no Colégio Anchieta dos jesuítas."
}
```

---
### `POST /api/personalidades`

Creates a new personality on database.

#### Parameters

`nome`: **Required**. Name of a personality.
`historia`: **Required**. A short description about the personality.

#### Example Request

```sh
$ curl --location 'localhost:8000/api/personalidades' \
--header 'Content-Type: application/json' \
--data '{
    "nome": "Alberto Santos Dumont",
    "historia": "Santos Dumont foi um aeronauta, esportista, autodidata e inventor brasileiro. Em 23 de outubro de 1906, ele voou cerca de sessenta metros a uma altura de dois a três metros com o Oiseau de Proie (francês para “ave de rapina”), no Campo de Bagatelle, em Paris. Em 12 de novembro de 1906, percorreu 220 metros a uma altura de seis metros com o Oiseau de Proie III"
}'
```

#### Example Response

```json
{
    "id": 3,
    "nome": "Alberto Santos Dumont",
    "historia": "Santos Dumont foi um aeronauta, esportista, autodidata e inventor brasileiro. Em 23 de outubro de 1906, ele voou cerca de sessenta metros a uma altura de dois a três metros com o Oiseau de Proie (francês para “ave de rapina”), no Campo de Bagatelle, em Paris. Em 12 de novembro de 1906, percorreu 220 metros a uma altura de seis metros com o Oiseau de Proie III"
}
```

---
### `DELETE /api/personalidades/:id`

Deletes a personality from database.

#### Parameters

`id`: **Required**. ID of a personality to delete.

#### Example Request

```sh
$ curl --location --request DELETE 'localhost:8000/api/personalidades/3'
```

#### Example Response

```json
{
    "id": 0,
    "nome": "",
    "historia": ""
}
```

---
### `PUT /api/personalidades/:id`

Update a personality data from database.

#### Parameters

`id`: **Required**. ID of a personality to update.

#### Example Request

```sh
$ curl --location --request PUT 'localhost:8000/api/personalidades/3' \
--header 'Content-Type: application/json' \
--data '{
    "historia": "Ele nasceu em 20 de julho de 1873 e faleceu em 23 de julho de 1932. Ele projetou, construiu e voou os primeiros balões dirigíveis com motor a gasolina. Ele também foi o primeiro a decolar a bordo de um avião impulsionado por um motor a gasolina"
}'
```

#### Example Response

```json
{
    "id": 3,
    "nome": "Alberto Santos Dumont",
    "historia": "Ele nasceu em 20 de julho de 1873 e faleceu em 23 de julho de 1932. Ele projetou, construiu e voou os primeiros balões dirigíveis com motor a gasolina. Ele também foi o primeiro a decolar a bordo de um avião impulsionado por um motor a gasolina"
}
```
