package league

import "fmt"

func Ordinal(number int) string {
	switch number % 100 {
	case 11, 12, 13:
		return fmt.Sprintf("%dth", number)
	default:
		switch number % 10 {
		case 1:
			return fmt.Sprintf("%dst", number)
		case 2:
			return fmt.Sprintf("%dnd", number)
		case 3:
			return fmt.Sprintf("%drd", number)
		default:
			return fmt.Sprintf("%dth", number)
		}
	}
}
