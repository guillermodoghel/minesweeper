package rest

import "net/http"

func IsStatusCode2XX(sc int) bool {
	if sc >= http.StatusOK && sc < http.StatusMultipleChoices {
		return true
	}
	return false
}

func IsStatusCode3XX(sc int) bool {
	if sc >= http.StatusMultipleChoices && sc < http.StatusBadRequest {
		return true
	}
	return false
}

func IsStatusCode4XX(sc int) bool {
	if sc >= http.StatusBadRequest && sc < http.StatusInternalServerError {
		return true
	}
	return false
}

func IsStatusCode5XX(sc int) bool {
	return sc >= http.StatusInternalServerError
}
