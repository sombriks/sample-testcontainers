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
- [htmx.org 2][htmx]
- [alpine.js 3][alpinejs]
- [xo 0.58][xo]
- [ava 6][ava]
- [supertest 7][supertest]
- and of course [@testcontainers/postgresql 10][postgres-testcontainers]

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

- Small but important reminder, this code is quite experimental and aims to
  showcase testcontainers, not the best Node.js practices. You where warned!
- Unlike _modern_ frontend development, we don't have a build step for code that
  runs on client-side. We're avoiding, however, download pure frontend artifacts
  and serve them as assets, because we might lose track of our dependencies,
  managed by npm. In spring version, there was [webjars][webjars], so we did a
  similar approach using [koa-static][koa-static] and [koa-mount][koa-mount].
- [Koa][koa]'s simplicity always amazes me. This is what express was supposed to
  be. Modular. Extensible. Easy to test.
- Thanks to [xo][xo] we have a simple way to have some code quality _standards_.
  Direct eslint configuration is quite unbearable and [other linter][standard]
  alternatives seems too strict.
- TestContainer setup is quite similar to what we saw on jvm version; it has,
  however, a distinct idiom to provide the init.sql script needed to create a
  known state in the database.
- Always remember that everything that executes in a node project has the point
  of execution, entrypoint, you name it, in the same folder as the package.json,
  the project root in other words. This is why in the `.withBindMounts` call the
  path to sql script in the other project jumps just one folder up and not two.
  This is why `koa-static` doesn't need to jump up no folder when mapping the
  frontend libs in _node_modules_.

[testcontainers]: https://testcontainers.com
[node]: https://nodejs.org
[docker]: https://docs.docker.com/engine/install
[koa]: https://koajs.com
[knex]: https://knexjs.org
[pug]: https://pugjs.org/api/getting-started.html
[bulma]: https://bulma.io/documentation/
[htmx]: https://htmx.org/docs/#introduction
[alpinejs]: https://alpinejs.dev
[xo]: https://github.com/xojs/xo
[ava]: https://avajs.dev/
[supertest]: https://github.com/ladjs/supertest
[standard]: https://standardjs.com/
[postgres-testcontainers]: https://testcontainers.com/modules/postgresql/
[package.json]: ./package.json
[webjars]: https://www.webjars.org/
[koa-static]: https://www.npmjs.com/package/koa-static
[koa-mount]: https://www.npmjs.com/package/koa-mount
[compose]: ../infrastructure/docker-compose.yml
