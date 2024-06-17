import ApiBuilder from 'koa-api-builder'
import { bodyParser } from '@koa/bodyparser'
import Koa from 'koa'
import Router from '@koa/router'

import { cabin } from './configs/logging.js'
import { db } from './configs/database.js'
import { pug } from './configs/views.js'
import { boardRequests } from './controllers/board-requests.js'
import { boardServices } from './services/board-services.js'

export const app = new Koa()

const service = boardServices({ db })
const controller = boardRequests({ service })

const router = new ApiBuilder({ router: new Router() })
  .path(b => {
    b.get('/', async ctx => { // simple redirect if the cookie is there or not
      const userMaybe = ctx.cookies.get('x-user-info')
      if (userMaybe) return ctx.redirect('/board')
      else return ctx.redirect('/login')
    })
    b.get('/board', controller.pages.board)
    b.get('/login', controller.pages.login)
    b.get('/table', controller.pages.table)
  }).build()

// configuring the koa app. the middleware order matters a lot
app.use(cabin.middleware)
pug.use(app)
app.use(bodyParser())
app.use(router.routes())
app.use(router.allowedMethods())
