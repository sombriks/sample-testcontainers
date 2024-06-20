import {resolve} from 'node:path';
import {PostgreSqlContainer} from '@testcontainers/postgresql';
import request from 'supertest';
import test from 'ava';
import {prepareApp} from './main.js';
import {prepareDatabase} from './configs/database.js';

test.before(async t => {
	t.context.postgres = await new PostgreSqlContainer('postgres:16.3-alpine3.20')
		.withDatabase(process.env.PG_DATABASE)
		.withUsername(process.env.PG_USERNAME)
		.withPassword(process.env.PG_PASSWORD)
		.withBindMounts([{
			source: resolve('../sample-kanban-jvm/src/test/resources/initial-state.sql'),
			target: '/docker-entrypoint-initdb.d/init.sql',
		}])
		.start();
	t.context.db = prepareDatabase(t.context.postgres.getConnectionUri());
	const {app} = prepareApp({db: t.context.db});
	t.context.app = app;
});

test.after.always(async t => {
	await t.context.db.destroy();
	await t.context.postgres.stop({timeout: 500});
});

test('app should be ok', async t => {
	const result = await request(t.context.app.callback()).get('/');
	t.is(result.status, 302);
	t.is(result.headers.location, '/board');
});

test('db should be ok', async t => {
	const result = await t.context.db.raw('SELECT 1 + 1');
	console.log(result)
	t.truthy(result)
})

test('should serve login and have users', async t => {
	const result = await request(t.context.app.callback()).get('/login');
	t.is(result.status, 200);
	t.regex(result.text, /Alice|Bob/);
});

