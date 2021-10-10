package app

type UserCache struct {
	cache   map[int64]*User
	service UserService
}

func NewUserCache(service UserService) *UserCache {
	return &UserCache{
		cache:   make(map[int64]*User),
		service: service,
	}
}

func (c *UserCache) User(id int64) (*User, error) {
	if u := c.cache[id]; u != nil {
		return u, nil
	}

	u, err := c.service.User(id)
	if err != nil {
		return nil, err
	} else if u != nil {
		c.cache[id] = u
	}
	return u, err
}
