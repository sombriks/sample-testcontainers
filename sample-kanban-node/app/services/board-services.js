/**
 * services for kanban
 *
 * @param {{db: import('knex') }} options payload with query builder
 *
 * @returns configured services
 */
export const boardServices = ({db}) => {
  return {
    async listUsers() {
      return db('kanban.person').orderBy("name")
    },
    async findUser(id) {
      return db('kanban.person').where({id}).first()
    },
    async listStatuses() {
      return db("kanban.status").orderBy("id")
    },
    async listTasks(q = "") {
      const tasks = await db("kanban.task")
        .whereILike("description", `%${q}%`)
        .orderBy("id")
      // TODO selects 1+N for related entities
      tasks.forEach(task => console.log(task))
      return tasks
    }
  }
}
