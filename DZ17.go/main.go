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

		//fmt.Println(markCar, modelCar, photoCar)
		// Cars := Car{
		// 	Mark:  markCar,
		// 	Model: modelCar,
		// 	Photo: photoCar,
		// }

		b := Car{Mark: markCar, Model: modelCar, Photo: photoCar}

		//fmt.Fprintf(w, jsonstring)
		//fmt.Println(jsonstring)
		if markCar != "" {

			dataFromFile, _ := ioutil.ReadFile("data.json")
			addCar := []Car{}

			json.Unmarshal(dataFromFile, &addCar)
			addCar = append(addCar, b)
			jsonData, _ := json.Marshal(addCar)
			//jsonstring := string(jsonData)

			//открываем файл для работы с ним
			file, _ := os.Create("data.json")

			defer file.Close()

			//записываем в файл данные

			file.WriteString(string(jsonData))

			//fmt.Println(Cars)
			// fmt.Println(addCar)
			// fmt.Println(jsonData)
			// fmt.Println(jsonstring)
			// fmt.Println(Cars)
		}

	})
	http.ListenAndServe(":8080", nil)
}
