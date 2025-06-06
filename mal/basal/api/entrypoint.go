package api

import (
	"fmt"
	"regexp"
	"strings"
)

type EntryPoint struct {
	Rule string `json:"rule"`
}

func (ep EntryPoint) Domain() string {
	// "Host(`api.xpool.top`) \u0026\u0026 PathPrefix(`/api/cloud-hashing-apis/version`)"
	regex := regexp.MustCompile(`(?:([a-z0-9-]+|\*)\.)?([a-z0-9-]{1,61})\.([a-z0-9]{2,7})`)
	domain := regex.Find([]byte(ep.Rule))
	return string(domain)
}

func (ep EntryPoint) ExportedPath() string {
	regex := regexp.MustCompile(`/[a-z0-9-/]+`)
	path := regex.Find([]byte(ep.Rule))
	return string(path)
}

func (ep EntryPoint) PathPrefix() (string, error) {
	path := ep.ExportedPath()
	paths := strings.Split(path, "/")

	const leastPathLen = 3
	if len(paths) < leastPathLen {
		return "", fmt.Errorf("invalid exported path")
	}

	return strings.Join(paths[0:3], "/"), nil
}

func (ep EntryPoint) ServiceKey() (string, error) {
	path := ep.ExportedPath()
	paths := strings.Split(path, "/")

	const leastPathLen = 3
	if len(paths) < leastPathLen {
		return "", fmt.Errorf("invalid exported path")
	}

	return paths[2], nil
}

func (ep EntryPoint) Path() (string, error) {
	path := ep.ExportedPath()
	paths := strings.Split(path, "/")

	const leastPathLen = 3
	if len(paths) < leastPathLen {
		return "", fmt.Errorf("invalid exported path")
	}

	return fmt.Sprintf("/%v", strings.Join(paths[3:], "/")), nil
}
