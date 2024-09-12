package geo

import (
	"context"
)

// EndpointBuilderSuggest defines functions that build urls geosuggest
type EndpointBuilderSuggest interface {
	GeosuggestURL(string) string
}

type SuggestResponseParserFactory func() SuggestResponseParser

// Suggest Response
type SuggestResponseParser interface {
    Addresses() (*Addresses, error)
}

// HTTPGeosuggest has EndpointBuilderSuggest and SuggestResponseParserFactory
type HTTPGeosuggest struct {
	EndpointBuilderSuggest
	SuggestResponseParserFactory
}

// ReverseGeocode returns address for location
func (g HTTPGeosuggest) Suggest(text string) (*Addresses, error) {
	responseParser := g.SuggestResponseParserFactory()

	ctx, cancel := context.WithTimeout(context.TODO(), DefaultTimeout)
	defer cancel()

	type sugResp struct {
		a *Addresses
		e error
	}
	ch := make(chan sugResp, 1)

	go func(ch chan sugResp) {
		if err := response(ctx, g.GeosuggestURL(text), responseParser); err != nil {
			ch <- sugResp{
				a: nil,
				e: err,
			}
		}

		addr, err := responseParser.Addresses()
		ch <- sugResp{
			a: addr,
			e: err,
		}
	}(ch)

	select {
	case <-ctx.Done():
		return nil, ErrTimeout
	case res := <-ch:
		return res.a, res.e
	}
}

