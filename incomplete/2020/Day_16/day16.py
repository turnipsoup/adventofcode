inp = [ x.strip() for x in open("./input.txt", "r").readlines() ]

rules = inp[:20]
my_ticket = inp[22]
tickets = inp[25:]

bad_values = []

# Build rules obj
rules_obj = {}

for rule in rules:
    rule_name, rule_vales = rule.split(":")
    rules_obj[rule_name] = {}

    i = 0
    for value in [ x.strip() for x in rule_vales.split("or") ]:
        lower_val, upper_val = value.split("-")
        rules_obj[rule_name][i] = [ x for x in range(int(lower_val), int(upper_val) + 1)]

        i +=1

def check_tickets(ticket_val, rules_object):
    for rule in rules_object:
        for rule_list in rules_object[rule]:
            if ticket_val in rules_object[rule][rule_list]:
                return True
    
    return False


for ticket in tickets:
    ticket_values = [ int(x) for x in ticket.split(",") ]

    for val in ticket_values:
        if check_tickets(val, rules_obj) == False:
            bad_values.append(val)

    
partone = sum(bad_values)
print(partone)
