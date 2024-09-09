# Greetings Module

Este es un pequeño proyecto en Go creado para aprender y experimentar con el lenguaje. El módulo proporciona funciones para generar saludos personalizados.

## Descripción

El módulo `greetings` ofrece las siguientes funciones:

- `Hello(name string) (string, error)`: Devuelve un saludo personalizado para el nombre proporcionado. Si el nombre es una cadena vacía, devuelve un error.
- `HelloNames(names []string) (map[string]string, error)`: Acepta una lista de nombres y devuelve un mapa donde cada nombre está asociado con un saludo personalizado. Si algún nombre en la lista es una cadena vacía, devuelve un error.
- `randomFormat() string`: Devuelve un formato de saludo al azar de una lista de opciones.

## Instalación

Para utilizar este módulo en tu proyecto, puedes utilizar este código:

```bash
git get -u github.com/santi28/greetings-mod
```

## Uso

A continuación, se muestra un ejemplo básico de cómo usar este módulo:

```go
package main

import (
    "fmt"
    "log"
    "greetings-mod/greetings"
)

func main() {
    message, err := greetings.Hello("Pablo")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(message)

    names := []string{"Pablo", "Ana", "Luis"}
    messages, err := greetings.HelloNames(names)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(messages)
}
```

## Pruebas

El módulo incluye pruebas básicas para validar su funcionamiento:

- `TestHelloName`: Verifica que `Hello(name)` devuelva un saludo que contenga el nombre proporcionado.
- `TestHelloEmpty`: Verifica que `Hello("")` devuelva un error cuando se le pasa un nombre vacío.

Para ejecutar las pruebas, utiliza el siguiente comando:

```bash
go test
```
