package inmemory

import "github.com/istsh/go-grpc-sample/app/entity/model"

// NewStore creates a new store.
func NewStore() *Store {
	return &Store{}
}

// Store is a memory store.
type Store struct {
	users         []*model.User
	userPasswords []*model.UserPassword
	userTokens    []*model.UserToken
}

// AddUsers adds some users.
func (s *Store) AddUsers(users ...*model.User) {
	s.users = append(s.users, users...)
}

// AddUserPasswords adds some user passwords.
func (s *Store) AddUserPasswords(userPasswords ...*model.UserPassword) {
	s.userPasswords = append(s.userPasswords, userPasswords...)
}

// AddUserTokens adds some user tokens.
func (s *Store) AddUserTokens(userTokens ...*model.UserToken) {
	s.userTokens = append(s.userTokens, userTokens...)
}
