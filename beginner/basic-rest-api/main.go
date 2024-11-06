package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Car struct {
	Name  string
	Model string
	Year  int
}

var Cars []Car

func GetAllCars(w http.ResponseWriter, r *http.Request) {
	cars := []Car{
		{
			Name:  "honda",
			Model: "civic",
			Year:  2024,
		},
		{
			Name:  "toyota",
			Model: "corolla",
			Year:  2022,
		},
	}

	for idx, car := range cars {
		Cars = append(Cars, car)
		fmt.Fprintf(w, "CAR %d: ", idx)
		fmt.Fprintf(w, "%s %s %d\n\n", car.Name, car.Model, car.Year)
	}
}

func CreateCar(w http.ResponseWriter, r *http.Request) {
	car := Car{
		Name:  "toyota",
		Model: "camry",
		Year:  2024,
	}

	b, err := json.Marshal(car)
	checkErr(err)

	body := bytes.NewBuffer(b)

	postURL := "https://jsonplaceholder.typicode.com/posts"
	resp, err := http.Post(postURL, "application/json", body)
	checkErr(err)
	defer resp.Body.Close()

	fmt.Fprintf(w, "Status received from server is: %s\n", resp.Status)
	fmt.Fprintf(w, "StatusCode received from server is: %d\n", resp.StatusCode)
	fmt.Fprintf(w, "Content Type received from Server is: %s\n", resp.Header["Content-Type"][0])
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", GetAllCars)
	mux.HandleFunc("/create", CreateCar)

	fmt.Println("Started server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
