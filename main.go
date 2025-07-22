package main

import (
    "fmt"
    "io"
    "net/http"

    "github.com/gorilla/mux"
)



func postHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("¡Bienvenido a la API con Gorilla Mux!")

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading the request")

		return
	}

	fmt.Printf("Contenido recibido: %s\n", string(body))
       w.Write([]byte("POST recibido correctamente"))
}

func main(){
	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Ruta no encontrada:", r.URL.Path, "Método:", r.Method)
		http.NotFound(w, r)
	})

	router.HandleFunc("/hello", postHandler).Methods("POST")

       router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Servidor Gorilla Mux en ejecución (GET /)")
	  }).Methods("GET")

	fmt.Println("Server listening on port 8090")

	err := http.ListenAndServe(":8090", router)
	if err != nil {
		fmt.Println(err.Error())
	}


}
