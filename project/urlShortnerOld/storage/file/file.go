package file

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"

	"github.com/luckyparakh/urlShortner/constants"
	"github.com/luckyparakh/urlShortner/encodedecode"
)

// UrlInfo is struct which contains info related to URLs
type UrlInfo struct {
	Id     uint64 `json:"id"`
	Link   string `json:"link"`
	Encode string `json:"encode"`
}

var urlInfos []UrlInfo

// IsEmpty checks emptiness of link provided by user
func (u *UrlInfo) IsEmpty() bool {
	return u.Link == ""
}

// GetAllUrlsFile returns details of all Urls present in file
func GetAllUrlsFile() []UrlInfo {
	var urlsFile []UrlInfo
	if bytedata, err := os.ReadFile("./resources/urlInfo.json"); err != nil {
		return nil
	} else if err := json.Unmarshal(bytedata, &urlsFile); err != nil {
		panic(err)
	}
	return urlsFile
}

// IsLinkPresentFile takes Url as input and returns true and link details, if link exists
// else returns false and empty string.
func IsLinkPresentFile(parUrl *UrlInfo) (bool, string) {
	//var urlsFile []UrlInfo
	urlsFile := GetAllUrlsFile()

	for _, urlFile := range urlsFile {
		if urlFile.Link == parUrl.Link {
			return true, constants.Scheme + "://" + constants.Prefix + "/" + urlFile.Encode
		}
	}
	return false, ""
}

// SaveLinkFile saves Url info into file and returns shortlink
func SaveLinkFile(parUrl UrlInfo) string {
	rand.Seed(time.Now().UnixNano())
	parUrl.Id = rand.Uint64()
	parUrl.Encode = encodedecode.EncodeLink(parUrl.Id)
	urlInfos = append(urlInfos, parUrl)
	f, err := os.Create("./resources/urlInfo.json")
	if err != nil {
		panic(err)
	}
	if byteData, err := json.Marshal(&urlInfos); err != nil {
		f.Close()
		panic(err)
	} else if _, err := f.Write(byteData); err != nil {
		f.Close()
		panic(err)
	}
	return constants.Scheme + "://" + constants.Prefix + "/" + parUrl.Encode
}
