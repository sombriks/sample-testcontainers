import { resolve } from 'node:path'
import { PostgreSqlContainer } from '@testcontainers/postgresql'
import request from 'supertest'
import test from 'ava'
import { prepareApp } from './main.js'
import { prepareDatabase } from './configs/database.js'
import { boardServices } from './services/board-services.js'
import { boardRoutes } from './routes/board-routes.js'
import { preparePostgres } from './configs/test-container.js'

test.before(async t => {
  // TestContainer setup
  t.context.postgres = await preparePostgres()

  // Application setup properly tailored for tests
  const database = prepareDatabase(t.context.postgres.getConnectionUri())
  const service = boardServices({ db: database })
  const controller = boardRoutes({ service })

  const { app } = prepareApp({ db: database, service, controller })

  // Context registering for proper teardown
  t.context.db = database
  t.context.app = app
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
  t.is(result, 2)
})

test('should serve login and have users', async t => {
  const result = await request(t.context.app.callback()).get('/login')
  t.is(result.status, 200)
  t.regex(result.text, /Alice|Bob|Caesar|Davide|Edward/)
})

test('should perform "login" (session cookie)', async t => {
  const result = await request(t.context.app.callback())
    .post('/login')
    .send({ userId: '1' }) // Alice
  t.is(result.status, 302)
  t.is(result.headers.location, '/board')
  t.truthy(result.header['set-cookie'])
  const [user] = result.header['set-cookie']
  t.regex(user, /Alice/)
})

test('should perform "logout" (remove session cookie)', async t => {
  const result = await request(t.context.app.callback())
    .get('/logout')
  t.is(result.status, 302)
  t.is(result.headers.location, '/login')
  t.truthy(result.header['set-cookie'])
  const [user] = result.header['set-cookie']
  t.regex(user, /x-user-info=;/)
})

test('should get board page', async t => {
  const result = await request(t.context.app.callback())
    .get('/board')
    .set('Cookie', 'x-user-info=name=Alice&id=1')
  t.is(result.status, 200)
  t.regex(result.text, /Welcome to the board, Alice/)
})

test('should get table page', async t => {
  const result = await request(t.context.app.callback())
    .get('/table')
    .set('Cookie', 'x-user-info=name=Alice&id=1')
  t.is(result.status, 200)
  t.regex(result.text, /Table for user Alice/)
})
