package books

import "fmt"

type Author string

type Book struct {
	ID            int
	Title         string
	Author        Author
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf("Title:\t\t%q\n"+
		"Author:\t\t%q\n"+
		"Published:\t%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	Book{
		ID:            1,
		Title:         "Hitchhiker's Guide to the Galaxy",
		Author:        "Douglas Adams",
		YearPublished: 1979,
	},
	Book{
		ID:            2,
		Title:         "The Hobbit",
		Author:        "J.R.R Tolkien",
		YearPublished: 2000,
	},
	Book{
		ID:            3,
		Title:         "Sapiens",
		Author:        "Author",
		YearPublished: 1980,
	},
	Book{
		ID:            4,
		Title:         "Pragmatic Programmer",
		Author:        "Adams",
		YearPublished: 1979,
	},
	Book{
		ID:            5,
		Title:         "The Alchemist",
		Author:        "Paulo Coelho",
		YearPublished: 2020,
	},
	Book{
		ID:            6,
		Title:         "The Economist",
		Author:        "Some Author",
		YearPublished: 2040,
	},
	Book{
		ID:            7,
		Title:         "Me",
		Author:        "You",
		YearPublished: 2020,
	},
	Book{
		ID:            8,
		Title:         "The Book",
		Author:        "Another Author",
		YearPublished: 2020,
	},
	Book{
		ID:            9,
		Title:         "Them",
		Author:        "Us",
		YearPublished: 2020,
	},
	Book{
		ID:            10,
		Title:         "Yule",
		Author:        "Mimi",
		YearPublished: 2020,
	},
}
