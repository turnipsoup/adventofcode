import itertools

inp = [ int(x.strip()) for x in open("./input", "r").readlines() ]

part1 = 0

def check_inp_slice(inp_slice):
    target_int = inp_slice[-1]
    check_nums = inp_slice[:-1]

    if (not any([x+y==target_int for x,y in itertools.combinations(check_nums, 2)])):
        return target_int
    else:
        return False


for i in range(len(inp)):
    t = i + 26
    check_slice = inp[i:t]
    
    if t == len(inp) + 1: # Dont need to continue after we reach the end of the input
        break

    z = check_inp_slice(check_slice)

    if z:
        part1 = z


part2 = None

for i in range(len(inp)):
    for j in range(i+2, len(inp)):
        ys = inp[i:j]
        if sum(ys)==part1:
            part2 = min(ys)+max(ys)

print(part1)
print(part2)