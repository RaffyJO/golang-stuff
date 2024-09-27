package basic

import (
	"learn-protobuf/protogen/basic"
	"log"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUserGroup() {
	users := []*basic.User{
		{
			Id:       1,
			Username: "Superman",
			IsActive: true,
			Password: []byte("123456"),
			Gender:   basic.Gender_GENDER_MALE,
			Emails:   []string{"superman@gmail.com", "superman@hotmail.com"},
		},
		{
			Id:       2,
			Username: "Cat Woman",
			IsActive: true,
			Password: []byte("123456"),
			Gender:   basic.Gender_GENDER_FEMALE,
			Emails:   []string{"catwoman@gmail.com", "catwoman@hotmail.com"},
		},
	}

	userGroup := basic.UserGroup{
		GroupId:     1,
		GroupName:   "Super Hero",
		Users:       users,
		Description: "this is a super hero group",
	}

	jsonBytes, err := protojson.Marshal(&userGroup)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(jsonBytes))
}
