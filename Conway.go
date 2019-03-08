/*
Conway's Game of Life in GoLang
*/
package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
)

// Gather information on the number of rows and columns and use it globally.
var rows, columns, maxGen int = 0, 0, 10
var currMap = [][] cell {}

type cell struct {
	isAlive bool	// is the cell alive?
	adjAlive int 	// how many neighbors are alive at the moment?
}

// Set the number of neighors alive for each cell
func GetAdjStatus() {
	sliceRange := [] int {-1, 0, 1}
	for rowIndex, eachRow := range currMap {
		for colIndex, _ := range eachRow {
			currMap[rowIndex][colIndex].adjAlive = 0	// clear # neighbors alive after each round
			for _,rowRange := range sliceRange {
				for _,colRange := range sliceRange {
					// bound checking
					if rowIndex + rowRange >= 0 && colIndex + colRange >=0 && rowIndex + rowRange < rows && colIndex + colRange < columns {
						if currMap[rowIndex + rowRange][colIndex + colRange].isAlive {
							currMap[rowIndex][colIndex].adjAlive ++
						}
					}
				}
			}
			if currMap[rowIndex][colIndex].isAlive{
				currMap[rowIndex][colIndex].adjAlive --
			}
		}
	}
}

/* Apply rules and update isAlive in type cell, based on # of adjacent neighbors
- if alive:
	-> dead if # of alive neighbors is < 2 or > 3
- if dead:
	-> comes back to life if # of alive neighbors is exactly 3
*/
func nextGen() {
	GetAdjStatus()
	for rowIndex, eachRow := range currMap {
		for colIndex, thisCell := range eachRow {
			if thisCell.isAlive{
				if thisCell.adjAlive < 2 || thisCell.adjAlive > 3 {
					currMap[rowIndex][colIndex].isAlive = false
				}
			} else {
				if thisCell.adjAlive == 3 {
					currMap[rowIndex][colIndex].isAlive = true
				}
			}
		}
	}
}

// Print out current map
func printMap() {
	for _, thisRow := range currMap {
		for _, thisCell := range thisRow {
			if thisCell.isAlive {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println("====================")
}

// Read from "life.txt" to get an inital map
func initializeData() {

	file, err := os.Open("life.txt")
	rowMap := [] cell {}
	currRow, currCol := 0, 0

	// Make sure that "life.txt" exists(that is the necessary input file)
	if err != nil{
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		currCol = 0
		rowMap = nil	// empty out slice
		for _,character := range line {
			if string(character) == "*" {
				rowMap = append(rowMap, cell {true, 0})
			} else {
				rowMap = append(rowMap, cell {false, 0})
			}
			currCol += 1
		}
		currMap = append(currMap, rowMap)
		currRow += 1
	}
	rows = currRow	// assumption, given that all lines are the same length
	columns  = currCol
	file.Close()
}

func main() {
	currGen := 1
	initializeData()
	fmt.Println("Initial World")
	printMap()
	for currGen <= maxGen {	// loop for 10 gens (you can set more generations by changing "maxGen")
		fmt.Println("Generation:", currGen)
		nextGen()
		printMap()
		currGen++
	}
}