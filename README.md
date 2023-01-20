# Uni School: User Microservice

user microservice is a service that manage user of uni school.

## Installation

- Step 1: Install all GRPC prerequisites on this [link](https://grpc.io/docs/languages/go/quickstart/#prerequisites) and Install Wire by running

```bash
go install github.com/google/wire/cmd/wire@latest
```

- Step 2: Run

```bash
go mod tidy
```

- Step 3: Copy `<dev|stag|prod|test>`.application.yaml.example to `<dev|stag|prod|test>`.application.yaml. `NOTE`: choose type dev

- Step 4: Run

```bash
make start-dev
```

## NOTE

- Number 1: make sure you do all step in installation guide
- Number 2: make sure in `<env>`.application.yaml variable passwordHashing.saltHash must have same value in file `<env>`.application.yaml on api-gateway
- Number 3: if you specify port of this grpc server, the server should be listening on `<server-host>`:`<server-port>`
