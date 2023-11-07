package main
import "fmt"
//for [инициализация счетчика]; [условие]; [изменение счетчика]{
    // действия
// }
func main() {
	var i = 1
	for ;i < 10; i++ {
		fmt.Println(i * i)
	}
}