package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type tree struct {
	nodes   []string
	fathers []int
	visited []bool
}

/*
	 G - H       J - K - L
	/           /

COM - B - C - D - E - F

	\
	 I

Cerco dalle foglie a COM, tenendo una lista di nodi già visitati
*/

func treeFromFile(fileName string) (orbits tree) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nodes := strings.Split(scanner.Text(), ")")
		// Check if the parent node already exists
		fatherIndex := slices.Index(orbits.nodes, nodes[0])
		if fatherIndex == -1 { // Parent node doesn't exist, add it
			orbits.nodes = append(orbits.nodes, nodes[0])
			orbits.fathers = append(orbits.fathers, -1) // Root has no parent
			orbits.visited = append(orbits.visited, false)
			fatherIndex = len(orbits.nodes) - 1 // Update index to the newly added parent
		}

		// Check if the child node already exists
		childIndex := slices.Index(orbits.nodes, nodes[1])
		if childIndex == -1 { // Child node doesn't exist, add it
			orbits.nodes = append(orbits.nodes, nodes[1])
			orbits.fathers = append(orbits.fathers, fatherIndex)
			orbits.visited = append(orbits.visited, false)
		} else {
			// If child exists, just update its father
			orbits.fathers[childIndex] = fatherIndex
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return
}

func (orbits tree) isVisited() bool {
	for _, v := range orbits.visited {
		if v == false {
			return false
		}
	}
	return true
}

func countOrbits(orbits tree) (count int) {
	for i := 0; i < len(orbits.nodes); i++ {
		currentNode := orbits.nodes[i]
		orbits.visited[i] = true
		j := i
		for currentNode != "COM" {
			count++
			if orbits.fathers[j] != -1 {
				j = orbits.fathers[j]
				currentNode = orbits.nodes[j]
			} else {
				count++
				break
			}
		}
	}
	return
}
func main() {
	orbitsMap := treeFromFile("Input.txt")
	fmt.Println(orbitsMap.nodes)
	fmt.Println(orbitsMap.fathers)
	fmt.Println(countOrbits(orbitsMap))
}
