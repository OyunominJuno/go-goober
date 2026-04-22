package api

// Define params for the API request
type CoinBalanceRequest struct {
	Username string
}

type CoinBalanceResponse struct {
	Code    int
	Balance int64
}

type ErrorResponse struct {
	Code    int
	Message string
}
