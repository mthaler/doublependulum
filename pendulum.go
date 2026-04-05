package main

import "math"

func x(x, t, deltat float64) float64 {
	omega1 := math.Sqrt(g / l1)
	omega2 := math.Sqrt(g / l2)
	// the equation of motion is best derived using the Lagrange formalism. The Euler-Lagrane equation for the equation of motion
	// is (d/dt)dL/d qdot_i - dL/dq_i = 0, where q_i are the generalized coordinates (in our case the angles)

	// The equation of motion is a coupled set of differential equations which I think cannot be solved analytically, thus we use
	// the Euler method to solve it numerically.
	xdot := -omega1*math.Sin(omega1*t) - omega2*math.Sin(omega2*t)
	return x + xdot + deltat
}
