package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`
			<html>
				<head>
					<title>Chat</title>
				</head>
				<body>
					Let's chat
				</body>
			</html>
		`))
	})

	// start the webserver
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAdnServer:", err)
	}
}