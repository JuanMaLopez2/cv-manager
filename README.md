Login endpoint
# Dependences
In order to run this project you need to execute the commands to obtain the required dependences.
```
go get github.com/joho/godotenv
go get gorm.io/gorm
go get github.com/felixge/httpsnoop
go get github.com/gorilla/handlers
go get github.com/gorilla/mux
go get github.com/jackc/pgpassfile
go get github.com/jackc/pgservicefile
go get github.com/jackc/pgx/v5
go get github.com/jackc/puddle/v2
go get github.com/jinzhu/inflection
go get github.com/jinzhu/now
go get github.com/lib/pq
go get golang.org/x/crypto
go get golang.org/x/sync
go get golang.org/x/text
go get gorm.io/driver/postgres
```
If there are some error with dependences, please execute.
```
go mod tidy 
```

# To run the endpoint.
Before running the server make sure the database is working.

Then, change the .env_example file name to .env and introduce the credentials.

Then, execute the comand below.
```
go run cmd\main.go
```
The server listens on port 8080.

If you want to send some POST, you can execute in POSTMAN.
```
http://localhost:8080/login
```
in the body of the POST, use the following format:
```
{
  "email": "email",
  "password": "password"
}
```
