package main
	import "net/http"
	import "log"
	import "fmt"
	import "text/template"
 
	type Inventory struct {
        Material string
        Count    uint
    }

func main() {


	//localhost:8080/bar

	println("Server starts")

	http.Handle("/resources/", http.StripPrefix("/resources",
	http.FileServer(http.Dir("./resources"))))

	http.HandleFunc("/bar", handler)
	http.HandleFunc("/query", queryhandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<html><h1>Hi there, I love %s!</h1><img src=\"/resources/jeff-photo2.jpg\"/></html>", r.URL.Path[1:])
}

func queryhandler(w http.ResponseWriter, r *http.Request) {
	println("query action called with r=", r)
	serverchoice := r.FormValue("serverchoice")
	query := r.FormValue("query")
	println("serverchoice=", serverchoice)
	println("query=", query)
    p := &Inventory{Material: "cotton", Count: 3}
    t, _ := template.ParseFiles("jeff.html.template")
    t.Execute(w, p)
}
