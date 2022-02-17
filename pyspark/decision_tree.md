https://www.kdnuggets.com/2020/01/decision-tree-algorithm-explained.html

### Problem Statement:  
which Input Column will be selected at a Node from the Dataset and its conditions if conntinuous/random values.

### Core Decision Tree Solution:
step 1 : make a 'Root' node from the dataset D. That Node holds the information of a column name and conditions applied to that column iff column values are not categorical otherwise conditions will be different class name. and based on that make branches 'child' node from that node.  
step 2 : Each 'child' will be a subset of D.  
step 3 : Apply a Statistical Method to 'Root' and its 'child' nodes and store it in for that specific 'Root' Node. {'RootType1':0.34343}  
step 4 : Tweak the 'Root' Node conditions and Column and do till step 3. {'RootType1':0.34343, 'RootType2':0.435443,...}  
step 5 : select the best 'Root' node type as 'Root' and then change the 'Root' to its Left child and do RootLeftRight  
what statistical method? can be Information Gain using Entropy, Standard Deviation Reduction using Standard Deviation, Reduction in Variance using Variance, Gini Index, Gini Ratio, Chi-square, etc.
```
# the TargetColumn is selected after filtering conditions/class in column1 when applied to (TargetColumn, Column1)

Max_Information_Gain_col(
    Information_Gain_col1(
        Entropy(TargetColumn), 
        Entropy(TargetColumn,Column1)
    ),
    ...
)

Max_Standard_Deviation_Reduction_col(
    SDR_col1(
        Standard_Deviation(TargetColumn),
        Standard_Deviation(TargetColumn,column1),
    ),
    ...
)

```


Algorithms:
1. ID3
2. C4.5
3. CART
4. CHAID
5. MARS
   

Attributes/Columns selections for root node or at different node level.
* for categorical data in Target column is easy to apply Entropy formula.
* for continuous data in Column but not in TargetColumn calculate boundary value and calculate information gain on those by considering.
$dataset.filter(Column1<boundaryValue1Column1).select(TargetColumn, Column1)$   
$dataset.filter(Column1>boundaryValue1Column1).select(TargetColumn, Column1)$  
$dataset.filter(Column1<boundaryValue2Column1).select(TargetColumn, Column1)$  
$...$
* for continuous data in target column? since Entropy can be calculated easily with categorical data. Use Reduction in Variance

for each column/attribute:

* Information gain  
   >ID3
    
    $Entropy(TargetColumn) = \sum_{i=1}^{c}-p(c_i)log_2p(c_i)$  
    $p(c_1)=\frac{c_1}{c_0+c_1+c_2...}$   
    $c_1$=numRowsHaving $c_1$ value in TargetColumn  
    $c$ = num of class in a TargetColumn  
    TargetColumn = a Target column in dataset    
    $Entropy(TargetColumn,Column1) = \sum_{c_i\in{Column1}}^{}p(c_i)Entropy($dataset.filter(Column1$==c_i).select(TargetColumn))$

    $InformationGain(TargetColumn,Column1)=Entropy(TargetColumn) - Entropy(TargetColumn,Column1)$
* Gini index
    >CART (Classification and Regression Tree) uses the Gini index method to create split points.

    $Gini(Column1) = 1-\sum_{i=1}^{c}(p(c_i))^2$
* Gain Ratio
    >C4.5, an improvement of ID3, uses Gain ratio

    $GainRatio(TargetColiumn,Column1)=\frac{InformationGain(TargetColumn, Column1)}{SplitInfo}$
* Reduction in Variance
    >When target column is not categorical.

    $variance(TargetColumn)=\frac{\sum(TargetColumn-\overline{TargetColumn})^2}{n}$

    $\overline{TargetColumn}=\frac{v_1+v_2+...}{2}$

    n = number of values

    $variance(TargetColumn, Column1)=\frac{\sum(TargetColumn-\overline{TargetColumn})^2}{n}$
* Chi-Square
    >The acronym CHAID stands for Chi-squared Automatic Interaction Detector

 


