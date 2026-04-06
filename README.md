# doublependulum
The equation of motion and the phase diagram of a pendulum

I used this video to derive the [equation of motion](https://www.youtube.com/watch?v=tc2ah-KnDXw)

Positions:

``` math
\begin{positions}
x_1 = l_1\cdot sin(\theta_1) \\
y_1 = -l_1\cdot cos(\theta_1) \\
x_2 =  l_1\cdot sin(\theta_1) +  l_2\cdot sin(\theta_2) \\
y_2 = -l_1\cdot cos(\theta_1) -l_2\cdot cos(\theta_2)
\end{positions}
```

Velocities:

```math
\dot{x}_1 = \dot{\theta_1} l_1 cos(\theta_1)

\dot{y}_1 = \dot{\theta_1} l_1 sin(\theta_1)

\dot{x}_2 = \dot{\theta_1} l_1 cos(\theta_1) + \dot{\theta_2} l_2 cos(\theta_2)

\dot{y}_2 = \dot{\theta_1} l_1 sin(\theta_1) + \dot{\theta_2} l_2 sin(\theta_2)
```

Equation of motion:

![Equation of Motion](eom.png)

The phase diagram is pretty boring, it is just an ellipse:

![Phase diagram](pd.png)