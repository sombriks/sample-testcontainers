import Knex from 'knex';

/**
 * Provision our query builder instance
 *
 * @param connection database connection string
 *
 * @returns {Knex} knex query builder instance
 */
export const prepareDatabase = (connection = process.env.PG_CONNECTION_URL) => Knex({
	client: 'pg',
	connection,
});
