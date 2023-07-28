package main

import "github.com/dev-sareno/ginamus/mrouter"

func main() {
	r := mrouter.GetRouter()
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
