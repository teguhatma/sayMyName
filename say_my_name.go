package saymyname

import "fmt"

func SayMyName(name, age string) string {
	return fmt.Sprintf("Hallo, I am %s and I am %s ", name, age)
}
