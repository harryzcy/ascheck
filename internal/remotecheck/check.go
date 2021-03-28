package remotecheck

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

const (
	sourceURL = "https://cdn.jsdelivr.net/gh/ThatGuySam/doesitarm@master/README.md"
)

var (
	pattern, _ = regexp.Compile(`\* \[(.*?)\]\((.*?)\) - (✅|✳️|⏹|🚫|🔶)`)
)

type appInfo struct {
	Name    string
	Website string
	Support string
}

// getInfo loads the list of reported app support from Does it ARM.
func getInfo() ([]appInfo, error) {
	resp, err := http.Get(sourceURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	matches := pattern.FindAllStringSubmatch(string(body), -1)

	var result []appInfo
	for _, match := range matches {
		info := appInfo{
			Name:    match[1],
			Website: match[2],
			Support: match[3],
		}
		result = append(result, info)
	}

	return result, nil
}
