import {resolve} from 'node:path';
import {PostgreSqlContainer} from '@testcontainers/postgresql';

/**
 * Helper to provision a postgresql for testing purposes
 *
 * @returns {Promise<StartedPostgreSqlContainer>} database container
 */
export const preparePostgres = async () => new PostgreSqlContainer('postgres:16.3-alpine3.20')
	.withDatabase(process.env.PG_DATABASE)
	.withUsername(process.env.PG_USERNAME)
	.withPassword(process.env.PG_PASSWORD)
	.withBindMounts([{
		source: resolve(process.env.PG_INIT_SCRIPT),
		target: '/docker-entrypoint-initdb.d/init.sql',
	}])
	.start();
