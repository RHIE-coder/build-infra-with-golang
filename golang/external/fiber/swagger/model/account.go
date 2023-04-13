package model

// Account model info
// @Description User account information
// @Description with user id and username
type Account struct {
	// ID this is userid
	ID      int    `json:"id"`
	Name    string `json:"name"` // This is Name
	Ignored int    `swaggerignore:"true"`
}
