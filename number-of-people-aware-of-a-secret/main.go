package main

import "fmt"

func peopleAwareOfSecret(n int, delay int, forget int) int {
	const mod = 1_000_000_007
	type person struct {
		delay  int
		forget int
	}

	people := []person{{delay - 1, forget - 1}}

	for day := 1; day < n; day++ {
		newPeople := []person{}
		nextPeople := []person{}

		for _, p := range people {
			if p.forget == 0 {
				continue // forgets the secret
			}
			if p.delay == 0 {
				// person shares
				newPeople = append(newPeople, person{delay - 1, forget - 1})
			} else {
				p.delay--
			}
			p.forget--
			nextPeople = append(nextPeople, p)
		}
		people = append(nextPeople, newPeople...)
	}

	return len(people) % mod
}

func main() {
	fmt.Println(peopleAwareOfSecret(6, 2, 4)) // 5
	fmt.Println(peopleAwareOfSecret(4, 1, 3)) // 6
	// fmt.Println(peopleAwareOfSecret(684, 18, 496)) :=> Number grows exponentially (millions+) :- Go panic with out of memory
	// I Will use Dynamic programming to track number of people that who knows secret at the end of n days
}
