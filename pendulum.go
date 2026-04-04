package main

import "math"

func x(x, t, deltat float64) float64 {
	omega := math.Sqrt(g / l)
	// it is not possible to solve the equation of motion analytically, we have to solve it numerically.
	// for x we get: x_t+1=x_t + dx / dt * Delta t
	xdot := -omega * math.Sin(omega*t)
	return x + xdot + deltat
}

func xdot(t float64) float64 {
	omega := math.Sqrt(g / l)
	return -(x0 * omega * math.Sin(omega*t))
}
