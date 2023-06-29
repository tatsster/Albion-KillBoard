package util

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

var paramPattern = regexp.MustCompile(`:([a-zA-Z][a-zA-Z0-9_]*)`)

type BaseURl string

// BuildURL create url for request
//
// The `path` argument can contains the path params.
// and the path param name in the `path` should follow the format
// :abc_xyz
//
// the `pathPararms` argument is map[string]string which contains
// key = path param name without colon leading for example:
// the param name in the `path` is :abc_xyz
// the key in `pathParams` argument should be `abc_xyz`
func (b BaseURl) BuildURL(path string, pathParams map[string]string, query url.Values) string {
	if len(pathParams) > 0 {
		// expand path params :param
		path = string(
			paramPattern.ReplaceAllFunc(
				[]byte(path), func(data []byte) []byte {
					if val, ok := pathParams[string(data[1:])]; ok {
						return []byte(val)
					}
					return data
				},
			),
		)
	}

	var queryString string
	if len(query) > 0 {
		queryString = fmt.Sprintf("?%s", query.Encode())
	}

	baseURL := strings.TrimLeft(string(b), "/")
	path = strings.Trim(path, "/")

	return fmt.Sprintf("%s/%s%s", baseURL, path, queryString)
}
