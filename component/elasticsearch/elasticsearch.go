package elasticsearch

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

type ClientES struct {
	client *elasticsearch.Client
	index  string
	alias  string
}

func NewClient(address string) (*ClientES, error) {
	config := elasticsearch.Config{
		Addresses: []string{
			address,
		},
	}

	client, err := elasticsearch.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("elasticsearch : failed to create client; %s", err.Error())
	}
	return &ClientES{
		client: client,
	}, nil
}

func (e *ClientES) CreateIndex(index string) error {
	e.index = index
	e.alias = index + "_alias"

	responseCheckIndexExitst, err := e.client.Indices.Exists([]string{e.index})
	defer responseCheckIndexExitst.Body.Close()
	if err != nil {
		return fmt.Errorf("eleasticsearch: error create index %w", err)
	}
	if responseCheckIndexExitst.StatusCode == 200 {
		return nil
	}
	if responseCheckIndexExitst.StatusCode == 404 {
		return fmt.Errorf("eleasticsearch: error in index existence response: %s", responseCheckIndexExitst.String())
	}
	responseCreateIndex, err := e.client.Indices.Create(e.index)

	if err != nil {
		return fmt.Errorf("eleasticsearch: cannot create index: %w", err)
	}
	if responseCreateIndex.IsError() {
		return fmt.Errorf("eleasticsearch: responseCreateIndex Error %s", responseCreateIndex.String())
	}
	responsePutAlias, err := e.client.Indices.PutAlias([]string{e.index}, e.alias)

	if err != nil {
		return fmt.Errorf("eleasticsearch: cannot Create Alias: %w", err)
	}
	if responsePutAlias.IsError() {
		return fmt.Errorf("eleasticsearch: Error reponse put alias %s", responsePutAlias.String())
	}
	return nil
}
