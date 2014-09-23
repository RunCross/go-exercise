package ser

import (
	"io"
	"net/http"
)

func Hello(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "test hello")
}
