package main

import "github.com/refluxwhw/nas/cmd/nas-server/app"

func main() {
	cmd := app.NewNasServerCommand()
	err := cmd.Execute()
	if err != nil {
		return
	}
}
