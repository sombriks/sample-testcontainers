/**
 * middleware that mimetizes a login filter.
 *
 * @param { import('koa').context } ctx Koa context
 * @param {Function} next handler to call next middleware
 * @returns {Promise<*>}
 */
export const fakeLoginCheck = async (ctx, next) => { // simple redirect if the cookie is there or not
  const userMaybe = ctx.cookies.get('x-user-info')
  if (userMaybe) return next(ctx)
  else return ctx.redirect('/login')
}
