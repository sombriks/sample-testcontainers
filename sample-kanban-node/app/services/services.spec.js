import {resolve} from 'node:path';
import test from 'ava';
import {PostgreSqlContainer} from '@testcontainers/postgresql';
import {prepareDatabase} from '../configs/database.js';
import {boardServices} from './board-services.js';

test.before(async t => {
	// Testcontainer setup
	t.context.postgres = await new PostgreSqlContainer('postgres:16.3-alpine3.20')
		.withDatabase(process.env.PG_DATABASE)
		.withUsername(process.env.PG_USERNAME)
		.withPassword(process.env.PG_PASSWORD)
		.withBindMounts([{
			source: resolve(process.env.PG_INIT_SCRIPT),
			target: '/docker-entrypoint-initdb.d/init.sql',
		}])
		.start();

	// Application setup properly tailored for tests
	const database = prepareDatabase(t.context.postgres.getConnectionUri());
	const service = boardServices({db: database});

	// Context registering for proper teardown
	t.context.db = database;
	t.context.service = service;
});

test.after.always(async t => {
	await t.context.db.destroy();
	await t.context.postgres.stop({timeout: 500});
});

test('should list people', async t => {
	const people = await t.context.service.listUsers();
	t.is(people.length, 5);
});

test('should list tasks', async t => {
	const tasks = await t.context.service.listTasks();
	t.is(tasks.length, 5);
});
