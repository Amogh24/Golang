package main
import(
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/",handler) // each http request will call the function handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil)) //standard error handling
}

 func handler(w http.ResponseWriter, r *http.Request){     //this function prints the path component 
 	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
 }


	