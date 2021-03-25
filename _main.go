package main

// import (
//     "fmt"
//     "net/http"
//     "log"

//     "github.com/julienschmidt/httprouter"
// )

// func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//     fmt.Fprint(w, "Welcome!\n")
// }

// func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//     fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
// }

// func main() {
//     router := httprouter.New()
//     router.GET("/", Index)
//     router.GET("/hello/:name", Hello)
// 	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if r.Header.Get("Access-Control-Request-Method") != "" {
// 			// Set CORS headers
// 			header := w.Header()
// 			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
// 			header.Set("Access-Control-Allow-Origin", "*")
// 		}
// 		// Adjust status code to 204
// 		w.WriteHeader(http.StatusNoContent)
// 	})
//     log.Fatal(http.ListenAndServe(":8888", router))
// }