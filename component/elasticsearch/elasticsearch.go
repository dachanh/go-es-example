package elasticsearch

import (
	"github.com/pkg/errors"

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
		return nil, errors.Wrap(err, "elasticsearch : failed to create client")
	}
	return &ClientES{
		client: client,
	}, nil
}

func (e *ClientES) CreateIndex(index string) {

}
