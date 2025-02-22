package testdata

import (
	"github.com/takeshisean/cckit/extensions/debug"
	"github.com/takeshisean/cckit/extensions/owner"
	"github.com/takeshisean/cckit/router"
	p "github.com/takeshisean/cckit/router/param"
	"github.com/takeshisean/cckit/state/testdata/schema"
)

const collection = "SampleCollection"

func NewBooksCC() *router.Chaincode {
	r := router.New(`books`)
	debug.AddHandlers(r, `debug`, owner.Only)

	r.Init(owner.InvokeSetFromCreator).
		Invoke(`bookList`, bookList).
		Invoke(`bookGet`, bookGet, p.String(`id`)).
		Invoke(`bookInsert`, bookInsert, p.Struct(`book`, &schema.Book{})).
		Invoke(`bookUpsert`, bookUpsert, p.Struct(`book`, &schema.Book{})).
		Invoke(`bookDelete`, bookDelete, p.String(`id`)).
		Invoke(`privateBookList`, privateBookList).
		Invoke(`privateBookGet`, privateBookGet, p.String(`id`)).
		Invoke(`privateBookInsert`, privateBookInsert, p.Struct(`book`, &schema.PrivateBook{})).
		Invoke(`privateBookUpsert`, privateBookUpsert, p.Struct(`book`, &schema.PrivateBook{})).
		Invoke(`privateBookDelete`, privateBookDelete, p.String(`id`))

	return router.NewChaincode(r)
}

func bookList(c router.Context) (interface{}, error) {
	return c.State().List(schema.BookEntity, &schema.Book{})
}

func bookInsert(c router.Context) (interface{}, error) {
	book := c.Param(`book`)
	return book, c.State().Insert(book)
}

func bookUpsert(c router.Context) (interface{}, error) {
	book := c.Param(`book`)
	return book, c.State().Put(book)
}

func bookGet(c router.Context) (interface{}, error) {
	return c.State().Get(schema.Book{Id: c.ParamString(`id`)})
}

func bookDelete(c router.Context) (interface{}, error) {
	return nil, c.State().Delete(schema.Book{Id: c.ParamString(`id`)})
}

func privateBookList(c router.Context) (interface{}, error) {
	return c.State().ListPrivate(collection, false, schema.PrivateBookEntity, &schema.PrivateBook{})
}

func privateBookInsert(c router.Context) (interface{}, error) {
	book := c.Param(`book`)
	err := c.State().Insert(book, "{}")
	if err != nil {
		return book, err
	}
	return book, c.State().InsertPrivate(collection, book)
}

func privateBookUpsert(c router.Context) (interface{}, error) {
	book := c.Param(`book`)
	err := c.State().Put(book, "{}")
	if err != nil {
		return book, err
	}
	return book, c.State().PutPrivate(collection, book)
}

func privateBookGet(c router.Context) (interface{}, error) {
	return c.State().GetPrivate(collection, schema.PrivateBook{Id: c.ParamString(`id`)})
}

func privateBookDelete(c router.Context) (interface{}, error) {
	c.State().Delete(schema.PrivateBook{Id: c.ParamString(`id`)})
	return nil, c.State().DeletePrivate(collection, schema.PrivateBook{Id: c.ParamString(`id`)})
}
