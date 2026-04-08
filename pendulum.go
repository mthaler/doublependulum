package main

// The equation of motion is from this video: https://www.youtube.com/watch?v=aTMJX1ZgMB0&t=433s
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

func x(prev data, t, deltat float64) data {
	//y1 := prev.y1 + deltat * y2
	//y2 := ((m1+m2)l1-m2l1cos^2(y1-y3))^-1 [-m2l1y2^2sin(y1-y3)cos(y1-y3)]+m2gsin(y3)cos(y1-y3)-m2ly4^2sin(y1-y3)-(m1+m2)gsin(y1)
	//y3 := prev.y3 + deltat * y4
	d := data{}
	return d
}
