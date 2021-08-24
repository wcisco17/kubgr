package book

const (
	PUBLISHED   = "PUBLISHED"
	UNPUBLISHED = "UNPUBLISHED"
)

type Book struct {
	Title     string
	Desc      string
	Published string
}
