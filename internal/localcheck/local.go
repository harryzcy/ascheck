package localcheck

import (
	"debug/macho"
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

func GetArchitectures(path string) (Architectures, error) {
	var (
		arch = Architectures{}
	)

	executable, err := getExecutableName(path)
	if err != nil {
		return arch, err
	}

	// binary file path
	binary := filepath.Join(path, "Contents", "MacOS", executable)

	fat, err := macho.OpenFat(binary)
	if err == nil {
		// file is Mach-O universal
		arch.LoadFromFat(fat.Arches)
	} else {
		// file is Mach-O
		f, err := macho.Open(binary)
		if err != nil {
			return arch, err
		}
		arch.Load(f.Cpu)
	}

	return arch, nil
}
