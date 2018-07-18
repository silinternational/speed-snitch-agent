package selfupdate

import (
	"github.com/silinternational/speed-snitch-agent"
	"os"
	"strings"
	"fmt"
	"runtime"
	"os/exec"
)

const DefaultFileMode = 0755
const SignedFileSuffix = ".sig"

const WindowsOS = "windows"
const WindowsServiceUpdater = "updateSpeedsnitch.bat"

func CopyFileOnLinux(sourcePath, targetPath string) error {

	cmd := exec.Command("cp", "-f", sourcePath, targetPath)
	err := cmd.Run()

	return err
}

func VerifySignature(downloadURL, downloadFileBase, downloadFile, workingDir string) error {
	signedDownloadURL := downloadURL + SignedFileSuffix
	signedFile := downloadFileBase + `-new.sig`
	err := agent.DownloadFile(workingDir+"/"+signedFile, signedDownloadURL, DefaultFileMode)
	if err != nil {
		return fmt.Errorf("Error downloading signed file: %s", err.Error())
	}

	publicKeys := getPublicKeys()
	err = agent.VerifyFileSignature(workingDir, downloadFile, signedFile, publicKeys)
	if err != nil {
		return fmt.Errorf("Error verifying the binary's signature: %s", err.Error())
	}

	return nil
}

// UpdateIfNeeded checks current version and config version and if different downloads the version from config
// If returns true, update occurred and process should be restarted, if false, check err, but if err is nil all is okay
func UpdateIfNeeded(
	currentVersion, configVersion, downloadURL string,
	verifySignature bool,
) (bool, error) {

	if currentVersion == configVersion {
		return false, nil
	}

	opSys := runtime.GOOS

	wd, _ := os.Getwd()
	downloadFileBase := getFilenameFromURL(downloadURL)
	downloadFile := downloadFileBase  + `-new`
	downloadPath := wd+"/"+downloadFile

	err := agent.DownloadFile(downloadPath, downloadURL, DefaultFileMode)
	if err != nil {
		return false, err
	}

	if verifySignature {
		err := VerifySignature(downloadURL, downloadFileBase, downloadFile, wd)
		if err != nil {
			errRm := os.Remove(downloadPath)
			if errRm != nil {
				return false, fmt.Errorf("Error verifying signature and error removing bad upgrade file: %s\n%s\n%s", downloadPath, err, errRm)
			}
			return false, fmt.Errorf("Error verifying signature of upgrade file: %s\n%s", downloadPath, err)
		}
	}

	if opSys == WindowsOS {
		cmd := exec.Command(WindowsServiceUpdater, configVersion)
		err := cmd.Run()
		if err != nil {
			return false, fmt.Errorf("Error calling the Windows updater script: %s\n%s", WindowsServiceUpdater, err.Error())
		}
	} else {
		execFilePath := wd + "/" + agent.ExeFileName
		err = CopyFileOnLinux(downloadPath, execFilePath)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

// getFilenameFromURL returns just the last part of a url after the last slash
func getFilenameFromURL(URL string) string {
	parts := strings.Split(URL, "/")
	return parts[len(parts)-1]
}


