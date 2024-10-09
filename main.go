package main

import (
	actions "sonata-cli/Actions"
	utils "sonata-cli/Utils"
	"sonata-cli/database"
)

func main() {
	database.CreateDb()
	utils.CheckOrCreateFolderLocal()
	actions.ShowMenuList()
}
