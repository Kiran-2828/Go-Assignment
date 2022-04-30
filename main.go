package main

import (
	"GoAssignment-1/handler"
	"GoAssignment-1/router"
)

func main() {
	router.List = router.GetCharacters()
	handler.HandleRequests()
}
