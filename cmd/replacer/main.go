package main

import (
	"os"

	"github.com/no-src/log"
	"github.com/no-src/replacer/cmd"
)

func main() {
	var err error
	defer func() {
		log.Close()
		if err != nil {
			os.Exit(1)
		}
	}()
	cmd.InitFlags()
	err = log.ErrorIf(cmd.RunWithFlags(), "running replacer error")
}
