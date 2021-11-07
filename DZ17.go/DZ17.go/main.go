package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type Car struct {
	Mark  string `json:"Mark"`
	Model string `json:"Model"`
	Photo string `json:
	"Photo"`
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		markCar := r.URL.Query().Get("mark")
		modelCar := r.URL.Query().Get("model")
		photoCar := r.URL.Query().Get("photo")

		b := Car{Mark: markCar, Model: modelCar, Photo: photoCar}
		if markCar != "" {

			dataFromFile, _ := ioutil.ReadFile("data.json")
			addCar := []Car{}

			json.Unmarshal(dataFromFile, &addCar)

			addCar = append(addCar, b)

			jsonData, _ := json.Marshal(addCar)

			file, _ := os.Create("data.json")

			defer file.Close()

			file.WriteString(string(jsonData))
		}

	})
	http.ListenAndServe(":8080", nil)
}
