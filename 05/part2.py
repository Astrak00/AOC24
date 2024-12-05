from collections import defaultdict

rules = defaultdict(list)
acc = 0

with open("input.txt", 'r') as f:
    for line in f:
        if '|' in line:
            n1, n2 = map(int, line.split('|'))
            rules[n1].append(n2)
        if "," in line:
            nums = list(map(int, line.split(',')))
            # Creates a copy, not just a pointer
            original_nums = nums[:]
            done = True
            while done:
                done = False
                for i in range(1, len(nums)):
                    if not nums[i] in rules[nums[i-1]]:
                        nums[i], nums[i-1] = nums[i-1], nums[i]
                        done = True
            if not original_nums == nums:
                acc += nums[len(nums)//2]

print(acc)
