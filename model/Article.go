package model

type (
	// Article ...
	Article struct {
		ID       int
		Title    string
		Body     string
		AuthorID int
		Author   *User
	}
)
