import test from 'ava'

import { bar } from './board.js'

test('board should be ok', t => {
  t.truthy(bar)
})
