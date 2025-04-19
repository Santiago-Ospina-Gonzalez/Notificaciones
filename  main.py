import threading
import time
import os
from datetime import datetime

def schedule_notification(label, title, message, notify_time):
    """
    Programa una notificación para ejecutarse en el momento especificado.
    """
    print(f"Notificación programada para las {notify_time}. Esperando...")
    while True:
        # Obtén la hora actual en cada iteración
        current_time = datetime.now().strftime("%H:%M")
        if current_time == notify_time:
            print(f"Hora actual: {current_time}. Ejecutando notificación...")
            # Ejecuta el script de Go
            os.system(f'go run mainNotify.go "{label}" "{title}" "{message}"')
            break
        time.sleep(20)  # Revisa cada 30 segundos

def main():
    print("=== Programador de Notificaciones ===")
    label = input("Ingrese el label (tópico): ")
    title = input("Ingrese el título: ")
    message = input("Ingrese el mensaje: ")
    notify_time = input("Ingrese la hora para enviar la notificación (HH:MM, formato 24 horas): ")

    # Crear un hilo para manejar la notificación
    notification_thread = threading.Thread(target=schedule_notification, args=(label, title, message, notify_time))
    notification_thread.start()

if __name__ == "__main__":
    main()