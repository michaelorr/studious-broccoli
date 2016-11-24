class Allergies(object):

    def __init__(self, allergy_score):
        mapping = [
            (128, 'cats'),
            (64, 'pollen'),
            (32, 'chocolate'),
            (16, 'tomatoes'),
            (8, 'strawberries'),
            (4, 'shellfish'),
            (2, 'peanuts'),
            (1, 'eggs'),
        ]

        self.lst = []
        for item in mapping:
            if allergy_score >= item[0]:
                allergy_score -= item[0]
                self.lst.append(item[1])

    def is_allergic_to(self, obj):
        return obj in self.lst
