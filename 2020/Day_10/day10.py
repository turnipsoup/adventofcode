inp = [ int(x.strip()) for x in open("./input.txt", "r").readlines() ]

sinp = [0] + sorted(inp)
sinp.append(max(sinp) + 3)


# Part 1
jolt_diffs = []
for i in range(len(sinp) - 1):
    t = i + 1
    jolt_diffs.append(sinp[t] - sinp[i])

ones = jolt_diffs.count(1)
threes = jolt_diffs.count(3)

print(f"Part 1: {ones * threes}")

# Part 2
# A smart cookie helped me here, this was a totally interesting solution and will help me think differently
# About arrays and pathways going forward.
# https://www.reddit.com/r/adventofcode/comments/ka8z8x/2020_day_10_solutions/gf9mvrh/
routes = {}

routes[0] = 1

for x in sinp[1:]:
    routes[x] = routes.get(x-1, 0) + routes.get(x-2, 0) + routes.get(x-3, 0)

print(f"Part 2: {routes[max(routes.keys())]}")
