package basic

import (
	"io/ioutil"
	"learn-protobuf/protogen/basic"
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func WriteProtoToJson(message proto.Message, fileName string) {
	bytes, err := protojson.Marshal(message)
	if err != nil {
		log.Fatal("Can not marshal proto message", err)
	}

	if err := ioutil.WriteFile(fileName, bytes, 0644); err != nil {
		log.Fatal("Can not write proto message to file", err)
	}
}

func ReadProtoFromJson(fileName string, message proto.Message) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Can not read proto message from file", err)
	}

	err = protojson.Unmarshal(bytes, message)
	if err != nil {
		log.Fatal("Can not unmarshal proto message", err)
	}
}

func WriteToJsonSample() {
	user := dummyUser()

	WriteProtoToJson(&user, "superman_file.json")
}

func ReadFromJsonSample() {
	var user basic.User

	ReadProtoFromJson("superman_file.json", &user)
	log.Println(&user)
}
