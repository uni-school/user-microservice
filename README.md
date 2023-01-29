# Uni School: User Microservice

user microservice is a service that manage user of uni school.

## Installation

- Step 1: Check Go Version (Must go1.19+)

```bash
go version
# go version go1.19
```

- Step 2: Install all GRPC prerequisites on this [link](https://grpc.io/docs/languages/go/quickstart/#prerequisites)

- Step 3: Install Wire by running

```bash
go install github.com/google/wire/cmd/wire@latest
```

- Step 4: Install Depedencies by Running

```bash
go mod tidy
```

- Step 5: Copy `<dev|stag|prod|test>`.application.yaml.example to `<dev|stag|prod|test>`.application.yaml. `NOTE`: choose type dev

- Step 6: Running the grpc server

```bash
make start-dev
```

## NOTE

- Number 1: make sure you do all step in installation guide
- Number 2: if you specify port of this grpc server, the server should be listening on `<server-host>`:`<server-port>`
