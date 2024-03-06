package pages

import (
	"fmt"
)

type Page2 struct {
}

func (*Page2) View() {
	fmt.Println("Page 2")
	fmt.Println("1 - page 1")
	fmt.Println("2 - welcome")
}

func (*Page2) Options() PageViewer {
	var i int
	fmt.Scanf("%d\n", &i)
	fmt.Println(i)
	if i == 1 {
		return &Page1{}
	} else {
		return &Welcome{}
	}
}
