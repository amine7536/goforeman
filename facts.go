package goforeman

import (
	"fmt"

	"github.com/amine7536/goforeman/context"
)

const factsBasePath = "api/hosts"

type FactsService interface {
	Get(context.Context, string, *ListOptions) (Facts, *Response, error)
}

type Facts map[string]map[string]interface{}

type rootFacts struct {
	Facts Facts `json:"results"`
}

type FactsServiceOp struct {
	client *Client
}

var _ FactsService = &FactsServiceOp{}

func (s *FactsServiceOp) Get(ctx context.Context, hostname string, opt *ListOptions) (Facts, *Response, error) {
	path := fmt.Sprintf("%s/%s/facts", factsBasePath, hostname)

	req, err := s.client.NewRequest(ctx, "GET", path, opt)
	if err != nil {
		return nil, nil, err
	}

	type respWithMeta struct {
		rootFacts
		ResponseMeta
	}

	root := new(respWithMeta)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, nil, err
	}

	if m := &root.ResponseMeta; m != nil {
		resp.Meta = m
	}

	return root.Facts, resp, nil
}
