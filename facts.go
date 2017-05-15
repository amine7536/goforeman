package goforeman

import (
	"fmt"

	"github.com/amine7536/goforeman/context"
)

const factsBasePath = "api/hosts"

type FactsService interface {
	Get(context.Context, string) (Facts, *Response, error)
}

type Facts map[string]map[string]interface{}

type rootFacts struct {
	Facts Facts `json:"results"`
}

type FactsServiceOp struct {
	client *Client
}

var _ FactsService = &FactsServiceOp{}

// func (d Facts) String() string {
// 	return Stringify(d)
// }

func (s *FactsServiceOp) Get(ctx context.Context, hostname string) (Facts, *Response, error) {

	path := fmt.Sprintf("%s/%s/facts", factsBasePath, hostname)

	req, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(rootFacts)
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, nil, err
	}

	return v.Facts, resp, nil
}
