package repository

// Transaction provides a set of repository reader/writer methods.
type Transaction interface {
	User() UserRepositoryModify
	UserPassword() UserPasswordRepositoryModify
	UserToken() UserTokenRepositoryModify
}

// Connection provides a set of repository reader methods.
type Connection interface {
	Close() error
	RunTransaction(f func(Transaction) error) error

	User() UserRepositoryAccess
	UserPassword() UserPasswordRepositoryAccess
	UserToken() UserTokenRepositoryAccess
}

// Repository provides an abstract connection method.
type Repository interface {
	//
	NewConnection() Connection
}
