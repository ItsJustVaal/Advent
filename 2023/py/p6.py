with open("../inputs/input.txt") as file:
    lines = [line.strip() for line in file.readlines()]

# P1
times = [int(num) for num in lines[0].split(" ")[1:] if num != '']
distances = [int(num) for num in lines[1].split(" ")[1:] if num != '']

final = []
for i, time in enumerate(times):
    sum = 0
    for x in range(0, time + 1):
        if (time - x) * x > distances[i]:
            sum += 1

    final.append(sum)

total = 1
for num in final:
    total = total * num

print(total)

# P2
times, distances = '', ''

for num in lines[0].split(" ")[1:]:
    if num != '':
        times = times + num

for num in lines[1].split(" ")[1:]:
    if num != '':
        distances = distances + num
times = int(times)
distances = int(distances)

sum = 0
for x in range(0, times + 1):
    if (times - x) * x > distances:
        sum += 1

print(sum)
