package notificaciones

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Notification representa la estructura de la notificación a enviar
type Notification struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

// SendNotification envía una notificación al servidor ntfy
func SendNotification(topic, title, message string) error {
	url := fmt.Sprintf("https://ntfy.sh/%s", topic)

	notification := Notification{
		Title:   title,
		Message: message,
	}

	payload, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("error al serializar la notificación: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("error al crear la solicitud HTTP: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error al enviar la solicitud: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("error en la respuesta del servidor: %s", resp.Status)
	}

	return nil
}
