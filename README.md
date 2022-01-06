The project `go-shopping` is a playground to practice writing an HTTP service in `GoLang` using `gin` and `gorm`.

The project represents a back-end microservice that serves HTTP calls to do basic CRUD (create/read/update/delete) operations
on the items in a *small magic shop* :)

The project was inspired by
[this article about Gin&Gorm](https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/)
and [this article about application layers](https://deliveroo.engineering/2019/05/17/testing-go-services-using-interfaces.html).
The implementation follows [official documentation](https://github.com/gin-gonic/gin)
and [this article](https://blog.hackajob.co/writing-and-testing-crud/).

Following [this article](https://deliveroo.engineering/2019/05/17/testing-go-services-using-interfaces.html),
the application is separated into 3 layers:
- _repositories_: interact directly with the database;
- _services_: contain the business logic;
- _handlers_: accept requests and builds the responses.

Run all tests in the project, but remove the db file first (workaround):
```
rm -f routers/items.db
go test ./...
```

If all tests pass, start the HTTP service with:
```
go run main.go
```
and the service will run on `localhost:8080`.

Additionally, in the root of the repository there are following files:
- `Go-Shopping.postman_collection.json` - a Postman collection that can be used for reference to send HTTP calls;
- `items.db` - a small SQLite Database to start using the project without the need to care about any pre-requisites;
- `items.sql` - a backup file containing SQL queries to create an SQLite Database if it does not exist.
