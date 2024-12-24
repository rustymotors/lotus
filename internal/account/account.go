package account

import (
	"errors"
	"sync"
)

type UserAccount struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	CustomerID string `json:"customerID"`
}

type userAccountRepository struct {
	accounts []UserAccount
}

var (
	lock = sync.Mutex{}
)

func (r *userAccountRepository) GetAccount(username string, password string) (*UserAccount, error) {
	lock.Lock()
	defer lock.Unlock()
	for _, account := range r.accounts {
		if account.Username == username && account.Password == password {
			return &account, nil
		}
	}
	return nil, errors.New("account not found")
}

func (r *userAccountRepository) AddAccount(account UserAccount) {
	lock.Lock()
	defer lock.Unlock()
	r.accounts = append(r.accounts, account)
}

func (r *userAccountRepository) init() {
	lock.Lock()
	defer lock.Unlock()
	r.accounts = []UserAccount{
		{
			Username:   "admin",
			Password:   "admin",
			CustomerID: "1234567890",
		},
	}
}

var (
	instance *userAccountRepository
)

func FetchUserAccountRepository() *userAccountRepository {
	if instance == nil {
		lock.Lock()
		instance = &userAccountRepository{}
		instance.init()
		defer lock.Unlock()
	}
	return instance
}
