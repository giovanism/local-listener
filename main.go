package main
 
import (
	"fmt"
	"net/http"
)
 
func main() {
	// handle route using handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        	fmt.Fprintf(w, "Hello, Local Citizen!\n")
    	})

	fmt.Println("Listening to localhost:8080")
    	http.ListenAndServe("localhost:8080", nil)
}
