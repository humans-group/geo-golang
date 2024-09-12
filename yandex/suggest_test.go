package yandex_test

import (
	"os"
	"testing"

	"github.com/codingsince1985/geo-golang/yandex"
	"github.com/stretchr/testify/assert"
)

var tokenSuggest = os.Getenv("YANDEX_SUGGEST_API_KEY")

func TestYandexGeoSuggest(t *testing.T) {
	ts := testServer(responseSugFound)
	defer ts.Close()
	suggest := yandex.Geosuggest(tokenSuggest, "3", "en_US", "geo", ts.URL+"/")
	addresses, err := suggest.Suggest("Burj")
	assert.NoError(t, err)
	assert.True(t, len(*addresses) > 0)
	assert.True(t, (*addresses)[0].FormattedAddress=="Burji")
	assert.True(t, (*addresses)[1].FormattedAddress=="Downtown Dubai")
	assert.True(t, (*addresses)[2].FormattedAddress=="City of Borjomi")
}

func TestYandexGeoSuggestNoResult(t *testing.T) {
	ts := testServer(responseSugNotFound)
	defer ts.Close()
	suggest := yandex.Geosuggest(tokenSuggest, "3", "en_US", "geo", ts.URL+"/")
	addresses, err := suggest.Suggest("$@#@")
	assert.Nil(t, err)
	assert.Nil(t, addresses)
}

const ( 
	responseSugFound = `{
  "suggest_reqid": "1726095099605352-4131673751-suggest-maps-yp-2",
  "results": [
    {
      "title": {
        "text": "Burji",
        "hl": [
          {
            "begin": 0,
            "end": 4
          }
        ]
      },
      "subtitle": {
        "text": "Kano"
      },
      "tags": [
        "locality"
      ],
      "distance": {
        "value": 1609715.142,
        "text": "998.02 mi"
      }
    },
    {
      "title": {
        "text": "Downtown Dubai"
      },
      "subtitle": {
        "text": "Dubai"
      },
      "tags": [
        "district"
      ],
      "distance": {
        "value": 6562464.289,
        "text": "4068.73 mi"
      }
    },
    {
      "title": {
        "text": "City of Borjomi",
        "hl": [
          {
            "begin": 8,
            "end": 12,
            "type": "MISPRINT"
          }
        ]
      },
      "subtitle": {
        "text": "край Самцхе-Джавахети"
      },
      "tags": [
        "locality"
      ],
      "distance": {
        "value": 6359917.658,
        "text": "3943.15 mi"
      }
    }
  ]
}`
responseSugNotFound = `{
  "suggest_reqid": "1726095099605352-4131673751-suggest-maps-yp-2"
}`
)