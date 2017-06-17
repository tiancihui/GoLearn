package main
import (
		"io"
		"log"
		"net/http"
)



func HttpHandler(w http.ResponseWriter, req *http.Request){
	io.WriteString(w, "hello world\r\n")
}



func main(){


	http.HandleFunc("/hello", HttpHandler)
	http.Handle("/", http.FileServer(http.Dir("./")))
		err := http.ListenAndServe(":12345", nil)
		if err != nil {
			log.Fatal("LestenAndServe", err)

		}
}
