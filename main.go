package main

import "gin_im/router"

func main() {
	eng := router.Router()
	eng.Run(":8080")
}