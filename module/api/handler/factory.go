package handler

// Factory ...
type Factory interface {
	User() UserHandler
}

type factory struct {
	userHandler UserHandler
}

// NewFactory ...
func NewFactory(userHandler UserHandler) Factory {
	return &factory{userHandler}
}

func (f *factory) User() UserHandler {
	return f.userHandler
}
