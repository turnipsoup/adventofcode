boarding_passes = [ x.strip() for x in open("./bpasses.txt", "r").readlines() ]

dir_map = {
    "F": "lower",
    "B": "upper",
    "L": "lower",
    "R": "upper"
}

starting_values = ( 0, 127 )

def get_bounds(current_min_val, current_max_val, direction):
    """
    Returns upper and lower portion of range as defined by challenge rules for
    seating. Takes "upper" or "lower" to support both F/B and L/R.
    """

    range_list = [ x for x in range(current_min_val, current_max_val + 1) ]
    halfway_mark = (len(range_list) // 2)

    if direction == "lower":
        target_range = range_list[:halfway_mark]
        return (target_range[0], target_range[-1])

    if direction == "upper":
        target_range = range_list[halfway_mark:]
        return (target_range[0], target_range[-1])

   
def check_pass(bpass, current_values):
    min_val = current_values[0]
    max_val = current_values[1]
    col_min_val = 0
    col_max_val = 7
    row = 0
    col = 0
    seat = 0

    fb = bpass[:-3]
    lr = bpass[-3:]

    for letter in fb:
        bounds = get_bounds(min_val, max_val, dir_map[letter])
        min_val = bounds[0]
        max_val = bounds[1]

        if min_val == max_val:
            row = min_val

    for letter in lr:
        bounds = get_bounds(col_min_val, col_max_val, dir_map[letter])
        col_min_val = bounds[0]
        col_max_val = bounds[1]

        if col_min_val == col_max_val:
            col = col_min_val

    seat = (row * 8) + col

    return (row, col, seat)

def find_my_seat(row_id_list):
    for i in range(min(id_list), max(id_list) + 1):
        if i not in row_id_list:
            return i


if __name__ == "__main__":

    # Part 1:
    id_list = []

    for bpass in boarding_passes:
        pass_info = check_pass(bpass, starting_values)
        id_list.append(pass_info[2])

    max_id = max(id_list)
    print(f"The max ID number in the list is {max(id_list)}")


    # Part 2
    my_seat = find_my_seat(id_list)
    print(f"The only empty seat is {my_seat}")

