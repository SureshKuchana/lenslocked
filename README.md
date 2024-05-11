# GO-Lang

## Alternative to (chi router) create routing

```go
type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request){
	switch r.URL.Path {
		case "/":
			homeHandlerFunc(w, r)
		case "/contact":
			contactHandler(w, r)
		default:
			http.Error(w, "Page not found", http.StatusNotFound)		
	}	
}

func main(){
	r := chi.NewRouter()
	fmt.Println("Starting the server on http://localhost:3000...")
	http.ListenAndServe(":3000", r)
}
```