import Pug from 'koa-pug';

export const pug = new Pug({
	viewPath: 'app/templates', // TODO use import.meta.url thing
});
