package session

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
)

type Session struct {
	CustomerId string `json:"customerID"`
	Ticket string `json:"ticket"`
}

type sessionRepository struct {
	sessions []Session
}

func (r *sessionRepository) GetSession(customerId string) *Session {
	lock.Lock()
	defer lock.Unlock()
	for _, session := range r.sessions {
		if session.CustomerId == customerId {
			return &session
		}
	}
	return nil
}

func (r *sessionRepository) AddSession(session Session) {
	r.sessions = append(r.sessions, session)
}

func (r *sessionRepository) init() {
	r.sessions = []Session{}
}

var (
	lock = sync.Mutex{}
	instance *sessionRepository
)

func FetchSessionRepository() *sessionRepository {
	if instance == nil {
		lock.Lock()
		instance = &sessionRepository{}
		instance.init()
		defer lock.Unlock()
	}
	return instance
}

func GenerateTicket(customerId string) (string, error) {
	u := make([]byte, 16)
	_, err := rand.Read(u)
	if err != nil {
		return "", err
	}

	u[8] = (u[8] | 0x80) & 0xBF // what does this do?
	u[6] = (u[6] | 0x40) & 0x4F // what does this do?

	ticket := hex.EncodeToString(u)

	FetchSessionRepository().AddSession(Session{
		CustomerId: customerId,
		Ticket: ticket,
	})

	return ticket, nil
}
