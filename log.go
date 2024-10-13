package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func logState(env *Environment) {
	entry, _ := env.TotalState()
	b, err := json.Marshal(entry)

	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}
