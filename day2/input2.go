package main

type Direction struct {
	Direction string
	Speed     int
}

var Input = []Direction{Direction{"forward", 2}, Direction{"down", 1}, Direction{"down", 7}, Direction{"forward", 6}, Direction{"down", 6}, Direction{"down", 6}, Direction{"forward", 6}, Direction{"down", 6}, Direction{"down", 9}, Direction{"up", 6}, Direction{"forward", 5}, Direction{"down", 1}, Direction{"down", 7}, Direction{"forward", 5}, Direction{"forward", 8}, Direction{"forward", 8}, Direction{"forward", 3}, Direction{"down", 3}, Direction{"down", 9}, Direction{"down", 1}, Direction{"forward", 5}, Direction{"up", 2}, Direction{"down", 6}, Direction{"forward", 9}, Direction{"up", 5}, Direction{"forward", 3}, Direction{"down", 9}, Direction{"down", 6}, Direction{"down", 2}, Direction{"down", 2}, Direction{"down", 5}, Direction{"up", 8}, Direction{"up", 1}, Direction{"down", 2}, Direction{"down", 7}, Direction{"forward", 1}, Direction{"down", 1}, Direction{"down", 4}, Direction{"down", 2}, Direction{"forward", 7}, Direction{"down", 9}, Direction{"forward", 3}, Direction{"up", 1}, Direction{"up", 6}, Direction{"forward", 2}, Direction{"forward", 2}, Direction{"down", 5}, Direction{"down", 5}, Direction{"forward", 1}, Direction{"forward", 5}, Direction{"down", 6}, Direction{"forward", 8}, Direction{"forward", 5}, Direction{"up", 8}, Direction{"down", 7}, Direction{"forward", 5}, Direction{"up", 9}, Direction{"down", 3}, Direction{"forward", 7}, Direction{"up", 8}, Direction{"down", 3}, Direction{"down", 3}, Direction{"up", 8}, Direction{"down", 1}, Direction{"down", 5}, Direction{"up", 8}, Direction{"down", 5}, Direction{"forward", 5}, Direction{"forward", 4}, Direction{"up", 8}, Direction{"forward", 6}, Direction{"down", 1}, Direction{"down", 3}, Direction{"forward", 8}, Direction{"forward", 6}, Direction{"forward", 2}, Direction{"down", 9}, Direction{"up", 9}, Direction{"forward", 6}, Direction{"up", 4}, Direction{"down", 9}, Direction{"forward", 3}, Direction{"forward", 4}, Direction{"forward", 4}, Direction{"up", 2}, Direction{"forward", 6}, Direction{"up", 7}, Direction{"forward", 4}, Direction{"down", 7}, Direction{"forward", 3}, Direction{"forward", 1}, Direction{"forward", 1}, Direction{"down", 9}, Direction{"up", 6}, Direction{"forward", 9}, Direction{"down", 1}, Direction{"up", 4}, Direction{"forward", 2}, Direction{"forward", 1}, Direction{"down", 3}, Direction{"down", 2}, Direction{"forward", 8}, Direction{"forward", 4}, Direction{"forward", 6}, Direction{"down", 3}, Direction{"down", 9}, Direction{"forward", 7}, Direction{"forward", 9}, Direction{"down", 4}, Direction{"up", 3}, Direction{"down", 7}, Direction{"down", 2}, Direction{"down", 8}, Direction{"forward", 7}, Direction{"down", 5}, Direction{"forward", 5}, Direction{"forward", 9}, Direction{"down", 7}, Direction{"down", 4}, Direction{"down", 2}, Direction{"up", 6}, Direction{"forward", 6}, Direction{"down", 7}, Direction{"down", 3}, Direction{"up", 7}, Direction{"forward", 4}, Direction{"down", 7}, Direction{"forward", 1}, Direction{"forward", 1}, Direction{"up", 1}, Direction{"up", 8}, Direction{"down", 7}, Direction{"forward", 8}, Direction{"up", 9}, Direction{"up", 2}, Direction{"forward", 8}, Direction{"forward", 2}, Direction{"forward", 3}, Direction{"forward", 2}, Direction{"down", 2}, Direction{"down", 4}, Direction{"forward", 1}, Direction{"forward", 1}, Direction{"down", 6}, Direction{"forward", 5}, Direction{"down", 9}, Direction{"down", 9}, Direction{"down", 3}, Direction{"forward", 6}, Direction{"forward", 5}, Direction{"down", 4}, Direction{"forward", 4}, Direction{"forward", 6}, Direction{"down", 9}, Direction{"down", 4}, Direction{"forward", 5}, Direction{"forward", 6}, Direction{"forward", 4}, Direction{"forward", 6}, Direction{"forward", 7}, Direction{"down", 6}, Direction{"down", 4}, Direction{"up", 8}, Direction{"down", 4}, Direction{"forward", 9}, Direction{"down", 3}, Direction{"down", 8}, Direction{"down", 5}, Direction{"up", 6}, Direction{"forward", 7}, Direction{"forward", 8}, Direction{"forward", 8}, Direction{"down", 5}, Direction{"up", 3}, Direction{"down", 3}, Direction{"forward", 4}, Direction{"forward", 4}, Direction{"up", 9}, Direction{"forward", 6}, Direction{"down", 1}, Direction{"down", 1}, Direction{"down", 7}, Direction{"down", 6}, Direction{"forward", 9}, Direction{"forward", 2}, Direction{"down", 5}, Direction{"up", 4}, Direction{"down", 5}, Direction{"forward", 8}, Direction{"down", 8}, Direction{"forward", 1}, Direction{"forward", 4}, Direction{"up", 2}, Direction{"down", 4}, Direction{"forward", 1}, Direction{"down", 6}, Direction{"forward", 1}, Direction{"forward", 3}, Direction{"up", 8}, Direction{"forward", 1}, Direction{"up", 8}, Direction{"down", 7}, Direction{"forward", 6}, Direction{"up", 5}, Direction{"down", 6}, Direction{"up", 8}, Direction{"down", 7}, Direction{"down", 6}, Direction{"forward", 7}, Direction{"forward", 9}, Direction{"forward", 2}, Direction{"down", 6}, Direction{"down", 4}, Direction{"up", 2}, Direction{"forward", 1}, Direction{"down", 8}, Direction{"forward", 9}, Direction{"down", 5}, Direction{"down", 7}, Direction{"forward", 6}, Direction{"up", 1}, Direction{"forward", 2}, Direction{"forward", 3}, Direction{"forward", 9}, Direction{"up", 3}, Direction{"forward", 7}, Direction{"up", 7}, Direction{"up", 7}, Direction{"down", 5}, Direction{"up", 7}, Direction{"down", 8}, Direction{"forward", 2}, Direction{"down", 4}, Direction{"down", 6}, Direction{"up", 2}, Direction{"down", 7}, Direction{"forward", 4}, Direction{"down", 2}, Direction{"forward", 3}, Direction{"forward", 7}, Direction{"forward", 4}, Direction{"forward", 1}, Direction{"up", 6}, Direction{"forward", 3}, Direction{"forward", 6}, Direction{"forward", 9}, Direction{"up", 9}, Direction{"down", 2}, Direction{"forward", 1}, Direction{"forward", 9}, Direction{"forward", 5}, Direction{"forward", 7}, Direction{"forward", 8}, Direction{"down", 1}, Direction{"down", 9}, Direction{"up", 4}, Direction{"forward", 5}, Direction{"up", 9}, Direction{"down", 9}, Direction{"down", 7}, Direction{"forward", 1}, Direction{"forward", 3}, Direction{"forward", 8}, Direction{"forward", 2}, Direction{"forward", 7}, Direction{"forward", 7}, Direction{"forward", 5}, Direction{"down", 8}, Direction{"down", 5}, Direction{"up", 9}, Direction{"forward", 2}, Direction{"down", 9}, Direction{"forward", 4}, Direction{"up", 9}, Direction{"up", 6}, Direction{"up", 2}, Direction{"up", 7}, Direction{"down", 2}, Direction{"forward", 8}, Direction{"forward", 9}, Direction{"down", 3}, Direction{"down", 6}, Direction{"down", 9}, Direction{"forward", 8}, Direction{"down", 4}, Direction{"down", 2}, Direction{"down", 8}, Direction{"down", 8}, Direction{"down", 9}, Direction{"up", 4}, Direction{"down", 5}, Direction{"forward", 4}, Direction{"forward", 7}, Direction{"down", 6}, Direction{"down", 3}, Direction{"forward", 9}, Direction{"up", 1}, Direction{"forward", 3}, Direction{"down", 1}, Direction{"down", 8}, Direction{"up", 1}, Direction{"down", 7}, Direction{"down", 5}, Direction{"forward", 2}, Direction{"forward", 3}, Direction{"up", 9}, Direction{"down", 7}, Direction{"up", 9}, Direction{"up", 8}, Direction{"up", 3}, Direction{"forward", 7}, Direction{"down", 4}, Direction{"forward", 8}, Direction{"forward", 9}, Direction{"down", 8}, Direction{"forward", 1}, Direction{"up", 7}, Direction{"up", 4}, Direction{"down", 2}, Direction{"forward", 6}, Direction{"up", 4}, Direction{"forward", 1}, Direction{"up", 1}, Direction{"forward", 1}, Direction{"forward", 2}, Direction{"up", 4}, Direction{"up", 2}, Direction{"up", 3}, Direction{"forward", 3}, Direction{"forward", 9}, Direction{"forward", 2}, Direction{"down", 1}, Direction{"forward", 9}, Direction{"down", 7}, Direction{"forward", 3}, Direction{"down", 3}, Direction{"up", 1}, Direction{"down", 2}, Direction{"forward", 9}, Direction{"down", 7}, Direction{"down", 5}, Direction{"up", 5}, Direction{"down", 2}, Direction{"down", 5}, Direction{"forward", 6}, Direction{"forward", 9}, Direction{"forward", 3}, Direction{"down", 5}, Direction{"down", 9}, Direction{"forward", 6}, Direction{"up", 3}, Direction{"down", 6}, Direction{"down", 8}, Direction{"up", 8}, Direction{"down", 1}, Direction{"forward", 2}, Direction{"down", 1}, Direction{"up", 5}, Direction{"down", 8}, Direction{"up", 4}, Direction{"down", 2}, Direction{"forward", 6}, Direction{"forward", 1}, Direction{"down", 8}, Direction{"down", 4}, Direction{"down", 2}, Direction{"forward", 4}, Direction{"down", 7}, Direction{"up", 9}, Direction{"forward", 6}, Direction{"forward", 5}, Direction{"down", 4}, Direction{"down", 5}, Direction{"up", 6}, Direction{"down", 6}, Direction{"forward", 4}, Direction{"forward", 8}, Direction{"up", 9}, Direction{"down", 8}, Direction{"forward", 3}, Direction{"down", 4}, Direction{"down", 9}, Direction{"up", 8}, Direction{"down", 1}, Direction{"forward", 7}, Direction{"down", 5}, Direction{"down", 1}, Direction{"forward", 7}, Direction{"up", 4}, Direction{"forward", 9}, Direction{"up", 6}, Direction{"forward", 8}, Direction{"forward", 1}, Direction{"up", 8}, Direction{"up", 8}, Direction{"up", 3}, Direction{"forward", 2}, Direction{"forward", 3}, Direction{"forward", 6}, Direction{"forward", 8}, Direction{"forward", 4}, Direction{"down", 7}, Direction{"up", 8}, Direction{"forward", 4}, Direction{"down", 3}, Direction{"down", 2}, Direction{"down", 5}, Direction{"down", 4}, Direction{"up", 2}, Direction{"up", 1}, Direction{"down", 3}, Direction{"down", 5}, Direction{"down", 3}, Direction{"up", 6}, Direction{"down", 8}, Direction{"down", 7}, Direction{"down", 9}, Direction{"forward", 2}, Direction{"down", 4}, Direction{"up", 1}, Direction{"forward", 4}, Direction{"down", 9}, Direction{"forward", 8}, Direction{"up", 3}, Direction{"down", 4}, Direction{"down", 1}, Direction{"up", 3}, Direction{"forward", 7}, Direction{"forward", 9}, Direction{"down", 5}, Direction{"up", 7}, Direction{"forward", 4}, Direction{"down", 4}, Direction{"up", 8}, Direction{"down", 2}, Direction{"up", 7}, Direction{"up", 3}, Direction{"forward", 7}, Direction{"down", 7}, Direction{"down", 4}, Direction{"forward", 6}, Direction{"up", 5}, Direction{"forward", 8}, Direction{"down", 4}, Direction{"down", 9}, Direction{"forward", 3}, Direction{"down", 6}, Direction{"forward", 3}, Direction{"down", 8}, Direction{"down", 5}, Direction{"up", 6}, Direction{"down", 8}, Direction{"down", 9}, Direction{"down", 4}, Direction{"down", 9}, Direction{"forward", 7}, Direction{"down", 4}, Direction{"down", 7}, Direction{"forward", 1}, Direction{"down", 8}, Direction{"forward", 9}, Direction{"forward", 1}, Direction{"down", 5}, Direction{"down", 3}, Direction{"up", 5}, Direction{"forward", 9}, Direction{"down", 7}, Direction{"down", 7}, Direction{"up", 3}, Direction{"up", 9}, Direction{"down", 4}, Direction{"down", 5}, Direction{"forward", 5}, Direction{"down", 5}, Direction{"up", 3}, Direction{"forward", 8}, Direction{"down", 5}, Direction{"forward", 1}, Direction{"down", 1}, Direction{"up", 6}, Direction{"forward", 1}, Direction{"forward", 4}, Direction{"up", 9}, Direction{"up", 5}, Direction{"down", 6}, Direction{"down", 7}, Direction{"forward", 8}, Direction{"down", 4}, Direction{"forward", 9}, Direction{"forward", 6}, Direction{"down", 7}, Direction{"down", 2}, Direction{"up", 1}, Direction{"down", 4}, Direction{"down", 5}, Direction{"up", 7}, Direction{"down", 7}, Direction{"down", 3}, Direction{"up", 8}, Direction{"forward", 8}, Direction{"down", 8}, Direction{"forward", 5}, Direction{"down", 6}, Direction{"down", 7}, Direction{"forward", 4}, Direction{"up", 7}, Direction{"forward", 9}, Direction{"down", 7}, Direction{"up", 7}, Direction{"forward", 2}, Direction{"down", 3}, Direction{"down", 7}, Direction{"up", 6}, Direction{"forward", 2}, Direction{"down", 1}, Direction{"forward", 1}, Direction{"up", 9}, Direction{"forward", 6}, Direction{"forward", 2}, Direction{"forward", 5}, Direction{"up", 1}, Direction{"forward", 3}, Direction{"down", 2}, Direction{"up", 2}, Direction{"forward", 4}, Direction{"up", 2}, Direction{"down", 2}, Direction{"down", 9}, Direction{"up", 9}, Direction{"forward", 2}, Direction{"down", 3}, Direction{"down", 2}, Direction{"forward", 9}, Direction{"forward", 9}, Direction{"down", 3}, Direction{"down", 6}, Direction{"forward", 3}, Direction{"down", 7}, Direction{"up", 6}, Direction{"down", 8}, Direction{"forward", 4}, Direction{"forward", 8}, Direction{"forward", 1}, Direction{"down", 7}, Direction{"down", 2}, Direction{"forward", 6}, Direction{"forward", 9}, Direction{"up", 6}, Direction{"up", 8}, Direction{"down", 4}, Direction{"forward", 9}, Direction{"forward", 9}, Direction{"forward", 1}, Direction{"forward", 4}, Direction{"forward", 3}, Direction{"up", 9}, Direction{"down", 8}, Direction{"down", 7}, Direction{"down", 4}, Direction{"down", 1}, Direction{"down", 9}, Direction{"down", 7}, Direction{"forward", 2}, Direction{"forward", 7}, Direction{"forward", 6}, Direction{"down", 7}, Direction{"forward", 2}, Direction{"forward", 1}, Direction{"forward", 8}, Direction{"forward", 9}, Direction{"forward", 2}, Direction{"down", 5}, Direction{"down", 9}, Direction{"up", 7}, Direction{"forward", 4}, Direction{"up", 4}, Direction{"up", 2}, Direction{"down", 1}, Direction{"down", 3}, Direction{"down", 4}, Direction{"forward", 1}, Direction{"up", 4}, Direction{"up", 1}, Direction{"up", 4}, Direction{"down", 6}, Direction{"down", 7}, Direction{"forward", 7}, Direction{"down", 4}, Direction{"down", 6}, Direction{"forward", 4}, Direction{"forward", 8}, Direction{"down", 2}, Direction{"down", 8}, Direction{"down", 1}, Direction{"forward", 1}, Direction{"down", 7}, Direction{"down", 1}, Direction{"down", 6}, Direction{"up", 2}, Direction{"down", 9}, Direction{"forward", 3}, Direction{"down", 7}, Direction{"down", 8}, Direction{"down", 9}, Direction{"down", 4}, Direction{"down", 7}, Direction{"down", 4}, Direction{"up", 1}, Direction{"forward", 1}, Direction{"forward", 2}, Direction{"up", 6}, Direction{"up", 6}, Direction{"up", 5}, Direction{"forward", 4}, Direction{"down", 3}, Direction{"down", 9}, Direction{"forward", 3}, Direction{"up", 5}, Direction{"down", 1}, Direction{"forward", 7}, Direction{"down", 6}, Direction{"down", 2}, Direction{"up", 3}, Direction{"up", 8}, Direction{"up", 5}, Direction{"forward", 4}, Direction{"down", 5}, Direction{"forward", 5}, Direction{"forward", 2}, Direction{"down", 3}, Direction{"forward", 3}, Direction{"forward", 9}, Direction{"down", 9}, Direction{"down", 9}, Direction{"down", 9}, Direction{"forward", 8}, Direction{"forward", 5}, Direction{"up", 1}, Direction{"down", 5}, Direction{"forward", 3}, Direction{"forward", 4}, Direction{"forward", 2}, Direction{"up", 1}, Direction{"up", 2}, Direction{"up", 8}, Direction{"down", 6}, Direction{"up", 4}, Direction{"forward", 3}, Direction{"down", 1}, Direction{"down", 6}, Direction{"down", 6}, Direction{"up", 7}, Direction{"forward", 7}, Direction{"down", 8}, Direction{"up", 2}, Direction{"up", 1}, Direction{"up", 5}, Direction{"forward", 4}, Direction{"down", 4}, Direction{"down", 8}, Direction{"down", 2}, Direction{"down", 5}, Direction{"down", 4}, Direction{"up", 4}, Direction{"up", 7}, Direction{"forward", 2}, Direction{"forward", 4}, Direction{"forward", 2}, Direction{"down", 6}, Direction{"down", 9}, Direction{"up", 3}, Direction{"up", 6}, Direction{"forward", 8}, Direction{"forward", 6}, Direction{"forward", 8}, Direction{"forward", 9}, Direction{"down", 7}, Direction{"forward", 9}, Direction{"up", 4}, Direction{"up", 5}, Direction{"down", 4}, Direction{"down", 7}, Direction{"down", 4}, Direction{"up", 8}, Direction{"down", 3}, Direction{"forward", 8}, Direction{"down", 2}, Direction{"down", 4}, Direction{"forward", 2}, Direction{"up", 6}, Direction{"up", 6}, Direction{"down", 9}, Direction{"forward", 6}, Direction{"forward", 8}, Direction{"down", 5}, Direction{"forward", 5}, Direction{"down", 3}, Direction{"down", 6}, Direction{"up", 9}, Direction{"forward", 2}, Direction{"forward", 8}, Direction{"up", 4}, Direction{"forward", 4}, Direction{"forward", 2}, Direction{"down", 4}, Direction{"forward", 3}, Direction{"down", 1}, Direction{"up", 4}, Direction{"down", 4}, Direction{"up", 7}, Direction{"forward", 2}, Direction{"forward", 8}, Direction{"down", 8}, Direction{"up", 7}, Direction{"up", 2}, Direction{"down", 7}, Direction{"down", 7}, Direction{"forward", 8}, Direction{"forward", 7}, Direction{"forward", 9}, Direction{"down", 4}, Direction{"down", 5}, Direction{"down", 9}, Direction{"down", 5}, Direction{"forward", 1}, Direction{"down", 5}, Direction{"up", 2}, Direction{"forward", 6}, Direction{"forward", 9}, Direction{"down", 3}, Direction{"forward", 5}, Direction{"down", 1}, Direction{"forward", 9}, Direction{"up", 2}, Direction{"forward", 1}, Direction{"up", 9}, Direction{"forward", 2}, Direction{"up", 2}, Direction{"forward", 4}, Direction{"forward", 3}, Direction{"up", 9}, Direction{"forward", 4}, Direction{"forward", 2}, Direction{"up", 7}, Direction{"up", 6}, Direction{"down", 3}, Direction{"forward", 7}, Direction{"forward", 1}, Direction{"forward", 6}, Direction{"down", 2}, Direction{"down", 8}, Direction{"forward", 5}, Direction{"down", 5}, Direction{"up", 5}, Direction{"forward", 5}, Direction{"down", 6}, Direction{"forward", 8}, Direction{"forward", 4}, Direction{"down", 2}, Direction{"up", 8}, Direction{"forward", 6}, Direction{"down", 2}, Direction{"down", 7}, Direction{"forward", 5}, Direction{"down", 4}, Direction{"down", 5}, Direction{"up", 6}, Direction{"down", 3}, Direction{"up", 6}, Direction{"down", 8}, Direction{"forward", 4}, Direction{"down", 8}, Direction{"down", 6}, Direction{"forward", 2}, Direction{"forward", 8}, Direction{"up", 8}, Direction{"up", 1}, Direction{"forward", 6}, Direction{"down", 2}, Direction{"down", 6}, Direction{"forward", 2}, Direction{"forward", 2}, Direction{"down", 8}, Direction{"forward", 7}, Direction{"up", 1}, Direction{"forward", 1}, Direction{"down", 8}, Direction{"forward", 3}, Direction{"down", 6}, Direction{"forward", 7}, Direction{"forward", 7}, Direction{"up", 7}, Direction{"down", 1}, Direction{"forward", 8}, Direction{"up", 7}, Direction{"forward", 8}, Direction{"forward", 8}, Direction{"forward", 2}, Direction{"down", 5}, Direction{"up", 7}, Direction{"forward", 1}, Direction{"forward", 9}, Direction{"down", 9}, Direction{"forward", 8}, Direction{"down", 6}, Direction{"down", 7}, Direction{"up", 4}, Direction{"down", 5}, Direction{"forward", 6}, Direction{"down", 8}, Direction{"forward", 5}, Direction{"up", 6}, Direction{"up", 4}, Direction{"forward", 8}, Direction{"forward", 5}, Direction{"forward", 1}, Direction{"forward", 3}, Direction{"up", 2}, Direction{"down", 3}, Direction{"down", 4}, Direction{"down", 2}, Direction{"forward", 3}, Direction{"forward", 2}, Direction{"up", 8}, Direction{"up", 1}, Direction{"forward", 5}, Direction{"up", 8}, Direction{"down", 1}, Direction{"up", 4}, Direction{"down", 6}, Direction{"forward", 1}, Direction{"down", 3}, Direction{"up", 8}, Direction{"down", 3}, Direction{"forward", 5}, Direction{"down", 7}, Direction{"forward", 3}, Direction{"down", 1}, Direction{"up", 8}, Direction{"up", 6}, Direction{"down", 4}, Direction{"down", 8}, Direction{"up", 2}, Direction{"forward", 1}, Direction{"forward", 8}, Direction{"down", 4}, Direction{"down", 8}, Direction{"down", 4}, Direction{"up", 7}, Direction{"down", 4}, Direction{"forward", 7}, Direction{"forward", 4}, Direction{"forward", 2}, Direction{"forward", 6}, Direction{"up", 4}, Direction{"down", 5}, Direction{"forward", 8}, Direction{"forward", 9}, Direction{"down", 1}, Direction{"down", 2}, Direction{"down", 2}, Direction{"forward", 6}, Direction{"forward", 1}, Direction{"forward", 4}, Direction{"down", 7}, Direction{"forward", 6}, Direction{"forward", 9}, Direction{"down", 5}, Direction{"forward", 9}, Direction{"up", 4}, Direction{"down", 2}, Direction{"down", 6}, Direction{"forward", 5}, Direction{"forward", 3}, Direction{"forward", 1}, Direction{"forward", 6}, Direction{"forward", 5}, Direction{"forward", 2}, Direction{"down", 9}, Direction{"forward", 1}, Direction{"up", 5}, Direction{"down", 2}, Direction{"down", 8}, Direction{"up", 5}, Direction{"down", 3}, Direction{"up", 1}, Direction{"forward", 5}, Direction{"forward", 8}, Direction{"forward", 2}, Direction{"down", 3}, Direction{"down", 7}, Direction{"up", 3}, Direction{"up", 8}, Direction{"forward", 5}, Direction{"up", 8}, Direction{"forward", 4}, Direction{"down", 7}, Direction{"forward", 4}, Direction{"down", 1}, Direction{"forward", 2}, Direction{"forward", 8}, Direction{"up", 9}, Direction{"up", 1}, Direction{"forward", 3}, Direction{"down", 6}, Direction{"up", 5}, Direction{"down", 9}, Direction{"forward", 8}, Direction{"forward", 3}, Direction{"down", 7}, Direction{"forward", 3}, Direction{"up", 9}, Direction{"down", 2}, Direction{"forward", 2}, Direction{"down", 6}, Direction{"down", 9}, Direction{"down", 4}, Direction{"down", 4}, Direction{"down", 7}, Direction{"up", 6}, Direction{"up", 2}, Direction{"down", 7}, Direction{"forward", 5}, Direction{"up", 1}, Direction{"down", 6}, Direction{"up", 2}, Direction{"forward", 8}, Direction{"up", 6}, Direction{"forward", 4}, Direction{"down", 2}, Direction{"up", 5}, Direction{"down", 6}, Direction{"down", 5}, Direction{"forward", 8}, Direction{"forward", 4}, Direction{"down", 5}, Direction{"forward", 5}, Direction{"forward", 3}, Direction{"down", 2}, Direction{"forward", 3}, Direction{"up", 3}, Direction{"down", 4}, Direction{"up", 2}, Direction{"forward", 9}, Direction{"up", 8}, Direction{"forward", 9}, Direction{"up", 5}, Direction{"up", 1}, Direction{"forward", 9}, Direction{"down", 8}, Direction{"down", 3}, Direction{"forward", 9}, Direction{"up", 4}, Direction{"down", 1}, Direction{"forward", 8}, Direction{"down", 6}, Direction{"down", 1}, Direction{"down", 2}, Direction{"down", 9}, Direction{"down", 1}, Direction{"forward", 5}, Direction{"down", 7}, Direction{"forward", 1}, Direction{"up", 6}, Direction{"down", 1}, Direction{"down", 6}, Direction{"forward", 6}, Direction{"forward", 9}, Direction{"down", 6}, Direction{"forward", 1}, Direction{"forward", 4}, Direction{"up", 4}, Direction{"forward", 1}, Direction{"up", 9}, Direction{"forward", 3}, Direction{"forward", 8}, Direction{"down", 3}, Direction{"down", 7}, Direction{"forward", 4}, Direction{"up", 1}, Direction{"down", 1}, Direction{"forward", 1}, Direction{"down", 3}, Direction{"down", 6}, Direction{"down", 9}, Direction{"down", 6}, Direction{"forward", 4}, Direction{"down", 1}, Direction{"up", 4}, Direction{"forward", 8}, Direction{"down", 6}, Direction{"up", 6}, Direction{"down", 5}, Direction{"up", 8}, Direction{"forward", 4}, Direction{"up", 7}, Direction{"forward", 4}, Direction{"forward", 2}, Direction{"down", 7}, Direction{"forward", 2}}