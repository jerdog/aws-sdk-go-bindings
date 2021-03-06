package aws

import "errors"

// SessionInput contains the input to be passed to New
type SessionInput struct {
	region string
}

// NewSessionInput returns a new *SessionInput
func NewSessionInput(region string) (*SessionInput, error) {

	if len(region) == 0 {
		return nil, errors.New(ErrNoRegionProvided)
	}

	svc := new(SessionInput)

	svc.region = region

	return svc, nil

}
