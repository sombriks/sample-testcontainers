/**
 * returns the configured requests to deliver the board pages and components
 * @param {{service}} options config object containing reference to the service
 * @returns {*&{components: {}, pages: {login(*): Promise<*>, board(*): Promise<void>, table(*): Promise<*>}}}
 */
export const boardRequests = ({ service }) => ({
  pages: {
    async board (ctx) {
      const model = {
        user: parseUser(ctx.cookies.get('x-user-info')),
        statuses: await service.listStatuses(),
        tasks: await service.listTasks()
      }
      await ctx.render('pages/board', model)
    },
    async login (ctx) {
      const model = {
        users: await service.listUsers()
      }
      await ctx.render('pages/login', model)
    },
    async table (ctx) {
      const model = {
        user: parseUser(ctx.cookies.get('x-user-info')),
        statuses: await service.listStatuses(),
        tasks: await service.listTasks()
      }
      await ctx.render('pages/table', model)
    }
  },
  components: {
    async addTask (ctx) {
      const { description, status } = ctx.request.body
      await service.addTask({ description, status_id: status })
      const model = {
        user: parseUser(ctx.cookies.get('x-user-info')),
        status: await service.findStatus(status),
        tasks: await service.listTasks()
      }
      await ctx.render('components/category-lanes', model)
    },
    async updateTask (ctx) {
      const { id } = ctx.request.params
      const { description, status } = ctx.request.body
      await service.updateTask({ id, description, status_id: status })
      const model = {
        user: parseUser(ctx.cookies.get('x-user-info')),
        status: await service.findStatus(status),
        task: await service.findTask(id)
      }
      await ctx.render('components/task-card', model)
    },
    async deleteTask (ctx) {
      const { id } = ctx.request.params
      const status = await service.findStatusByTaskId(id)
      await service.deleteTask(id)
      const model = {
        user: parseUser(ctx.cookies.get('x-user-info')),
        tasks: await service.listTasks(),
        status
      }
      await ctx.render('components/category-lanes', model)
    },
    async removePerson (ctx) {
      const { id, personId } = ctx.request.params
      await service.removePerson({ id, personId })
      const task = await service.findTask(id)
      const user = parseUser(ctx.cookies.get('x-user-info'))
      await ctx.render('components/task-members', { user, task })
    },
    async joinTask (ctx) {
      const { id } = ctx.request.params
      const user = parseUser(ctx.cookies.get('x-user-info'))
      await service.joinTask({ task_id: id, person_id: user.id })
      const task = await service.findTask(id)
      await ctx.render('components/task-members', { user, task })
    }
  }
})

/**
 * helper to crack open user
 *
 * @param cookie user encoded
 * @returns {{[p: string]: string} | null}
 */
const parseUser = cookie => {
  if (!cookie) return null
  const [kId, kName] = cookie.split('&')
  const kvId = kId.split('=')
  const kvName = kName.split('=')
  return { [kvId[0]]: kvId[1], [kvName[0]]: kvName[1] }
}
