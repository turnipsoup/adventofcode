inp = [ x.strip() for x in open("./dir.txt", "r").readlines() ]

angle_map = {
    0: "N",
    90: "E",
    180: "S",
    270: "W",
    360: "N"
}

cur_pos = {
    "Y": 0,
    "X": 0
}

w_pos = {
    "Y": 1,
    "X": 10
}

movements = []

cur_deg = 90

# Going to need two functions -
# 1. If instruction is a rotation, update the value of the cur_deg variable by the rotation amount. If negative, return cur_deg + 360
# 2. If instruction is a movement, log the movement direction in the movement tracker list. If the instruction is "F", copy
# The previous movement direction and add the new additional movement to the tracker list

def calc_rotate(cur_deg, rot_int):
    """
    If instruction is a rotation, update the value of the cur_deg variable by the rotation amount. If negative, return cur_deg + 360
    """
    direction, degrees = rot_int[0], int(rot_int[1:])

    if cur_deg == 360:
        cur_deg = 0

    if cur_deg + degrees > 360:
        cur_deg = (cur_deg + degrees) - cur_deg
        return cur_deg

    if direction == "L":
        cur_deg -= degrees
    if direction == "R":
        cur_deg += degrees

    if cur_deg < 0:
        return cur_deg + 360

    return cur_deg

def alt_calc_rotate(cur_deg, rot_int):
    """
    Tried this after doing some research after my math was wrong
    https://github.com/sophiebits/adventofcode/blob/main/2020/day12.py
    This entire thing is about learning and fun, so I chose to learn!
    """
    direction, degrees = rot_int[0], int(rot_int[1:])

    if direction == "L":
        cur_deg -= degrees
    if direction == "R":
        cur_deg += degrees

    return cur_deg


def find_f_dir(cur_def): 
    """
    Tried this after doing some research after my math was wrong
    https://github.com/sophiebits/adventofcode/blob/main/2020/day12.py
    This entire thing is about learning and fun, so I chose to learn!
    """

    action = ""
    
    if cur_def % 360 == 90:
        action = 'E'
    if cur_def % 360 == 0:
        action = 'N'
    if cur_def % 360 == 270:
        action = 'W'
    if cur_def % 360 == 180:
        action = 'S'

    return action



def calc_move(cur_deg, movements_list, move_int):
    """
    If instruction is a movement, log the movement direction in the movement tracker list. If the instruction is "F", copy
    The previous movement direction and add the new additional movement to the tracker list
    """
    direction, distance = move_int[0], int(move_int[1:])

    ml = movements_list[:]

    if direction != "F":
        ml.append([direction, distance])
    else:
        new_dir = find_f_dir(cur_deg)
        ml.append([new_dir, distance])

    return ml


def calc_manhat(movements_list, cur_p):
    """
    Calculate the Manhattan distance the ship has travelled after all of the movements are recorded, updating the cur_pos object.
    """

    for move in movements_list:
        direction = move[0]
        distance = move[1]
        
        if direction == "S":
            cur_p["Y"] -= distance
        
        if direction == "N":
            cur_p["Y"] += distance
        
        if direction == "W":
            cur_p["X"] -= distance

        if direction == "E":
            cur_p["X"] += distance

    return abs(cur_p["X"]) + abs(cur_p["Y"])



for inst in inp:
    letter = inst[0]
    mv = inst[1:]
    
    if letter == "L" or letter == "R":
        cur_deg = alt_calc_rotate(cur_deg, inst)
    else:
        movements = calc_move(cur_deg, movements, inst)

partone = calc_manhat(movements, cur_pos)

print(partone)