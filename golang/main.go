package main

import (
	"github.com/dev-sareno/ginamus/web"
	"github.com/dev-sareno/ginamus/webworker"
	"os"
	"strings"
)

func main() {
	argsAll := strings.Join(os.Args, " ")
	if strings.Contains(argsAll, "webworker") {
		webworker.Run()
	} else {
		web.Run()
	}
}
