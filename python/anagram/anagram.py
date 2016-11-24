from collections import Counter


def detect_anagrams(word, candidates):
    return [w for w in candidates if Counter(w.lower()) == Counter(word.lower()) and word.lower() != w.lower()]
