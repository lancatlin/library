package search

import "github.com/lancatlin/library/pkg/model"

type AccountSet struct {
	m        map[uint]bool
	accounts []model.Account
}

func NewAccountSet() AccountSet {
	return AccountSet{
		m: map[uint]bool{},
	}
}

func (s *AccountSet) Add(accts []model.Account) {
	for _, a := range accts {
		if _, ok := s.m[a.ID]; !ok {
			s.m[a.ID] = true
			s.accounts = append(s.accounts, a)
		}
	}
}

func (s *AccountSet) List() []model.Account {
	return s.accounts
}
