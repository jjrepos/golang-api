package dtos

type CreateBookDto struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookDto struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}