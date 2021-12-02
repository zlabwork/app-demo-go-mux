package app

type UserCache struct {
	cache map[int64]*User
	srv   UserService
}

func NewUserCache(service UserService) *UserCache {
	return &UserCache{
		cache: make(map[int64]*User),
		srv:   service,
	}
}

func (uc *UserCache) Get(id int64) (*User, error) {
	if u := uc.cache[id]; u != nil {
		return u, nil
	}

	user, err := uc.srv.Get(id)
	if err != nil {
		return nil, err
	} else if user != nil {
		uc.cache[id] = user
	}
	return user, err
}
