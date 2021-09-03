package server

import (
	"path"
	"strings"
)

// shiftPath splits path on the next /
// e.g. "api/user/password" => "api", "user/password"
func shiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}
