package selfupdate

import (
	"fmt"
	"github.com/silinternational/speed-snitch-agent"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetLatestVersion(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	respJson := `{
	"url": "https://api.github.com/repos/silinternational/speed-snitch-agent/releases/6433121",
	"assets_url": "https://api.github.com/repos/silinternational/speed-snitch-agent/releases/6433121/assets",
	"upload_url": "https://uploads.github.com/repos/silinternational/speed-snitch-agent/releases/6433121/assets{?name,label}",
	"html_url": "https://github.com/silinternational/speed-snitch-agent/releases/tag/3.2",
	"id": 6433121,
	"tag_name": "3.2",
	"target_commitish": "master",
	"name": "Support multi-level releases",
	"draft": false,
	"author": {
		"login": "fillup",
		"id": 556105,
		"avatar_url": "https://avatars3.githubusercontent.com/u/556105?v=4",
		"gravatar_id": "",
		"url": "https://api.github.com/users/fillup",
		"html_url": "https://github.com/fillup",
		"followers_url": "https://api.github.com/users/fillup/followers",
		"following_url": "https://api.github.com/users/fillup/following{/other_user}",
		"gists_url": "https://api.github.com/users/fillup/gists{/gist_id}",
		"starred_url": "https://api.github.com/users/fillup/starred{/owner}{/repo}",
		"subscriptions_url": "https://api.github.com/users/fillup/subscriptions",
		"organizations_url": "https://api.github.com/users/fillup/orgs",
		"repos_url": "https://api.github.com/users/fillup/repos",
		"events_url": "https://api.github.com/users/fillup/events{/privacy}",
		"received_events_url": "https://api.github.com/users/fillup/received_events",
		"type": "User",
		"site_admin": false
	},
	"prerelease": false,
	"created_at": "2017-05-18T17:46:02Z",
	"published_at": "2017-05-18T17:46:42Z",
	"assets": [],
	"tarball_url": "https://api.github.com/repos/silinternational/speed-snitch-agent/tarball/3.2",
	"zipball_url": "https://api.github.com/repos/silinternational/speed-snitch-agent/zipball/3.2",
	"body": ""
}`

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, respJson)
	})

	expected := "3.2"
	version, _ := GetLatestVersion(server.URL)

	if version != expected {
		t.Errorf("Returned version, %s, not what was expected: %s", version, expected)
	}
}

func TestGetDownloadUrlForBinaryVersion(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	respJson := `{
	"url": "https://api.github.com/repos/silinternational/speed-snitch-agent/releases/6433121",
	"assets_url": "https://api.github.com/repos/silinternational/speed-snitch-agent/releases/6433121/assets",
	"upload_url": "https://uploads.github.com/repos/silinternational/speed-snitch-agent/releases/6433121/assets{?name,label}",
	"html_url": "https://github.com/silinternational/speed-snitch-agent/releases/tag/3.2",
	"id": 6433121,
	"tag_name": "3.2",
	"target_commitish": "master",
	"name": "Support multi-level releases",
	"draft": false,
	"author": {
		"login": "fillup",
		"id": 556105,
		"avatar_url": "https://avatars3.githubusercontent.com/u/556105?v=4",
		"gravatar_id": "",
		"url": "https://api.github.com/users/fillup",
		"html_url": "https://github.com/fillup",
		"followers_url": "https://api.github.com/users/fillup/followers",
		"following_url": "https://api.github.com/users/fillup/following{/other_user}",
		"gists_url": "https://api.github.com/users/fillup/gists{/gist_id}",
		"starred_url": "https://api.github.com/users/fillup/starred{/owner}{/repo}",
		"subscriptions_url": "https://api.github.com/users/fillup/subscriptions",
		"organizations_url": "https://api.github.com/users/fillup/orgs",
		"repos_url": "https://api.github.com/users/fillup/repos",
		"events_url": "https://api.github.com/users/fillup/events{/privacy}",
		"received_events_url": "https://api.github.com/users/fillup/received_events",
		"type": "User",
		"site_admin": false
	},
	"prerelease": false,
	"created_at": "2017-05-18T17:46:02Z",
	"published_at": "2017-05-18T17:46:42Z",
	"assets": [],
	"tarball_url": "https://api.github.com/repos/silinternational/speed-snitch-agent/tarball/3.2",
	"zipball_url": "https://api.github.com/repos/silinternational/speed-snitch-agent/zipball/3.2",
	"body": ""
}`

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, respJson)
	})

	latestVersion, _ := GetLatestVersion(server.URL)
	expected := "https://github.com/silinternational/speed-snitch-agent/raw/3.2/dist/linux/arm64/speedsnitch"
	binaryURL, _ := getDownloadUrlForBinaryVersion(agent.RepoURL, latestVersion, "linux", "arm64")
	if binaryURL != expected {
		t.Errorf("Download url not what was expected. Received %s, expected %s", binaryURL, expected)
	}
}
