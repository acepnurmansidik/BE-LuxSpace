package review

import "LuxSpace/app/v1/user"

type Review struct {
	ID        int
	Rate      int
	Comments  string
	ProductID int
	UserID    int
	Images    []ReviewImages
	User      user.User
}

type ReviewImages struct {
	ID       int
	Name     string
	ReviewId int
}
