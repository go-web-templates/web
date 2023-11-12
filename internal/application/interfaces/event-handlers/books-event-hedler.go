package eventhandlers

import (
	events "github.com/go-web-templates/web/internal/domain/events/books"
)

type BooksEventHandler interface {
	Handle(event events.BookEvent) 
}
