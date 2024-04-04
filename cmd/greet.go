package cmd

import (
	"fmt"

	"github.com/fatih/color"
)

func greet() {
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

	fmt.Print(red.Sprint(nameASCIIFirst))
	fmt.Print(boldYellow.Sprint(nameASCIIDescription))
	fmt.Println(red.Sprint(nameASCIILast))
}
