package main

import "fmt"

func main() {
	var sname, college, div string
	var rno, go int,

	fmt.Printf("Enter Student name:")
	fmt.Scanln(&sname)
	fmt.Printf("Enter division:")
	fmt.Scanln(&div)
	fmt.Printf("Enter college name:")
	fmt.Scanln(&college)
	fmt.Printf("Enter Student roll no:")
	fmt.Scanln(&rno)

	fmt.Printf("Student name:%s\nDivision:%s\nCollege name:%s\nRoll no:%d", sname, div, college, rno)
}
