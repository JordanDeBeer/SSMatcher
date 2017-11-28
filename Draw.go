package main

import (
	"fmt"
	"math"
	"math/rand"
    "os"
	"time"
)

type Kid struct {
	Name     string
	Siblings []string
	Draw     string
	Drawn    bool
}

func SetDraw(k *Kid, name string) {
	k.Draw = name
}

func SetDrawn(k *Kid, boolean bool) {
	k.Drawn = boolean
}

var Kids []Kid

func init() {
    Kids = []Kid{
        Kid{Name: "kid1", Siblings: []string{"kid2"}, Draw: "", Drawn: false},
        Kid{Name: "kid2", Siblings: []string{"kid1"}, Draw: "", Drawn: false},
        Kid{Name: "kid3", Siblings: []string{""}, Draw: "", Drawn: false},
        Kid{Name: "kid4", Siblings: []string{""}, Draw: "", Drawn: false},
        Kid{Name: "kid5", Siblings: []string{""}, Draw: "", Drawn: false},
        Kid{Name: "kid6", Siblings: []string{""}, Draw: "", Drawn: false},
    }
    if math.Mod(float64(len(Kids)), 2) != 0 {
        fmt.Println("Not even!")
        os.Exit(1)
    }
}

func MatchKids() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < len(Kids); i++ {
		kid := &Kids[i]
		for kid.Draw == "" {
			draw_number := r1.Intn(len(Kids))
			draw := Kids[draw_number]
			// Comparison to make sure:
			// the kid doesn't draw their own name,
			// the name hasn't already been drawn,
			// and the name they draw isn't a sibling
			if draw.Name != kid.Name && draw.Drawn == false {
				for s := range kid.Siblings {
					if draw.Name != kid.Siblings[s] {
						SetDraw(&Kids[i], draw.Name)
						SetDrawn(&Kids[draw_number], true)
					}
				}
			} else if draw.Name == Kids[len(Kids)-1].Name && Kids[len(Kids)-1].Drawn == false {
				for j := 0; j < len(Kids); j++ {
					SetDraw(&Kids[j], "")
					SetDrawn(&Kids[j], false)
				}
			MatchKids()
			}
		}
	}
}

func main() {
	MatchKids()
	for i := 0; i < len(Kids); i++ {
		fmt.Println(Kids[i].Name, "has drawn", Kids[i].Draw)
	}
}
