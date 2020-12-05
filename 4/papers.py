# Load data
passports = [ x.strip() for x in open("./passports.txt", "r").readlines() ]


# Define required subjects
col_nom = {
    "byr": "Mandatory",
    "iyr": "Mandatory",
    "eyr": "Mandatory",
    "hgt": "Mandatory",
    "hcl": "Mandatory",
    "ecl": "Mandatory",
    "pid": "Mandatory",
    "cid": "Optional"
}


def split_indv_docs(raw_pp_list):
    """
    For part 1.
    Takes a list of all passports and then splits them into individual passports
    """
    rtr_pp_list = []

    current_pp = []

    for pp in raw_pp_list:
        if len(pp) > 1: # If line is not blank, it is part if a passport
            current_pp.append(pp)
        else: # if line is blank, move on to next passport
            rtr_pp_list.append(" ".join(current_pp))
            current_pp = []

    rtr_pp_list.append(" ".join(current_pp))

    return rtr_pp_list

def detect_subjects(pp_row, subj_names):
    """
    For part 1.
    Takes a passport and determines if it has all of the proper fields
    """
    for col in subj_names:
        if col not in pp_row and subj_names[col] == "Mandatory":
            return False
    
    return True


def make_pp_obj(pp_row):
    """
    For part 2. Turn each passport into an object for easier
    data validation
    """
    rtn_obj = {}

    for field in pp_row.split():
        field_vals = field.split(":")[1]
        field_name = field.split(":")[0]

        rtn_obj[field_name] = field_vals

    return rtn_obj

def validate_pp_obj(pp_obj):
    """
    For part 2. Validate pp_object against new data rules.
    """
    byr = pp_obj["byr"]
    iyr = pp_obj["iyr"]
    eyr = pp_obj["eyr"]
    hgt = pp_obj["hgt"]
    hcl = pp_obj["hcl"]
    ecl = pp_obj["ecl"]
    pid = pp_obj["pid"]
    #cid = pp_obj["cid"]

    # BYR
    if len(byr) != 4 or int(byr) < 1920 or int(byr) > 2002:
        return False

    # IYR
    if len(iyr) != 4 or int(iyr) < 2010 or int(iyr) > 2020:
        return False

    # EYR
    if len(eyr) != 4 or int(eyr) < 2020 or int(eyr) > 2030:
        return False

    # HGT
    if "cm" not in hgt and "in" not in hgt:
        return False

    if "cm" in hgt:
        hgt_val = int(hgt.removesuffix("cm"))
        if hgt_val < 150 or hgt_val > 193:
            return False

    if "in" in hgt:
        hgt_val = int(hgt.removesuffix("in"))
        if hgt_val < 59 or hgt_val > 76:
            return False

    # HCL
    if hcl[0] != "#":
        return False

    try:
        int(hcl[1:], 16)
    except:
        return False

    # ECL
    if ecl not in [ "amb", "blu", "brn", "gry", "grn", "hzl", "oth" ]:
        return False

    # PID
    if len(pid) != 9:
        return False

    return True

def run_part_1():
    total_valid = 0
    split_passports = split_indv_docs(passports)

    for document in split_passports:
        if detect_subjects(document, col_nom):
            total_valid += 1

    print(f"Part 1: There are {total_valid} valid passports.")

def run_part_2():
    split_passports = split_indv_docs(passports)
    valid_pp = []
    valid_total = 0

    for document in split_passports:
        if detect_subjects(document, col_nom):
            valid_pp.append(document)

    for valid_doc in valid_pp:
        pp_obj = make_pp_obj(valid_doc)
        if validate_pp_obj(pp_obj):
            valid_total += 1

    print(f"Part 2: There are {valid_total} valid passports.")

if __name__ == "__main__":
    run_part_1()
    run_part_2()



