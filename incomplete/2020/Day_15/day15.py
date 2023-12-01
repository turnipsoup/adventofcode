nums = [5,2,8,16,18,0,1]

# Need a function that will go back and find the previous turn that the prior number was spoken
# Current turn will always be i, previous turn is i - 1, and i is in range(2021)
# Then append this value to nums[] and do it again for that newly appended number

def ind_list_gen(nums_list, number):
    """
    https://thispointer.com/python-how-to-find-all-indexes-of-an-item-in-a-list/
    """
    last_spoke = number
    ind_pos_list = []
    ind_pos = 0

    
    while True:
        try:
            ind_pos = nums.index(last_spoke, ind_pos)
            ind_pos_list.append(ind_pos)
            ind_pos += 1
        except:
            break

    return ind_pos_list


for i in range(30000001):
    # Skip the first one, its going to be 0
    if i == 0:
        nums.append(0)
        continue

    turn = i + 1
    if i == len(nums) - 1:
        i_count = ind_list_gen(nums, nums[i])

        if len(i_count) > 1:
            app_num = i_count[-1] - i_count[-2]
            nums.append(app_num)

        else:
            nums.append(0)

    
    if turn == 30000000:
        print(f"Turn: {turn} - {nums[i]}")
            
    