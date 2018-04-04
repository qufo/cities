package main
import (
	"net/http"
	"github.com/qufo/IniReader"
)

func main() {
	config := IniReader.NewIniReader("./config.ini")
	http.Handle("/", http.FileServer(http.Dir(config.Get("directory"))))
	http.ListenAndServe(":"+config.Get("port"), nil)
}