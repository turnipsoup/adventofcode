inp = [ x.strip() for x in open("./inp.txt", "r").readlines() ]

stamp = int(inp[0])
busses = inp[1].split(",")

# Take each bus and make a list of all the times it can go until it is large enough
# Append this list to a map with key of bus number
# Find the first integer that is larger than our stamp in each list
# Find the smallest value of this list
# return (smallest value larger than stamp) - stamp * bus_id

bus_map = {}
e_map = {}


for i in busses:
    if i != "x":
        i = int(i)
        range_list = [0]
        for x in range(100000):
            range_list.append(x*i)
        
        bus_map[i] = range_list

for bus in bus_map:
    bus_list = [ x for x in bus_map[bus] ]
    for i in bus_list:
        if i > stamp:
            e_map[bus] = i
            break

e_list = [ e_map[x] for x in e_map ]

for e in e_map:
    if e_map[e] == min(e_list):
        wait_time =  e_map[e] - stamp
        print(wait_time * e)
