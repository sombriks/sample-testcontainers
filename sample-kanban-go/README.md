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
- [air 1.52][air] for live reload
- and of course [testcontainers for go][go-testcontainers]

Make sure you have $HOME/go/bin in your $PATH.

## How to build

```bash
go build .
```

## How to test

## How to run

## Noteworthy

- The project structure is [not idiomatic to go][go-project-structure] but
  follows [SOLID][solid] standards as much as possible. It's a deliberate choice
  because go documented recommendation is a bad design and can lead quickly to
  mental fatigue.
- 

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
