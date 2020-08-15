package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sprinkelcell/CLI-Task-Manager/cmd"
	"github.com/sprinkelcell/CLI-Task-Manager/db"
	"github.com/sprinkelcell/CLI-Task-Manager/dir"
)

func main() {
	//cmd.RootCmd.Execute()
	homeDir, err := dir.DirWindows()
	must(err)
	dbPath := filepath.Join(homeDir, "tasks.db")
	
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
