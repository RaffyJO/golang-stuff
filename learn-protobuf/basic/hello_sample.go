package basic

import (
	"learn-protobuf/protogen/basic"
	"log"
)

func BasicHello() {
	h := basic.Hello{
		Name: "Raffy Jamil Octavialdy",
	}

	log.Println(&h)
}
