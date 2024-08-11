package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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

	var data []interface{}
	er = json.Unmarshal(s, &data)
	if er != nil {
		fmt.Println(er)
		return
	}

	for i, im := range data {
		m := im.(map[string]interface{})
		fmt.Println(i, m["name"].(string), m["mail"].(string), m["tel"].(string))
	}

}
