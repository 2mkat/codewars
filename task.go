package main
import (
	"fmt"
	"strings"
)

func FindShort(s string) int {
	res := 100
	for _, word := range strings.Fields(s){
		if len(word) <= res{
			res = len(word)
		}
	}

	return res
}

func main(){

	fmt.Println(FindShort("turnses hgng es"))

}