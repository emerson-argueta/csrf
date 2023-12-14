//go:build !go1.7
// +build !go1.7

package csrf

import "net/http"

func addNosurfContext(r *http.Request) *http.Request {
	return r
}
