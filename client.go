package main

import (
	"github.com/xanzy/go-gitlab"
)

type Client struct {
	base     *gitlab.Client
	Releases *ReleasesService
}

func NewClient(token string, options ...gitlab.ClientOptionFunc) (*Client, error) {
	client, err := gitlab.NewClient(token, options...)
	if err != nil {
		return nil, err
	}
	return newClient(client), nil
}

func newClient(base *gitlab.Client) *Client {
	client := &Client{
		base: base,
	}
	client.Releases = newRelease(client)
	return client
}
