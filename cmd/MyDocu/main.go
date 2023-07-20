package main

import (
	"github.com/yosa12978/MyDocu/internal/api"
	"github.com/yosa12978/MyDocu/internal/config"
)

func init() {
	//config.NewJsonConfigParser(os.Args[1]).Parse()
	config.NewJsonConfigParser("config.json").Parse()
}

func main() {
	api.Listen()
}
