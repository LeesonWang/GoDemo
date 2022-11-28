package httpDemo

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Httptest() {
	// 客户端
	client := &http.Client{}

	url := "www.baidu.com"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	response, err := client.Do(request)

	stdout := os.Stdout
	io.Copy(stdout, response.Body)

	code := response.StatusCode

	fmt.Println(code)
}
