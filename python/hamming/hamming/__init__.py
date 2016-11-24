from itertools import izip


def distance(first, second):
    distance = 0
    for a, b in izip(first, second):
        if a != b:
            distance += 1
    return distance
