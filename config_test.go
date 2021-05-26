package gpwebpay

import (
	"testing"
)

func TestValidate(t *testing.T) {

	config := Config{}
	if validationError := config.validate(); validationError != nil {
		// This won't be true as soon as `validate` implements an actual validation;
		// this check is here just to get us started with writing tests.
		t.Error("Config is always valid. Expected nil, got", validationError)
	}

}
