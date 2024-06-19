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
        tasks: await service.listTasks(),
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
        tasks: await service.listTasks(),
      }
      await ctx.render('pages/table', model)
    }
  },
  components: {}
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
