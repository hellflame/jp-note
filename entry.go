package main

import (
	"embed"
	"fmt"

	"github.com/hellflame/jp-note/api"
	"github.com/hellflame/jp-note/utils"
)

//go:embed pages
var pages embed.FS

//go:embed data
var data embed.FS

func main() {
	args := utils.ParseArgs()
	if args != nil {
		address := args.GetAddress()
		if !args.NoBrowser {
			utils.OpenExplore(fmt.Sprintf("http://%s", address))
		}
		api.RegisterAPI(&pages, &data).Start(address)
	}
}
