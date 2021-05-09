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

func NumHandler (w http.ResponseWriter, r *http.Request) {
	vars := PageVars{
		Title: "Verify Phone Number",
		Phno: "",
	}
	render(w, "numverify.html", vars);
}

func NumInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		//fmt.Println(r.FormValue("num"));
		http.Redirect(w, r, "/", http.StatusSeeOther);
		return;
	}
	//fmt.Println(r.FormValue("num"));

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
	lines := strings.Split(string(rawBytes), "\r\n");
	key := lines[0];

	phone := r.FormValue("num");

	safePhone := url.QueryEscape(phone);

	url := fmt.Sprintf("http://apilayer.net/api/validate?access_key=%s&number=%s", key,safePhone);

	// Make request
	req, err := http.NewRequest("GET", url, nil);
	if err != nil {
		log.Fatal("NewRequest: ", err);
		return;
	}

	// create a Client
	client := &http.Client{};

	// Send the request via client
	// Do sends an HTTP request and returns an HTTP response
	resp, err := client.Do(req);
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Close resp.Body when done reading
	// Defer the closing of the body
	defer resp.Body.Close();

	// Fill the record with the data from the JSON
	var record PageVars;

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err);
	}
	/*
	vars := PageVars{
		Title : "Verify Phone Number",
		Phno : phone,
		Valid : record.Valid,
		LocalFormat : record.LocalFormat,
		InternationalFormat : record.InternationalFormat,
		CountryPrefix : record.CountryPrefix,
		CountryCode : record.CountryCode,
		CountryName : record.CountryName,
		Location : record.Location,
		Carrier : record.Carrier,
		LineType : record.LineType,
	}
	*/
	render(w, "numverify.html", record);
}