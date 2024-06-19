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

/**
 *
 * @param {service} options payload containing services to do business with
 *
 * @returns {function(*): Promise<*>} koa middleware
 */
export const fakeLogin = ({ service }) => async ctx => {
  const { userId } = ctx.request.body
  const user = await service.findUser(userId)
  ctx.cookies.set('x-user-info', `name=${user.name}&id=${user.id}`)
  return ctx.redirect('/board')
}

export const fakeLogout = ({ service }) => async ctx => {
  ctx.cookies.set('x-user-info')
  return ctx.redirect('/login')
}
