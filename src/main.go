package main

import (
	"crane/config"
	"crane/src/http"
	"strconv"
)

func main()  {

	e := http.CreateServer()

	e.Logger.Fatal(e.Start(config.Host + ":" + strconv.Itoa(config.PORT)))
}