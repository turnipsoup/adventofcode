bank_data = [ int(x.strip()) for x in open("./bankdata.txt", "r").readlines() ]

# Part 1
for transact1 in bank_data:
    for transact2 in bank_data:
        if transact1 + transact2 == 2020:
            print(f"1: {transact1}, 2: {transact2} - Multiplied: {transact1*transact2}")
            

# Part 2

for transact1 in bank_data:
    for transact2 in bank_data:
        for transact3 in bank_data:
            if transact1 + transact2 + transact3 == 2020:
                print(f"1: {transact1}, 2: {transact2}, 3: {transact3} - Multiplied: {transact1*transact2*transact3}")