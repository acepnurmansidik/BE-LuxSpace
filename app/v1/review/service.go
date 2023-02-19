package review

type Service interface {
	CreateUserReview(inputData CreateReviewInput) (Review, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateUserReview(inputData CreateReviewInput) (Review, error) {
	// mapping datanya
	formatReview := Review{}
	formatReview.Comments = inputData.Comments
	formatReview.Rate = inputData.Rate
	formatReview.ProductID = inputData.ProductID
	formatReview.UserID = inputData.UserId
	// simpan datanya
	review, err := s.repository.SaveUserReview(formatReview)
	if err != nil {
		return review, err
	}

	return review, nil
}
