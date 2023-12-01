### Load Data

inp = [(x.split()[0], int(x.split()[1])) for x in open("./input.txt", "r").readlines()]

## Part One
horz = 0
depth = 0

for command in inp:
    match command:
        case ("forward", x):
            horz += x
        case ("down", x):
            depth += x
        case ("up", x):
            depth -= x
            

print(horz * depth)

## Part Two
horz = 0
depth = 0
aim = 0

for command in inp:
    match command:
        case ("forward", x):
            horz += x
            depth += aim*x
        case ("down", x):
            #depth += x
            aim += x
        case ("up", x):
            #depth -= x
            aim -= x

print(horz*depth)