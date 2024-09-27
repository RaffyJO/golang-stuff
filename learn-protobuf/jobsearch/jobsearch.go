package jobsearch

import (
	"encoding/json"
	"learn-protobuf/protogen/basic"
	"learn-protobuf/protogen/dummy"
	"learn-protobuf/protogen/jobsearch"
)

func JobSearchSoftware() {
	js := jobsearch.JobSoftware{
		JobSoftwareId: 1,
		Application: &basic.Application{
			Version:   "1.0.0",
			Name:      "The Application",
			Platforms: []string{"Windows", "Linux", "MacOS"},
		},
	}

	jsonBytes, err := json.Marshal(&js)
	if err != nil {
		panic(err)
	}

	println(string(jsonBytes))
}

func JobSearchCandidate() {
	jc := jobsearch.JobCandidate{
		JobCandidateId: 1,
		Application: &dummy.Application{
			ApplicationId:       1,
			ApplicationFullName: "Raffy Jamil",
			Phone:               "1234567890",
			Email:               "raffy@gmail.com",
		},
	}

	jsonBytes, err := json.Marshal(&jc)
	if err != nil {
		panic(err)
	}

	println(string(jsonBytes))
}
