/**
 * services for kanban
 *
 * @param {{db: import('knex') }} options payload with query builder
 *
 * @returns configured services
 */
export const boardServices = ({ db }) => {
  return {
    async listUsers () {
      return db('kanban.person')
    },
    async findUser (id) {
      return db('kanban.person').where({ id }).first()
    }
  }
}
