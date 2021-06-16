package main

import (
	"fmt"
	"github.com/kellydunn/golang-geo"
	"math"
)

func main() {
		//106.47142,
		//-6.219
	// Make a few points
	//p := geo.NewPoint(107.6176, -6.97384)
	//p2 := geo.NewPoint(97.6054916, 1.3146232)
	//	106.5083517,
	//	-6.2582864
	p := geo.NewPoint(6.13963, 106.52921)
	p2 := geo.NewPoint(-6.219, 106.47142)

	// find the great circle distance between them
	dists := p.GreatCircleDistance(p2)
	dist := math.Round(dists/0.01) * 0.01
	if dist > 10 {
		fmt.Println("jarak lebih dari 10km")
	}else{
		fmt.Println("jarak kurang dari 10km")
	}
	fmt.Println("great circle distance: ", dists)

}

