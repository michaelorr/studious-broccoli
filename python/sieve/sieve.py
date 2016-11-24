def sieve(limit):
    possibles = range(2, limit + 1)
    for poss in possibles:
        count = 2
        while poss * count <= limit:
            try:
                possibles.remove(poss * count)
            except ValueError:
                pass
            count += 1
    return possibles
