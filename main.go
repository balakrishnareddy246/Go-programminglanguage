package main 

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"

)
type Article struct {
	  Title string 'json:"Title"'
	  Ttle stirng 'json:"desc"'
	  Content string 'json:"content"'
}
type Articles []Article
func allArticles(w http.ResponseWriter, r *http.Request) {
	  articles := Articles{
		   Article{Ttile:"Test Title", Desc: "Test Description", Cntent: "Hello world"},
	  }
fmt.Println("Endpoint Hit: All Articles Endpoint")
json.NewEncoder(w).Endcode(articles)
	}
	func homepage(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintlf(w).Encode(articles)
	}
	func handlerRequests(0 {
		http.HandFunc{"/", homepage}
		htt.HandlerFunc("/articles", allArticles)
		log.Fatal(http.ListenAndServe(":8081", nil))
	}
	func main() {
		handleRequest()
	}