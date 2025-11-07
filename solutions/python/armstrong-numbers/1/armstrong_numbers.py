def is_armstrong_number(number):
    digits = [int(i) for i in str(number)]
    multiplier = len(digits)
    armnum = sum(i ** multiplier for i in digits)
    res = False
    if armnum == number:
        res = True

    return res
