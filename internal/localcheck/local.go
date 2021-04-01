package localcheck

import (
	"bufio"
	"debug/macho"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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
	executableName, err := getExecutableName(path)
	if err != nil {
		return Architectures{}, err
	}

	// binary file path
	executable := filepath.Join(path, "Contents", "MacOS", executableName)

	return getExecutableArchitectures(executable)
}

func getExecutableArchitectures(path string) (Architectures, error) {
	var (
		arch = Architectures{}
	)

	// file is a Mach-O universal file
	fat, err := macho.OpenFat(path)
	if err == nil {
		arch.LoadFromFat(fat.Arches)
		return arch, nil
	}

	// file is a Mach-O file
	f, err := macho.Open(path)
	if err == nil {
		arch.Load(f.Cpu)
		return arch, nil
	}

	// file is a text file
	if IsTextFile(path) {
		interpreter, ok := getInterpreterPath(path)
		if !ok {
			return arch, errors.New("unable to get executable path")
		}
		return getExecutableArchitectures(interpreter)
	}

	return arch, errors.New("unknown file type")
}

func getInterpreterPath(filename string) (path string, ok bool) {
	f, err := os.Open(filename)
	if err != nil {
		return "", false
	}
	defer f.Close()

	// read the first line of the file; ensure that it starts with Shebang
	reader := bufio.NewReader(f)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	if !strings.HasPrefix(line, "#!") {
		return "", false
	}

	line = line[2:] // skip Shebang
	if strings.HasPrefix(line, "/usr/bin/env") {
		line = line[13:] // skip logical path

		interpreter := strings.SplitN(line, " ", 2)[0]
		path, err := exec.LookPath(interpreter)
		if err != nil {
			return "", false
		}

		return path, true
	}

	path = strings.SplitN(line, " ", 2)[0]

	return path, true
}
