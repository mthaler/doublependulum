package main

import "math"

func x(x, t, deltat float64) float64 {
	omega1 := math.Sqrt(g / l1)
	omega2 := math.Sqrt(g / l2)
	// The equation of motion is from this video: https://www.youtube.com/watch?v=aTMJX1ZgMB0&t=433s
	//
	// the equation of motion is best derived using the Lagrange formalism. The Euler-Lagrane equation for the equation of motion
	// is (d/dt)dL/d qdot_i - dL/dq_i = 0, where q_i are the generalized coordinates (in our case the angles)
	//
	// we get:
	//
	// (m1 + m2)l1^2 th1dotdot + m2 l1 l2 th2dotdot cos(th1-th2) + m2 l1 l2 th2dotdot sin(th1 - th2) + (m1 + m2) g l1 sin(th1) = 0
	// and
	// l2th2dotdot + l1th1dotdot cos(th1-th2) - l1 th1dot^2 sin(th1-th2) + g sin(th2) = 0
	//
	// The equation of motion is a coupled set of differential equations which I think cannot be solved analytically, thus we use
	// the Euler method to solve it numerically.
	// The Euler method only works for first order differential equations. The equation of motion is a a coupled set of two second
	// order differential equations, thus we first have to transform it to first order differential equations
	//
	// we make the following substituions:
	// y1 = theta1
	// y2 = theta1dot
	// y3 = theta2
	// y4 = theta2dot
	//
	// then we get:
	//
	// th1 = th1dot
	// th1 dot = [(m1+m2)l1-m2l1cos^2(th1-th2)]^-1 [-m2l1th1dot^2sin(th1-th2)cos(th1-th2)]+m2gsin(th2)cos(th1-th2)-m2l2th2dot^2sin(th1-th2-(m1+m2)gsin(th1)
	// th2 = th2dot
	// th2dot = [l2 - m2l2/(m1 + m2)cos^2(th1-th2)]^-1[l1 th1dot^2sin(th1-th2)-gsin(th2)+th2dot^2sin(th1-th2)cos(th1-th2)+gsin(th1)cos(th1-th2)]
	xdot := -omega1*math.Sin(omega1*t) - omega2*math.Sin(omega2*t)
	return x + xdot + deltat
}
