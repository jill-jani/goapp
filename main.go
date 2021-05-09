package main

import (
	"fmt" 
	"net/http" 
	"os"
	"html/template"
	"log" //for logging errors
)

//all page variables
type PageVars struct {
	Title               string;
	Phno                string;
	Valid               bool   `json:"valid"`
	Number              string `json:"number"`
	LocalFormat         string `json:"local_format"`
	InternationalFormat string `json:"international_format"`
	CountryPrefix       string `json:"country_prefix"`
	CountryCode         string `json:"country_code"`
	CountryName         string `json:"country_name"`
	Location            string `json:"location"`
	Carrier             string `json:"carrier"`
	LineType            string `json:"line_type"`
	Name                string `json:"name"`
	MainValues          Main1  `json:"main"`
}

type Main1 struct {
	Temp                     float64
	Pressure                 float64
	Humidity                 float64
	Temp_max                 float64
	Temp_min                 float64
}

func main() {
	http.HandleFunc("/",IndexHandler);
	http.HandleFunc("/Weather",WeatherHandler);
	http.HandleFunc("/NumVerify",NumHandler);
	http.HandleFunc("/NumInfo",NumInfo);
	http.HandleFunc("/WeatherInfo",WeatherInfo);

	fs := http.FileServer(http.Dir("static"));
    http.Handle("/static/", http.StripPrefix("/static/", fs));

	http.ListenAndServe(getPort(), nil);
}

func getPort() string {
	p := os.Getenv("PORT");
	if p!= "" {
		return ":" + p;
	}
	return ":8080";
}

func render(w http.ResponseWriter, tmpl string, pv PageVars) {
	tmpl = fmt.Sprintf("templates/%s",tmpl);

	t , err := template.ParseFiles(tmpl);

	if err != nil {
		log.Print("Template Parsing error: ", err);
	}

	err = t.Execute(w,pv);
	if err != nil {
		log.Print("Template executing error: ", err);
	}
}
