seats_arr = [ x.strip() for x in open("./input.txt", "r").readlines() ]

# Convert into mutable lists in a matrix
seats_mat = []
for i in seats_arr:
    seats_mat.append([x for x in i])

# Get final index to see the maximum
# Used for functions below
row_len = len(seats_mat[0])
mat_len = len(seats_mat)

# Ecexution plan:
# for each seat check seat-1 and seat+1 in the same array as the seat, plus the array below and above the seat
# in the matrix

def count_row(symbol, row, check_seats):
    """
    Returns total occurances of a given symbol in a row. Allows you to skip the
    central seat, since an occupant would not look to see if their own seat
    was empty.
    """
    symbol_count = 0
    for seat in check_seats:
        if row[seat] == symbol:
            symbol_count += 1

    return symbol_count

def get_adjacent_items(check_item, len_int):
    """
    Returns a list of row/seat indexes to scrutinize
    The currently being investigated item is *always* item [0] in the returned array.
    """

    # Return false if the check_row int is too large, just in case.
    if check_item >= len_int:
        return False

    send_rows = []
    send_rows.append(check_item)

    for b in [ check_item - 1, check_item + 1 ]:
        if b >= 0 and b < len_int:
            send_rows.append(b)  

    return send_rows

def seat_modify(matrix, adj_rows, adj_same_row, curr_seat=(0,0)):
    """
    Determines, based on the rules in the challenge, if a seat changes state from
    occupied (#) to unoccupied (L). curr_seat is a tuple containing (row, column) or (i, t)
    """

    current_symbol = matrix[curr_seat[0]][curr_seat[1]]
    curr_adj_slice = matrix[adj_rows[0]:[adj_rows[-1] + 1]]

    total_occupied = 0
    total_unoccupied = 0

    for row in curr_adj_slice:
        
        total_occupied += count_total_symbols(curr_adj_slice[row], "#")
        total_unoccupied += count_total_symbols(curr_adj_slice[row], "L")

    for row in adj_rows:
        for seat in adj_same_row:
            if row == curr_seat[0] and seat == curr_seat[1]:
                return current_symbol
            
            if seat == "L" and 
                pass

def compare_matrices(new_mat, old_mat):
    """
    Compares the old seats matrix to the new one. The challenges states that if nothing changes, then
    we need to take action for the final result. Returns True/False if the same/different, respectively.
    """

    for i in range(len(new_mat)):
        if "".join(new_mat[i]) != "".join(old_mat[i]):
            return True
    
    return False

def count_total_symbols(matrix_row, symbol):
    total_occupied_seats = 0

    for seat in matrix_row:
        if seat == symbol:
            total_occupied_seats += 1

    return total_occupied_seats
            

# Boolean to check if the matrix is the same or not as the previous iteration
matrix_change = True

# For reassigning for comparison
old_matrix = seats_mat

while matrix_change:
    new_matrix = []

    # For each row in the matrix
    for i in range(0, mat_len):

        # Append the new seat icon to this list for building the new matrix
        new_row = []

        # For each seat in this row
        for t in range(0, row_len):

            adj_rows = get_adjacent_items(i, mat_len)
            adj_same_row = get_adjacent_items(t, row_len)

            new_row.append(old_matrix[i][t])

        
        # Add new row to new Matrix
        new_matrix.append(new_row)

    # Compare logic here....
    matrix_change = compare_matrices(new_matrix, old_matrix)

    old_matrix = new_matrix

part_one = 0
for row in old_matrix:
        part_one += count_total_symbols(row, "#")
print(part_one)