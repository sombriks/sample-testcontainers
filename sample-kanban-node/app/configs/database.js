import Knex from 'knex';

/**
 * Provision our query builder instance
 *
 * @param dsn connection string
 *
 * @returns {Knex<any, unknown[]>} knex query builder instance
 */
export const prepareDatabase = (dsn = process.env.PG_CONNECTION_URL) => Knex({
	connection: dsn,
	client: 'pg',
});
