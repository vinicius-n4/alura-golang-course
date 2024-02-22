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
$ curl
```

#### Example Response

```json
{
  "id": "1",
  "name": "Vinicius"
}
```

---
### `GET /api/items/${id}`

A simple description about what the API method does.

#### Parameters

`id`: **Required**. Id of item to fetch

#### Example Request

```sh
$ curl
```

#### Example Response

```json
{
  "id": "1",
  "name": "Vinicius"
}
```