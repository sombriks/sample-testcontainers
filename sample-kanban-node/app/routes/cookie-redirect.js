/**
 * Middleware that mimetizes a login filter.
 *
 * @param { import('koa').context } ctx Koa context
 * @param {Function} next handler to call next middleware
 * @returns {Promise<*>}
 */
export const cookieCheck = async (context, next) => { // Simple redirect if the cookie is there or not
	const userMaybe = context.cookies.get('x-user-info');
	if (userMaybe) {
		return next(context);
	}
	return context.redirect('/login');
};

/**
 *
 * @param {service} options payload containing services to do business with
 *
 * @returns {function(*): Promise<*>} koa middleware
 */
export const fakeLogin = ({service}) => async context => {
	const {userId} = context.request.body;
	const user = await service.findPerson(userId);
	context.cookies.set('x-user-info', `name=${user.name}&id=${user.id}`);
	return context.redirect('/board');
};

// eslint-disable-next-line no-unused-vars
export const fakeLogout = ({service}) => async context => {
	context.cookies.set('x-user-info');
	return context.redirect('/login');
};
