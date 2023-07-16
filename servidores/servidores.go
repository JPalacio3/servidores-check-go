package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
SIN CONCURRENCIA:
La no concurrencia significa que las tareas se ejecutan de manera secuencial y no simultánea.
No se utilizan goroutines ni canales para paralelizar la ejecución.
El programa se ejecuta línea por línea, sin intercalación de tareas.
*/

func revisarServidor(servidor string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, "No está disponible")
	} else {
		fmt.Println(servidor, "Se encuentra Online")
	}
}

func main() {
	inicio := time.Now()
	servidores := []string{
		"http://oregoom.com/",
		"https://github.com/JPalacio3/Escuela_Desarrollo_de_Software_Basico/tree/master/GO/src/figuras",
		"https://www.udemy.com/",
		"https://facebook.com/",
		"htps://youtube.com/",
		"https://google.com/",
	}

	for _, servidor := range servidores {
		revisarServidor(servidor)
	}

	tiempoTranscurrido := time.Since(inicio)
	fmt.Println("Tiempo de ejecución:", tiempoTranscurrido)
}

/*
SALIDA:
http://oregoom.com/ Se encuentra Online
https://github.com/JPalacio3/Escuela_Desarrollo_de_Software_Basico/tree/master/GO/src/figuras Se encuentra Online
https://www.udemy.com/ Se encuentra Online
https://facebook.com/ Se encuentra Online
htps://youtube.com/ No está disponible
https://google.com/ Se encuentra Online
Tiempo de ejecución: 2.4043432s

*/
