package main

import "fmt"
import "redisgo/cache"

func main()  {
	_int, err := cache.Incr("asdsadasd")
	fmt.Println(_int, err)
	fmt.Println("aaaaa")
}