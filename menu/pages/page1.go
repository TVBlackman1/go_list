package pages

import (
	"fmt"
)

type Page1 struct {
}

func (*Page1) View() {
	fmt.Println("Page 1")
	fmt.Println("1 - page 2")
	fmt.Println("2 - welcome")
}

func (*Page1) Options() PageViewer {
	var i int
	fmt.Scanf("%d\n", &i)
	fmt.Println(i)
	if i == 1 {
		return &Page2{}
	} else {
		return &Welcome{}
	}
}
