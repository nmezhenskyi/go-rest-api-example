# Go REST API Example

This project showcases a minimalistic REST API written in Go using only the standard library. It serves as an overview of core Go concepts.

The program relies on a custom router, in-memory storage, and `net/http` package. The purpose of this API is to manage the resource *Wine*.

## Disclaimer

This example is not a production ready code. Many techniques used in the project have an introductory purpose. They aim to demonstrate fundamental concepts in a readable and easy-to-understand manner.
Namely, the custom router used by the application is inefficient compared to more complex solutions such as [httprouter](https://github.com/julienschmidt/httprouter). Likewise, in-memory hash table is not a persistent way to store data. However, this project serves as a good example on how to structure a Go application, manage dependencies, and unit tests.

## Functionality

The API supports 6 HTTP end-points:

- `GET "/"`
- `GET "/api/wine"`
- `GET "/api/wine/:id"`
- `POST "/api/wine"`
- `PUT "/api/wine/:id"`
- `DELETE "/api/wine/:id"`

The underlying resource is *Wine* with the following schema:

```json
{
   "id": {
      "type": "string",
      "example": "1",
   },
   "name": {
      "type": "string",
      "example": "Example Wine"
   },
   "category": {
      "type": "string",
      "example": "Some category"
   },
   "label": {
      "type": "string",
      "example": "Some label description"
   },
   "volume": {
      "type": "string",
      "example": "0.7"
   },
   "region": {
      "type": "string",
      "example": "Some region"
   },
   "producer": {
      "type": "string",
      "example": "Some producer"
   },
   "year": {
      "type": "number",
      "example": 2022
   },
   "alcohol": {
      "type": "string",
      "example": "11.5%"
   },
   "price": {
      "type": "string",
      "example": "12.50"
   }
}
```

## Project Structure

- `/cmd`
  - `/restapi/main.go`: entry point that begins the execution of the program
- `/internal`
  - `/router`: custom http router
  - `/storage`: wrapper for a hash table to use as in-memory storage
  - `/webserver`: http webserver implementation and route handlers
  - `/model`: contains domain model *Wine*
- `/data`
  - `/data_sample.json`: sample data for the webserver

## Commands

Compile:

- `go build -o ./bin/restapi ./cmd/restapi`

Start the server:

- `./bin/restapi` or `go run ./cmd/restapi`

Run tests:

- `go test ./internal/webserver`

Alternatively you can use Makefile

## Further Recommendations

- [Mat Ryer talks about best practices for writing HTTP Web Services in Go.](https://www.youtube.com/watch?v=rWBSMsLG8po&list=PLhp8mtvB6UEDCuvKT6ESWgJ0udJtWa09g&index=1&t=822s)
- [Detailed example of how to set up a file structure for a Go application.](https://github.com/golang-standards/project-layout)
- [Gregory Schier explains how to build a simple router in Go.](https://codesalad.dev/blog/how-to-build-a-go-router-from-scratch-3)
- [Fast and efficient router based on a Radix Tree.](https://github.com/julienschmidt/httprouter)

## Contribute

Feel free to send a pull request with new features or bugfix, and to open a new issue with recommendations or suggestions.

## License

This source code is licensed under MIT License.
