package helpers

import (
	"net/http"
)

func ResponseHasError(response *http.Response) bool {
	switch response.StatusCode {
	case http.StatusOK:
		return false
	case http.StatusCreated:
		return false
	case http.StatusFound:
		return false
	case http.StatusNotModified:
		return false
	case http.StatusBadRequest:
		return true
	case http.StatusUnauthorized:
		return true
	case http.StatusNotFound:
		return true
	case http.StatusMethodNotAllowed:
		return true
	case 429: // Too many requests
		return true
	case http.StatusInternalServerError:
		return true
	case http.StatusServiceUnavailable:
		return true
	}

	return false
}
