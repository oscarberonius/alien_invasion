package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type city struct {
	name, north, west, south, east string
}

func readFile(path string) []string {
	//Read file, return array of strings, one for each line
	var lines []string

	file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("Cannot open file, err: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var line = scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Cannot scan file, err: %v", err)
	}

	return lines
}

func buildCities(path string) []city {
	// Load city data into required data structures
	var cities []city

	var mapData []string = readFile(path)

	for i := 0; i < len(mapData); i++ { // Every line

		var cols = strings.Fields(mapData[i]) // Split on spaces

		var name string = cols[0]
		var north string = ""
		var west string = ""
		var south string = ""
		var east string = ""

		for j := 1; j < len(cols); j++ {
			// cols[j] == str(direction)=str(name) Right now
			var s = strings.Split(cols[j], "=")
			var direction, name = s[0], s[1]

			switch {
			case direction == "north":
				north = name
			case direction == "west":
				west = name
			case direction == "south":
				south = name
			case direction == "east":
				east = name
			}
		}
		var c = city{name, north, west, south, east}
		cities = append(cities, c)
	}
	return cities
}

func main() {
	fmt.Printf("Read file and create cities\n")
	var cities = buildCities("cities.txt")
	fmt.Printf("Cities created: \n First city: %+v \n Second city: %+v\n", cities[0], cities[1])

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
