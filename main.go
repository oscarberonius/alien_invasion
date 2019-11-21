package main

import "fmt"

func main() {
	fmt.Printf("Hello world\n")

	//Text file format:
	// kba north=gbg west=uk south=skne east=sthlm

	//1. Read file, create cities. 2 data structures for cities, 1 list of all city objets and 1 hash map with cities-aliens key-value pairs.
	//2. Generate aliens, assign random cities.
	//3. 1 Step, iterate through all aliens and get random neighbours of their cities.
	//4. Update cities hashmap with each new city for each alien.
	//5. Remove all cities from hash map that contain more than 1 alien.
	//6. Iterate through all cities and check if each neighbour is a key in hash map. If not, set to ""
	//7. Do this while cities remain or 10000 steps have been made.
	//8. Write to file
}
