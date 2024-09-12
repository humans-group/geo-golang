// Package yandex is a geo-golang based Yandex Maps Location API
package yandex

import (
	"fmt"

	"github.com/humans-group/geo-golang"
)

type (
	//suggest
	SuggestResponse struct {
		SuggestReqID string         `json:"suggest_reqid"`
		Results      []SuggestResult `json:"results"`
	}
	
	SuggestResult struct {
		Title    Title    `json:"title"`
		Subtitle Subtitle `json:"subtitle,omitempty"`
		Tags     []string `json:"tags"`
		Distance Distance `json:"distance"`
	}
	
	Title struct {
		Text string    `json:"text"`
		HL   []HLRange `json:"hl,omitempty"`
	}
	
	HLRange struct {
		Begin int `json:"begin"`
		End   int `json:"end"`
	}
	
	Subtitle struct {
		Text string `json:"text"`
	}
	
	Distance struct {
		Value float64 `json:"value"`
		Text  string  `json:"text"`
	}
)

// Geosuggest constructs Yandex geosuggest service
func Geosuggest(apiKey string, baseURLs ...string) geo.Geosuggest {
	return geo.HTTPGeosuggest{
		EndpointBuilderSuggest:       baseURL(getSuggestURL(apiKey, baseURLs...)),
		SuggestResponseParserFactory: func() geo.SuggestResponseParser { return &SuggestResponse{} },
	}
}

func getSuggestURL(apiKey string, baseURLs ...string) string {
	if len(baseURLs) > 0 {
		return baseURLs[0]
	}
	URL := fmt.Sprintf("https://suggest-maps.yandex.ru/v1/suggest?apikey=%s", apiKey)
	return URL
}

func (r *SuggestResponse) Addresses() (*geo.Addresses, error) {
	if (r.Results == nil) || (len(r.Results) == 0) {
		return nil, nil
	}
	var addrs geo.Addresses
	for _, result := range r.Results {
		if result.Title.Text != "" {
            addrs = append(addrs, geo.Address{FormattedAddress: result.Title.Text,})
        }
    }
	return &addrs, nil
}

