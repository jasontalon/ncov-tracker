package main

import "github.com/jasontalon/ncov-tracker/api"

func main() {
	a := api.NewApp()

	a.StartServer()
}
