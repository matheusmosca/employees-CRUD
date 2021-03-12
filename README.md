# Employees CRUD

### Requirements
- All CRUD operations were implemented
- REST API built with Golang

### Technologies :computer:
- React JS
- SASS
- Material UI (Only used for inputs and date pickers)
- Golang
- PostgreSQL
- Docker compose
- Eslint

### Development setup :construction:
Requirements:
- node.js 14.15.4
- Golang 1.15.6
- Docker

#### Running The Backend
Change your directory to the `backend` folder and do the following commands:
- Run `docker-compose up`
- Run `go mod tidy` and `go run api/main.go`
- Run `go test ./...` to see some unit tests implemented

### Running the Frontend
Inside the `frontend` directory, use `yarn` to install all the dependencies and execute `yarn start`. Access http://localhost:3000 to see the application.
