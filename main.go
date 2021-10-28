package main

import (
	"fmt"

	"github.com/andy1341/lets-go-chat/pkg/hasher"
)

func main() {
	fmt.Println(hasher.HashPassword("asd"))
	fmt.Println(hasher.CheckPasswordHash("asd", "passhash"))
}
