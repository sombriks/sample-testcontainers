/**
 * returns the configured requests to deliver the board pages and components
 * @param {{service}} options config object containing reference to the service
 * @returns {*&{components: {}, pages: {login(*): Promise<*>, board(*): Promise<void>, table(*): Promise<*>}}}
 */
export const boardRequests = (options) => ({
  ...options,
  pages: {
    async board (ctx) {
      await ctx.render('pages/board')
    },
    async login (ctx) {
      await ctx.render('pages/login')
    },
    async table (ctx) {
      await ctx.render('pages/table')
    }
  },
  components: {}
})
