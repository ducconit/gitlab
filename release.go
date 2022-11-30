package main

import (
	"github.com/xanzy/go-gitlab"
)

type ReleasesService struct {
	base   *gitlab.ReleasesService
	client *Client
}

func newRelease(client *Client) *ReleasesService {
	return &ReleasesService{
		base:   client.base.Releases,
		client: client,
	}
}

func (r *ReleasesService) LatestRelease(pid interface{}, opt *gitlab.ListReleasesOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Release, *gitlab.Response, error) {
	releases, resq, err := r.client.base.Releases.ListReleases(pid, opt, options...)
	if err != nil {
		return nil, nil, err
	}
	release := new(gitlab.Release)
	if len(releases) == 0 {
		return release, resq, nil
	}
	release = releases[0]
	return release, resq, nil
}
