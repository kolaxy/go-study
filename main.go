package main
import "fmt"
//for [инициализация счетчика]; [условие]; [изменение счетчика]{
    // действия
// }
func main() {
	sum := 0
	for i := 1; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)
}