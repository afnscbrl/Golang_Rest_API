package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type jsonFile struct {
	Id       int     `json:"id"`
	Describe string  `json:"describe"`
	Value    float64 `json:"value"`
	Date     string  `json:"date"`
}

func testGet(endpoint string) bool {
	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return false
	}
	defer response.Body.Close()

	resp, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return false
	}

	fmt.Println(string(resp))
	//VERIFICAR O RESP.HEADER PRA VER SE RETORNA 200
	return true
}

func testPost(endpoint string, jf jsonFile) {
	//ADD RETORNO BOOL

	json_data, err := json.Marshal(jf)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	response, err := http.Post(endpoint, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer response.Body.Close()

	var res jsonFile

	json.NewDecoder(response.Body).Decode(&res)
	fmt.Println(res)
}

func testDelete(endpoint string) {

	response, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(response)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll((resp.Body))
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))

}

func testPut(endpoint string, jf jsonFile, id int) {
	client := &http.Client{}
	jf.Id = id
	json_data, err := json.Marshal(jf)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	response, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(json_data))
	response.Header.Set("Content-Type", "application/json")
	response.ParseForm()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	resp, err := client.Do(response)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))

}

func main() {
	values := jsonFile{
		Describe: "test_Put5",
		Value:    15,
		Date:     "01-06-1989",
	}

	method := os.Args[1]
	endp := os.Args[2]
	switch {
	case method == "Get":
		{
			testGet("http://localhost:8000/api/" + endp)
		}
	case method == "Post":
		{
			testPost("http://localhost:8000/api/"+endp, values)
		}
	case method == "Delete":
		{
			testDelete("http://localhost:8000/api/" + endp)
		}
	case method == "Put":
		{
			id := os.Args[3]
			int_id, _ := strconv.Atoi(id)
			testPut("http://localhost:8000/api/"+endp+"/"+id, values, int_id)
		}
	}

}
