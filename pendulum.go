package main

import "math"

func x(x, t, deltat float64) float64 {
	omega1 := math.Sqrt(g / l1)
	omega2 := math.Sqrt(g / l2)
	// it is not possible to solve the equation of motion analytically, we have to solve it numerically.
	// for x we get: x_t+1=x_t + dx / dt * Delta t
	xdot := -omega1*math.Sin(omega1*t) - omega2*math.Sin(omega2*t)
	return x + xdot + deltat
}
