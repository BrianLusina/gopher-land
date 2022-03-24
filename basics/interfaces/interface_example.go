package interfaces

import (
	"fmt"
	"time"
)

type Entry interface {
	Title() string
}

type Book struct {
	Name      string
	Author    string
	Published time.Time
}

func (b Book) Title() string {
	return fmt.Sprintf("%s by %s (%s)", b.Name, b.Author, b.Published.Format("Jan 2006"))
}

type Movie struct {
	Name     string
	Director string
	Year     int
}

func (m Movie) Title() string {
	return fmt.Sprintf("%s (%d)", m.Name, m.Year)
}

func Display(e Entry) string {
	return e.Title()
}
