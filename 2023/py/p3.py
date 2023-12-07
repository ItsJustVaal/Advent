# go through each row looking for symbols
# if found, look up, down, left right
# if num found, look left and right of num, append
# if not look diag, repeat
# sum final array
# Part 1
with open("../inputs/p3input.txt") as file:
    lines = [line.strip() for line in file.readlines()]


def checkUpDown(currIndex, row, rowIndex):
    if not checkColOutOfBounds(rowIndex):
        if row[currIndex].isdigit():
            return True
        return False


def checkDiagLeft(currIndex, row):
    if row[currIndex - 1].isdigit():
        return checkLeft(currIndex - 1, row)


def checkDiagRight(currIndex, row):
    if row[currIndex + 1].isdigit():
        return checkRight(currIndex, row)


def checkLeft(currIndex, row):
    if checkIfSymbol(row[currIndex]):
        tmpStr = ""
    else:
        tmpStr = row[currIndex]
    while currIndex >= 0:
        if not checkRowOutOfBounds(currIndex, row):
            if row[currIndex - 1].isdigit():
                tmpStr += row[currIndex - 1]
                currIndex -= 1
            else:
                break
        else:
            break
    return tmpStr[::-1]


def checkRight(currIndex, row):
    tmpStr = ""
    if currIndex + 1 > len(row):
        return
    while currIndex < len(row):
        if not checkRowOutOfBounds(currIndex, row):
            if row[currIndex + 1].isdigit():
                tmpStr += row[currIndex + 1]
                currIndex += 1
            else:
                break
        else:
            break
    return tmpStr


def checkIfSymbol(char):
    if char == '*':
        # if not char.isdigit() and not char == '.' and (ord(char) <= 64 or ord(char) >= 91 and ord(char) <= 96 or ord(char) >= 123 and ord(char) <= 126):
        return True
    return False


def checkRowOutOfBounds(currIndex, row):
    if currIndex == 0 or currIndex == len(row) - 1:
        return True
    return False


def checkColOutOfBounds(currRow):
    if currRow == 0 or currRow >= len(lines):
        return True
    return False

# P1
# nums = []
# for i, line in enumerate(lines):
#     for j, char in enumerate(line):
#         if checkIfSymbol(char):
#             checkUp = checkUpDown(j, lines[i - 1], i)
#             checkDown = checkUpDown(j, lines[i + 1], i)
#             if checkUp:
#                 checkleft = checkLeft(j, lines[i - 1])
#                 checkright = checkRight(j, lines[i - 1])

#                 nums.append(int(checkleft + checkright))
#             else:
#                 checkLeftDiag = checkDiagLeft(j, lines[i - 1])
#                 checkRightDiag = checkDiagRight(j, lines[i - 1])
#                 if checkLeftDiag:
#                     nums.append(int(checkLeftDiag))
#                 if checkRightDiag:
#                     nums.append(int(checkRightDiag))

#             if checkDown:
#                 checkleft = checkLeft(j, lines[i + 1])
#                 checkright = checkRight(j, lines[i + 1])

#                 nums.append(int(checkleft + checkright))
#             else:
#                 checkLeftDiag = checkDiagLeft(j, lines[i + 1])
#                 checkRightDiag = checkDiagRight(j, lines[i + 1])
#                 if checkLeftDiag:
#                     nums.append(int(checkLeftDiag))
#                 if checkRightDiag:
#                     nums.append(int(checkRightDiag))

#             checkleft = checkLeft(j, line)
#             checkright = checkRight(j, line)
#             if checkleft:
#                 nums.append(int(checkleft))
#             if checkright:
#                 nums.append(int(checkright))


# P2
nums = []
for i, line in enumerate(lines):
    tmp = []
    for j, char in enumerate(line):
        if checkIfSymbol(char):
            checkUp = checkUpDown(j, lines[i - 1], i)
            checkDown = checkUpDown(j, lines[i + 1], i)
            if checkUp:
                checkleft = checkLeft(j, lines[i - 1])
                checkright = checkRight(j, lines[i - 1])

                tmp.append(int(checkleft + checkright))
            else:
                checkLeftDiag = checkDiagLeft(j, lines[i - 1])
                checkRightDiag = checkDiagRight(j, lines[i - 1])
                if checkLeftDiag:
                    tmp.append(int(checkLeftDiag))
                if checkRightDiag:
                    tmp.append(int(checkRightDiag))
            if checkDown:
                checkleft = checkLeft(j, lines[i + 1])
                checkright = checkRight(j, lines[i + 1])

                tmp.append(int(checkleft + checkright))
            else:
                checkLeftDiag = checkDiagLeft(j, lines[i + 1])
                checkRightDiag = checkDiagRight(j, lines[i + 1])
                if checkLeftDiag:
                    tmp.append(int(checkLeftDiag))
                if checkRightDiag:
                    tmp.append(int(checkRightDiag))
            checkleft = checkLeft(j, line)
            checkright = checkRight(j, line)
            if checkleft:
                tmp.append(int(checkleft))
            if checkright:
                tmp.append(int(checkright))
            print(tmp)
            if len(tmp) == 2:
                print(f"LEN IS 2 {tmp}")
                nums.append(tmp[0] * tmp[1])
            tmp = []

print(sum(nums))
