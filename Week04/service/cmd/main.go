package main

import(
	"internal"
	"log"
)

func main(){
	log.Fatal(internal.NewApp().Run())
}
