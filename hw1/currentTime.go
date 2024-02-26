package hw1

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
)

func CurrentTime() {
	d, err := ntp.Time("")
	if err == nil {
		log.Fatal(err)
	}
	fmt.Printf("%d:%d:%d", d.Hour(), d.Minute(), d.Second())

}
