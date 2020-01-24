package store

func (s Store) CreateUser(email string) error {
	// TODO: actually connect to database
	log.Info("Creating user " + email)
	return nil
}
