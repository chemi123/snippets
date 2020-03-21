# -*- coding: utf-8 -*-

import sys

argvs = sys.argv
argc = len(argvs)

if argc != 2:
    print('need an arg')
    exit(1)

print(argc)
print(argvs)
