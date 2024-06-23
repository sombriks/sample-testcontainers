# sample-kanban-go

Application showcasing how to configure a postgresql database for testing
purposes using [TestContainers][testcontainers]

## Dependencies

- [go 1.22][go]
- [docker 25][docker]
- [echo 4.12][echo]
- [goqu 9.19][goqu] query builder
- [gomponents 0.20][gomponents]
- [bulma 1.0][bulma]
- [htmx.org 2.0][htmx]
- [alpine.js 3.14][alpinejs]
- [testify 1.9][testify]
- [godotenv 1.5][godotenv] to simplify configuration
- and of course [testcontainers for go][go-testcontainers]

Make sure you have $HOME/go/bin in your $PATH.

## How to build

```bash
go build .
```

## How to test

```bash
go test -cover -v ./...
```

## How to run

Make sure that you have a valid and running postgresql accepting the proper
[credentials configured in the .env file][env]. If you don't want to configure a
postgresql yourself, check the [docker compose file][compose] in the
infrastructure folder.

If you build the project, run the binary:

```bash
./sample-kanban-go
```

But you simply can run the entrypoint:

```bash
go run main.go
```

## Noteworthy

- The project structure is [not idiomatic to go][go-project-structure] but
  follows [SOLID][solid] standards as much as possible. It's a deliberate choice
  because go documented recommendation is a bad design and can lead quickly to
  mental fatigue.
- Thanks to [godotenv][godotenv], we have an environment-aware configuration
  strategy almost as flexible as the jvm version.
- GO packages are very different from classpath or node ESM/CJS. For instance,
  they are their own point of execution, so the project root is not what you
  expect it to be.
- Like the node version, we're not using any [DI container][di] in this sample.
  Because of that, mind you configuration phase, write testable code, provide as
  much inversion of control points as possible while providing reasonable
  default values.
- [goqu][goqu] query builder resembles [knex][knex] query builder a little and, 
  just like we did on node project, it's up to us solve the select [N+1][n+1]
  cases, which affects mainly the `task` and `message` models.
- Unlike other two projects, the template language isn't a markup language but
  golang itself. The [gomponents][gomponents] library isn't exactly a new idea
  but looks funny see how it goes in this exercise.
- Another discrepancy of this implementation is the use of static resources to
  serve unversioned frontend libraries. In jvm version there where
  [webjars][webjars], in node version we served versioned libraries directly
  from [node_modules][node_modules, nothing similar is available for go projects
  by the time of this writing.
- Note that `//go:embed static` and `// go:embed static` are not the same thing.
- There is a tool called [air][air] that delivers similar experience on go
  projects that [nodemon][nodemon] delivers on node projects. It's completely
  optional and independent of project dependencies but it worth the
  configuration effort.

[testcontainers]: https://testcontainers.com/
[go]: https://go.dev/
[docker]: https://docs.docker.com/engine/install
[echo]: https://echo.labstack.com/
[goqu]: https://github.com/doug-martin/goqu
[gomponents]: https://www.gomponents.com/
[bulma]: https://bulma.io/documentation/
[htmx]: https://htmx.org/docs/#introduction
[alpinejs]: https://alpinejs.dev
[testify]: https://github.com/stretchr/testify
[godotenv]: https://github.com/joho/godotenv
[air]: https://github.com/air-verse/air
[go-testcontainers]: https://golang.testcontainers.org/modules/postgres/
[go-project-structure]: https://go.dev/doc/modules/layout
[solid]: https://en.wikipedia.org/wiki/SOLID
[env]: ./.env
[compose]: ../infrastructure/docker-compose.yml
[di]: https://martinfowler.com/articles/injection.html
[knex]: https://knexjs.org
[n+1]: https://stackoverflow.com/a/39696775/420096
[webjars]: https://www.webjars.org/
[node_modules]: https://docs.npmjs.com/cli/v7/configuring-npm/folders#node-modules
[nodemon]: https://nodemon.io/
