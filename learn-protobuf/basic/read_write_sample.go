package basic

import (
	"io/ioutil"
	"learn-protobuf/protogen/basic"
	"log"

	"google.golang.org/protobuf/proto"
)

func WriteProtoToFile(message proto.Message, fileName string) {
	bytes, err := proto.Marshal(message)
	if err != nil {
		log.Fatal("Can not marshal proto message", err)
	}

	if err := ioutil.WriteFile(fileName, bytes, 0644); err != nil {
		log.Fatal("Can not write proto message to file", err)
	}
}

func ReadProtoFromFile(fileName string, message proto.Message) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Can not read proto message from file", err)
	}

	err = proto.Unmarshal(bytes, message)
	if err != nil {
		log.Fatal("Can not unmarshal proto message", err)
	}
}

func dummyUser() basic.User {
	address := basic.Address{
		Street:  "123 Main Street",
		City:    "Anytown",
		Country: "USA",
		Coordinate: &basic.Address_Coordinate{
			Latitude:  123.456,
			Longitude: 789.123,
		},
	}

	communicationChannel := randomCommnunicationChannel()
	skillRating := map[string]uint32{
		"durability": 10,
		"speed":      20,
	}

	return basic.User{
		Id:                   1,
		Username:             "Superman",
		IsActive:             true,
		Password:             []byte("123456"),
		Gender:               basic.Gender_GENDER_MALE,
		Emails:               []string{"superman@gmail.com", "superman@hotmail.com"},
		Address:              &address,
		CommunicationChannel: &communicationChannel,
		SkillRating:          skillRating,
	}
}

func WriteToFileSample() {
	user := dummyUser()

	WriteProtoToFile(&user, "superman_file.bin")
}

func ReadFromFileSample() {
	var user basic.User

	ReadProtoFromFile("superman_file.bin", &user)
	log.Println(&user)
}
