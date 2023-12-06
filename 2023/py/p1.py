# numMap = {
#     "one": 1,
#     "two": 2,
#     "three": 3,
#     "four": 4,
#     "five": 5,
#     "six": 6,
#     "seven": 7,
#     "eight": 8,
#     "nine": 9,
#     "1": 1,
#     "2": 2,
#     "3": 3,
#     "4": 4,
#     "5": 5,
#     "6": 6,
#     "7": 7,
#     "8": 8,
#     "9": 9,
# }


# file = open("../inputs/p1input.txt")
# # first, last, sum = 0, 0, 0

# input = []
# for line in file.readlines():
#     input.append(line.rstrip("/\n"))


# def checkDict(str):
#     for key in numMap.keys():
#         if key in str:
#             return key
#     return None
# # P1
# # for x in input:
# #     for y in x[0]:
# #         if first == 0 and y.isdigit():
# #             first = y
# #         elif y.isdigit():
# #             last = y
# #     if last == 0:
# #         last = first
# #     sum += int((first + last))
# #     last, first = 0, 0


# # P2
# this isnt my solution, I wanted to see how far off I was from the answer, i was off by 5
# import regex


# def spam(spam):
#     """spam"""
#     spam = regex.findall(r'(\d|one|two|three|four|five|six|seven|eight|nine)',
#                          spam, overlapped=True)
#     print(spam)
#     spam = [{'one': '1', 'two': '2', 'three': '3', 'four': '4', 'five': '5',
#              'spam': 'spam', 'six': '6', 'seven': '7', 'eight': '8', 'nine': '9'}.get
#             (spam, spam)
#             for spam in spam]
#     return int(spam[0] + spam[-1])


# for item in map(spam, open("../inputs/p1input.txt")):
#     print(item)
