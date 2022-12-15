package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./example/example.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Make adjacency list.
	adjList := make([][]int, 0)

	var (
		inBlock       int
		curBlock      int
		prevBlock     int
		numPrevBlocks = make([]int, 0)
		prevsStack    = make([]int, 0)
	)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		// If string contains main() then we found start node.
		if strings.Contains(line, "func main() {") {
			adjList = append(adjList, []int{})
			// That's our first block
			inBlock = 0
			// and current block if start
			curBlock = 0
		}

		// If we found line that starts with "if" then it's an another node.
		if strings.HasPrefix(line, "if") {
			// We add nest level
			inBlock++

			if len(numPrevBlocks) < inBlock {
				numPrevBlocks = append(numPrevBlocks, 0)
			}
			numPrevBlocks[inBlock-1]++

			// Make new block
			curBlock = len(adjList)
			// We in if block then we have previous block
			if inBlock > 1 {
				prevsStack = append(prevsStack, prevBlock)
			}
			prevBlock = curBlock - 1

			adjList = append(adjList, []int{})

			// Obviously our current block is achieved from previous one, i.e. inBlock-1
			//fmt.Printf("if: adding %d to %d\n", curBlock, prevBlock)
			adjList[prevBlock] = append(adjList[prevBlock], curBlock)
		} else if strings.HasPrefix(line, "} else if") {
			// We on the same nest level
			inBlock--
			inBlock++

			numPrevBlocks[inBlock-1]++

			// Make new block
			curBlock = len(adjList)
			adjList = append(adjList, []int{})

			// Obviously our current block is achieved from previous one, i.e. inBlock-1
			//fmt.Printf("else if: adding %d to %d\n", curBlock, prevBlock)
			adjList[prevBlock] = append(adjList[prevBlock], curBlock)
		} else if strings.HasPrefix(line, "} else") {
			// We on the same nest level
			inBlock--
			inBlock++

			numPrevBlocks[inBlock-1]++

			// Make new block
			curBlock = len(adjList)
			adjList = append(adjList, []int{})

			// Obviously our current block is achieved from previous one, i.e. inBlock-1
			//fmt.Printf("else: adding %d to %d\n", curBlock, prevBlock)
			adjList[prevBlock] = append(adjList[prevBlock], curBlock)
		} else if line == "}" {
			// We on the above nest level
			inBlock--
			if inBlock == -1 {
				break
			}

			// And our current block is len(adjList), i.e. new block
			curBlock = len(adjList)

			// We exit if-else block so our prevBlock is back
			n := len(prevsStack) - 1
			if n < 0 {
				prevBlock = 0
			} else {
				prevBlock = prevsStack[n]
				prevsStack = prevsStack[:n] // Pop
			}

			adjList = append(adjList, []int{})
			// Obviously our current block is achieved from all in previous if-else set.
			for i := numPrevBlocks[inBlock]; i > 0; i-- {
				//fmt.Printf("block-end: adding %d to %d\n", curBlock, curBlock-i)
				adjList[curBlock-i] = append(adjList[curBlock-i], curBlock)
			}
			numPrevBlocks[inBlock] = 0
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Connect end node to start node.
	endBlock := len(adjList) - 1
	adjList[endBlock] = append(adjList[endBlock], 0)

	/* for k, v := range adjList {
		fmt.Println(k, ":", v)
	} */

	var (
		E int
		N int = len(adjList)
		P int
	)
	for _, v := range adjList {
		E += len(v)
	}

	comps := make([]int, len(adjList))
	for v := 0; v < len(adjList); v++ {
		if comps[v] == 0 {
			P++
			dfs(v, P, comps, adjList)
		}
	}

	fmt.Printf("E: %d, N: %d, P: %d\n", E, N, P)
	fmt.Println("M:", E-N+P)
}

func dfs(v int, num int, comps []int, adjList [][]int) {
	comps[v] = num
	for _, u := range adjList[v] {
		if comps[u] == 0 {
			dfs(u, num, comps, adjList)
		}
	}
}
