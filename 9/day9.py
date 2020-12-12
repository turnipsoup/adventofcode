inp = [ int(x.strip()) for x in open("./input", "r").readlines() ]

lowerb = 0
upperb = 26

def check_list(input_list):
    target_int = input_list[-1]
    target_nums = input_list[:-1]

    sum_nums = []

    for t in range(len(target_nums)):
        for y in range(len(target_nums)):
            if t == y:
                continue

            if t + y == target_int:
                sum_nums.append(t + y)

    if len(sum_nums) == 0:
        return target_int
    
    return False




for i in range(len(inp)):

    try:
        target_list = inp[lowerb:upperb]
        x = check_list(target_list)
        if x:
            print(x)

        lowerb += 1
        upperb += 1
    except:
        pass

