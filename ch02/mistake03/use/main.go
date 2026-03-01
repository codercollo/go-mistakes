package main

import (
	"fmt"
	"net/http"
)

// VALID USE: init is acceptable here because:
// 1. Cannot fail - http.HandleFunc only panics if handler is nil (not the case)
// 2. No global variables being set
// 3. No impact on unit tests - just register static routes
// 4. Pure static configuration - no external dependencies
func init() {
	redirect := func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	//Registering static routes - cannot fail, no side effects on tests
	http.HandleFunc("/blog", redirect)
	http.HandleFunc("/blog/", redirect)

	static := http.FileServer(http.Dir("static"))
	http.Handle("/favicon.ico", static)
	http.Handle("/fonts.css", static)
	http.Handle("/fonts/", static)

}

func main() {
	fmt.Println("server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
