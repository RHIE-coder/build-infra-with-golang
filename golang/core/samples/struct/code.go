package main

import "fmt"

type Data struct {
	Prim string
	Ref  []int
	Sub  SubData
}

type SubData struct {
	StrAddr *string
}

/* fmt.Println(&origin)의 결과 */
func (d *Data) String() string {
	return "this is my data"
}

func (d Data) ChangeSub(str string) {
	*d.Sub.StrAddr = str
}

/* fmt.Println(origin)의 결과 */
// func (d Data) String() string {
// 	return "this is my data"
// }

func (d *Data) Clone() Data {
	d2 := new(Data)
	*d2 = *d
	return *d2
}

func USAGE_TYPE_OF_struct() {
	str := "hi"

	origin := Data{
		Prim: "hello",
		Ref:  []int{1, 2, 3},
		Sub: SubData{
			StrAddr: &str,
		},
	}

	fmt.Println(origin)

	copy := origin.Clone()

	copy.Prim = "world"
	copy.Ref = append(copy.Ref, 4)
	*copy.Sub.StrAddr = "bye"
	fmt.Println(*origin.Sub.StrAddr, " : ", *copy.Sub.StrAddr) // bye  :  bye
	origin.ChangeSub("re-bye")
	fmt.Println(*origin.Sub.StrAddr, " : ", *copy.Sub.StrAddr) // re-bye  :  re-bye

	fmt.Println(&origin) // this is my data

	fmt.Printf("origin address: %p\n", &origin)                      // 0xc00007e150
	fmt.Printf("copy address: %p\n", &copy)                          // 0xc00000e028
	fmt.Printf("origin ref address: %p\n", &origin.Ref)              // 0xc00007e160
	fmt.Printf("copy ref address: %p\n", &copy.Ref)                  // 0xc00007e190
	fmt.Printf("origin sub address: %p\n", &origin.Sub)              // 0xc00007e178
	fmt.Printf("copy sub address: %p\n", &copy.Sub)                  // 0xc00007e1a8
	fmt.Printf("origin sub str address: %p\n", (origin.Sub.StrAddr)) // 0xc000010250
	fmt.Printf("copy sub str address: %p\n", (copy.Sub.StrAddr))     // 0xc000010250
	fmt.Printf("%+v\n", origin)                                      // {Prim:hello Ref:[1 2 3]   Sub:{StrAddr:0xc000010250}}
	fmt.Printf("%+v\n", copy)                                        // {Prim:world Ref:[1 2 3 4] Sub:{StrAddr:0xc000010250}}

}

type User struct {
	username string
	age      int
}

func (u *User) Ref() {
	fmt.Printf("%p\n", u)
}

func (u User) Val() {
	fmt.Printf("%p\n", &u)
}

func INFO_DIFFERENCE_BETWEEN_REFERENCE_AND_VALUE() {
	u1 := User{}
	fmt.Printf("%p\n", &u1) // 0xc00000c030
	u1.Ref()                // 0xc00000c030
	u1.Ref()                // 0xc00000c030
	u1.Ref()                // 0xc00000c030
	u1.Val()                // 0xc00000c048
	u1.Val()                // 0xc00000c060
	u1.Val()                // 0xc00000c078

	u2 := &User{}
	fmt.Printf("%p\n", u2)  // 0xc00000c090
	fmt.Printf("%p\n", &u2) // 0xc00000e030
	u2.Ref()                // 0xc00000c090
	u2.Ref()                // 0xc00000c090
	u2.Ref()                // 0xc00000c090
	u2.Val()                // 0xc00000c0a8
	u2.Val()                // 0xc00000c0c0
	u2.Val()                // 0xc00000c0d

	fmt.Println("----------")
	u1a := u1
	u1b := u1
	u2a := u2
	u2b := u2
	u2aa := *u2
	u2bb := *u2

	fmt.Printf("%p\n", &u1a)  // 0xc0000ac0d8
	fmt.Printf("%p\n", &u1b)  // 0xc0000ac0f0
	fmt.Printf("%p\n", u2a)   // 0xc0000ac078
	fmt.Printf("%p\n", u2b)   // 0xc0000ac078
	fmt.Printf("%p\n", &u2aa) // 0xc00000c120
	fmt.Printf("%p\n", &u2bb) // 0xc00000c138

}

func main() {
	// USAGE_TYPE_OF_struct()
	// INFO_DIFFERENCE_BETWEEN_REFERENCE_AND_VALUE()
}
