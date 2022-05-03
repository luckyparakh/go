// memory package handles storage (memory) related operation.
package memory

import (
	"encoding/json"
	"math/rand"
	"time"
	"sync"
	"os"
	"fmt"

	"github.com/luckyparakh/urlShortner/constants"
	"github.com/luckyparakh/urlShortner/encodedecode"
)

// UrlInfo is struct which contains info related to URLs
type UrlInfo struct {
	Id     uint64 `json:"id"`
	Link   string `json:"link"`
	Encode string `json:"encode"`
}

type UrlInfos struct{
	urls []UrlInfo
	mu sync.Mutex
	Ch chan [2]string
}

// IsEmpty checks emptiness of link provided by user
func (u *UrlInfo) IsEmpty() bool {
	return u.Link == ""
}

var urlInfos []UrlInfo

func NewUrlInfos() *UrlInfos{
	ui := &UrlInfos{Ch:make(chan [2]string,10),urls:make([]UrlInfo,0)}
	ui.readAllUrlsFile()
	return ui
}

// IsLinkPresent takes Url as input and returns true and link details, if link exists
// else returns false and empty string.
func (uis *UrlInfos) IsLinkPresent(parUrl *UrlInfo) {
	fmt.Println("Check if link present")
	var output [2]string
	for _, urlInfo := range urlInfos {
		if urlInfo.Link == parUrl.Link {
			fmt.Println("Link present")
			output[0] = "found"
			output[1] = urlInfo.Encode
			uis.Ch <- output
			//uis.Ch <- true
			break
			//return true, constants.Scheme + "://" + constants.Prefix + "/" + urlInfo.Encode
		}
	}
	//uis.Ch <- false
	//uis.Ch <- [2]string{"", ""}
	output[0] = ""
	output[1] = ""
	uis.Ch <- output
}

// SaveLink saves Url info into slice and returns shortlink
func (uis *UrlInfos) SaveLink(parUrl UrlInfo) string {
	rand.Seed(time.Now().UnixNano())
	parUrl.Id = rand.Uint64()
	parUrl.Encode = encodedecode.EncodeLink(parUrl.Id)
	uis.mu.Lock()
	uis.urls = append(uis.urls, parUrl)
	uis.saveLinkFile()
	uis.mu.Unlock()
	return constants.Scheme + "://" + constants.Prefix + "/" + parUrl.Encode
}

// GetAllUrls returns details of all Urls present in slice
func (uis *UrlInfos) GetAllUrls() []UrlInfo {
	return uis.urls
}

// GetAllUrlsFile returns details of all Urls present in file
func (uis *UrlInfos) readAllUrlsFile(){
	var urlsFile []UrlInfo
	if bytedata, err := os.ReadFile("./resources/urlInfo.json"); err != nil {
		panic(err)
	} else if err := json.Unmarshal(bytedata, &urlsFile); err != nil && len(bytedata)>0 {
		panic(err)
	}
	for _,v := range urlsFile{
		fmt.Println(v)
		uis.urls = append(uis.urls, UrlInfo{v.Id,v.Link,v.Encode})
	}
}

// SaveLinkFile saves Url info into file
func (uis *UrlInfos) saveLinkFile() {
	f, err := os.Create("./resources/urlInfo.json")
	if err != nil {
		panic(err)
	}
	if byteData, err := json.Marshal(uis.urls); err != nil {
		f.Close()
		panic(err)
	} else if _, err := f.Write(byteData); err != nil {
		f.Close()
		panic(err)
	}
}

