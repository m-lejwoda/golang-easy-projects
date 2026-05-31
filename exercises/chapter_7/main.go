package main

import (
	"fmt"
	"io"
	"os"
	"slices"
)

func main() {
	liga := League{
		Teams: []Team{},
		Wins:  make(map[string]int),
	}

	liga.MatchResult("Gryfy", 5, "Smoki", 3)
	liga.MatchResult("Wilki", 2, "Gryfy", 4)
	liga.MatchResult("Smoki", 6, "Wilki", 1)

	liga.Ranking()
	RankPrinter(liga, os.Stdout)
}

type Ranker interface {
	Ranking() []string
}

type Team struct {
	TeamName string
	LastName string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l League) MatchResult(firstTeam string, FirstNumOfPoints int, secondTeam string, SecondNumOfPoints int) {
	if FirstNumOfPoints > SecondNumOfPoints {
		l.Wins[firstTeam] += 1
	} else if SecondNumOfPoints > FirstNumOfPoints {
		l.Wins[secondTeam] += 1
	}
}

func (l League) Ranking() []string {
	keys := make([]string, 0, len(l.Wins))
	for k := range l.Wins {
		keys = append(keys, k)
	}
	slices.SortFunc(keys, func(a, b string) int {
		return l.Wins[b] - l.Wins[a]
	})

	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, l.Wins[k])
	}

	return keys
}

func RankPrinter(ranker Ranker, writer io.Writer) {
	ranks := ranker.Ranking()
	for _, rank := range ranks {
		_, err := fmt.Fprintln(writer, rank)
		if err != nil {
			return
		}
	}
}
