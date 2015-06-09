package twilio

func NewTwilioClient(accountSid string, authToken string) (twilioClient *TwilioClient) {
	twilioClient = &TwilioClient{
		AccountSid: accountSid,
		AuthToken:  authToken,
	}

	return
}
