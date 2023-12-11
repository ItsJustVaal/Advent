with open("../inputs/p5input.txt") as file:
    lines = [line.strip() for line in file.readlines()]
    seeds = [int(seed) for seed in lines[0].split(": ")[1].strip().split(" ")]


maps = {}
currMap = ""
lowest = 0

for line in lines[2:]:
    if 'map' in line:
        maps[line] = []
        currMap = line
    elif line:
        maps[currMap].extend(line.split("\n"))

# Part 1
for seed in seeds:
    next, tmp = seed, 0
    for map in maps.items():
        print(f"CURRENT MAP: {map[0]}")
        for values in map[1]:
            print(f"START OF VALUES {values}")
            dest = int(values.split('\n')[0].split(" ")[0])
            source = int(values.split('\n')[0].split(" ")[1])
            rng = int(values.split('\n')[0].split(" ")[2])
            if next in range(source, source + rng + 1):
                tmp = next - source
                if dest + tmp <= dest + rng:
                    next = dest + tmp
                    break
                else:
                    next = seed
            else:
                print(f"NEXT IS {next}")
                print(f"END OF VALUES {values}")
                continue
            print(f"NEXT IS {next}")
            print(f"END OF VALUES {values}")
    if lowest == 0 or next <= lowest:
        lowest = next
    print(f"END OF SEED {seed}, LOWEST IS {lowest}")

print(lowest)

# Part 2
