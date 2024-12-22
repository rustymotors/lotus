package authlogin

import (
	"fmt"
	"net/http"
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

func HandleAuthLogin(r *http.Request, w http.ResponseWriter) {
	fmt.Println("AuthLogin")
	request := AuthLoginRequest{}
	request.Username = r.URL.Query().Get("username")
	request.Password = r.URL.Query().Get("password")

	switch request.Username {
	case "admin":
		response := AuthLoginResponse{
			Valid:  true,
			Ticket: "1234567890",
		}
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%v", response)
	default:
		response := AuthLoginResponse{
			Valid:      false,
			Ticket:     "",
			ReasonCode: "INV-100",
			ReasonText: "Oh Dear",
			ReasonUrl:  "http://www.google.com",
		}
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%v", response)
	}
}