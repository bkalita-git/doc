## suport vector machine
* hard svm (hyperplane which can separate two class without errors)
* soft svm (hyperplane with some acceptable errors)




Terms  
* hyperplane/straight-line
* margin: margin line on left and right parallel to hyperplane.
* marginal distance : distance between the margin lines.
* support vectors : points passing through the marginal plane/line
* svm kernel: 2d points to high dimension for Non-linear separable.
* maximal margin classifier (hard margin)
* support vector classifier (soft margin)
* support vector machine (kernel functions)
  
## cross validation using kernel function
let's say we select "polynomial" as kernel function. and we will find different hyperplane for different degrees of polynomial. for example, if our training dataset has only one column, that is 1D, then we will first find hyperplane for it by setting d=1. after set we will make another column by taking the first_col and square it, here d=2, like wise we will increase d value and after that we will select the best hyperplane.
example of diff kernel:
* polynomial kernel $(a\times b+r)^d$
  * we will add many polynomial kernel for r=0 and d to infinity.
* radial kernel $e^{-\gamma(a-b)^2}$

## kernel trick
we actually transformed data from 1 dimension to higher which is extra computation. with kernel trick we can do cross validation without actually transforming data to higher dimension.


## relation between 2 vectors


Linear separable
Non-linear separable


|id|x1|x2|x3|...|xn|label|
|--|--|--|--|---|--|-----|
|0 |1 |3 |4 |...|58|yes  |
|1 |3 |2 |6 |...|24|no   |
|. |. |. |. |...|..|...  |

equation of hyperplane:  
$w^Tx+w_0=0$  

```
where x is a matrix [x1, x2, x3, .. xn] of size (nx1)
x1,x2.. are attribute/column of dataset. 
consider it as we have plotted x1,x2..xn in a n dimensional graph with their label for every row. Now, we want to draw a hyperplane so it can separate yes and no label points. right now we are doing classification on 2 class. but in reality there can be multiple class, for example, [yes, no, notsure] for this 3 hyperplane will be made to separate (yes,no), (no,notsure) and (yes, notsure)

next, we will add an extra column/attribute to the dataset and this value can be 1 for every row. why? because to solve non linear separable points we added an extra dimension.

Anyway, after plotting those points, for (yes,no) we will draw a initial hyperplane. but how? we will use the above formula of hyperplane. but what will be the value of w,x and w0? w will be a initialized vector and w0 will be a initialized value. for each row x will vary.


let's consider (yes,no) as (1,-1) and consider label as y

our problem is to find support vectors for yes and no class.
``` 
now for every row $i$:  
if $y_i(w^Tx_i+w_0)-1==0$ then $x_i$ is support vector and save $w$ and $b$  
else if $y_i(w^Tx_i+w_0)-1>0$ then save the parameter $w$ and $w_0$   
else if $y_i(w^Tx_i+w_0)-1<0$ then update parameters $w$ and $w_0$

>what in the above steps are doing that first we make a hyperplane and we did that by initializing $w$ and $w_0$ with random numbers.***(to make a hyperplane we only need $w$ and $w_0$ and put in $w^Tx+w_0=0$ where $x$ matrix values will be $-\infin,.-1,0,1,..\infin$ for every dimension in $x$)*** now the question, is this a correct one? if not then we update $w$ and $w_0$ and how we know that the hyperplane is correct or incorrect? simple, if the hyperplane is the correct one then for a row from dataset for a specific label putting in a formula will be greater or equal to 0.

equation of positive marginal plane:  
    $w^Tx_2+b=1$  
equation of negative marginal plane:  
    $w^Tx_1+b=-1$  
distance between two marginal plane:  
    $w^T(x_2-x_1)=2$    
    $\frac{w^T}{||w||}(x_2-x_1) = \frac{2}{||w||}$  

maximize $\frac{2}{||w||}$ same as minimize $\frac{||w||^2}{2}$

## optimization of a function with constraints
```
we can optimize a function using gradient descent but it will not work if there are any rules/constraints applied to that function.
so,
step1: convert the constraints optimization function to non constraint optimization function using lagrangian multiplier. it lets you combine two equation, (the function and the constraint) into one. It can handle inequality constraint.
```


## prediction after getting $w$ and $w_0$
here $y_i$ is a class from "label" column  
if $w^Tx+w_0>=1$ then $y_i=1$
  else if $w^Tx+w_0<=-1$ then $y_i=-1$  


## REFERENCES
1. https://ankitnitjsr13.medium.com/math-behind-support-vector-machine-svm-5e7376d0ee4d
2. Statistical learning theory by vladmir 
3. http://web.mit.edu/6.034/wwwbob/svm-notes-long-08.pdf