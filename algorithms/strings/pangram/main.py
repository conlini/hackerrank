alphas = {}

s = raw_input()
pangram=False
for i in s:

    if i != ' ':
        if i.lower() not in alphas:
            alphas[i.lower()] = 1
            if len(alphas) == 26:
                pangram=True
                break
print(alphas, len(alphas))
if pangram:
    print("pangram")
else:
    print("not pangram")
