/**
 * services for kanban
 *
 * @param {{db}} options
 * @returns configured services
 */
export const boardServices = ({db}) => {
  return {
    async listUsers() {
      return db("kanban.person")
    }
  }
}
