package api

import (
	"fmt"
	"errors"
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	//	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type FizzBuzzRequest struct {
	Int1  int
	Int2  int
	Limit int
	Str1  string
	Str2  string
}

func (fbr FizzBuzzRequest) Validate() error {
	return validation.ValidateStruct(&fbr,
		validation.Field(&fbr.Int1, validation.Required),
		validation.Field(&fbr.Int2, validation.Required),
		validation.Field(&fbr.Limit, validation.Required),
		validation.Field(&fbr.Str1, validation.Required, validation.Length(1, 100)),
		validation.Field(&fbr.Str2, validation.Required, validation.Length(1, 100)),
	)
}

// String represents the 
func (fbr FizzBuzzRequest) String() string {
	return fmt.Sprintf("int1 : %v , str1 : '%v' / int2 : %v , str2 : '%v' / limit : %v", fbr.Int1,fbr.Str1,fbr.Int2,fbr.Str2,fbr.Limit)
}

func createFizzBuzzRequest(r *http.Request) (*FizzBuzzRequest, error) {
	querystring := r.URL.Query()
	var err error

	errz := make([]error, 0)

	i1, err := strconv.Atoi(querystring.Get("int1"))

	if err != nil {
		errz = append(errz, errors.New("int1 is not an integer"))
	}

	i2, err := strconv.Atoi(querystring.Get("int2"))
	if err != nil {
		errz = append(errz, errors.New("int2 is not an integer"))
	}

	limit, err := strconv.Atoi(querystring.Get("limit"))
	if err != nil {
		errz = append(errz, errors.New("limit is not an integer"))
	}

	if len(errz) > 0 {
		var msg string = ""
		for _, e := range errz {
			msg = msg + " " + e.Error()
		}

		return nil, errors.New(msg)
	}

	fbr := FizzBuzzRequest{
		Int1:  i1,
		Int2:  i2,
		Limit: limit,
		Str1:  querystring.Get("str1"),
		Str2:  querystring.Get("str2"),
	}

	return &fbr, nil
}
