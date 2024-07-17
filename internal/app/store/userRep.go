package store


type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model) (*model.User, error)  {
	return nil, nil 
}