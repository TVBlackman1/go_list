package pages

import (
	"fmt"
)

type Welcome struct {
}

func (*Welcome) View() {
	fmt.Println("Welcome")
	fmt.Println("1 - page 1")
	fmt.Println("2 - page 2")
}

func (*Welcome) Options() PageViewer {
	var i int
	fmt.Scanf("%d\n", &i)
	fmt.Println(i)
	if i == 1 {
		return &Page1{}
	} else {
		return &Page2{}
	}
}
