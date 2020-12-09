bootcode = [ x.strip() for x in open("input.txt", "r").readlines() ]

# Part 1
# Do the math with the accumulator
def acc_work(split_inst, accumulator, next_inst):
    curr_acc_action = split_inst[1][0]
    curr_acc_action_value = int(split_inst[1][1:])

    if curr_acc_action == "+":
        accumulator += curr_acc_action_value
    elif curr_acc_action == "-":
        accumulator -= curr_acc_action_value

    next_inst += 1
    
    return accumulator, next_inst

# Do the jmp math
def jmp_work(split_inst, next_inst):
    curr_jmp_action = split_inst[1][0]
    curr_jmp_action_value = int(split_inst[1][1:])

    if curr_jmp_action == "+":
        next_inst += curr_jmp_action_value
    if curr_jmp_action == "-":
        next_inst -= curr_jmp_action_value

    return next_inst

# Just increment next_instruction index by 1
def nop_work(split_inst, next_inst):
    next_inst += 1
    return next_inst


def check_prog(bootcode_list):
    accumulator = 0
    next_inst = 0
    has_run = set()

    while True:
        if next_inst == len(bootcode_list):
            return accumulator
        if next_inst in has_run:
            # Uncomment the following two lines for part 1
            #print(f"The accumulator is at value {accumulator} right before it runs a duplicate instruction.")
            #exit(0)
            return None

        has_run.add(next_inst)

        split_inst = bootcode[next_inst].split()
        curr_inst = split_inst[0]

        if curr_inst == "acc":
            accumulator, next_inst = acc_work(split_inst, accumulator, next_inst)
        elif curr_inst == "jmp":
            next_inst = jmp_work(split_inst, next_inst)
        elif curr_inst == "nop":
            next_inst = nop_work(split_inst, next_inst)

# Part 1 execution
#check_prog(bootcode)

# Part 2 execution

for i in range(len(bootcode)):
    program = bootcode[:]

    if program[i].startswith("jmp"):
        program[i] = program[i].replace("jmp", "nop")
    elif program[i].startswith("nop"):
        program[i] = program[i].replace("nop", "jmp")
    
    y = check_prog(program)

    if y:
        print(y)



