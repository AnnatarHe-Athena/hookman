package main

import (
	"fmt"
	"os"
	"time"

	"github.com/AnnatarHe-Athena/hookman/jobs"
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
	setupLogFile()
	jobs.UpdateUsersFor1week()
}
