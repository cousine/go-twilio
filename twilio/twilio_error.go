package twilio

import (
	"encoding/json"
	"fmt"
	"io"
)

type TwilioError struct {
	Status   int    `json:"status"`
	Message  string `json:"message"`
	Code     int    `json:"code"`
	MoreInfo string `json:"more_info"`
}

func NewLocalTwilioError(message string) (twilioError TwilioError) {
	return TwilioError{
		Status:  422,
		Message: message,
		Code:    -1,
	}
}

func NewTwilioError(rawJson io.ReadCloser) (twilioError TwilioError) {
	twilioError = TwilioError{}

	jsonDecoder := json.NewDecoder(rawJson)
	err := jsonDecoder.Decode(&twilioError)
	if err != nil {
		return NewLocalTwilioError(err.Error())
	}

	return
}

func (twilioError TwilioError) Error() string {
	return fmt.Sprintf("go-twilio: Client: Exception %v: %v", twilioError.Code, twilioError.Message)
}
