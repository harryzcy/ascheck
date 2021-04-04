package localcheck

import (
	"os"
	"path/filepath"

	"howett.net/plist"
)

type executableDecoded struct {
	CFBundleExecutable string
}

func getExecutableName(path string) (string, error) {
	plistFile := filepath.Join(path, "Contents", "Info.plist")

	f, err := os.Open(plistFile)
	if err != nil {
		return "", err
	}
	defer f.Close()

	decoder := plist.NewDecoder(f)
	var plistDecoded executableDecoded
	err = decoder.Decode(&plistDecoded)
	if err != nil {
		return "", err
	}

	return plistDecoded.CFBundleExecutable, err
}
