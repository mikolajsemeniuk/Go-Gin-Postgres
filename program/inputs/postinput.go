package inputs

type PostInput struct {
	Title       string `json:"title" validate:"nonzero,max=10"`
	Description string `json:"description" validate:"min=4,max=20"`
}
