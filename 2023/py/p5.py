from numpy import Infinity


with open("../inputs/p5input.txt") as file:
    lines = [line.strip() for line in file.readlines()]
    seedsInit = list(map(int, lines[0].split(" ")[1:]))

seeds = [(seedsInit[x], seedsInit[x + 1]) for x in range(0, len(seedsInit), 2)]
print(seeds)

maps = []
currMap = ""
lowest = 0

i = 2
while i < len(lines):
    start, _, finish = lines[i].split(" ")[0].split("-")
    maps.append([])

    i += 1
    while i < len(lines) and not lines[i] == "":
        dStart, sStart, rng = map(int, lines[i].split())
        maps[-1].append((dStart, sStart, rng))
        i += 1

    maps[-1].sort(key=lambda x: x[1])
    i += 1

# Part 1
for seed in seeds:
    next, tmp = seed, 0
    for map in maps:
        print(f"CURRENT MAP: {map}")
        for dest, source, rng in map:
            if next in range(source, source + rng + 1):
                tmp = next - source
                if dest + tmp <= dest + rng:
                    next = dest + tmp
                    break
                else:
                    next = seed
            else:
                continue

    if lowest == 0 or next <= lowest:
        lowest = next


print(lowest)

# # Part 2

# This was way more difficult and im still not sure I totally understand what is going on


def newMap(low, high, map):
    newMap = []
    for dest, source, rng in map:

        end = source + rng - 1
        shift = dest - source
        if not (end < low or source > high):
            newMap.append((max(source, low), min(end, high), shift))

    for x, item in enumerate(newMap):
        left, right, dist = item

        yield (left + dist, right + dist)

        if x < len(newMap) - 1 and newMap[x + 1][0] > right + 1:
            yield (right + 1, newMap[x + 1][0] - 1)

    if len(newMap) == 0:
        yield (low, high)
        return

    if newMap[0][0] != low:
        yield (low, newMap[0][0] - 1)
    if newMap[-1][1] != high:
        yield (newMap[-1][1] + 1, high)


lowest = Infinity

for start, range in seeds:
    curr = [(start, start + range - 1)]
    new = []

    for map in maps:
        for low, high in curr:
            for newM in newMap(low, high, map):
                new.append(newM)

        curr, new = new, []

    for low, high in curr:
        lowest = min(lowest, low)

print(lowest)
