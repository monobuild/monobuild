package monobuild

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

func (configuration *buildConfiguration) configurationIsValid() bool {
	validate := validator.New()

	err := validate.Struct(configuration)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			logrus.Errorf("%s is %s", err.StructField(), err.ActualTag())
		}
		return false
	}
	return true
}
