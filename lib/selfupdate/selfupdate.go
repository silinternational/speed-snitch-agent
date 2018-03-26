package selfupdate

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fillup/semver"
	"github.com/silinternational/speed-snitch-agent"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const DefaultFileMode = 0755

// GetLatestVersion calls Github API to get lastest version
// Returns latest version as string and error or nil
func GetLatestVersion(repoURL string) (string, error) {
	client := http.Client{
		Timeout: time.Second * 30,
	}

	resp, err := client.Get(repoURL)
	if err != nil {
		log.Fatal("Unable to get latest version for self update, err: ", err)
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Unable to read all body from get latest version api response: ", err)
		return "", err
	}

	var latest GithubLatestRelease
	err = json.Unmarshal(body, &latest)
	if err != nil {
		log.Fatal("Unable to unmarshal body into GithubLatestRelease: ", err, string(body))
		return "", err
	}

	return latest.TagName, nil
}

func UpgradeIfNeeded(currentVersion, repoURL, platform, arch string) (bool, error) {
	latest, err := GetLatestVersion(repoURL)
	if err != nil {
		log.Print(err)
		return false, err
	}

	isNewer, err := semver.IsNewer(currentVersion, latest)
	if err != nil {
		log.Print(err)
		return false, err
	}

	if isNewer {
		downloadURL, _ := getDownloadUrlForBinaryVersion(repoURL, latest, platform, arch)
		filename, _ := os.Getwd()
		err = agent.DownloadFile(filename+"/downloaded.tar.gz", downloadURL, DefaultFileMode)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil
}

func getDownloadUrlForBinaryVersion(repoURL, version, platform, arch string) (string, error) {
	if repoURL == "" {
		return "", errors.New("repoURL is required")
	} else if version == "" {
		return "", errors.New("version is required")
	} else if platform == "" {
		return "", errors.New("platform is required")
	} else if arch == "" {
		return "", errors.New("arch is required")
	}
	return fmt.Sprintf("%s/raw/%s/dist/%s/%s/speedsnitch", repoURL, version, platform, arch), nil
}

type GithubLatestRelease struct {
	URL             string `json:"url"`
	AssetsURL       string `json:"assets_url"`
	UploadURL       string `json:"upload_url"`
	HTMLURL         string `json:"html_url"`
	ID              int    `json:"id"`
	TagName         string `json:"tag_name"`
	TargetCommitish string `json:"target_commitish"`
	Name            string `json:"name"`
	Draft           bool   `json:"draft"`
	Author          struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
	Prerelease  bool          `json:"prerelease"`
	CreatedAt   time.Time     `json:"created_at"`
	PublishedAt time.Time     `json:"published_at"`
	Assets      []interface{} `json:"assets"`
	TarballURL  string        `json:"tarball_url"`
	ZipballURL  string        `json:"zipball_url"`
	Body        string        `json:"body"`
}
