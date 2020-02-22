package main

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/pwinning1991/task/cmd"
	"github.com/pwinning1991/task/db"
)

func main() {
	home, _ := homedir.Dir()
	dbpath := filepath.Join(home, "tasks.db")
	must(db.Init(dbpath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
