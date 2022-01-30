package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	ExcuteTestsIncome()

}

type jsonFile struct {
	Id       int     `json:"id"`
	Describe string  `json:"describe"`
	Value    float64 `json:"value"`
	Date     string  `json:"date"`
	Category string  `json:"category"`
}

func TestGet(endpoint string) string {
	response, err := http.Get(endpoint)
	if err != nil {
		//fmt.Println(err)
		log.Fatal(err)
		return response.Status
	}
	defer response.Body.Close()

	// resp, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	//log.Fatal(err)
	// 	return response.Status
	// }

	//fmt.Println("------ TestGet ------")
	//fmt.Println(string(resp))
	//fmt.Println("response Status : ", response.Status)
	//fmt.Println("response Headers : ", response.Header)
	//fmt.Println("response Body : ", response.Body)
	return response.Status
}

func TestPost(endpoint string, jf jsonFile) string {

	json_data, err := json.Marshal(jf)
	if err != nil {
		//fmt.Println(err)
		log.Fatal(err)
	}
	response, err := http.Post(endpoint, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		//fmt.Println(err)
		log.Fatal(err)
		return response.Status
	}
	defer response.Body.Close()

	var res jsonFile

	json.NewDecoder(response.Body).Decode(&res)
	// fmt.Println("------ TestPost ------")
	// fmt.Println(res)
	// fmt.Println("response Status : ", response.Status)
	// fmt.Println("response Headers : ", response.Header)
	// fmt.Println("response Body : ", response.Body)
	return response.Status
}

func TestDelete(endpoint string) string {

	response, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		//fmt.Println(err)
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(response)
	if err != nil {
		//fmt.Println(err)
		log.Fatal(err)
		return resp.Status
	}
	defer resp.Body.Close()
	// respBody, err := ioutil.ReadAll((resp.Body))
	// if err != nil {
	// 	fmt.Println(err)
	// 	log.Fatal(err)
	// 	return resp.Status
	// }

	// fmt.Println("------ TestDelete ------")
	// fmt.Println("response Status : ", resp.Status)
	// fmt.Println("response Headers : ", resp.Header)
	// fmt.Println("response Body : ", string(respBody))
	return resp.Status

}

func TestPut(endpoint string, jf jsonFile, id int) string {
	client := &http.Client{}
	jf.Id = id
	json_data, err := json.Marshal(jf)
	if err != nil {
		//fmt.Println(err)
		log.Fatal(err)
	}

	response, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(json_data))
	response.Header.Set("Content-Type", "application/json")
	response.ParseForm()
	if err != nil {
		//fmt.Println(err)
		log.Fatal(err)
	}

	resp, err := client.Do(response)
	if err != nil {
		//fmt.Println(err)
		log.Fatal(err)
		return resp.Status
	}

	defer resp.Body.Close()

	// respBody, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	log.Fatal(err)
	// 	return resp.Status
	// }

	// fmt.Println("------ TestPut ------")
	// fmt.Println("response Status : ", resp.Status)
	// fmt.Println("response Headers : ", resp.Header)
	// fmt.Println("response Body : ", string(respBody))
	return resp.Status

}

func ExcuteTestsIncome() {
	tests_ok := 0
	tests_fail := 0
	var response string
	var values jsonFile

	//---- TEST INCOME -----
	fmt.Println("-------- Income Tests --------")
	fmt.Println("")
	//	case method == "Get":
	response = TestGet("http://localhost:8000/api/receitas")
	if response != "200 OK" {
		tests_fail += 1
		log.Println("Test GET all fail, expected 200 get", response)
	} else {
		tests_ok += 1
		log.Println("Test GET all ok...")
	}
	response = TestGet("http://localhost:8000/api/receitas/2")
	if response != "200 OK" {
		tests_fail += 1
		log.Println("Test GET by id fail, expected 200 get", response)
	} else {
		tests_ok += 1
		log.Println("Test GET by id ok...")
	}
	response = TestGet("http://localhost:8000/api/receitas/0")
	if response != "404 Not Found" {
		tests_fail += 1
		log.Println("Test GET by id fail, expected 404 get", response)
	} else {
		tests_ok += 1
		log.Println("Test GET by id 'NOT FOUND' ok...")
	}

	//	case method == "Post":
	values = jsonFile{
		Id:       2,
		Describe: "teste01",
		Value:    15,
		Date:     "2002-01-01",
	}
	response = TestPost("http://localhost:8000/api/receitas", values)
	if response != "400 Bad Request" {
		tests_fail += 1
		log.Println("Test POST with ID fail, expected 400 get", response)
	} else {
		tests_ok += 1
		log.Println("Test POST with ID 'Bad Request' ok...")
	}

	values = jsonFile{
		Describe: "teste01",
		Value:    15,
		Date:     "2002-15-32",
	}
	response = TestPost("http://localhost:8000/api/receitas", values)
	if response != "400 Bad Request" {
		tests_fail += 1
		log.Println("Test POST with wrong date fail, expected 400 get", response)
	} else {
		tests_ok += 1
		log.Println("Test POST with wrong date 'Bad Request' ok...")
	}

	values = jsonFile{
		Describe: "TRE",
		Value:    15000,
		Date:     "2022-10-29",
	}
	response = TestPost("http://localhost:8000/api/receitas", values)
	if response != "409 Conflict" {
		tests_fail += 1
		log.Println("Test POST with same describe in month, expected 409 get", response)
	} else {
		tests_ok += 1
		log.Println("Test POST with same describe in month 'Conflict' ok...")
	}

	//	case method == "Delete":
	response = TestDelete("http://localhost:8000/api/receitas/1")
	if response != "404 Not Found" {
		tests_fail += 1
		log.Println("Test DELETE fail, expected 404 get", response)
	} else {
		tests_ok += 1
		log.Println("Test DELETE 'NOT FOUND' ok...")
	}

	//	case method == "Put":
	values = jsonFile{
		Describe: "TRE",
		Value:    15000,
		Date:     "2022-10-29",
	}
	response = TestPut("http://localhost:8000/api/receitas/1", values, 1)
	if response != "404 Not Found" {
		tests_fail += 1
		log.Println("Test EDIT fail, expected 404 get", response)
	} else {
		tests_ok += 1
		log.Println("Test EDIT 'NOT FOUND' ok...")
	}

	fmt.Printf("Total tests fails: %d", tests_fail)
	fmt.Println("Total tests ok: ", tests_ok)
}
