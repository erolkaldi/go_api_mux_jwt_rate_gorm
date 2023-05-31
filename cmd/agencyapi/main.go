package main

import "github.com/erolkaldi/agency/pkg/app"

func main() {
	a := app.App{}
	if a.InitializeDB() {
		a.Routes()
		a.Run()
	}

}
