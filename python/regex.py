import re

# よく使うものをとりあえず2パターンだけ

pattern = 'hell'
content = r'hello world'

# simple
res1 = re.match(pattern, content)
if res1:
    print('match')
else:
    print('not match')

# compile (faster)
comp = re.compile(pattern)
res2 = comp.match(content)
print(res2)
