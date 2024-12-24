package authlogin

import (
	"fmt"
	"net/http"

	"github.com/rustymotors/lotus/internal/account"
	"github.com/rustymotors/lotus/internal/session"
)


type AuthLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthLoginResponse struct {
	Valid      bool   `json:"valid"`
	Ticket     string `json:"ticket"`
	ReasonCode string `json:"reasonCode"`
	ReasonText string `json:"reasonText"`
	ReasonUrl  string `json:"reasonUrl"`
}

func (r AuthLoginResponse) String() string {
	switch r.Valid {
	case true:
		return "Valid=TRUE\nTicket=" + r.Ticket
	case false:
		return "reasonCode=" + r.ReasonCode + "\nreasonText=" + r.ReasonText + "\nreasonUrl=" + r.ReasonUrl
	}
	panic("Valid must be true or false")
}

func processAuthLogin(request AuthLoginRequest) AuthLoginResponse {
	response := invalidResponse("INV-100", "Oh Dear", "http://www.google.com")

	accountRepo := account.FetchUserAccountRepository()

	account, err := accountRepo.GetAccount(request.Username, request.Password)
	if err != nil {
		return response
	}
	ticket, err := session.GenerateTicket(account.CustomerID)
	if err != nil {
		return response
	}

	response = validResponse(ticket)
	return response
}

func HandleAuthLogin(r *http.Request, w http.ResponseWriter) {
	fmt.Println("AuthLogin")
	request := AuthLoginRequest{}
	request.Username = r.URL.Query().Get("username")
	request.Password = r.URL.Query().Get("password")

	response := processAuthLogin(request)

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "%v", response)

}


func validResponse(ticket string) AuthLoginResponse {
	return AuthLoginResponse{
		Valid:  true,
		Ticket: ticket,
	}
}

func invalidResponse(reasonCode, reasonText, reasonUrl string) AuthLoginResponse {
	return AuthLoginResponse{
		Valid:      false,
		ReasonCode: reasonCode,
		ReasonText: reasonText,
		ReasonUrl:  reasonUrl,
	}
}