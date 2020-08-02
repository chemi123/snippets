# -*- coding: utf-8 -*-

import sys
import random

if __name__ == "__main__":
    #if len(sys.argv) != 2:
    #    print("Usage: python3 create_dataset.py ${dataset_num}")
    #    exit(1)
    
    #if not sys.argv[1].isdecimal():
    #    print("input number for dataset_num")
    #    exit(1)
    
    n = random.randint(1, 100)
    q = random.randint(1, 1000)
    print(n, end=" ")
    print(q)
    for i in range(q):
        print(random.randint(0, 1), end=" ")
        print(random.randint(0, n), end=" ")
        print(random.randint(0, n))
