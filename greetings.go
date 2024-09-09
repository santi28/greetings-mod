package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("el nombre no puede ser vacío")
	}

	// Devuelve un saludo que incluye el nombre de la persona "Hola, <nombre>. Bienvenido!"
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func HelloNames(names []string) (map[string]string, error) {
	// Uitlizamos make para crear un mapa que asigna strings a strings
	messages := make(map[string]string)

	// Iteramos a través de los valores recibidos
	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hola, %v. Bienvenido!",
		"¡Hola %v! ¿Cómo estás?",
		"¡Hola %v! ¿Qué tal?",
	}

	return formats[rand.Intn(len(formats))]
}
