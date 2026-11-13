package ch01

import "fmt"

type User struct {
	ID            int
	Score         float64
	Name, address string
}

func (u User) hello() {
	fmt.Println("My name is ", u.Name)
}

func (u User) UpdateID(newID int) {
	u.ID = newID
}

func (u *User) PtrUpdateID(newID int) {
	u.ID = newID
}

// nested struct
type TreeNode struct {
	Value int
	// Left TreeNode - this doesn't work because it cases self-reference
	Left  *TreeNode
	Right *TreeNode
}

func StructRun() {
	var u1, u2 User
	u1 = User{ID: 32, address: "WA", Name: "Golang"}
	fmt.Println(u1.Name)
	u2 = User{32, 34, "WA", "Golang"}
	u2.hello()

	// typed variable
	// this method applied to the scenario that the struct will only be used once
	// within the current scope
	var student struct {
		Name string
		Age  int
	}
	student.Name = u2.Name
	student.Age = 16
	fmt.Println(student.Name, student.Age)

	u3 := &u2 // get u2's address
	u3.Name = "Altered"
	fmt.Println(u2.Name)

	u4 := new(User) // u4 now is address. new creates empty struct, then returns address
	// print struct
	fmt.Printf("%v\n", u4)   // print values
	fmt.Printf("%+v\n", u4)  // print field name and value
	fmt.Printf("%+#v\n", u4) // print struct and field name and value

	// difference between struct method and ptr struct method
	u5 := User{ID: 345}
	fmt.Printf("original id: %d\n", u5.ID)
	u5.UpdateID(123)
	fmt.Printf("struct method update id: %d\n", u5.ID)
	u5.PtrUpdateID(789)
	fmt.Printf("ptr method update id: %d\n", u5.ID)
}
