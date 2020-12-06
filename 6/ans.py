answers = open("./answers.txt", "r").read().split("\n\n")

# Part 1
one_line_answers = [ x.replace("\n", "") for x in answers ]
total_yes = 0

for answer in one_line_answers:
    total_yes += len(set(answer))

print(f"There are a total of {total_yes} yes answered questions.")

# Part 2
total_shared_yes = 0

for answer in answers:
    print(answer)
    shared_ans = ""

    split_ans = [ x.strip() for x in answer.split("\n") if len(x) > 0 ]
    print(split_ans)
    total_people = len(split_ans)
    print(total_people)

    uniq_qs = set(answer)

    for q in uniq_qs:
        q_count = answer.count(q)

        if q_count == total_people:
            shared_ans += q

    

    if len(shared_ans) > 0:
        total_shared_yes += len(shared_ans)

print(f"There are {total_shared_yes} answers shared in groups.")
