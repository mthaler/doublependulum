package main

import "math"

func angles(prev data, t, deltat float64) data {
	// to simplify things we make the approximation that the previous value of a varialbe is approximately the current value of the variable
	y1 := prev.y1 + deltat*prev.y2
	y2 := prev.y2 + deltat*(-m2*l1*prev.y2*prev.y2*math.Sin(y1-prev.y3)*math.Cos(y1-prev.y3)+m2*g*math.Sin(prev.y3)*math.Cos(y1-prev.y3)-m2*l2*prev.y4*prev.y4*math.Sin(y1-prev.y3)-(m1+m2)*g*math.Sin(y1))/((m1+m2)*l1-m2*l1*math.Cos(y1-prev.y3)*math.Cos(y1-prev.y3))
	y3 := prev.y3 + deltat*prev.y4
	y4 := prev.y4 + deltat*(l1*y2*y2*math.Sin(y1-y3)-g*math.Sin(y3)+m2*l2/(m1+m2)*prev.y4*prev.y4*math.Sin(y1-y3)*math.Cos(y1-y3)+g*math.Sin(y1)*math.Cos(y1-y3))/(l2-m2*l2/(m1+m2)*math.Cos(y1-y3)*math.Cos(y1-y3))
	d := data{y1: y1, y2: y2, y3: y3, y4: y4}
	return d
}
