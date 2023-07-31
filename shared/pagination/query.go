package pagination

import "net/http"

func ParseQueryParams(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}
