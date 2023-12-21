package color

import "fmt"

// 31 = red
// 32 = green
// 33 = yellow
// 34 = blue
// 35 = magenta
// 0 = reset
func Red(text string) string {
	return fmt.Sprintf("\033[31m%s\033[0m", text)
}
func Blue(text string) string {
	return fmt.Sprintf("\033[34m%s\033[0m", text)
}
func Green(text string) string {
	return fmt.Sprintf("\033[32m%s\033[0m", text)
}
func Yellow(text string) string {
	return fmt.Sprintf("\033[33m%s\033[0m", text)
}
func Magenta(text string) string {
	return fmt.Sprintf("\033[35m%s\033[0m", text)
}
