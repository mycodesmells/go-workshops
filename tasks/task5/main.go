package main

import "fmt"

// Table is where you are walking or jumping, apparently.
type Table struct {
	broken bool
}

func walkOnTable(t Table) {
	fmt.Println("Walking on the table...")
	t.broken = true
}

func jumpOnTable(t *Table) {
	fmt.Println("Jumping on the table...")
	t.broken = true
}

func printTableState(t Table) {
	fmt.Printf("Table broken? %v\n", t.broken)
}

func main() {
	t := Table{}
	printTableState(t)
	walkOnTable(t)
	printTableState(t)
	jumpOnTable(&t)
	printTableState(t)
}
