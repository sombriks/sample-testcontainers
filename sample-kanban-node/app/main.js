import ApiBuilder from 'koa-api-builder'
import { bodyParser } from '@koa/bodyparser'
import Koa from 'koa'
import mount from 'koa-mount'
import Router from '@koa/router'
import serve from 'koa-static'

import { cabin } from './configs/logging.js'
import { prepareDb } from './configs/database.js'
import { pug } from './configs/views.js'
import { boardRequests } from './routes/board-requests.js'
import { boardServices } from './services/board-services.js'
import { fakeLogin, fakeLoginCheck, fakeLogout } from './routes/cookie-redirect.js'

const _db = prepareDb()
const _service = boardServices({ db: _db })
const _controller = boardRequests({ service: _service })

/**
 * function to return a properly configured Koa application
 *
 * @param {import('knex')} db query builder reference
 * @param service our service module to do business with
 * @param controller our routes/requests/controllers to validate and handle things
 *
 * @returns {{app: Application, db: import('knex')}}  app and db with all
 * the goodies we will need
 */
export const prepareApp = ({
  db = _db,
  service = _service,
  controller = _controller
}) => {
  const app = new Koa()

  // dynamic api planning
  const router = new ApiBuilder({ router: new Router() })
    .path(b => {
      b.get('/', async ctx => ctx.redirect('/board'))
      b.get('/board', fakeLoginCheck, controller.pages.board)
      b.path('/login', b => {
        b.get(controller.pages.login)
        b.post(fakeLogin({ service }))
      })
      b.get('/logout', fakeLogout({ service }))
      b.get('/table', fakeLoginCheck, controller.pages.table)
      b.path('/task', fakeLoginCheck, b => {
        b.post(controller.components.addTask)
        b.path('/:id', b => {
          b.put(controller.components.updateTask)
          b.del(controller.components.deleteTask)
          b.del('/person/:personId', controller.components.removePerson)
          b.post('/join', controller.components.joinTask)
          b.post('/comments', controller.components.addComment)
        })
      })
    }).build()

  // static assets
  const alpinejs = serve('node_modules/alpinejs/dist')
  const bulma = serve('node_modules/bulma/css')
  const htmx = serve('node_modules/htmx.org/dist')
  const hxDsInc = serve('node_modules/hx-dataset-include/lib')
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
  app.use(mount('/hx-ds-inc', hxDsInc))
  app.use(mount('/ionicons', ionicons))

  return { app, db }
}
