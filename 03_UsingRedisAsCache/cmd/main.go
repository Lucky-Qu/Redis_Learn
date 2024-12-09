package main

import route "Redis_Learn/03_UsingRedisAsCache/routes"

func main() {
	route.RunRouter(route.NewRouter())
}
