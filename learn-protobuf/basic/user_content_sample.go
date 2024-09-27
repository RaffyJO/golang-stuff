package basic

import (
	"learn-protobuf/protogen/basic"
	"log"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicWriteUserContentV1() {
	userContent := basic.UserContent{
		UserContentId: 1,
		Slug:          "/this-is-v1",
		// Title:         "5 young rich people in the world",
		HtmlContent: "<p> This is the content of the user content </p>",
		// AuthorId:      1,
	}

	WriteProtoToFile(&userContent, "user_content_v1.bin")
}

func BasicReadUserContentV1() {
	log.Println("Reading user content v1")
	var userContent basic.UserContent

	ReadProtoFromFile("user_content_v1.bin", &userContent)
	log.Println(&userContent)

	jsonBytes, err := protojson.Marshal(&userContent)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(jsonBytes))
}

func BasicWriteUserContentV2() {
	userContent := basic.UserContent{
		UserContentId: 2,
		Slug:          "/this-is-v2",
		// Title:         "5 young rich people in the world version 2",
		HtmlContent: "<p> This is the content of the user content version 2 </p>",
		// AuthorId:      1,
		// Category: "NEWS",
	}

	WriteProtoToFile(&userContent, "user_content_v2.bin")
}

func BasicReadUserContentV2() {
	log.Println("Reading user content v2")
	var userContent basic.UserContent

	ReadProtoFromFile("user_content_v2.bin", &userContent)
	log.Println(&userContent)

	jsonBytes, err := protojson.Marshal(&userContent)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(jsonBytes))
}

func BasicWriteUserContentV3() {
	userContent := basic.UserContent{
		UserContentId: 3,
		Slug:          "/this-is-v3",
		// Title:         "5 young rich people in the world version 2",
		HtmlContent: "<p> This is the content of the user content version 3 </p>",
		// AuthorId:      1,
		// Category:    "NEWS",
		// SubCategory: "PEOPLE",
	}

	WriteProtoToFile(&userContent, "user_content_v3.bin")
}

func BasicReadUserContentV3() {
	log.Println("Reading user content v3")
	var userContent basic.UserContent

	ReadProtoFromFile("user_content_v3.bin", &userContent)
	log.Println(&userContent)

	jsonBytes, err := protojson.Marshal(&userContent)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(jsonBytes))
}
