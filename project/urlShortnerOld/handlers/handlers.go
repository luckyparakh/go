// handlers package handles incoming REST Calls
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/luckyparakh/urlShortner/storage/file"
	"github.com/luckyparakh/urlShortner/storage/memory"
)

// ListAllShortLinkHandler parses incoming request and
// returns list of all URLs in JSON format.
func ListAllShortLinkHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all URLs")
	w.Header().Set("content-type", "application/json")
	// Get List of All URLs
	json.NewEncoder(w).Encode(memory.GetAllUrls())
}

// ListAllShortLinkFileHandler parses incoming request and
// returns list of all URLs in JSON format by reading a JSON file.
func ListAllShortLinkFileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all URLs from file")
	w.Header().Set("content-type", "application/json")
	// Get List of All URLs
	json.NewEncoder(w).Encode(file.GetAllUrlsFile())
}

// CreateShortLinkHandler parses incoming request
// and returns either existing link details or newly created link details.
func CreateShortLinkHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create short link")
	var info memory.UrlInfo
	w.Header().Set("content-type", "application/json")

	//Error Scenario, if body is nil.
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Please provide some data.")
		return
	}
	json.NewDecoder(r.Body).Decode(&info)

	//Error Scenario, if link provided by user is nil.
	if info.IsEmpty() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Url provided is empty.")
		return
	}

	// Parse link provided by user.
	// If link is not in correct URL format, then return.
	_, err := url.ParseRequestURI(info.Link)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Provide correct URL.")
		return
	}

	//If link is present then return same shortlink.
	isLinkPresent, shortLink := memory.IsLinkPresent(&info)
	if isLinkPresent {
		json.NewEncoder(w).Encode(shortLink)
		return
	}

	//Create new shortlink, if link is not present.
	json.NewEncoder(w).Encode(memory.SaveLink(info))
}

func CreateShortLinkFileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create short link from file")
	var info file.UrlInfo
	w.Header().Set("content-type", "application/json")

	//Error Scenario, if body is nil.
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Please provide some data.")
		return
	}
	json.NewDecoder(r.Body).Decode(&info)

	//Error Scenario, if link provided by user is nil.
	if info.IsEmpty() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Url provided is empty.")
		return
	}

	// Parse link provided by user.
	// If link is not in correct URL format, then return.
	_, err := url.ParseRequestURI(info.Link)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Provide correct URL.")
		return
	}

	//If link is present then return same shortlink.
	isLinkPresent, shortLink := file.IsLinkPresentFile(&info)
	if isLinkPresent {
		json.NewEncoder(w).Encode(shortLink)
		return
	}

	//Create new shortlink, if link is not present.
	json.NewEncoder(w).Encode(file.SaveLinkFile(info))
}
