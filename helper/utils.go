package helper

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"os"
)

func Generate4CharsPassword(input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	output := fmt.Sprintf("%x", h.Sum(nil))

	return output[len(output)-4:]
}

func RenderHTML(w http.ResponseWriter, r *http.Request, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(w, r, file.Name(), fi.ModTime(), file)
}
