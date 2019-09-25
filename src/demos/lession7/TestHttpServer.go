package lession7

import (
	"../../controller"
	"net/http"
)

func main() {
	http.HandleFunc("/", controller.IndexHandler)
	http.ListenAndServe("127.0.0.1:8000", nil)
}
