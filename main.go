package main
import "fmt"

func main() {
	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	users1 := initialUsers[2:6]
	users2 := initialUsers[:4]
	users3 := initialUsers[3:]

	fmt.Println(users1, users2, users3)
}