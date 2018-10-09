package main

import (
	"net"
	"time"
	"crypto/tls"
	"net/http"
	"strings"
	"log"
	"net/http/httputil"
	"fmt"
	"io/ioutil"
)

func main()  {
	tr := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 	30 *time.Second,
			KeepAlive: 	30 *time.Second,
		}).Dial,
		TLSClientConfig:&tls.Config{InsecureSkipVerify:true},
	}
	client := &http.Client{Transport: tr, Timeout: 60*time.Second}

	req, err := http.NewRequest("POST","https://www.baidu.com",strings.NewReader("hello"))
	if err != nil {
		log.Fatal("new request err", err)
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	reqMsg, _ := httputil.DumpRequestOut(req, true)
	fmt.Println("请求报文:", string(reqMsg))

	rsp, err := client.Do(req)
	if err != nil {
		log.Fatal("request err", err)
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		log.Fatal("rsp err",rsp.Status)
	}

	rspMsg, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Fatal("read rsp msg err", err)
	}

	fmt.Println("应答报文:", string(rspMsg))
}
