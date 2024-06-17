import fs from "fs"

const result = fs.readdirSync(".", {withFileTypes: true})
console.log(result)
