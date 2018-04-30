package selfupdate

import (
	"github.com/silinternational/speed-snitch-agent"
	"os"
	"strings"
	"os/exec"
	"fmt"
)

const DefaultFileMode = 0755

// UpdateIfNeeded checks current version and config version and if different downloads the version from config
// If returns true, update occurred and process should be restarted, if false, check err, but if err is nil all is okay
func UpdateIfNeeded(currentVersion, configVersion, downloadURL string) (bool, error) {

	if currentVersion != configVersion {
		wd, _ := os.Getwd()
		downloadFile := getFilenameFromURL(downloadURL) + `-new`
		err := agent.DownloadFile(wd+"/"+downloadFile, downloadURL, DefaultFileMode)
		if err != nil {
			return false, err
		}

		execFilePath := wd+"/"+agent.ExeFileName

		cmd := exec.Command("cp", "-f", wd+"/"+downloadFile, execFilePath)
		err = cmd.Run()
		if err != nil {
			return true, fmt.Errorf("Error copying new version of file: %s\n\t%s", execFilePath, err.Error())
		}

		cmd = exec.Command("chown", "pi:pi", execFilePath)
		err = cmd.Run()
		if err != nil {
			return true, fmt.Errorf("Error changing ownership of new version of executable: %s\n\t%s", execFilePath, err.Error())
		}

		return true, nil
	}

	return false, nil
}

// getFilenameFromURL returns just the last part of a url after the last slash
func getFilenameFromURL(URL string) string {
	parts := strings.Split(URL, "/")
	return parts[len(parts)-1]
}
