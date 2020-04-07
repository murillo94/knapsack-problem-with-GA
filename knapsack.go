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

	for i := 0; i < lifeCycle; i++ {
    weights = createWeightBackpacks(populations)
    rewards = createRewardsBackpacks(populations)
    selections = zip(populations, weights, rewards)
    selections = filter(lambda x: x[1] <= 30, selections)
    selections = sorted(list(selections), key=lambda x: x[2],  reverse=True)
    selections = selections[:int(generations/2)]
    population := []int{}

		for i := 2; i < selections; i-- {
				index = random.randint(0, len(selections) - 1)
				a = selections[index][0]
				selections.pop(index)

				index = random.randint(0, len(selections) - 1)
				b = selections[index][0]
				selections.pop(index)

				for i in range(int(size/2)):
					child = []
					for i in range(6):
						if random.randint(0, 1) == 1:
							child.append(a[i])
						else:
							child.append(b[i])
					childs.append(child)
		}
		populations = filter(lambda x: calculateWeight(x) <= 30, childs)
	}
  return sorted(list(populations), key=lambda x: calculateReward(x),  reverse=True)
}
