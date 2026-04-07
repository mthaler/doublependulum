# doublependulum
The equation of motion and the phase diagram of a pendulum

I used this video to derive the [equation of motion](https://www.youtube.com/watch?v=tc2ah-KnDXw)

Positions:

``` math
\begin{aligned}
& x_1 = l_1\cdot sin(\theta_1)\\
& y_1 = -l_1\cdot cos(\theta_1)\\
& x_2 =  l_1\cdot sin(\theta_1) +  l_2\cdot sin(\theta_2)\\
& y_2 = -l_1\cdot cos(\theta_1) -l_2\cdot cos(\theta_2)
\end{aligned}
```

Velocities:

```math
\begin{aligned}
& \dot{x}_1 = \dot{\theta}_1 l_1 cos(\theta_1)\\
& \dot{y}_1 = \dot{\theta}_1 l_1 sin(\theta_1)\\
& \dot{x}_2 = \dot{\theta}_1 l_1 cos(\theta_1) + \dot{\theta}_2 l_2 cos(\theta_2)\\
& \dot{y}_2 = \dot{\theta}_1 l_1 sin(\theta_1) + \dot{\theta}_2 l_2 sin(\theta_2)
\end{aligned}
```

Potential energy:

```math
\begin{aligned}
V &= m_1gy_1+m_2gy_2 \\
  &= -(m_1+m_2) g l_1 cos(\theta_1)-m_2 g l_2 cos(\theta_2)
\end{aligned}
```

Kinetic energy:

```math
T = \frac{1}{2}m_1 l_1^2 \dot{\theta}_1 + \frac{1}{2}m_2(l_1^2\dot{\theta}_1^2 + l_2^2\dot{\theta}_2^2 + 2 l_1 l_2 \theta_1 \theta_2 cos(\theta_1 - \theta_2))
```

The equation of motion is best derived using the Lagrange formalism:

```math
L = T - U
```

The Euler-Lagrange equation for the equation of motion is 

```math
\frac{d}{dt}\frac{dL}{d \dot{q}_i} - \frac{dL}{d q_i}  = 0 
```

where $q_i$ are the generalized coordinates (in our case the angles). We get:

```math
\begin{aligned}
(m_1 + m_2)l_1^2 \ddot{\theta}_1 + m_2 l_1 l_2 \ddot{\theta}_2 cos(\theta_1-\theta_2) + m_2 l_1 l_2 \ddot{\theta}_2 sin(\theta_1 - \theta_2) + (m_1 + m_22) g l_2 sin(\theta_1) &= 0 \\
\ddot{\theta_2} + \ddot{\theta_1} cos(\theta_1-\theta_2) - l_1 \dot{\theta}_1^2 sin(\theta_1-\theta_2) + g sin(\theta_2) &= 0
\end{aligned}
```

The equation of motion is a coupled set of differential equations which I think cannot be solved analytically, thus we use the Euler method to solve it numerically. The Euler method only works for first order differential equations. The equation of motion is a a coupled set of two second order differential equations, thus we first have to transform it to first order differential equations. Using the following substitutions

```math
\begin{aligned}
y_1 &= \theta_1\\
y_2 &= \dot{\theta}_1\\
y_3 &= \theta_2\\
y_4 &= \dot{\theta}_2
\end{aligned}
```

the result [is](https://www.youtube.com/watch?v=aTMJX1ZgMB0&t=433s):

```math
\begin{aligned}
\dot{y}_1 &= y_2
\dot{y}_2 &= \frac{[-m2l1y2^2sin(y1-y3)cos(y1-y3)]+m2gsin(y3)cos(y1-y3)-m2ly4^2sin(y1-y3)-(m1+m2)gsin(y1)}{[(m_1+m_2)l_1-m_2 l_1 cos^2(y1-y3)]} 
\end{aligned}
```

Equation of motion:

![Equation of Motion](eom.png)

The phase diagram is pretty boring, it is just an ellipse:

![Phase diagram](pd.png)