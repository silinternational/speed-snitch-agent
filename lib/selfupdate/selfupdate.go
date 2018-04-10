package selfupdate

import (
	"github.com/silinternational/speed-snitch-agent"
	"os"
	"strings"
	"os/exec"
)

const DefaultFileMode = 0755

// UpdateIfNeeded checks current version and config version and if different downloads the version from config
// If returns true, update occurred and process should be restarted, if false, check err, but if err is nil all is okay
func UpdateIfNeeded(currentVersion, configVersion, downloadURL string) (bool, error) {

	if currentVersion != configVersion {
		wd, _ := os.Getwd()
		wd += "/"
		newFilename := getFilenameFromURL(downloadURL)
		err := agent.DownloadFile(wd+newFilename, downloadURL, DefaultFileMode)
		if err != nil {
			return false, err
		}

		cmd := exec.Command("cp", wd+newFilename, wd+agent.ExeFileName)
		cmd.Run()

		return true, nil
	}

	return false, nil
}

// getFilenameFromURL returns just the last part of a url after the last slash
func getFilenameFromURL(URL string) string {
	parts := strings.Split(URL, "/")
	return parts[len(parts)-1]
}

