import random

BACKPACK = [["Sleeping bag", 15, 15], ["Rope", 3, 7], ["Switchblade", 2, 10],
            ["Torch", 5, 5], ["Bottle", 9, 8], ["Food", 20, 17]]


def getInitialPopulation():
    population = []
    for i in range(6):
        population.insert(0, (random.randint(0, 1)))
    return population


def getTotalWeight(backpack):
    weight = 0
    for idx, i in enumerate(backpack):
        weight += BACKPACK[idx][1] * i
    return weight


def getTotalReward(backpack):
    reward = 0
    for idx, i in enumerate(backpack):
        reward += BACKPACK[idx][2] * i
    return reward


def weightBackpacks(backpacks):
    return map(getTotalWeight, backpacks)


def rewardsBackpacks(backpacks):
    return map(getTotalReward, backpacks)


def getSelection(selections):
    minimum = 1
    index = random.randint(0, len(selections) - minimum)
    a = selections[index][0]
    selections.pop(index)

    index = random.randint(0, len(selections) - minimum)
    b = selections[index][0]
    selections.pop(index)

    return a, b


def createGenerations():
    lifeCycle = 5
    generations = 50
    populations = []

    for i in range(generations):
        populations.insert(0, (getInitialPopulation()))

    for cycle in range(lifeCycle):
        selections = filter(lambda x: x[1] <= 30, zip(
            populations,
            weightBackpacks(populations),
            rewardsBackpacks(populations)
        ))
        selections = sorted(
            list(selections), key=lambda x: x[2], reverse=True
        )
        selections = selections[:int(generations/2)]
        sub_populations = []

        while len(selections) > 2:
            a, b = getSelection(selections)

            for i in range(int(generations/2)):
                child = []
                for i in range(6):
                    if random.randint(0, 1) == 1:
                        child.insert(0, (a[i]))
                    else:
                        child.insert(0, (b[i]))
                sub_populations.insert(0, (child))
        populations = filter(lambda x: getTotalWeight(x)
                             <= 30, sub_populations)
    return sorted(list(populations), key=lambda x: getTotalReward(x),  reverse=True)


def getBestGeneration():
    generation = (createGenerations())

    return list(
        zip(
            generation,
            weightBackpacks(generation),
            rewardsBackpacks(generation)
        )
    )[0]


def main():
    print("the best items to choose is: {result}".format(
        result=getBestGeneration())
    )


if __name__ == "__main__":
    main()