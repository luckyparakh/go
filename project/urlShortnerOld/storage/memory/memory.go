// memory package handles storage (memory) related operation.
package memory

import (
	"math/rand"
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

// IsEmpty checks emptiness of link provided by user
func (u *UrlInfo) IsEmpty() bool {
	return u.Link == ""
}

var urlInfos []UrlInfo

// IsLinkPresent takes Url as input and returns true and link details, if link exists
// else returns false and empty string.
func IsLinkPresent(parUrl *UrlInfo) (bool, string) {

	for _, urlInfo := range urlInfos {
		if urlInfo.Link == parUrl.Link {
			return true, constants.Scheme + "://" + constants.Prefix + "/" + urlInfo.Encode
		}
	}
	return false, ""
}

// SaveLink saves Url info into slice and returns shortlink
func SaveLink(parUrl UrlInfo) string {
	rand.Seed(time.Now().UnixNano())
	parUrl.Id = rand.Uint64()
	parUrl.Encode = encodedecode.EncodeLink(parUrl.Id)
	urlInfos = append(urlInfos, parUrl)
	return constants.Scheme + "://" + constants.Prefix + "/" + parUrl.Encode
}

// GetAllUrls returns details of all Urls present in slice
func GetAllUrls() []UrlInfo {
	//Read from file load in memory
	return urlInfos
}
