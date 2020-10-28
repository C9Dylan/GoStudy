package main

import "fmt"

func main() {
	// get brand name
	fmt.Println(getBrandName(1))
	fmt.Println(getBrandNameWithoutBreak(3))
	// fall through
	fallThrough()
}

// Get brand name
func getBrandName(brandCode int) string {
	brandName := ""
	switch brandCode {
		case 1 : {
			brandName = "Adidas"
			break
		}
		case 2 :
			brandName = "Nike"
			break
		default :
			brandName = "other"
	}
	return brandName
}

// Get brand name without break keyword (Not like java, if it's true a condition, will continue executing code)
func getBrandNameWithoutBreak(brandCode int) string {
	brandName := ""
	switch brandCode {
	case 1 : brandName = "Adidas"
	case 2 : brandName = "Nike"
	default : brandName = "other"
	}
	return brandName
}

// fall through
func fallThrough() {
	switch {
		case true:
			fmt.Println("Case 1 is true")
			fallthrough
		case false: fmt.Println("Case 2 is true")
	}
}