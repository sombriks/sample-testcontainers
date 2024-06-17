import test from 'ava'

import { app } from './main.js'

test('app should be ok', t => {
  t.truthy(app)
})
