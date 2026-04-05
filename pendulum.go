package main

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
// y1dot = y2
// y2dot = [(m1+m2)l1-m2l1cos^2(y1-y3)]^-1 [-m2l1y2^2sin(y1-y3)cos(y1-y3)]+m2gsin(y3)cos(y1-y3)-m2ly4^2sin(y1-y3)-(m1+m2)gsin(y1)
// y3dot = y4
// y4dot = [l2 - m2l2/(m1 + m2)cos^2(y1-y3)]^-1[l1 y2^2sin(y1-y3)-gsin(y3)+m2l2/(m1 +2)y4^2sin(y1-y3)cos(y1-y3)+gsin(y1)cos(y1-y3)]

func x(y1prev, y2prev, y3prev, y4prev, t, deltat float64) (float64, float64, float64, float64) {
	y1 := y1prev + deltat * y2
	y2 := ((m1+m2)l1-m2l1cos^2(y1-y3))^-1 [-m2l1y2^2sin(y1-y3)cos(y1-y3)]+m2gsin(y3)cos(y1-y3)-m2ly4^2sin(y1-y3)-(m1+m2)gsin(y1)
	y3 := 
	return 0.0, 0.0, 0.0, 0.0
}
