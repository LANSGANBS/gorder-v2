package main

import (
	"github.com/LANSGANBS/gorder-v2/config"
	"log"
)

func init() {
	if err := config.NewviperConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Printf("%v", viper.Get("order"))
}
