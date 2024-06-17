import test from 'ava'

import { boardServices } from './board-services.js'

test('board should be ok', t => {
  t.truthy(boardServices)
})
