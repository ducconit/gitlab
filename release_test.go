package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestReleasesService_LatestRelease(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/releases",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, exampleReleaseListResponse)
		})

	release, _, err := client.Releases.LatestRelease(1, nil)
	if err != nil {
		t.Error(err)
	}
	if release.TagName != exampleLatestTagName {
		t.Errorf("expected tag %s, got %s", exampleLatestTagName, release.TagName)
	}
}
