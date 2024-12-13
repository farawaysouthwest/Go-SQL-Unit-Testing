## SQL Testing in a Go Application


### Introduction
This is a simple Go application that connects to a MySQL database and performs some basic operations. 
The purpose of this application is to demonstrate how to write tests for a Go application that interacts with a database. 
The application uses the `gorm` package to interact with the MySQL database.

### creating mocks for your own interfaces

The excellent `gomock` package can be used to generate mocks for your own interfaces.
`go get go.uber.org/mock/gomock`

Install the mockgen tool with the following command:
`go install go.uber.org/mock/mockgen@latest`

Then you can generate a mock for your own interface with the following command:
`mockgen -source=path/to/your/interface.go -destination=path/to/your/mock.go`