package main

import (
	"fmt"
	"log"

	"github.com/alexflint/go-arg"
)

var version = "v0.5.0"

type config struct {
	SourceDir string `arg:"required" help:"filesystem path to read files relative from"`
	AppKey    string `arg:"env:MY_APP_KEY"`
}

func (config) Version() string {
	return "my_app " + version
}

func main() {
	fmt.Println("Do something useful here...")

	log.SetPrefix("my_app: ")
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	var cfg config

	p := arg.MustParse(&cfg)

	if cfg.AppKey == "" {
		p.Fail("Must provide AppKey wither in OS env MY_APP_KEY or --appkey.")

	}
}
