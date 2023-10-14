# Go Project Organization

```text
├── LICENSE
├── README.md
├── config.go
├── go.mod
├── go.sum
├── clientlib
│   ├── lib.go
│   └── lib_test.go
├── cmd
│   ├── modlib-client
│   │   └── main.go
│   └── modlib-server
│       └── main.go
├── internal
│   └── auth
│       ├── auth.go
│       └── auth_test.go
└── serverlib
    └── lib.go
```

## Folders and files

### /cmd

This folder contains the main application entry point files

### /pkg

This folder contains code which is okay for other services to consume, this may include API clients, util functions

### /internal

This holds the private library code. It is specific to the function of the service and not shared with other services

### go.mod

It defines the module's path and it's dependencies. Similar to package.json

### go.sum

? Research required

## Packages

All go code is organised into packages and its simply consists of several .go files. A package is the entry point to access Go code.

- **Best Practice to use multiple files**

  - Feel free to separate you code into as many files as possible.
  - Aim to ensure it is easy to navigate your way around
  - Loosely couple sections of the service or application

- **Keep types close** <br/> It is often a good practice to keep the core types grouped at the top of a file.

- **Organise by responsibility** <br/>
  In other languages we organise types by models or types but in go we organise code by their functional responsibilities.

- **Start to Use Godoc** <br/>
  Godoc extracts and generates documentation for Go programs. It runs as a web server and presents the documentation as a web page.

- Dont't write your business logic to main.go
