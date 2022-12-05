package main

import (
	"github.com/arlomcwalter/pylai/cmd"
	"github.com/arlomcwalter/pylai/database"
)

func main() {
	database.Init()
	defer database.Shutdown()
	cmd.Execute()
}
