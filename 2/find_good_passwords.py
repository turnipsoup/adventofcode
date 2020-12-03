passwd = [ x.strip() for x in open("./passwords.txt", "r").readlines() ]


# Part 1
def get_pass_info(passwd_line):
    """
    Devide up each password into a usable object
    """

    splitpass = passwd_line.split(" ")

    nums = splitpass[0].split("-")

    return_obj = {
        "min_num": nums[0],
        "max_num": nums[1],
        "letter": splitpass[1][0],
        "pw": splitpass[2]
    }

    return return_obj

def examine_passwd(passwd_obj):
    """
    Compare the rules within the password object to the target password.
    Return True or False

    """
    target_letter = passwd_obj["letter"]
    min_num = passwd_obj["min_num"]
    max_num = passwd_obj["max_num"]
    pw = passwd_obj["pw"]

    letter_cnt = pw.count(target_letter)

    if letter_cnt >= int(min_num) and letter_cnt <= int(max_num):
        return True
    else:
        return False


def get_total_good_passwords(passwd_list):
    """
    Take the raw input from the file and iterate over all of it,
    returning the total "good" passwords
    """

    total_good = 0

    for pw in passwd_list:
        if new_rule_examine_oassword(get_pass_info(pw)):
            total_good += 1

    print(f"There are a total of {total_good} good passwords.")


# Part 2

def new_rule_examine_oassword(passwd_obj):
    """
    Checks to ensure the new password rules are followed
    Returns True or False

    """
    target_letter = passwd_obj["letter"]
    pos1 = int(passwd_obj["min_num"]) - 1
    pos2 = int(passwd_obj["max_num"]) - 1
    pw = passwd_obj["pw"]


    total_occur = [ pw[pos1], pw[pos2] ].count(target_letter)

    if total_occur > 0 and total_occur < 2:
        return True


get_total_good_passwords(passwd)