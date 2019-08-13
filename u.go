package testGo
import (
	"fmt"
	"github.com/panjf2000/ants"
)
var g *ants.PoolWithFunc
func init(){
	g,_=ants.NewPoolWithFunc(50, func(i interface{}) {
		fu := i.(func())
		fu()
	})
}
func P(){
	g.Invoke(func() {
		fmt.Println("exxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	})
}