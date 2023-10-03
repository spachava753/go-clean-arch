# REST API with Go

Some best practices involved when building REST API in Go. 

## Key Concepts
- **DDD**: This project covers the idea of Domain Driven Design, which splits and organizes code based on its responsibility to a specific domain 
- **Go Kit**: This project utilizes the go-kit architecture of constructing microservices. This architecture ensures that your business logic is isolated from microservice concerns like observability, retries, circuit breaking, and transport type.
- **12 factor app**: You'll learn how to create microservices that are agnostic to the way they are deployed, through creating configurations and containerizing the microservice
- **Generics**: As part of the newer releases of Go, generics have been added. This project utilizes generics for better code reuse

## Prerequisites
- Install [Docker](https://www.docker.com/get-started) and [Docker Compose](https://docs.docker.com/compose/install/)
- [Download and install Go](https://go.dev/doc/install)
- [Insomnia API Client](https://insomnia.rest/)
- Completion of the workshop part 1
- Knowledge about interfaces

## Directory Structure

Here's a brief explanation of the main directories and files in this project:

**Directories:**

- `internal/`: This is a specially recognized directory in Go. Any code defined inside internal is not importable outside the module. This is important when you are creating a library or app as it allows to control what others can use when importing your code. In our case, we don't want anyone using our code
- `internal/api`: Contains the code for API-first Open API based development
- `internal/api/gen`: Contains the Open API generated server code
- `internal/domain`: Contains code that implements DDD
- `internal/domain/stream`: Contains code that contains business logic and boilerplate for the Streams domain
- `internal/pkg`: Contains code that is helpful the app, but directly related to business logic

**Files:**

- `api.yaml`: The OpenAPI specification for the API.
- `internal/api/gen/gen.go`: Contains `go:generate` directive to generate `gen/api.gen.go`
- `go.mod` and `go.sum`: Used by Go's module system to manage dependencies.
- `main.go`: The entry point of the application.
- `tools.go`: Defines tools used by the application, so that their versions are managed by `go.mod`. 

## How to Run

### Generate code

```bash
go generate
```

This generates the code necessary for the application.

### Build the application

```bash
go build -o streams
```

* The `-o streams` option specifies the output file name for the executable.

### Run the application

```bash
./streams
```

The application will now be running and listening for incoming requests.

### Change configuration

Edit the `config.yaml` file and rerun the app to see how app behavior changes **without** rebuilding the app  

### Making Requests

Once your application is running, you may interact with your API by making HTTP requests to it. This repo contains an Insomnia Request Collection that can be imported into the [Insomnia API Client](https://insomnia.rest/). 

To begin, this repository only contains business logic for creating a resource. The GET and DELETE methods must be implemented as part of this workshop. In the meantime, requests made to these routes will return a 501 (Not Implemented) status code.


## TODO

Take the time to work through the following tasks. Each of them contain learning opportunities for various syntax, data structures, and patterns. Many of these are commonly used in Go projects.

- [ ] Write new middleware for timeouts
- [ ] Write new transport for Graphql using gqlgen (optional)
