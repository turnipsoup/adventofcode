# Thanks to this user for pointing out that I needed to NOT STRIP THE NEW LINES
# FOR SOME DAMN REASON: https://github.com/djotaku/adventofcode/blob/main/2020/Day_3/solution_1.py

tree_map = [ x for x in open("./map.txt", "r").readlines() ]

current_slope = ( 3, 1 ) # Part 1
slope_array = [ # Part 2
    ( 1, 1 ),
    ( 3, 1 ),
    ( 5, 1 ),
    ( 7, 1 ),
    ( 1, 2 )
]

def increment_slope(current_coords, slope_nums):
    """
    Increment current coorindates by the slope passed to the function
    """

    return [ current_coords[0] + slope_nums[0], current_coords[1] + slope_nums[1]]


def find_tree(tree_map_row, next_index):
    """
    Will take the tree map and find out if there is a tree at the 
    passed coordinates. Returns True or False if the spot is a tree
    """

    map_object = tree_map_row[next_index]

    if map_object == "#":
        return True
    else:
        return False


def read_map(tree_map, curr_coords, curr_slope):
    """
    Iterates over all of the rows in the map and returns how many trees
    The rider will hit
    """

    total_trees = 0
    next_index = current_coords[0] + curr_slope[0]

    for row in tree_map[1:]: # Skip first row

        if find_tree(row, next_index):
            total_trees += 1
        next_index = next_index + curr_slope[0]

        if next_index > len(row) - 2:
            next_index = next_index - len(row) + 1


    return total_trees


def read_map_skips(tree_map, curr_coords, curr_slope):
    """
    Iterates over all of the rows in the map and returns how many trees
    The rider will hit
    """

    total_trees = 0
    next_index = current_coords[0] + curr_slope[0]
    next_line = current_coords[1] + curr_slope[1]

    if curr_slope[1] != 2:
        for row in tree_map[1:]: # Skip first row

            if find_tree(row, next_index):
                total_trees += 1
            next_index = next_index + curr_slope[0]

            if next_index > len(row) - 2:
                next_index = next_index - len(row) + 1
    else:
        for row in tree_map[2::2]: # Skip first two rows and check every other row

            if find_tree(row, next_index):
                total_trees += 1
            next_index = next_index + curr_slope[0]

            if next_index > len(row) - 2:
                next_index = next_index - len(row) + 1


    return total_trees

        

if __name__ == "__main__":

    total_trees_array = []
    tree_array_product = 1

    for target_slope in slope_array:
        current_coords = [ 0, 0 ]

        total_trees = read_map_skips(tree_map, current_coords, target_slope)
        total_trees_array.append(total_trees)
        print(f"For slope {target_slope} there are a total of {total_trees} trees on this path.")


    for tree_count in total_trees_array:
        tree_array_product *= tree_count

    print(f"The product of all of the total trees for all slopes is {tree_array_product}")

