package main

import (
	"github.com/dev-sareno/ginamus/web"
	"github.com/dev-sareno/ginamus/worker"
	"os"
	"strings"
)

func main() {
	argsAll := strings.Join(os.Args, " ")
	if strings.Contains(argsAll, "worker") {
		worker.Run()
	} else {
		web.Run()
	}
}
