package cmd

import (
	"fmt"

	"github.com/fatih/color"
)

type Wiggumizer struct {
	Params Params
}

func NewWigomiser() (Wiggumizer, error) {
	var w Wiggumizer

	p, err := newParams()
	if err != nil {
		return w, err
	}

	w.Params = p
	w.greet()
	return w, nil
}

func (w Wiggumizer) greet() {
	nameASCIIFirst := `
	__       __  __                                              __                     
	|  \  _  |  \|  \                                            |  \                    
	| $$ / \ | $$ \$$  ______    ______   __    __  ______ ____   \$$ ________   ______  
	| $$/  $\| $$|  \ /      \  /      \ |  \  |  \|      \    \ |  \|        \ /      \ 
	| $$  $$$\ $$| $$|  $$$$$$\|  $$$$$$\| $$  | $$| $$$$$$\$$$$\| $$ \$$$$$$$$|  $$$$$$\
	| $$ $$\$$\$$| $$| $$  | $$| $$  | $$| $$  | $$| $$ | $$ | $$| $$  /    $$ | $$    $$
	| $$$$  \$$$$| $$| $$__| $$| $$__| $$| $$__/ $$| $$ | $$ | $$| $$ /  $$$$_ | $$$$$$$$
	| $$$    \$$$| $$ \$$    $$ \$$    $$ \$$    $$| $$ | $$ | $$| $$|  $$    \ \$$     \
	 \$$      \$$ \$$ _\$$$$$$$ _\$$$$$$$  \$$$$$$  \$$  \$$  \$$ \$$ \$$$$$$$$  \$$$$$$$
			 |  \__| $$|  \__| $$   `

	nameASCIIDescription := "Web Traffic 4nalizer"

	nameASCIILast := `
			  \$$    $$ \$$    $$
			   \$$$$$$   \$$$$$$
	`

	// Define the colors
	red := color.New(color.FgBlue)
	boldYellow := color.New(color.FgYellow, color.Bold).Add(color.Underline)

	fmt.Println(red.Sprint(nameASCIIFirst))
	fmt.Print(boldYellow.Sprint(nameASCIIDescription))
	fmt.Println(red.Sprint(nameASCIILast))
}
