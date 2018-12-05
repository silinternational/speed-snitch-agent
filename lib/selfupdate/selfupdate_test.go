package selfupdate

import "testing"

func TestGetFilenameFromURL(t *testing.T) {
	fixtures := []struct {
		URL      string
		Filename string
	}{
		{
			URL:      "https://domain.com/path/file",
			Filename: "file",
		},
		{
			URL:      "https://domain.com/file.sh",
			Filename: "file.sh",
		},
		{
			URL:      "https://windowssucks.com/file.exe",
			Filename: "file.exe",
		},
	}

	for _, fixture := range fixtures {
		filename := getFilenameFromURL(fixture.URL)
		if filename != fixture.Filename {
			t.Errorf("Filenames do not match from url (%s) and expected (%s)", filename, fixture.Filename)
			t.Fail()
		}
	}
}

func TestUpdateIfNeeded(t *testing.T) {
	// Test if update not needed
	updated, err := UpdateIfNeeded("1.0.0", "1.0.0", "https://dummy", false)
	if updated != false || err != nil {
		t.Errorf("UpdateIfNeeded did not respond as expected for equal versions")
		t.Fail()
	}

	// Test update needed but bad url
	updated, err = UpdateIfNeeded("1.0.0", "2.0.0", "https://dummy", false)
	if updated != false || err == nil {
		t.Errorf("UpdateIfNeeded should have returned an error for a bad url")
		t.Fail()
	}
}

func TestCopyFileOnLinux(t *testing.T) {
	sourcePath := "./testSourceFile"
	targetPath := "./testTargetFile"

	err := CopyFileOnLinux(sourcePath, targetPath)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
}
