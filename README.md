The project `go-shopping` is a playground to practice writing an HTTP service in `GoLang` using `gin` and `gorm`.

The project represents a back-end microservice that serves HTTP calls to do basic CRUD (create/read/update/delete) operations
on the items in a *small magic shop* :)

The implementation of the project follows [this article](https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/).

Start the HTTP service with:
```
$ go run main.go
```
and the service will run on `localhost:8080`.

In the root of the repository there are following files:
- `Go-Shopping.postman_collection.json` - a Postman collection that can be used for reference to send HTTP calls;
- `items.db` - a small SQLite Database to start using the project without the need to care about any pre-requisites;
- `items.sql` - a backup file containing SQL queries to create an SQLite Database if it does not exist.
