line = open("examples/day5.txt").read().splitlines()[0]

oldline = None
while oldline != line:
    oldline = line
    for i in range(0,26):
        line = line.replace(chr(ord("a") + i) + chr(ord("A") + i),"")
        line = line.replace(chr(ord("A") + i) + chr(ord("a") + i),"")

print("Part1:")
print(line)
