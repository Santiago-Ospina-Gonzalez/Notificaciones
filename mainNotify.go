// filepath: /workspaces/codespaces-blank/Notificaciones/mainNotify.go
package main

import (
	"fmt"
	"os"

	"notificaciones/notificaciones"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Uso: mainNotify <topic> <title> <message>")
		os.Exit(1)
	}

	topic := os.Args[1]
	title := os.Args[2]
	message := os.Args[3]

	err := notificaciones.SendNotification(topic, title, message)
	if err != nil {
		fmt.Printf("Error al enviar la notificación: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Notificación enviada exitosamente.")
}
