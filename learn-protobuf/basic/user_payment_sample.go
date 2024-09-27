package basic

import (
	"learn-protobuf/protogen/basic"
	"log"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicReadUserPayment() {
	log.Println("Reading user payment")
	var userPayment basic.UserPayment

	ReadProtoFromFile("user_content_v1.bin", &userPayment)
	log.Println(&userPayment)

	jsonBytes, err := protojson.Marshal(&userPayment)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(jsonBytes))
}
