package selfupdate

import (
	"github.com/silinternational/speed-snitch-agent"
	"os"
	"strings"
	"os/exec"
	"fmt"
	"golang.org/x/crypto/openpgp"
)

const DefaultFileMode = 0755
const GPGKeyFileName = "gpgPub.key"

// UpdateIfNeeded checks current version and config version and if different downloads the version from config
// If returns true, update occurred and process should be restarted, if false, check err, but if err is nil all is okay
func UpdateIfNeeded(currentVersion, configVersion, downloadURL, signedDownloadURL string, verifySignature bool) (bool, error) {

	if currentVersion != configVersion {
		wd, _ := os.Getwd()
		downloadFileBase := getFilenameFromURL(downloadURL)
		downloadFile := downloadFileBase  + `-new`
		err := agent.DownloadFile(wd+"/"+downloadFile, downloadURL, DefaultFileMode)
		if err != nil {
			return false, err
		}

		if verifySignature {
			signedFile := downloadFileBase + `-new.sig`
			err := agent.DownloadFile(wd+"/"+signedFile, signedDownloadURL, DefaultFileMode)
			if err != nil {
				return false, fmt.Errorf("Error downloading signed file: %s", err.Error())
			}

			err = VerifyUpdateSignature(wd, downloadFile, signedFile)
			if err != nil {
				return false, err
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

func VerifyUpdateSignature(directory, targetFile, signedFile string) error {
	keyFilePath := directory + "/" + GPGKeyFileName

	_, err := os.Stat(keyFilePath)
	if os.IsNotExist(err) {
		return nil
	}

	keyRingReader, err := os.Open(keyFilePath)
	if err != nil {
		return err
	}

	signature, err := os.Open(signedFile)
	if err != nil {
		return err
	}

	verificationTarget, err := os.Open(targetFile)
	if err != nil {
		return err
	}

	keyring, err := openpgp.ReadArmoredKeyRing(keyRingReader)
	if err != nil {
		return fmt.Errorf("Error Reading Armored Key Ring: %s", err.Error())
	}

	_, err = openpgp.CheckArmoredDetachedSignature(keyring, verificationTarget, signature)
	if err != nil {
		return err
	}

	return nil
}

