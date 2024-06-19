import Knex from 'knex'

export const prepareDb = (dsn = process.env.PG_CONNECTION_URL) => Knex({
  connection: dsn,
  client: 'pg'
})
