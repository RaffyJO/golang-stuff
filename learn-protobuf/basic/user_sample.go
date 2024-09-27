package basic

import (
	"encoding/json"
	"learn-protobuf/protogen/basic"
	"log"
	"math/rand"

	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/genproto/googleapis/type/latlng"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func BasicUser() {
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
		"python": 10,
		"golang": 20,
	}

	user := basic.User{
		Id:                   1,
		Username:             "Raffy",
		IsActive:             true,
		Password:             []byte("123456"),
		Gender:               basic.Gender_GENDER_MALE,
		Emails:               []string{"raffy@gmail.com", "raffy@hotmail.com"},
		Address:              &address,
		CommunicationChannel: &communicationChannel,
		SkillRating:          skillRating,
		LastLoginTimestamp:   timestamppb.Now(),
		BirthDate:            &date.Date{Year: 2000, Month: 5, Day: 27},
		LastKnownLocation: &latlng.LatLng{
			Latitude:  -6.29847717,
			Longitude: 106.8290577,
		},
	}

	jsonBytes, err := protojson.Marshal(&user)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(jsonBytes))
}

func ProtoToJsonUser() {
	user := basic.User{
		Id:       1,
		Username: "raffy",
		IsActive: true,
		Password: []byte("123456"),
		Gender:   basic.Gender_GENDER_MALE,
		Emails:   []string{"raffy@gmail.com", "raffy@hotmail.com"},
	}

	jsonBytes, err := protojson.Marshal(&user)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(jsonBytes))
}

func JsonToProtoUser() {
	json := `
	{
		"id":1,
		"username":"raffy",
		"is_active":true,
		"password":"123456",
		"gender":1,
		"emails":[
			"raffy@gmail.com",
			"raffy@hotmail.com"
		]
	}`

	var user basic.User
	err := protojson.Unmarshal([]byte(json), &user)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(&user)
}

func randomCommnunicationChannel() anypb.Any {
	paperMail := basic.PaperMail{
		PaperMailAddress: "Paper Mail Address",
	}

	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "Social Media Platform",
		SocialMediaUsername: "Social Media Username",
	}

	instantMessaging := basic.InstantMessaging{
		InstantMessagingProduct:  "Instant Messaging Product",
		InstantMessagingUsername: "Instant Messaging Username",
	}

	var a anypb.Any

	switch r := rand.Intn(10) % 3; r {
	case 0:
		anypb.MarshalFrom(&a, &paperMail, proto.MarshalOptions{})
	case 1:
		anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})
	default:
		anypb.MarshalFrom(&a, &instantMessaging, proto.MarshalOptions{})
	}

	return a
}

func BasicUnmarshalAnyKnown() {
	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "Social Media Platform",
		SocialMediaUsername: "Social Media Username",
	}

	var a anypb.Any
	anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})

	// Known type (Social Media)
	sm := basic.SocialMedia{}
	err := a.UnmarshalTo(&sm)
	if err != nil {
		log.Fatal(err)
	}

	jsonBytes, err := json.Marshal(&sm)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(jsonBytes))
}

func BasicUnmarshalAnyNotKnown() {
	a := randomCommnunicationChannel()

	var unmarshaled protoreflect.ProtoMessage

	unmarshaled, err := a.UnmarshalNew()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Unmarshal as", unmarshaled.ProtoReflect().Descriptor().Name())

	jsonBytes, err := proto.Marshal(unmarshaled)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(jsonBytes))
}

func BasicUnmarshalAnyIs() {
	a := randomCommnunicationChannel()

	// pm for paper mail
	pm := basic.PaperMail{}

	if a.MessageIs(&pm) {
		log.Println("Message is Paper Mail")
		err := a.UnmarshalTo(&pm)
		if err != nil {
			log.Fatal(err)
		}

		jsonBytes, err := protojson.Marshal(&pm)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(string(jsonBytes))
	} else {
		log.Println("Message is not Paper Mail, but ", a.TypeUrl)
	}

}

func BasicOneOf() {
	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "Social Media Platform",
		SocialMediaUsername: "Social Media Username",
	}

	ecomm := basic.User_SocialMedia{
		SocialMedia: &socialMedia,
	}

	user := basic.User{
		Id:                    1,
		Username:              "Raffy",
		IsActive:              true,
		ElectronicCommChannel: &ecomm,
	}

	jsonBytes, err := protojson.Marshal(&user)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(jsonBytes))
}
