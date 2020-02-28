package statistic

import (
	"bufio"
	"bytes"
	"fradyspace.com/go-server-practice/utils"
	"io"
	"net/http"
	"os"
)

var logger *utils.Logger

func init() {
	logger = utils.NewLogger("statistic")
}

func MainController(w http.ResponseWriter, req *http.Request) {
	logger.Printf("called!")
	switch req.Method {
	case http.MethodGet:
		getStaticBasePage(w, req)
	case http.MethodPost:
	}
}

func getStaticBasePage(w http.ResponseWriter, req *http.Request) {
	dir, err := os.Getwd()
	if err != nil {
		logger.Errorf(err.Error())
		w.WriteHeader(500)
		return
	}
	html, err := os.Open(dir + "/app/web/statistic")
	if err != nil {
		logger.Errorf(err.Error())
		w.WriteHeader(404)
		return
	}
	reader := bufio.NewReader(html)
	var buf bytes.Buffer
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			logger.Errorf(err.Error())
			w.WriteHeader(500)
			return
		}
		buf.WriteString(str)
	}

	_, err = w.Write(buf.Bytes())
	if err != nil {
		logger.Errorf(err.Error())
		return
	}
}
