/**
 * services for kanban
 *
 * @param {{db: import('knex') }} options payload with query builder
 *
 * @returns configured services
 */
export const boardServices = ({ db }) => {
  // knex is a query builder, it's up to us to handle all 1+N selects
  async function loadTaskDeps (task) {
    task.status = await db('kanban.status').where({ id: task.status_id }).first()
    task.people = await db('kanban.person')
      .whereIn('id', db('kanban.task_person')
        .select('person_id').where({ task_id: task.id }))
    task.messages = await db('kanban.message').where({ task_id: task.id })
    for (const message of task.messages) {
      message.person = await db('kanban.person').where({ id: message.person_id }).first()
      message.task = task
    }
  }

  return {
    async listUsers () {
      return db('kanban.person').orderBy('name')
    },
    async findUser (id) {
      return db('kanban.person').where({ id }).first()
    },
    async listStatuses () {
      return db('kanban.status').orderBy('id')
    },
    async findStatus (id) {
      return db('kanban.status').where({ id }).first()
    },
    async findStatusByTaskId (id) {
      return db('kanban.status')
        .where('id',
          db('kanban.task')
            .select('status_id')
            .where({ id }).first())
        .first()
    },
    async listTasks (q = '') {
      const tasks = await db('kanban.task')
        .whereILike('description', `%${q}%`)
        .orderBy('id')
      for (const task of tasks) {
        await loadTaskDeps(task)
      }
      return tasks
    },
    async findTask (id) {
      const task = await db('kanban.task').where({ id }).first()
      await loadTaskDeps(task)
      return task
    },
    async addTask ({ status_id, description }) {
      return db('kanban.task')
        .insert({ description, status_id }).returning('*')
    },
    async updateTask ({ status_id, description, id }) {
      return db('kanban.task')
        .update({ description, status_id })
        .where({ id }).returning('*')
    },
    async deleteTask (id) {
      return db('kanban.task')
        .where({ id })
        .delete().returning('*')
    },
    async removePerson ({ id, personId }) {
      return db('kanban.task_person')
        .where({ person_id: personId, task_id: id })
        .delete().returning('*')
    },
    async joinTask ({ task_id, person_id }) {
      return db('kanban.task_person')
        .insert({ task_id, person_id })
        .returning('*')
    }
  }
}
