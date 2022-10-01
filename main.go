package main

import "simple-todo/routers"

func main() {
	const PORT = ":8585"
	routers.StartServer().Run(PORT)
}
