package snippetconverter

import (
	"encoding/json"
	"fmt"
	"net/http"

	customerrors "snippetapp/customerrors"
	snippetmodel "snippetapp/domain/snippet/model"
)

type Converter struct{}

func NewSnippetConverter() *Converter {
	return &Converter{}
}

func (c *Converter) ConvertCreateRequestToSnippet(request *http.Request) (*snippetmodel.CreateRequest, error) {
	decoder := json.NewDecoder(request.Body)
	params := &snippetmodel.CreateRequest{}
	err := decoder.Decode(params)
	if err != nil {
		return nil, customerrors.NewBadRequestError(fmt.Errorf("error parsing params %v", err))
	}
	return params, err
}
