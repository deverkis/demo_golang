package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	//message := fmt.Sprintf("Hi, %v. Welcome!", name)
	message := fmt.Sprintf(randomFormat(), name)
	//message := fmt.Sprint(randomFormat()) for test failure
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// Go executes init functions automatically at program startup
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi,%v. Welcome!",
		"Great to see you, %v!",
		"Hail,%v! Well met!",
	}
	// php $formats[array_random($formats)]
	// go format[rand.Intn(len(formats))]
	return formats[rand.Intn(len(formats))]
}
