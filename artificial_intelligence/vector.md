## lines
* $w_1x_1+w_2x_2+w_0=0$ is a line. where $w^T$ vector is $[w_1 w_2]$ associates with direction $x_1$ and $x_2$ and $w_0$ is a scalar quantity which will determine the intercept at $x_2$ axis. Now, if we multiply $w^T$ vector and $w_0$ with a scalar quantity the line will be same, you can try it in graph. i.e. $w_1x_1+w_2x_2+w_0=0$ is same as $24w_1x_1+24w_2x_2+24w_0=0$
* $w_1x_1+w_2x_2+w_0=1$ is a line on right side of $w_1x_1+w_2x_2+w_0=0$ and parallel to it. The R.H.S. i.e. for this example is $1$ defines how far on the right side. again the L.H.S. $w^T$ and $w_0$ also defines how far it will be from $w_1x_1+w_2x_2+w_0=0$ line.
* $w_1x_1+w_2x_2+w_0=-1$ is a line on left side of $w_1x_1+w_2x_2+w_0=0$ and parallel to it.
  
## Minimum distance between two parallel lines
* to calculate the minimum distance between two parallel lines $w_1x_1+w_2x_2+w_0=0$ and $w_1x_1+w_2x_2+w_0=1$ we choose any one point $(x_1^*, x_2^*)$ on line $w_1x_1+w_2x_2+w_0=1$ and then find perpendicular distance from that point to the line $w_1x_1+w_2x_2+w_0=0$.

## Minimum distance between a point on a line and a line
* the minimum distance between a line/hyperplane $w_1x_1+w_2x_2+w_0=0$ and a vector $(x_1^*, x_2^*)$ on line $w_1x_1+w_2x_2+w_0=1$ is defined as $\gamma$ [gamma]
$\gamma=\frac{|w_1x_1^*+w_2x_2^*+w_0|}{\sqrt{w_1^2+w_2^2}}=\frac{1}{\sqrt{w_1^2+w_2^2}}=\frac{1}{||w||}$
why 1? because the point $(x_1^*, x_2^*)$ is on line line $w_1x_1+w_2x_2+w_0=1$ and look at the R.H.S. is 1.

## Minimum distance between a point at any position and a line  
![](line2.drawio.svg)  
the minimum distance between a line/hyperplane $w_1x_1+w_2x_2+w_0=0$ and a vector $(x_1^*, x_2^*)$ is defined as $\gamma$ [gamma]
$\gamma=\frac{|w_1x_1^*+w_2x_2^*+w_0|}{\sqrt{w_1^2+w_2^2}}$

## Position of a point w.r.t. a line
the line is $w_1x_1+w_2x_2+w_0=0$ and let's say the point is $(x_1^*, x_2^*)$  
$d=w_1x_1^*+w_2x_2^*+w_0$

$d$ becomes $0$ if $(x_1^*, x_2^*)$ is a point on the hyperplane.  
$d$ becomes $+ve$ if $(x_1^*, x_2^*)$ is a point on right side of the hyperplane.  
$d$ becomes $-ve$ if $(x_1^*, x_2^*)$ is a point on left side of the hyperplane.

## plot 2 equation and then set value of 1st equation till it touches 2nd equation.