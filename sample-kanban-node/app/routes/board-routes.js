/**
 * Returns the configured requests to deliver the board pages and components
 * @param {{service}} options config object containing reference to the service
 * @returns {*&{components: {}, pages: {login(*): Promise<*>, board(*): Promise<void>, table(*): Promise<*>}}}
 */
export const boardRoutes = ({service}) => ({
	pages: {
		async board(context) {
			const model = {
				user: parseUser(context.cookies.get('x-user-info')),
				statuses: await service.listStatuses(),
				tasks: await service.listTasks(),
			};
			await context.render('pages/board', model);
		},
		async login(context) {
			const model = {
				users: await service.listUsers(),
			};
			await context.render('pages/login', model);
		},
		async table(context) {
			const model = {
				user: parseUser(context.cookies.get('x-user-info')),
				statuses: await service.listStatuses(),
				tasks: await service.listTasks(),
			};
			await context.render('pages/table', model);
		},
	},
	components: {
		async addTask(context) {
			const {description, status} = context.request.body;
			await service.addTask({description, statusId: status});
			const model = {
				user: parseUser(context.cookies.get('x-user-info')),
				status: await service.findStatus(status),
				tasks: await service.listTasks(),
			};
			await context.render('components/category-lanes', model);
		},
		async updateTask(context) {
			const {id} = context.request.params;
			const {description, status} = context.request.body;
			await service.updateTask({id, description, statusId: status});
			const model = {
				user: parseUser(context.cookies.get('x-user-info')),
				status: await service.findStatus(status),
				task: await service.findTask(id),
			};
			await context.render('components/task-card', model);
		},
		async deleteTask(context) {
			const {id} = context.request.params;
			const status = await service.findStatusByTaskId(id);
			await service.deleteTask(id);
			const model = {
				user: parseUser(context.cookies.get('x-user-info')),
				tasks: await service.listTasks(),
				status,
			};
			await context.render('components/category-lanes', model);
		},
		async removePerson(context) {
			const {id, personId} = context.request.params;
			await service.removePerson({id, personId});
			const task = await service.findTask(id);
			const user = parseUser(context.cookies.get('x-user-info'));
			await context.render('components/task-members', {user, task});
		},
		async joinTask(context) {
			const {id} = context.request.params;
			const user = parseUser(context.cookies.get('x-user-info'));
			await service.joinTask({taskId: id, personId: user.id});
			const task = await service.findTask(id);
			await context.render('components/task-members', {user, task});
		},
		async addComment(context) {
			const {id} = context.request.params;
			const {content} = context.request.body;
			const user = parseUser(context.cookies.get('x-user-info'));
			await service.addComment({taskId: id, personId: user.id, content});
			const task = await service.findTask(id);
			await context.render('components/task-comments', {user, task});
		},
	},
});

/**
 * Helper to crack open user
 *
 * @param cookie user encoded
 * @returns {{[p: string]: string} | null}
 */
const parseUser = cookie => {
	if (!cookie) {
		return null;
	}

	const [kId, kName] = cookie.split('&');
	const kvId = kId.split('=');
	const kvName = kName.split('=');
	return {[kvId[0]]: kvId[1], [kvName[0]]: kvName[1]};
};
