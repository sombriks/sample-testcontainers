import { prepareApp } from './app/main.js'

// entrypoint
const start = async ({ app, db } = prepareApp({})) => {
  // check db, run migrates, anything
  await db.raw('select 1 + 1')

  app.listen(process.env.PORT, () => {
    console.log(`http://localhost:${process.env.PORT}`)
  })
}

start()
