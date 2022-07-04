package main

import "math/rand"

func RandomChoice(prisonerNumber int, boxes []int, allowedTries int) (bool, int) {
	foundOwnNumber := false
	numBoxes := len(boxes)
	choices := make([]int, numBoxes)
	for i := 0; i < numBoxes; i++ {
		choices[i] = i
	}
	rand.Shuffle(numBoxes, func(i, j int) {
		iVal := choices[i]
		choices[i] = choices[j]
		choices[j] = iVal
	})
	log("Prisoner %v will run through choices: %v", prisonerNumber, choices[:allowedTries])
	choice := 0
	for t := 0; t < allowedTries; t++ {
		choice = choices[t]
		contents := boxes[choice]
		log("Prisoner %v looks in box %v and finds number %v", prisonerNumber, choice, contents)
		if contents == prisonerNumber {
			foundOwnNumber = true
			break
		}
	}
	return foundOwnNumber, choice
}

func SequenceChoice(prisonerNumber int, boxes []int, allowedTries int) (bool, int) {
	foundOwnNumber := false
	choice := prisonerNumber
	for t := 0; t < allowedTries; t++ {
		contents := boxes[choice]
		log("Prisoner %v looks in box %v and finds number %v", prisonerNumber, choice, contents)
		if contents == prisonerNumber {
			foundOwnNumber = true
			break
		}
		choice = contents
	}
	return foundOwnNumber, choice
}

func RunExperiment(numBoxes int, numPrisoners int, prisonerActions func(int, []int, int) (bool, int)) bool {
	log("%v boxes - %v prisoners", numBoxes, numPrisoners)
	log("Assigning numbers to boxes...")
	boxes := make([]int, numBoxes)
	for i := 0; i < numBoxes; i++ {
		boxes[i] = i
	}
	rand.Shuffle(numBoxes, func(i, j int) {
		iVal := boxes[i]
		boxes[i] = boxes[j]
		boxes[j] = iVal
	})
	log("boxes: %v", boxes)

	log("Prisoners begin choosing...")
	allowedTries := numBoxes / 2
	p := 0
	for p = 0; p < numPrisoners; p++ {
		foundOwnNumber, choice := prisonerActions(p, boxes, allowedTries)
		if foundOwnNumber {
			log("Prisoner %v found their number in box %v", p, choice)
		} else {
			log("Prisoner %v did not find their number after %v tries", p, allowedTries)
			break
		}
	}

	if p == numPrisoners {
		log("They live")
	} else {
		log("They die")
	}

	return p == numPrisoners
}
