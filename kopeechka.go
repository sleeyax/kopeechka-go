package kopeechka

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
)

const (
	apiUrl     = "https://api.kopeechka.store"
	apiVersion = "2.0"
)

type Api struct {
	// Your API key from kopeechka.store.
	token string

	// The response type you want to receive from the API.
	// Defaults to Json.
	responseType ResponseType

	client http.Client
}

func New(token string) *Api {
	return &Api{client: http.Client{}, token: token, responseType: Json}
}

func (api *Api) SetResponseType(responseType ResponseType) {
	api.responseType = responseType
}

func (api *Api) SetToken(token string) {
	api.token = token
}

func (api *Api) getBaseUrl(endpoint string, data interface{}) (string, error) {
	v, err := query.Values(data)
	if err != nil {
		return "", err
	}

	v.Set("api", apiVersion)
	v.Set("type", string(api.responseType))
	v.Set("token", api.token)

	return fmt.Sprintf("%s/%s?%s", apiUrl, endpoint, v.Encode()), nil
}

func (api *Api) GetBalance() (BalanceResponse, error) {
	url, err := api.getBaseUrl("user-balance", nil)
	if err != nil {
		return BalanceResponse{}, err
	}

	res, err := api.client.Get(url)
	if err != nil {
		return BalanceResponse{}, err
	}
	defer res.Body.Close()

	var balanceResponse BalanceResponse
	if err = json.NewDecoder(res.Body).Decode(&balanceResponse); err != nil {
		return BalanceResponse{}, err
	}

	return balanceResponse, nil
}

func (api *Api) OrderMail(request OrderMailRequest) (OrderMailResponse, error) {
	if request.Site == "" {
		return OrderMailResponse{}, MissingRequiredParameterError
	}

	url, err := api.getBaseUrl("mailbox-get-email", request)

	res, err := api.client.Get(url)
	if err != nil {
		return OrderMailResponse{}, err
	}
	defer res.Body.Close()

	var response OrderMailResponse
	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return OrderMailResponse{}, err
	}

	return response, nil
}

func (api *Api) GetMessage(request MessageRequest) (MessageResponse, error) {
	if request.OrderId == "" {
		return MessageResponse{}, MissingRequiredParameterError
	}

	url, err := api.getBaseUrl("mailbox-get-message", request)

	res, err := api.client.Get(url)
	if err != nil {
		return MessageResponse{}, err
	}
	defer res.Body.Close()

	var response MessageResponse
	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return MessageResponse{}, err
	}

	return response, nil
}

func (api *Api) CancelMail(request CancelMailRequest) (CancelMailResponse, error) {
	if request.OrderId == "" {
		return CancelMailResponse{}, MissingRequiredParameterError
	}

	url, err := api.getBaseUrl("mailbox-cancel", request)

	res, err := api.client.Get(url)
	if err != nil {
		return CancelMailResponse{}, err
	}
	defer res.Body.Close()

	var response CancelMailResponse
	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return CancelMailResponse{}, err
	}

	return response, nil
}

// TODO: Re-order mail (/Reorder mail)
// TODO: Find orderId from mail (/Find ID activation of mail)
// Docs: https://faq.kopeechka.store/api_page/
