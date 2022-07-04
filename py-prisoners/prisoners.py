#!/usr/bin/env python3

import argparse
import random
import time

parser = argparse.ArgumentParser(description = "Prisoner Box Choice Simulation")
parser.add_argument("boxes", type = int,
                     help = "How many numbered boxes are in the box room.")
parser.add_argument("prisoners", type = int,
                     help = "How many prisoners will enter the box room.")
parser.add_argument("runs", type = int,
                     help = "How many times to repeat the test.")
parser.add_argument("method", choices=['random', 'sequence'])

def main():
    args = parser.parse_args()
    print(args.boxes, args.prisoners, args.runs, args.method)

    if args.method == "random":
        experimentFn = RandomChoice
    else:
        experimentFn = SequenceChoice

    random.seed(time.time())
    live = 0
    die = 0
    for _ in range(args.runs):
        if RunExperiment(args.boxes, args.prisoners, experimentFn):
            live+=1
        else:
            die+=1
    escapePercent = 100 * float(live) / float(live+die)
    print(f"%.2f%% escapes"% escapePercent)

def AssignNumbersToBoxes(numBoxes, numPrisoners):
    print(f"%d boxes - %d prisoners"% (numBoxes, numPrisoners))
    print("Assigning numbers to boxes...")
    boxes = [i for i in range(numBoxes)]
    random.shuffle(boxes)
    print(f"boxes: %s"% boxes)
    return boxes

def RandomChoice(prisonerNumber, boxes, allowedTries):
    choices = [i for i in range(len(boxes))]
    random.shuffle(choices)
    print(f"Prisoner %d will run through choices: %s"% (prisonerNumber, choices[:allowedTries]))
    for t in range(allowedTries):
        choice = choices[t]
        contents = boxes[choice]
        print(f"Prisoner %d looks in box %d and finds number %d"% (prisonerNumber, choice, contents))
        if contents == prisonerNumber:
            return True, choice
    return False, 0

def SequenceChoice(prisonerNumber, boxes, allowedTries):
    choice = prisonerNumber
    for _ in range(allowedTries):
        contents = boxes[choice]
        print(f"Prisoner %d looks in box %d and finds number %d"% (prisonerNumber, choice, contents))
        if contents == prisonerNumber:
            return True, choice
        choice = contents
    return False, 0

def RunExperiment(numBoxes, numPrisoners, prisonerActions):
    boxes = AssignNumbersToBoxes(numBoxes, numPrisoners)
    print("Prisoners begin choosing...")
    allowedTries = int(numBoxes / 2)
    allFoundTheirNumber = True
    for p in range(numPrisoners):
        foundOwnNumber, choice = prisonerActions(p, boxes, allowedTries)
        if foundOwnNumber:
            print(f"Prisoner %d found their number in box %d"% (p, choice))
        else:
            print(f"Prisoner %d did not find their number after %d tries"% (p, allowedTries))
            allFoundTheirNumber = False
            break

    if allFoundTheirNumber:
        print("They live")
    else:
        print("They die")

    return allFoundTheirNumber


if __name__ == "__main__":
    main()