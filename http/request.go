package http

import "strings"

type Request struct {
	method  string
	path    string
	version string
	headers map[string]string
	body    string
}

func ParseRequest(requestBuffer []byte) (*Request, error) {
	lines := strings.Split(string(requestBuffer), "\n")

	method, path, version, err := parseFirstLine(lines[0])
	if err != nil {
		return nil, err
	}

	body := "" // TODO

	headers := make(map[string]string) // TODO

	return &Request{
		method:  method,
		path:    path,
		version: version,
		headers: headers,
		body:    body,
	}, nil
}

func parseFirstLine(line string) (string, string, string, error) {
	parts := strings.Split(line, " ")

	return parts[0], parts[1], parts[2], nil
}
