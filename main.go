package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type city struct {
	name, north, west, south, east string
}

func (c city) getRandomNeighbour() string {
	// Get a random neighbour of a city. Return name of city of no paths exist
	var a []string
	var city string

	if c.north != "" {
		a = append(a, c.north)
	}

	if c.west != "" {
		a = append(a, c.west)
	}

	if c.south != "" {
		a = append(a, c.south)
	}

	if c.east != "" {
		a = append(a, c.east)
	}

	var neighbours = len(a)

	if neighbours > 0 {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		city = a[r.Intn(neighbours)]
	} else {
		city = c.name //Not very elegant but whatevs
	}

	return city
}

func (c *city) removePaths(cityMap map[string][]int) {
	// Removes any paths in a city object if that path is no longer in the city map

	_, nok := cityMap[c.north]
	if !nok {
		c.north = ""
	}

	_, wok := cityMap[c.west]
	if !wok {
		c.west = ""
	}

	_, sok := cityMap[c.south]
	if !sok {
		c.south = ""
	}

	_, eok := cityMap[c.east]
	if !eok {
		c.east = ""
	}
}

func removeAllPaths(c []city, cm map[string][]int) {
	// Iterate through all city objects and remove destroyed paths

	for i := 0; i < len(c); i++ {
		c[i].removePaths(cm)
	}
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

func buildCities(path string) ([]city, map[string][]int) {
	// Load city data into required data structures
	var cities []city
	var cityMap = map[string][]int{} // map(City name - alien id)

	var mapData []string = readFile(path)

	for i := 0; i < len(mapData); i++ { // Every line

		var cols = strings.Fields(mapData[i]) // Split on spaces

		var name string = cols[0]
		cityMap[name] = []int{}
		var north string = ""
		var west string = ""
		var south string = ""
		var east string = ""

		for j := 1; j < len(cols); j++ {
			// cols[j] == str(direction)=str(name) Right now
			var s = strings.Split(cols[j], "=")
			var direction, n = s[0], s[1]

			switch {
			case direction == "north":
				north = n
				var c = city{name: n, south: name}
				cities = append(cities, c)
			case direction == "west":
				west = n
				var c = city{name: n, east: name}
				cities = append(cities, c)
			case direction == "south":
				south = n
				var c = city{name: n, north: name}
				cities = append(cities, c)
			case direction == "east":
				east = n
				var c = city{name: n, west: name}
				cities = append(cities, c)
			}
			cityMap[n] = []int{}

		}
		var c = city{name, north, west, south, east}
		cities = append(cities, c)
	}
	return cities, cityMap
}

func spawnAliens(cityMap map[string][]int, numAliens int, currentAlien int) {
	// Iterate through map and add alien IDs to each city
	// Not completely random as maps aren't continually reshuffled, but will do for now

	for k, v := range cityMap {
		cityMap[k] = append(v, currentAlien)
		currentAlien++
		if currentAlien >= numAliens {
			break
		}
	}

	// If there are more aliens than cities, keep spawning 'randomly'
	if currentAlien < numAliens {
		spawnAliens(cityMap, numAliens, currentAlien)
	}

}

func step(cityMap map[string][]int, cities []city) {
	var targetMap = make(map[string][]int)

	// Clear positions in copy before refilling with new random positions
	for k := range cityMap {
		targetMap[k] = nil
	}

	// Retrieve random neighbours and fill new map
	for name, aliens := range cityMap {
		// Add exception check (1) here
		var city = getCity(name, cities)
		if city.name != "" {
			// Alien enters neighbouring city
			var randomNeighbour = city.getRandomNeighbour()
			targetMap[randomNeighbour] = append(targetMap[randomNeighbour], aliens...)
		}
	}

	// All aliens have taken 1 step, update state
	cityMap = targetMap

	// Check for fights
	checkFights(cityMap, cities)

}

func getCity(name string, cities []city) city {
	// Super ugly way of finding city object given city name, optimize later

	for _, c := range cities {
		if c.name == name {
			return c
		}
	}
	return city{} // Should not be possible, add exception check (1) or find better way
}

func checkFights(cityMap map[string][]int, cities []city) {

	// Remove all cities and aliens where more than 2 aliens end up

	for cityName, aliens := range cityMap {
		if len(aliens) > 1 {
			fmt.Printf("%s has been destroyed by aliens: %v", cityName, aliens)
			delete(cityMap, cityName)

			// Also remove city from city list. Ugly, make better than O(n^2) if needed
			for _, c := range cities {
				if c.name == cityName {
					c = city{}
				}
			}
		}
	}
}

func main() {
	// fmt.Printf("Read file and create cities\n")
	// var _, cityMap = buildCities("cities.txt")
	// fmt.Printf("Citymap: %v\n", cityMap)
	// spawnAliens(cityMap, 30, 0)
	// fmt.Printf("30 aliens spawned: %+v\n", cityMap)

	// var cities, cityMap = buildCities("cities.txt")
	//fmt.Printf("Cities built\n cities: %+v \n cityMap: %+v \n", cities, cityMap)
	//spawnAliens(cityMap, 5, 0)
	// fmt.Printf("\n5 aliens spawned. cityMap: \n %+v \n cities: \n %+v\n", cityMap, cities)

	//for i := 0; i < 20; i++ {
	//step(cityMap, cities)

	//fmt.Printf("\nOne step taken, cityMap: \n%+v \n", cityMap)
	//}

	// fmt.Printf("Cities created: \n First city: %+v \n Second city: %+v\n", cities[0], cities[1])
	// fmt.Printf("CityMap created: %+v\n", cityMap)
	//var neighbour = cities[0].getRandomNeighbour()
	//fmt.Printf("Random neighbour of city %+v : %s\n", cities[0], neighbour)

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
