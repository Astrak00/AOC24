orderings = []
updates = []

with open("input.txt", 'r') as f:
    before_double_split = True
    for line in f: 
        if before_double_split:
            if line == '\n':
                before_double_split = False
            else:
                line = list(map(int, line.strip().split('|')))
                orderings.append(line)
        else:
            line = list(map(int, line.strip().split(',')))
            updates.append(line)

def check_ordering(rules: list[list[int]], pages: list[int]) -> tuple[bool, int]:
    for pageIdx in range(len(pages)):
        before, elem, after = pages[:pageIdx], pages[pageIdx], pages[pageIdx+1:]
        for rule in rules:
            if rule[1] == elem:
                if rule[0] in after:
                    return (False, -1)
            if rule[0] == elem:
                if rule[1] in before:
                    return (False, -1)

    return (True, pages[len(pages)//2])
    
    
        
        
acc = 0
for update in updates:
    temp_orderings = []

    # Finding all the orderings that apply to that set of updates
    for order in orderings:
        if order[0] in update or order[1] in update:
            temp_orderings.append(order)

    check, val = check_ordering(temp_orderings, update)
    if check:
        acc += val
    
print(acc)