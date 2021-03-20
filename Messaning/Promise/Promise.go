package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fanliao/go-promise"
)

func main() {
	p := promise.NewPromise()
	p.OnSuccess(func(v interface{}) {
		// fmt.Println("My V is : ", v)
		fmt.Println("Success")
	}).OnFailure(func(v interface{}) {
		fmt.Println("Failrur")

	}).OnComplete(func(v interface{}) {
		fmt.Println("Complete")
	}).OnCancel(func() {
		fmt.Println("Cancel")
	})

	go func() {
		url := "http://example.com/"

		resp, err := http.Get(url)
		defer resp.Body.Close()
		if err != nil {
			p.Reject(err)
		}

		body, err1 := ioutil.ReadAll(resp.Body)
		if err1 != nil {
			p.Reject(err)
		}

		p.Resolve(string(body))
	}()

	p.SetTimeout(5000)
	result, err := p.Get()
	if err != nil {
		fmt.Println("err : ", err)
	}
	fmt.Println("Result is ", result)
}
