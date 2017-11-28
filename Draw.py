#!/usr/bin/env python3

import sys
import random


kids = [
    {"name": "kid1", "siblings": ["kid2"], "draw": "", "drawn": False},
    {"name": "kid2", "siblings": ["kid1"], "draw": "", "drawn": False},
    {"name": "kid3", "siblings": [], "draw": "", "drawn": False},
    {"name": "kid4", "siblings": [], "draw": "", "drawn": False},
    {"name": "kid5", "siblings": [], "draw": "", "drawn": False},
    {"name": "kid6", "siblings": [], "draw": "", "drawn": False}
    ]


def MatchKids(kids):
    if len(kids) % 2 != 0:
        print("Uneven number of kids. Please enter an even number of kids")
        sys.exit(1)
    else:
        for kid in kids:
            while(kid["draw"] == ""):
                draw = random.choice(kids)
                if draw["name"] != kid["name"] and draw["name"] not in kid["siblings"] and draw["drawn"] is False:
                    kid["draw"] = draw["name"]
                    draw["drawn"] = True
                elif draw["name"] == kids[-1]["name"] and kids[-1]["drawn"] == False:
                    for kid in kids:
                        kid["draw"] = ""
                        kid["drawn"] = False
                    MatchKids(kids)


if __name__ == '__main__':
        MatchKids(kids)
#        for kid in kids:
#            print(kid["name"], "has drawn", kid["draw"])
