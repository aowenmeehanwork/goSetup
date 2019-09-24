package server

import "fmt"

func Init() {

	fmt.Sprintf("Starting on localhost, port 3000")
	r := NewRouter()
	r.Run(":3000")

}
