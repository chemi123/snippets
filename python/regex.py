import re

pattern = 'hell'

res = re.match(pattern, r'hellow world')
if res:
    print('match')
else:
    print('not match')
