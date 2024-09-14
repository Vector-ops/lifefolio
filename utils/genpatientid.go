package utils

import (
	"fmt"
	"strings"
	"time"
)

func GenPatientId(firstName, lastName string) string {
	prefix := "LFP"
	name := fmt.Sprintf("%s%s", strings.Split(firstName, "")[0], strings.Split(lastName, "")[0])
	year, month, day := time.Now().Date()

	patientId := fmt.Sprintf("%s%s%d%d%d", prefix, name, year, month, day)

	return patientId
}
