import {resolve} from 'node:path';
import test from 'ava';
import {PostgreSqlContainer} from '@testcontainers/postgresql';
import {prepareDatabase} from '../configs/database.js';
import {boardServices} from './board-services.js';
import { preparePostgres } from '../configs/test-container.js'

test.before(async t => {
	// TestContainer setup
	t.context.postgres = await preparePostgres()

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
