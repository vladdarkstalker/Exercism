def is_valid_truangle(sides):
    a, b, c = sides
    return all(i > 0 for i in sides) and a + b >= c and b + c >= a and a + c >= b

def equilateral(sides):
    if not is_valid_truangle(sides):
        return False
    return sides[0] == sides[1] == sides[2]

def isosceles(sides):
    if not is_valid_truangle(sides):
        return False
    return sides[0] == sides[1] or sides[1] == sides[2] or sides[0] == sides[2]

def scalene(sides):
    if not is_valid_truangle(sides):
        return False
    return sides[0] != sides[1] and sides[1] != sides[2] and sides[0] != sides[2]
