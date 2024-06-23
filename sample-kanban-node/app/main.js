import ApiBuilder from 'koa-api-builder';
import {bodyParser} from '@koa/bodyparser';
import Koa from 'koa';
import mount from 'koa-mount';
import Router from '@koa/router';
import serve from 'koa-static';
import {cabin} from './configs/logging.js';
import {prepareDatabase} from './configs/database.js';
import {pug} from './configs/views.js';
import {boardRoutes} from './routes/board-routes.js';
import {boardServices} from './services/board-services.js';
import {fakeLogin, cookieCheck, fakeLogout} from './routes/cookie-redirect.js';

// Default components
const _database = prepareDatabase();
const _service = boardServices({db: _database});
const _controller = boardRoutes({service: _service});

/**
 * Function to return a properly configured Koa application
 *
 * @param {import('knex')} db query builder reference
 * @param service our service module to do business with
 * @param controller our routes/requests/controllers to validate and handle things
 *
 * @returns {{app: Application, db: import('knex')}}  app and db with all the goodies we will need
 */
export const prepareApp = ({
	db: database = _database,
	service = _service,
	controller = _controller,
} = {}) => {
	const app = new Koa();

	// Dynamic api planning
	const router = new ApiBuilder({router: new Router()})
		.path(b => {
			b.get('/', async context => context.redirect('/board'));
			b.get('/board', cookieCheck, controller.pages.board);
			b.path('/login', b => {
				b.get(controller.pages.login);
				b.post(fakeLogin({service}));
			});
			b.get('/logout', fakeLogout({service}));
			b.get('/table', cookieCheck, controller.pages.table);
			b.path('/task', cookieCheck, b => {
				b.post(controller.components.addTask);
				b.path('/:id', b => {
					b.put(controller.components.updateTask);
					b.del(controller.components.deleteTask);
					b.del('/person/:personId', controller.components.removePerson);
					b.post('/join', controller.components.joinTask);
					b.post('/comments', controller.components.addComment);
				});
			});
		}).build();

	// Static assets
	const alpinejs = serve('node_modules/alpinejs/dist');
	const bulma = serve('node_modules/bulma/css');
	const htmx = serve('node_modules/htmx.org/dist');
	const hxDsIncrement = serve('node_modules/hx-dataset-include/lib');
	const ionicons = serve('node_modules/ionicons/dist/');
	const _static = serve('app/static');

	// Configuring the koa app. the middleware order matters a lot
	app.use(cabin.middleware);
	pug.use(app);
	app.use(bodyParser());
	app.use(router.routes());
	app.use(router.allowedMethods());
	app.use(mount('/', _static));
	app.use(mount('/alpinejs', alpinejs));
	app.use(mount('/bulma', bulma));
	app.use(mount('/htmx.org', htmx));
	app.use(mount('/hx-ds-inc', hxDsIncrement));
	app.use(mount('/ionicons', ionicons));

	return {app, db: database};
};
