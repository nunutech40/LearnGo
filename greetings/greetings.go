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

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func MuslimHello(name string) string {
	message := fmt.Sprintf(RandomFormat(), name)
	return message
}