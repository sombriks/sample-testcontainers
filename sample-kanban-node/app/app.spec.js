import { resolve } from 'node:path'
import { PostgreSqlContainer } from '@testcontainers/postgresql'
import request from 'supertest'
import test from 'ava'
import { prepareApp } from './main.js'
import { prepareDatabase } from './configs/database.js'
import { boardServices } from './services/board-services.js'
import { boardRoutes } from './routes/board-routes.js'

test.before(async t => {
  // testcontainer setup
  t.context.postgres = await new PostgreSqlContainer('postgres:16.3-alpine3.20')
    .withDatabase(process.env.PG_DATABASE)
    .withUsername(process.env.PG_USERNAME)
    .withPassword(process.env.PG_PASSWORD)
    .withBindMounts([{
      source: resolve(process.env.PG_INIT_SCRIPT),
      target: '/docker-entrypoint-initdb.d/init.sql',
    }])
    .start()

  // application setup properly tailored for tests
  const db = prepareDatabase(t.context.postgres.getConnectionUri())
  const service = boardServices({ db })
  const controller = boardRoutes({ service })
  const { app } = prepareApp({ db, service, controller })

  // context registering for proper teardown
  t.context.app = app
  t.context.db = db
})

test.after.always(async t => {
  await t.context.db.destroy()
  await t.context.postgres.stop({ timeout: 500 })
})

test('app should be ok', async t => {
  const result = await request(t.context.app.callback()).get('/')
  t.is(result.status, 302)
  t.is(result.headers.location, '/board')
})

test('db should be ok', async t => {
  const { rows: [{ result }] } = await t.context.db.raw('SELECT 1 + 1 as result')
  t.truthy(result)
  t.is(2, result)
})

test('should serve login and have users', async t => {
  const result = await request(t.context.app.callback()).get('/login')
  t.is(result.status, 200)
  t.regex(result.text, /Alice|Bob/)
})

