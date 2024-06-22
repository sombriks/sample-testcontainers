import {resolve} from 'node:path';
import Pug from 'koa-pug';

export const pug = new Pug({
	viewPath: resolve('./app/templates'),
});
