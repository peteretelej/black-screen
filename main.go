package main

import (
	"flag"
	"net/http"
)

var listen = flag.String("listen", ":8080", "http listen address")

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9521", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte(`<!DOCTYPE html>
	<html> <head> <title> </title> 
	<style> html,body{
		display:block; margin:0;padding:0;
		background-color:#000;
	}
	body{
		display:block; width:100%;
		height:20px;
		max-height:100%;
	}
	</style>
	</head>
	<body>
	</body>
	</html>
	`))

}
