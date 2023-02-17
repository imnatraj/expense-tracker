package helper

import (
	"encoding/json"
	"errors"
	"net/http"

	"imnatraj/expense-tracker/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func Init() {
	Validate = validator.New()
}

type request struct {
	*http.Request
}

func Req(w *http.Request) *request {
	return &request{w}
}

// request param
// this will return null is no param exist
func (r *request) P(name string) (result string) {
	result = chi.URLParam(r.Request, name)
	return
}

// request query param
// this will return null if no query param exist
func (r *request) Q(name string) (result string) {
	res := r.URL.Query().Get(name)
	if len(res) > 0 {
		result = res
	}
	return
}

// get request body, returns error if struct not mapped properly
func (r *request) B(body interface{}) error {
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return err
	}
	err = Validate.Struct(body)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		var errStr string
		for _, err := range err.(validator.ValidationErrors) {
			errStr += err.Error() + ", "
		}
		return errors.New(errStr)
	}
	return nil
}

func (r *request) MetaData() (*models.MetaData, error) {
	ctx := r.Context()
	sessionUser, ok := ctx.Value(models.RequestDataMetaData).(*models.MetaData)
	if !ok {
		return nil, errors.New("invalid meta data")
	}
	return sessionUser, nil
}
