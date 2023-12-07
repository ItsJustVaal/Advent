with open("../inputs/p4input.txt") as file:
    lines = [line.strip() for line in file.readlines()]


# P1
sum = 0
for line in lines:
    tmpNum = 0
    tmp = line.split(": ")
    winners = tmp[1].split(" | ")[0].split(" ")
    myNums = tmp[1].split(" | ")[1].split(" ")

    for num in myNums:
        if num and num in winners:
            if tmpNum == 0:
                tmpNum += 1
            else:
                tmpNum += tmpNum
    sum += tmpNum
print(sum)

# P2
cardMap = {x: 1 for x in range(1, len(lines) + 1)}

sum = 0
for i, line in enumerate(lines):
    tmp = line.split(": ")
    winners = tmp[1].split(" | ")[0].split(" ")
    myNums = tmp[1].split(" | ")[1].split(" ")

    numWinners = 0
    for num in myNums:
        if num and num in winners:
            numWinners += 1
    for x in range(numWinners):
        cardMap[i + x + 2] += 1

    copies = cardMap[i + 1] - 1
    while copies != 0:
        for x in range(numWinners):
            cardMap[i + x + 2] += 1
        copies -= 1


for k, v in cardMap.items():
    sum += v

print(sum)
