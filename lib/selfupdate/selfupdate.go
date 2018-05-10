package selfupdate

import (
	"github.com/silinternational/speed-snitch-agent"
	"os"
	"strings"
	"os/exec"
	"fmt"
)

const DefaultFileMode = 0755
const SignedFileSuffix = ".sig"

// UpdateIfNeeded checks current version and config version and if different downloads the version from config
// If returns true, update occurred and process should be restarted, if false, check err, but if err is nil all is okay
func UpdateIfNeeded(
	currentVersion, configVersion, downloadURL string,
	verifySignature bool,
) (bool, error) {

	if currentVersion != configVersion {
		wd, _ := os.Getwd()
		downloadFileBase := getFilenameFromURL(downloadURL)
		downloadFile := downloadFileBase  + `-new`
		err := agent.DownloadFile(wd+"/"+downloadFile, downloadURL, DefaultFileMode)
		if err != nil {
			return false, err
		}

		if verifySignature {
			signedDownloadURL := downloadURL + SignedFileSuffix
			signedFile := downloadFileBase + `-new.sig`
			err := agent.DownloadFile(wd+"/"+signedFile, signedDownloadURL, DefaultFileMode)
			if err != nil {
				return false, fmt.Errorf("Error downloading signed file: %s", err.Error())
			}

			err = agent.VerifyFileSignature(wd, downloadFile, signedFile)
			if err != nil {
				return false, fmt.Errorf("Error verifying the binary's signature: %s", err.Error())
			}
		}

		execFilePath := wd+"/"+agent.ExeFileName

		cmd := exec.Command("cp", "-f", wd+"/"+downloadFile, execFilePath)
		err = cmd.Run()
		if err != nil {
			return true, fmt.Errorf("Error copying new version of file: %s\n\t%s", execFilePath, err.Error())
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


