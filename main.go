package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

var locations = map[string]string{
	"Polska": "Warsaw",
	"Niemcy": "Berlin",
	"Francja": "Paris",
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<html><body><h1>Wybierz kraj</h1><form action='/weather' method='POST'>")
	fmt.Fprint(w, "<select name='country'>")
	for country := range locations {
		fmt.Fprintf(w, "<option value='%s'>%s</option>", country, country)
	}
	fmt.Fprint(w, "</select><br><br><input type='submit' value='Pokaż pogodę'></form></body></html>")
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	country := r.FormValue("country")
	city, exists := locations[country]
	if !exists {
		http.Error(w, "Nieprawidłowy wybór", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric",
		city, apiKey,
	)

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		http.Error(w, "Błąd pobierania danych pogodowych", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var weather WeatherResponse
	json.Unmarshal(body, &weather)

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<html><body><h1>Pogoda w %s</h1><p>Temperatura: %.1f°C</p><a href='/'>Powrót</a></body></html>", weather.Name, weather.Main.Temp)
}

func main() {
	author := "Bohdan Maikut"
	port := "8080"
	fmt.Printf("Aplikacja uruchomiona: %s\n", time.Now().Format(time.RFC1123))
	fmt.Printf("Autor: %s\n", author)
	fmt.Printf("Nasłuchiwanie na porcie: %s\n", port)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/weather", weatherHandler)
	http.ListenAndServe(":"+port, nil)
}
