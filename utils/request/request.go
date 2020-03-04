package request

import (
	"bytes"
	"io"
	"net/http"
)

func ReadJsonBody(r *http.Request) []byte {
	var buf bytes.Buffer
	for {
		b := make([]byte, 512)
		l, err := r.Body.Read(b)
		buf.Write(b[:l])
		if err == io.EOF {
			break
		}
	}
	return buf.Bytes()
}
