  package last

  import (
    "net/http"
    "github.com/egocentri/go-webcalc/package_calc/calculation.go"
    "github.com/egocentri/go-webcalc/package_web/handler.go"
  func main() {
	  http.HandleFunc("/api/v1/calculate", handler.calculateHandler)
  	port := "8080"
  	err := http.ListenAndServe(":"+port, nil)
  	if err != nil {
	  	fmt.Printf("Error starting server: %s\n", err)
	}
}  
