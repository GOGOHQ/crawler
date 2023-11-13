package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://blog.csdn.net/weixin_43864567/article/details/124130041"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("fetch url error:%v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		res, _ := io.ReadAll(resp.Body)
		fmt.Printf("data:%v", string(res))
	}

}
