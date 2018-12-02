package e

import (
	"log"
)

func Ignor(err error) {}

func Log(err error) {
	if err != nil {
		go func() {
			log.Print(err)
		}()
	}
}

func Panic(err error) {
	if err != nil {
		go func() {
			panic(err)
		}()
	}
}
