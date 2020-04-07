package knapsack

import (
	"fmt"
	"math/rand"
)

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func getInitialPopulation() {
	population := []int{}
	for i := 0; i < len(population); i++ {
		population[i] = randomInt(0, 1)
	}
}

func getFirstPopulation(size int) {
	population := []int{}
	for i := 0; i < size; i++ {
			slice[i] = getInitialPopulation()
	}
}

func calculateWeight(backpack []) {
	weight := 0
	for i := 0; i < len(backpack); i++ {
		weight += BACKPACK_DEFAULT[i][1] * i
	}
}

func calculateReward(backpack []) {
	weight := 0
	for i := 0; i < len(backpack); i++ {
		weight += BACKPACK_DEFAULT[i][2] * i
	}
}

func createWeightInBackpacks(backpacks []) {
	return calculateWeight(backpacks)
}

func createRewardsInBackpacks(backpacks []) {
	return calculateReward(backpacks)
}

func createGenerations() {
	lifeCycle := 3
	generations = 50
	populations = getFirstPopulation(generations)


}
