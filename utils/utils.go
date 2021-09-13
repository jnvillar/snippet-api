package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	customerrors "snippetapp/customerrors"
)

func ReturnJSONResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func ReturnJSONResponseWithStatus(w http.ResponseWriter, status int, response interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func ReturnError(w http.ResponseWriter, err error) {
	notFound := &customerrors.NotFound{}
	if errors.As(err, &notFound) {
		writeError(err.(*customerrors.NotFound).GetCode(), w, err)
		return
	}
	badRequest := &customerrors.BadRequest{}
	if errors.As(err, &badRequest) {
		writeError(err.(*customerrors.BadRequest).GetCode(), w, err)
		return
	}
	writeError(http.StatusInternalServerError, w, fmt.Errorf("internal server error"))
}

func writeError(status int, w http.ResponseWriter, err error) {
	w.WriteHeader(status)
	ReturnJSONResponse(w, &customerrors.ErrorResponse{Error: err.Error()})
}
