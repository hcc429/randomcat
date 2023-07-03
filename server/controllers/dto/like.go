package dto

type Likes struct {
	Likes []Like `binding:"required"`
}

type Like struct {
	URL    string `json:"url"`
	Amount int    `json:"amount"`
}
