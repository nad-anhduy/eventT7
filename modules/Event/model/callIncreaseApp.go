package model

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CallAppIncrease(uid, id string) {

	url := fmt.Sprintf("https://cohon-itc.mservice.com.vn:8000/portal/game-history/%s", id)
	method := "POST"
	log.Printf(`SoulT7 - [%s] - Request call increase app with url: %s`, uid, url)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	log.Printf(`SoulT7 - [%s] - Request call increase app: %s`, uid, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("authorization", "e7be79d0-7842-492c-8e5c-5eb3fbceac23")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Printf(`SoulT7 - [%s] - Status fail : %s`, uid, res.Status)
		log.Printf(`SoulT7 - [%s] - Increase fail with err: %s`, uid, err)
		fmt.Println(err)
		return
	}
	log.Printf(`SoulT7 - [%s] -  Increase fail success: %s`, uid, string(body))
}
