package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type MyStruct struct {
	Name     string    `json:"myname"`
	DateTime time.Time `json:"mytime"`
}

func main() {
	m := MyStruct{"djohnson", time.Now()}
	fmt.Printf("%#v\n", m)
	if buf, err = json.Marshal(m); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", string(buf))

	var m1 MyStruct
	if err = json.Unmarshal(buf, &m1); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", m1)
}
