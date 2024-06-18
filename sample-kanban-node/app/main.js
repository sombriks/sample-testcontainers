import ApiBuilder from 'koa-api-builder'
import { bodyParser } from '@koa/bodyparser'
import Koa from 'koa'
import mount from 'koa-mount'
import Router from '@koa/router'
import serve from 'koa-static'

import { cabin } from './configs/logging.js'
import { db } from './configs/database.js'
import { pug } from './configs/views.js'
import { boardRequests } from './routes/board-requests.js'
import { boardServices } from './services/board-services.js'
import { fakeLoginCheck } from './routes/cookie-redirect.js'

export const app = new Koa()

const service = boardServices({ db })
const controller = boardRequests({ service })

// dynamic api planning
const router = new ApiBuilder({ router: new Router() })
  .path(b => {
    b.get('/', async ctx => ctx.redirect('/board'))
    b.get('/board', fakeLoginCheck, controller.pages.board)
    b.get('/login', controller.pages.login)
    b.get('/table', fakeLoginCheck, controller.pages.table)
  }).build()

// static assets
const alpinejs = serve('node_modules/alpinejs/dist')
const bulma = serve('node_modules/bulma/css')
const htmx = serve('node_modules/htmx.org/dist')
const ionicons = serve('node_modules/ionicons/dist/')
const _static = serve('app/static')

// configuring the koa app. the middleware order matters a lot
app.use(cabin.middleware)
pug.use(app)
app.use(bodyParser())
app.use(router.routes())
app.use(router.allowedMethods())
app.use(mount('/', _static))
app.use(mount('/alpinejs', alpinejs))
app.use(mount('/bulma', bulma))
app.use(mount('/htmx.org', htmx))
app.use(mount('/ionicons', ionicons))
