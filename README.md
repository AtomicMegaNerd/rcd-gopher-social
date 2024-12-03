# RCD Gopher Social

This is a Go app that simulates a social network. This is for a backend engineering with Go
course on Udemy.

[https://udemy.com/course/backend-engineering-with-go/](https://udemy.com/course/backend-engineering-with-go/)

## Libraries

The following interesting Go libraries are used in this app:

- [Chi](https://github.com/go-chi/chi) a router for building Go services.
- [PQ](https://github.com/lib/pq) a PostgreSQL driver for Go.

## Tools

### Air

[https://github.com/air-verse/air](https://github.com/air-verse/air)

This project uses `air` to automatically reload the server when changes are made. To install it, run the following command:

```bash
go install github.com/air-verse/air@latest
```

#### Configuration

Checkout `.air.toml` for the configuration. The `air` command will look for this file in the root of the project.

Assuming you have `$GOPATH/bin` in your `$PATH`, you can run the following command to start the server:

```bash
air
```

Each time you save a file, the server will automatically reload.

### Taskfile

[https://taskfile.dev](https://taskfile.dev)

Think of this as Go's (much more modern) version of make. See [./Taskfile.yml](./Taskfile.yml) for the available commands.

To install it, run the following command:

```bash
brew install go-task/tap/go-task
```

To build:

```bash
task build
```

To run the tests:

```bash
task test
```

### Direnv

[https://github.com/direnv/direnv](https://github.com/direnv/direnv)

This project uses `direnv` to manage environment variables. To use it, create a `.envrc` file in the root of the project with the following content:

```bash
export ADDR=":3000"
export DB_ADDR="postgres://admin:adminpassword@localhost/social?sslmode=disable"
```

Then, run the following command to allow the `.envrc` file:

```bash
direnv allow
```

### Migrate

[https://github.com/golang-migrate/migrate](https://github.com/golang-migrate/migrate)

This project uses `migrate` to manage database migrations. To install it, run the following command:

```bash
brew install golang-migrate
```

The [Taskfile.yml](Taskfile.yml) file has some commands to help with migrations. The Taskfile
also contains the commands for reference.

Create migrations (found in ./cmd/migrate/migrations):

```bash
task migrate-create-users
task migrate-create-posts
```

Run migration to upgrade:

```bash
task migrate-up
```

Run migration to downgrade:

```bash
task migrate-down
```

### Rainfrog

[https://github.com/achristmascarl/rainfrog](https://github.com/achristmascarl/rainfrog)

Right now Rust has to be installed first. Also `$HOME/.cargo/bin` has to be in the PATH.

```bash
brew install rustup
rustup-init
```

Then install rainfrog:

```bash
cargo install rainfrog
```

To run Rainfrog, use the following command:

```bash
rainfrom --url $DB_ADDR
```
