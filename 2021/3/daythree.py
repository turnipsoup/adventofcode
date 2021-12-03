### Load Input
inp = [ x for x in open("./input.txt", "r").readlines() ]

## Part One
from collections import Counter

def most_frequent(List):
    occurence_count = Counter(List)
    return occurence_count.most_common(1)[0][0]

def get_gamma_epsilon(letter, inp_int):
    return_str = ""
    if letter == "gamma":
        for i in inp_int:
            return_str += inp_int[i]

    if letter == "epsilon":
        for i in inp_int:
            if inp_int[i] == "0":
                return_str += "1"

            if inp_int[i] == "1":
                return_str += "0"

    return return_str

all_numbers = {0: [], 1: [], 2: [], 3: [], 4: [], 5: [], 6: [], 7: [], 8: [], 9: [], 10: [], 11: []}

for binary in inp:
    for i in range(len(binary) - 1):
        all_numbers[i].append(binary[i])

common = {}

for i in range(len(inp[0]) - 1):
    common[i] = most_frequent(all_numbers[i])


gamma = get_gamma_epsilon("gamma", common)
epsilon = get_gamma_epsilon("epsilon", common)

print("Gamma: ", gamma)
print("Epsilon: ", epsilon)
print("Power Consumption: ", int(gamma, 2) * int(epsilon, 2))

## Part Two


