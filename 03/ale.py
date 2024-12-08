
with open("input.txt", "r") as file:
    a = file.readlines()

input = ' '.join(a)

a += "don't()"
new_string = ""
do_position = 0
while do_position != -1:
    dont_pos = input.find("don't()", do_position, len(input))
    new_string += input[do_position:dont_pos]
    do_position = input.find("do()", dont_pos, len(input))

with open("sal.sal", "w") as file:
    file.write(new_string)