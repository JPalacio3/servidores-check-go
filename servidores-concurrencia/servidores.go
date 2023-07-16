package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
CONCURRENCIA:
La concurrencia permite la ejecución simultánea de múltiples tareas.
Se utilizan goroutines para realizar tareas en paralelo.
Los canales facilitan la comunicación y sincronización entre goroutines.
Los canales en Go son mecanismos de comunicación para sincronizar y transferir datos entre goroutines.
Permiten la comunicación segura y coordinada entre diferentes partes del programa.
Se definen con el operador `chan` seguido del tipo de datos transmitidos.
Los datos se envían y reciben utilizando las operaciones `<-`.
Ayudan a prevenir condiciones de carrera y facilitan la coordinación en programas concurrentes.
Las goroutines son unidades de ejecución ligeras en Go.
Permiten la concurrencia y ejecución en paralelo.
Son más eficientes y rápidas que los hilos tradicionales.

SE UTILIZA LA PALABRA RESERVADA "go" PARA INDICAR LA CONCURRECNIA EN LA EJECUCIÓN DE LA APLICACIÓN
*/

// Función para revisar disponibilidad de servidores
func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		canal <- servidor + "No está disponible" // Implementación del canal
	} else {
		canal <- servidor + "Se encuentra Online"
	}
}

func main() {
	inicio := time.Now()

	// Creación del canal
	canal := make(chan string)

	servidores := []string{
		"http://oregoom.com/",
		"https://github.com/JPalacio3/Escuela_Desarrollo_de_Software_Basico/tree/master/GO/src/figuras",
		"https://www.udemy.com/",
		"https://facebook.com/",
		"htps://youtube.com/",
		"https://google.com/",
	}

	for _, servidor := range servidores {
		go revisarServidor(servidor, canal) // "go" es la palabra reservada que crea la concurrencia
	}

	// Se debe crear un ciclo para poder imprimir los resultados de las iteraciones de los canales
	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}

	tiempoTranscurrido := time.Since(inicio)
	fmt.Println("Tiempo de ejecución:", tiempoTranscurrido)
}

/*
SALIDA:
htps://youtube.com/No está disponible
https://www.udemy.com/Se encuentra Online
https://github.com/JPalacio3/Escuela_Desarrollo_de_Software_Basico/tree/master/GO/src/figurasSe encuentra Online
https://facebook.com/Se encuentra Online
https://google.com/Se encuentra Online
http://oregoom.com/Se encuentra Online
Tiempo de ejecución: 754.6543ms
*/
