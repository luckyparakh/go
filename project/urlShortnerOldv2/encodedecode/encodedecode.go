// encodedecode package encode and decode links which helps in generation of shortlink
package encodedecode

import (
	"errors"
	"math"
	"strings"
)

const alphaNum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const alphaNumLen = uint64(len(alphaNum))

// EncodeLink takes randomly genereted ID and 
// returns its base62 encoded version.
func EncodeLink(id uint64) string {
	var encoded string

	// Base62 encoding
	for ; id > 0; id = id / alphaNumLen {
		encoded = encoded + string(alphaNum[id%alphaNumLen])
	}
	return encoded
}

// DecodeLink takes base 62 encode-coverprofile=coverage.outd string and 
// returns decoded version.
func DecodeLink(encoded string) (uint64, error) {
	var decoded uint64

	// Base62 decoding
	for index, val := range encoded {
		positionInAlphanum := strings.Index(alphaNum, string(val))
		if positionInAlphanum == -1 {
			return 0, errors.New("Invalid Character," + string(val))
		} else {
			decoded += uint64(positionInAlphanum) * uint64(math.Pow(float64(alphaNumLen), float64(index)))
		}
	}
	return decoded, nil
}
