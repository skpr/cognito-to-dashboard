package federation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
)

func GetSignInLink(credentials aws.Credentials, issuer, dashboardURL string, duration int) (string, error) {
	// Get console federated login link
	federationURL := url.URL{
		Scheme: "https",
		Host:   "signin.aws.amazon.com",
		Path:   "/federation",
	}

	sessionParams := map[string]string{
		"sessionId":    credentials.AccessKeyID,
		"sessionKey":   credentials.SecretAccessKey,
		"sessionToken": credentials.SessionToken,
	}

	jsonParams, err := json.Marshal(sessionParams)
	if err != nil {
		return "", fmt.Errorf("failed marshalling session params: %w", err)
	}

	query := federationURL.Query()
	query.Add("Action", "getSigninToken")
	query.Add("SessionDuration", strconv.Itoa(duration))
	query.Add("Session", string(jsonParams))
	federationURL.RawQuery = query.Encode()

	response, err := http.Get(federationURL.String())
	if err != nil {
		return "", fmt.Errorf("failed getting federation URL: %w", err)
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed reading response body: %w", err)
	}

	data := map[string]string{}
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return "", fmt.Errorf("failed unmarshalling response body: %w", err)
	}

	signInToken := data["SigninToken"]

	federationURL = url.URL{
		Scheme: "https",
		Host:   "signin.aws.amazon.com",
		Path:   "/federation",
	}

	query = federationURL.Query()
	query.Add("Action", "login")
	query.Add("Issuer", issuer)
	query.Add("Destination", dashboardURL)
	query.Add("SigninToken", signInToken)

	federationURL.RawQuery = query.Encode()

	return federationURL.String(), nil
}
