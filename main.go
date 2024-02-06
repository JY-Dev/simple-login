package main

import (
	"simple-login/apis/route"
)

func main() {

	r := route.NewRouter()

	r.Run()
}
