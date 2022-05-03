// handlers package handles incoming REST Calls
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/luckyparakh/urlShortner/constants"
	"github.com/luckyparakh/urlShortner/storage/file"
	"github.com/luckyparakh/urlShortner/storage/memory"
)

var memObj = memory.NewUrlInfos()

// ListAllShortLinkHandler parses incoming request and
// returns list of all URLs in JSON format.
func ListAllShortLinkHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all URLs")
	w.Header().Set("content-type", "application/json")
	// Get List of All URLs
	json.NewEncoder(w).Encode(memObj.GetAllUrls())
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
	go memObj.IsLinkPresent(&info)

	//Read from channel, return link if already present else creates new shortlink and then return new link.
	for {
		if i := <-memObj.Ch; i[0] == "found" {
			fmt.Println("Infinite loop", i)
			shortLink := constants.Scheme + "://" + constants.Prefix + "/" + i[1]
			json.NewEncoder(w).Encode(shortLink)
			return
		} else {
			fmt.Println("Infinite loop else", i)
			json.NewEncoder(w).Encode(memObj.SaveLink(info))
			return
		}
	}

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
