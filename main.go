package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"log/slog"
	"net/http"
	// ...existing code...
)

func main() {
	http.HandleFunc("/", handleRoot)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	slog.Info("Request", "ip", r)

	visitorInfo, err := getVisitorInfo(ip)
	if err != nil {
		http.Error(w, "Failed to get visitor info", http.StatusInternalServerError)
		return
	}

	slog.Info("Visitor info", "visitorInfo", visitorInfo)

	// Assuming visitorInfo contains location data
	directionsMap := getDirectionsMap(visitorInfo.Asname)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<html><body><h1>Visitor Info</h1><p>%v</p><h2>Map</h2><img src='%s' /></body></html>", visitorInfo, directionsMap)
}

func getVisitorInfo(ip string) (VisitorInfo, error) {
	resp, err := http.Get("http://ip-api.com/json/" + ip)
	if err != nil {
		return VisitorInfo{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return VisitorInfo{}, err
	}

	var visitorInfo VisitorInfo
	err = json.Unmarshal(body, &visitorInfo)
	if err != nil {
		return VisitorInfo{}, err
	}

	return visitorInfo, nil
}

func getDirectionsMap(location string) string {
	// Replace with actual implementation to get map with directions
	return fmt.Sprintf("https://maps.googleapis.com/maps/api/staticmap?center=%s&zoom=14&size=400x400&key=YOUR_API_KEY", location)
}

type VisitorInfo struct {
	Query         string  `json:"query"`
	Status        string  `json:"status"`
	Continent     string  `json:"continent"`
	ContinentCode string  `json:"continentCode"`
	Country       string  `json:"country"`
	CountryCode   string  `json:"countryCode"`
	Region        string  `json:"region"`
	RegionName    string  `json:"regionName"`
	City          string  `json:"city"`
	District      string  `json:"district"`
	Zip           string  `json:"zip"`
	Lat           float64 `json:"lat"`
	Lon           float64 `json:"lon"`
	Timezone      string  `json:"timezone"`
	Offset        int     `json:"offset"`
	Currency      string  `json:"currency"`
	Isp           string  `json:"isp"`
	Org           string  `json:"org"`
	As            string  `json:"as"`
	Asname        string  `json:"asname"`
	Mobile        bool    `json:"mobile"`
	Proxy         bool    `json:"proxy"`
	Hosting       bool    `json:"hosting"`
}
