## Quadratic Optimization of a function without constraints
$f(x)=(3x-4)^2$ minimize $f(x)$ we can apply gradient descent here.
## Quadratic Optimization of a function with constraints
$f(w)=\frac{||w||^2}{2}$ minimize $f(w)$ but the constraints is -   
$y_i(w^Tx_i+w_0)>=1$ where $i$ is row of a dataset and $y_i$ is label (+1 or -1) and $x_i$ is a column vector of input columns for $i^{th}$ row.
we can apply 
* Lagrangian multiplier for n inequality constraints
  >only works with equality constraints. it converts constraints problem to unconstraints function.so, we have n inequality constraints here which are that every $i^{th}$ row follows the constraint above.

    $L(w,\alpha,w_0)=\frac{||w||^2}{2}-(\sum_{i=1}^{n}{\alpha_i (y_i(w^Tx_i+w_0)-1) })$  
    minimize $L(w,\alpha,w_0)$ such that $\alpha_i>=0$  
    Now:  
    $\frac{\delta L}{\delta w}=w-\sum_{i=1}^n{\alpha_i y_ix_i}=0$  
    so, $w=\sum_{i=1}^n{\alpha_i y_ix_i}$   ---(1)

    $\frac{\delta L}{\delta w_0}=-\sum_{i=1}^n{\alpha_i y_i}=0$ ---(2)

    Plugin (1) and (2) in $L$  
    $L=\sum{\alpha_i}-\frac{1}{2}\sum_{i=1}\sum_{j=1}\alpha_i \alpha_j y_iy_jx_i\dot{x_j}$  
    or we can pass $x_i,x_j$ through kernel function and rewrite the formula as   
    $L=\sum{\alpha_i}-\frac{1}{2}\sum_{i=1}\sum_{j=1}\alpha_i \alpha_j y_iy_jk(x_i,x_j)$  
    and we need to maximize L

## Non Quadratic


Quadratic Programming Solver
Dual Formulation

## Optimization Problem
optimization problem has 

* cost function
* equality constraints
* inequality constraints

if cost, equality, inequality functions are Linear then Linear Program.

Linear Functions:
    $w^Tx+w_0=0$

## Convex Optimization 
1987: algorithm for convex optimization problem 
convex concept w.r.t.
* sets
  * a set is convex if for any pair of points falls entierly in the set
* functions
  * a function f is convex if its epigraph i.e. the region above the graph of the function is a convex set
* Optimization
  * an optimization problem is said to be convex if the objective function (cost function) is convex, the function for inequality contraint is convex and then equality constrait function to be Linear.

## Principle of duality
seeing the same problem from two different perspective, primal one and dual one. 


## function
$z=f(x,y)=x^2y$, $z$ is a function of $x$ and $y$ variable.
difference between 
$x^2y$ and $x^2y=2$

## REFERENCES
  1. https://www.svm-tutorial.com/2016/09/duality-lagrange-multipliers/