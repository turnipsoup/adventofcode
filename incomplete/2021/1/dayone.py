## Load dataset
depths = [int(x) for x in open("./input_one.txt").readlines()]

## Part One
increased = 0

for i in range(len(depths)):
    if i != 0:
        if depths[i] > depths[i-1]:
            increased += 1

print(increased)

## Part Two
increased = 0
prev_sum = 0

for i in range(len(depths)):
    if i == 0: 
        prev_sum = sum(depths[i:i+3])
        continue
        
    try:
        curr_sum = sum(depths[i:i+3])

        if curr_sum > prev_sum: 
            increased += 1

        prev_sum = curr_sum

    except:
        pass

print(increased)

