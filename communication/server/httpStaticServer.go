package server

import (
	"net/http"
	"io/ioutil"
	"path"
	"fmt"
)

type StaticHttpHandler struct {
	Root string
}

func (this *StaticHttpHandler) ServeHTTP(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	filePath := httpRequest.URL.Path;
	if filePath == "/" {
		filePath = "/index.html"
	}
	file, err := ioutil.ReadFile(path.Join(this.Root, filePath))

	if err != nil {
		http.NotFound(httpResponse, httpRequest)
		return
	}

	fmt.Printf("serve static:%s\n", filePath)
	httpResponse.Write(file)
}
