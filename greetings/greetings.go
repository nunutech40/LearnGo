package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)	

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomFormat() string {
	formats := []string{
		"Hi, %v. Wellcome!",
		"Great to see you, %v.",
		"Asalamualaikum %v. ya ahli Surga, insya Allah",
	}
	return formats[rand.Intn(len(formats))]
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("nama kosong")
	}
	message := fmt.Sprintf(RandomFormat(), name)
	return message, nil
}

func MuslimHello(name string) string {
	message := fmt.Sprintf(RandomFormat(), name)
	return message
}