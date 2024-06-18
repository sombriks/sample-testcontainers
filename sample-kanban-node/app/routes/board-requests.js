/**
 * returns the configured requests to deliver the board pages and components
 * @param {{service}} options config object containing reference to the service
 * @returns {*&{components: {}, pages: {login(*): Promise<*>, board(*): Promise<void>, table(*): Promise<*>}}}
 */
export const boardRequests = ({service}) => ({
  pages: {
    async board (ctx) {
      await ctx.render('pages/board')
    },
    async login (ctx) {
      const model = {
        users: await service.listUsers()
      }
      await ctx.render('pages/login', model)
    },
    async table (ctx) {
      await ctx.render('pages/table')
    }
  },
  components: {}
})
