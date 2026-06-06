package main

import "taskapi/internal/router"

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
