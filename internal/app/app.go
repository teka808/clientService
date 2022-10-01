package app

import (
	"fmt"
	"sync"
)

var (
	once sync.Once
	app  *instance
)

type instance struct {
	title string
}

func Instance(title string) *instance {
	once.Do(func() {
		app := new(instance)
		app.title = title
		fmt.Println(app.title)
	})
	return app
}
