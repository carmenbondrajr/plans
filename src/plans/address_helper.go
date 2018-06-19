package plans

import (
	"strings"
)

// This implementation requires the address string to end with the zip code.
// You could use a Regex here as well, but the 'perfect' Zip code regex is not trivial.
func GetZipFromAddress(address string) string {
	lastSpace := strings.LastIndex(address, " ") + 1
	zipCode := address[lastSpace:]
	lastDash := strings.LastIndex(zipCode, "-")

	if lastDash == -1 {
		return zipCode
	}

	return zipCode[0:lastDash]
}
