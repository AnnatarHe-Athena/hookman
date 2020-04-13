package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/AnnatarHe-Athena/hookman/service"
	"github.com/sirupsen/logrus"
)

func setupLogFile() {
	filename := fmt.Sprintf("/var/log/hookman/log-%s.log", time.Now().Format(time.RFC3339))
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(f)
}

func main() {
	// setupLogFile()
	// jobs.UpdateUsersFor1week()
	update()
}

type t struct {
	UserName string `json:"username"`
	UID      string `json:"uid"`
}

func update() {
	f, _ := ioutil.ReadFile("./mapping.json")

	var mapping []t

	json.Unmarshal(f, &mapping)

	for _, d := range mapping {
		err := service.TempUpdateUserDomainToUid(d.UserName, d.UID)
		logrus.Println(err)
	}
}
