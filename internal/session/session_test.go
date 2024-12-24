package session

import (
	"testing"
)

func TestGenerateTicket(t *testing.T) {
	customerId := "customer123"
	ticket, err := GenerateTicket(customerId)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(ticket) == 0 {
		t.Errorf("expected a ticket to be generated, got an empty string")
	}

	repo := FetchSessionRepository()
	session := repo.GetSession(customerId)
	if session == nil {
		t.Fatalf("expected session to be found for customerId %s, got nil", customerId)
	}

	if session.Ticket != ticket {
		t.Errorf("expected ticket %s, got %s", ticket, session.Ticket)
	}
}

func TestGenerateTicketEmptyCustomerId(t *testing.T) {
	customerId := ""
	ticket, err := GenerateTicket(customerId)
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if ticket != "" {
		t.Errorf("expected an empty ticket, got %s", ticket)
	}
}

func TestFetchSessionRepository(t *testing.T) {
	repo1 := FetchSessionRepository()
	repo2 := FetchSessionRepository()

	if repo1 != repo2 {
		t.Errorf("expected the same instance of sessionRepository, got different instances")
	}
}

func TestAddAndGetSession(t *testing.T) {
	repo := FetchSessionRepository()
	customerId := "customer456"
	ticket := "ticket456"

	repo.AddSession(Session{
		CustomerId: customerId,
		Ticket: ticket,
	})

	session := repo.GetSession(customerId)
	if session == nil {
		t.Fatalf("expected session to be found for customerId %s, got nil", customerId)
	}

	if session.Ticket != ticket {
		t.Errorf("expected ticket %s, got %s", ticket, session.Ticket)
	}
}