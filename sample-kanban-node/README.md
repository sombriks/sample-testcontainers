# sample-kanban-node

Application showcasing how to configure a postgresql database for testing
purposes using [TestContainers][testcontainers]

## Dependencies

- [node 20][node]
- [docker 25][docker]
- [koa 2][koa]
- [knex 3][knex]
- [pug 3][pug]
- [bulma 1][bulma]
- [htmx.org 1.19][htmx]
- [alpine.js 3][alpinejs]
- [ava 6][ava]
- [standard 17][standard]
- and of course [@testcontainers/postgresql 10][pgtc]

There are a few others, check the [package.json][package.json] for details.

## How to build

Download npm dependencies:

```bash
npm i
```

## How to test

```bash
npm run test
```

## How to run

Make sure you have a running postgresql available. There is a convenient
[docker compose][compose] file in the infrastructure folder that will provision
the sample kanban database for you.

```bash
npm run dev
```

## Noteworthy

- Unlike _modern_ frontend development, we don't have a build step for code that
  runs on client-side. We're avoiding, however, download pure frontend artifacts
  and serve them as assets, because we might lose track or our dependencies,
  managed by npm. In spring version, there was [webjars][webjars], so we did a
  similar approach using [koa-static][koa-static] and [koa-mount][koa-mount].
- [Koa][koa]'s simplicity always amazes me. This is what express was supposed to
  be. Modular. Extensible. Easy to test.
- Thanks to [standard][standard] we have a simple way to have some code quality
  _standards_.

[testcontainers]: https://testcontainers.com
[node]: https://nodejs.org
[docker]: https://docs.docker.com/engine/install
[koa]: https://koajs.com
[knex]: https://knexjs.org
[pug]: https://pugjs.org/api/getting-started.html
[bulma]: https://bulma.io/documentation/
[htmx]: https://htmx.org/docs/#introduction
[alpinejs]: https://alpinejs.dev
[ava]: https://avajs.dev/
[standard]: https://standardjs.com/
[pgtc]: https://testcontainers.com/modules/postgresql/
[package.json]: ./package.json
[webjars]: https://www.webjars.org/
[koa-static]: https://www.npmjs.com/package/koa-static
[koa-mount]: https://www.npmjs.com/package/koa-mount
[compose]: ../infrastructure/docker-compose.yml
