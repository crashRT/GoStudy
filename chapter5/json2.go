package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Mydata struct {
	Name string
	Mail string
	Tel  string
}

func (m *Mydata) Str() string {
	return "<\"" + m.Name + "\" " + m.Mail + " " + m.Tel + ">"
}

func main() {

	p := "https://tuyano-dummy-data.firebaseio.com/mydata.json"

	re, er := http.Get(p)
	if er != nil {
		fmt.Println(er)
		return
	}
	defer re.Body.Close()

	s, er := io.ReadAll(re.Body)
	if er != nil {
		fmt.Println(er)
		return
	}

	var items []Mydata
	er = json.Unmarshal(s, &items)
	if er != nil {
		fmt.Println(er)
		return
	}

	for i, im := range items {
		fmt.Println(i, im.Str())
	}

}
