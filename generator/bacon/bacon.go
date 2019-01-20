package bacon

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/friendsofgo/bacon-ipsum/generator"
)

// Bacon represent the filter to use on bacon-ipsum api
type baconIpsum struct {
	Type      TextType
	Paras     int
	Sentences int
	WithLorem bool
}

type TextType string

const (
	// AllMeat represent that with this type only generate text with meat
	AllMeat TextType = "all-meat"
	// MeatAndFiller for meat mixed with miscellaneous ‘lorem ipsum’ filler.
	MeatAndFiller TextType = "meat-and-filler"
)

func (t TextType) Validate() bool {
	validTypes := []TextType{
		AllMeat,
		MeatAndFiller,
	}

	for _, typ := range validTypes {
		if typ == t {
			return true
		}
	}
	return false
}

const apiUrl = "https://baconipsum.com/api"

// NewBaconIpsum create an a generator for type bacon ipsum
func NewBaconIpsum(t TextType, p, s int, withLorem bool) generator.Generator {
	return &baconIpsum{
		Type:      t,
		Paras:     p,
		Sentences: s,
		WithLorem: withLorem,
	}
}

func (b *baconIpsum) formatUrl() string {
	finalUrl := apiUrl

	finalUrl += fmt.Sprintf("?type=%s", b.Type)
	if b.Paras > 0 && b.Sentences <= 0 {
		finalUrl += fmt.Sprintf("&paras=%d", b.Paras)
	} else if b.Sentences > 0 {
		finalUrl += fmt.Sprintf("&sentences=%d", b.Sentences)
	}

	if b.WithLorem {
		finalUrl += fmt.Sprintf("&start-with-lorem=1")
	}
	return finalUrl
}

// GenerateText generate a text based on bacon-ipsum API
func (b *baconIpsum) GenerateText() string {
	req, err := http.NewRequest("GET", b.formatUrl(), nil)
	if err != nil {
		log.Fatal("Error trying connect to bacon-ipsum API")
	}

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		log.Fatal("Error trying to obtain a valid response")
	}
	defer resp.Body.Close()

	// The bacon api return response into a array
	var text []string

	if err := json.NewDecoder(resp.Body).Decode(&text); err != nil {
		log.Fatal("Error reading the body")
	}
	return strings.Join(text, "\n")
}
