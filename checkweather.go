package main

import (
	"fmt"
	"os"
	"io/ioutil" //for reading keys
	"strings"
	"net/http"
	"net/url" //for api
	"encoding/json" //for parsing json
	"log" //for logging errors
)

func WeatherHandler (w http.ResponseWriter, r *http.Request) {
	vars := PageVars{
		Title: "Check Weather",
	}
	render(w, "weather.html", vars);
}

func WeatherInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		//fmt.Println(r.FormValue("num"));
		http.Redirect(w, r, "/", http.StatusSeeOther);
		return;
	}
	//fmt.Println(r.FormValue("num"));
	cityName := r.FormValue("city");
	safeCityName := url.QueryEscape(cityName);
	// reading key from keys.txt
	fileIO, err := os.OpenFile("keys.txt", os.O_RDWR, 0600);
	if err != nil {
		panic(err);
	}
	defer fileIO.Close()
	rawBytes, err := ioutil.ReadAll(fileIO);
	if err != nil {
		panic(err);
	}
	lines := strings.Split(string(rawBytes), "\n");
	key := lines[1];

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", safeCityName, key);

	// Build the request
	req, err := http.NewRequest("GET", url, nil);
	if err != nil {
		log.Fatal("NewRequest: ", err);
		return;
	}

	// Client
	client := &http.Client{};

	// Send the request via client
	resp, err := client.Do(req);
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	defer resp.Body.Close();

	// Fill the record with the data from the JSON
	var record PageVars;
	/*
	resp = string(resp.Body);
	json.Unmarshal([]byte(resp), &record);
	fmt.Println(record);
	*/
	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	render(w, "weather.html", record);
}