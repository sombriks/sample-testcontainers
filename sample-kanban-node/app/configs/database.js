import Knex from 'knex'

export const db = Knex({
  connection: process.env.PG_CONNECTION_URL,
  client: 'pg'
})
