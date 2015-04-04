[![wercker status](https://app.wercker.com/status/e26eee26a6a67411aed47b3006e04877/m "wercker status")](https://app.wercker.com/project/bykey/e26eee26a6a67411aed47b3006e04877)

# apiGO

This is a Mock Web Service for iTEC 2015

This API will be build using the Go programming language.

# Configuration steps for the API

1. Install Go and set up your [GOPATH](http://golang.org/doc/code.html#GOPATH)

2. Install [MongoDb](https://scotch.io/tutorials/an-introduction-to-mongodb#installation-and-running-mongodb)

3. Create a database named __apiGO_db__ and then create an user for the database using the following command in **mongodb shell**:
>###`db.createUser( { user: "apiGOAdmin", pwd: "apiGOPass", roles: [ { role: "readWrite", db: "apiGO_db" } ] } )`

4. Install all the project dependencies by navigating to the project directory _($GOPATH/src/apiGO)_ and using the following command in the **cmd/console/terminal**:
>###`go get -t ./...`

5. For testing purposes, create another database named __apiGO_db_test__, but don't create a user for it like for the main database.
In order for the tests to run, you need to set the following environment variables correctly:

> **APIGO_TESTDB_NAME** = apiGO_db_test

> **APIGO_TESTDB_CONN** = mongodb://localhost:27017/apiGO_db_test

If you don't want to use the terminal for creating the databases, you can always use an useful app with GUI, such as **Robomongo**