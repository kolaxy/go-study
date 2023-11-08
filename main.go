package main
import "fmt"

func main() {
	initialUsers := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	initialUsers = append(initialUsers, "NEW")
	fmt.Println(initialUsers)
}