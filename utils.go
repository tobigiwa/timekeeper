package main

import (
	"log"

	"github.com/charmbracelet/bubbles/list"
	"github.com/olekukonko/ts"
)

func getwindowSize() ts.Size {
	var (
		size ts.Size
		err  error
	)

	if size, err = ts.GetSize(); err != nil {
		log.Fatal(err)
	}

	return size

}

type Dimension struct {
	Height int
	Width  int
}

func getListItemDim(t list.Model) Dimension {
	return Dimension{Height: t.Height(), Width: t.Width()}
}
