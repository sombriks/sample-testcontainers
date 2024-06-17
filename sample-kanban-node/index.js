import { app } from './app/main.js'
import { db } from './app/configs/database.js'

// entrypoint
const start = async () => {
  // check db, run migrates, anything
  await db.raw('select 1 + 1')

  app.listen(process.env.PORT, () => {
    console.log(`http://localhost:${process.env.PORT}`)
  })
}

start()
