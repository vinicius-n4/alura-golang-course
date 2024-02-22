# alura-api-go-gin

## Run locally

1. Start the PostgreSQL database:
```sh
$ docker-compose up
```

2. Run the server locally on port `:8080`:
```sh
$ go run main.go
```

## API Documentation

This API documentation provides information on the endpoints available for this service.

All URLs referenced in the documentation have the following base: `https://localhost:8080`

## Endpoints

### `GET /students`

Retrieves a list of all students recorded on database.

#### Parameters

None

#### Example Request

```sh
$ curl --location 'localhost:8080/students'
```

#### Example Response

```json
{
    "ID": 1,
    "CreatedAt": "2024-02-21T23:46:00.986449Z",
    "UpdatedAt": "2024-02-22T00:06:37.991631Z",
    "DeletedAt": null,
    "name": "Vinicius",
    "cpf": "12345678901",
    "rg": "123456789"
}
```

---
### `GET /:name`

Get a greeting for a given name.

#### Parameters

None

#### Example Request

```sh
$ curl --location 'localhost:8080/Vinicius'
```

#### Example Response

```json
{
    "API says": "Hello, Vinicius, how are you doing?"
}
```

---
### `POST /students`

Creates a register for a student.

### Parameters

name (required): The name of the student.
cpf (required): The CPF of the student.
rg (required): The RG of the student.

### Example Request

```sh
$ curl --location 'localhost:8080/students' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Vinicius",
    "cpf": "09876543210",
    "rg": "098765432"
}'
```

### Example Response

```json
{
    "ID": 1,
    "CreatedAt": "2024-02-22T01:08:02.325413659Z",
    "UpdatedAt": "2024-02-22T01:08:02.325413659Z",
    "DeletedAt": null,
    "name": "Vinicius",
    "cpf": "09876543210",
    "rg": "098765432"
}
```

---
### `GET /students/:id`

Retrieves a student by ID.

### Parameters

id (required): The database ID of the student.

### Example Request

```sh
$ curl --location 'localhost:8080/students/1'
```

### Example Response

```json
{
    "ID": 1,
    "CreatedAt": "2024-02-21T23:46:00.986449Z",
    "UpdatedAt": "2024-02-22T00:06:37.991631Z",
    "DeletedAt": null,
    "name": "Vinicius",
    "cpf": "12345678901",
    "rg": "123456789"
}
```

---
### `DELETE /students/:id`

Deletes a student's register by ID.

### Parameters

id (required): The database ID of the student.

### Example Request

```sh
$ curl --location --request DELETE 'localhost:8080/students/1'
```

### Example Response

```json
{
    "data": "Student successfuly deleted"
}
```

---
### `PATCH /students/:id`

Update a student's register by ID.

### Parameters

id (required): The database ID of the student.

### Example Request

```sh
$ curl --location --request PATCH 'localhost:8080/students/1' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Vinicius",
    "cpf": "12345678901",
    "rg": "123456789"
}'
```

### Example Response

```json
{
    "ID": 1,
    "CreatedAt": "2024-02-22T01:17:08.944001182Z",
    "UpdatedAt": "2024-02-22T01:17:08.944001182Z",
    "DeletedAt": null,
    "name": "Vinicius",
    "cpf": "12345678901",
    "rg": "123456789"
}
```

---
### `GET /students/cpf/:cpf`

Retrieves a student's register by CPF.

### Parameters

cpf (required): The CPF of the student recorded on database.

### Example Request

```sh
$ curl --location 'localhost:8080/students/cpf/12345678901'
```

### Example Response

```json
{
    "ID": 1,
    "CreatedAt": "2024-02-22T01:17:08.944001Z",
    "UpdatedAt": "2024-02-22T01:17:08.944001Z",
    "DeletedAt": null,
    "name": "Vinicius",
    "cpf": "12345678901",
    "rg": "123456789"
}
```

---
### `GET /index`

Retrieves an HTML with a list of students recorded on database.

### Parameters

None

### Example Request

```sh
$ curl --location 'localhost:8080/index'
```

### Example Response

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <link rel="stylesheet" href="../assets/index.css">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Students</title>
</head>
<body>
    <div class="widget-wrap">
        <h1>API - Gin e Golang</h1>
        <div class="widget-wrap">
            <ul id="the-list">
                <li>Vinicius</li>
                <li>Costa</li>
            </ul>
        </div>
</body>
</html>
```

### Preview Response

![image](https://github.com/vinicius-n4/alura-golang-course/assets/58186122/5d6cde7b-518d-4396-809e-35ad403c63d8)
