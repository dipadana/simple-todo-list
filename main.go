package main

import "simple-todo/routers"

func main() {
	const PORT = ":3000"
	routers.StartServer().Run(PORT)
}
