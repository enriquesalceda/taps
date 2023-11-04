package domain

import "errors"

type FieldCensus struct {
	ID            string
	FirstLastName string
	LastLastName  string
	FirstName     string
	DOB           string
	State         string
	Gender        string
	Number        int
}

func (c FieldCensus) Validate() error {
	validationErrors := []error{}

	if c.ID == "" {
		validationErrors = append(validationErrors, errors.New("No ID"))
	}

	if c.FirstLastName == "" {
		validationErrors = append(validationErrors, errors.New("No FirstLastName"))
	}

	if c.LastLastName == "" {
		validationErrors = append(validationErrors, errors.New("No LastLastName"))
	}

	if c.FirstName == "" {
		validationErrors = append(validationErrors, errors.New("No FirstName"))
	}

	if c.DOB == "" {
		validationErrors = append(validationErrors, errors.New("No DOB"))
	}

	if c.Gender == "" {
		validationErrors = append(validationErrors, errors.New("No Gender"))
	}

	if c.State == "" {
		validationErrors = append(validationErrors, errors.New("No State"))
	}

	if c.Number == 0 {
		validationErrors = append(validationErrors, errors.New("No Number"))
	}

	if len(validationErrors) > 0 {
		return errors.Join(validationErrors...)
	}

	return nil
}
