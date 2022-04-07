## concept
* Weak Learner
  * A model which is trained with not much sufficient data such that subset of feature columns and subset of rows are used to train.
* Strong Learner
  * multiple weak learner together can be called as Strong Learner or A model which is trained with more data.
* Base Model/learners
  * those models which are used first to feed the test/train data in.
* Meta Model/learners
  * A model which takes data from Base models output to train
* Heterogeneous Weak Learner
  * each Model can be of different types for example, $M_1$ can be SVM and $M_2$ can be decision Tree
* Homogeneous Weak learner
  * each Model $M_i$ is of same type. for example all models are LinearRegressor
* stump
  * Decision tree with depth 1
## Types of Ensemble Learning:
* ## Bagging  
  we have train dataset TRAIN_D and a test dataset TEST_D and we have max sample size $S$ for each model $M_i$. we will build and train $M_i$ models.
  1. for every Model $M_i$ do row sampling with replacement of size $S$ from TRAIN_D such that each small dataset from $M_i$ are not same and train $M_i$
  2. feed TEST_D to $M_n$ models and select the result for each sample in TEST_D where the result comes maximum number of times.
  * ### Types
    1. Bagging meta-estimator
    2. Random forest
        * with decision tree

* ## Stacking
  * ### Hold out method (BLENDING)  
    Dataset can be divided by % of test and train purpose. For example
    ```
    MAIN_TRAIN_D, MAIN_TEST_D = Dataset.randomSplit(80%,20%)
    
    TRAIN_D_L1,TEST_D_L1 = MAIN_TRAIN_D.split(80%,20%)

    lr = LinearRegression(featuresCol="x",labelCol="y",predictionCol="lr_prediction")
    dt = DecisionTreeRegressor(featuresCol="x",labelCol="y",predictionCol="dt_prediction")
    kr = KNNRegressor(featuresCol="x",labelCol="y",predictionCol="kr_prediction")
    va = VectorAssembler(inputCols=["lr_prediction","dt_prediction","kr_prediction"],outputCol="base_features")
    pipeline = Pipeline(lr,dt,kr,va)
    base_model = pipeline.fit(TRAIN_D_L1)
    base_prediction_df=base_model.transform(TEST_D_L1)
    

    rfr = RandomForestRegressor(featuresCol="base_features",labelCol="y", predictionCol="meta_prediction")

    meta_model = rfr.fit(base_prediction_df)


    ## test
    meta_model.transform(base_models.transform(MAIN_TEST_D))
    ```
  * ### K-Fold method (STACKING)
    ```
    MAIN_TRAIN_D, MAIN_TEST_D = Dataset.randomSplit(80%,20%)
    TRAIN_D,TEST_D = D.split(80%,20%)

    heterogeneous_train_models = [LinearRegressor, DecesionTreeRegressor,KNNRegressor]
    meta_train_model = RandomForestRegressor
    k=5
    D_ARRAY = TRAIN_D.split(k times equally)
    for i in len(D_ARRAY):
        TEST_D_i  = D_ARRAY[i] 
        TRAIN_D_i = D_ARRAY[all element except i]
        heterogeneous_train(TRAIN_D_i)
        NEW_predictions_D_ARRAY.append(predictions = heterogeneous_test(TEST_D_i), for_position=i)

    heterogeneous_train(TRAIN_D)
    (TEST_D)
    
    ```
  
* ## Boosting  
  we have train dataset TRAIN_D and a test dataset TEST_D and we have $n$ steps to perform and for each step we have max sample size $S$ and for each step there will a model we will build and train called $M_i$ 
  1. Build a model Model $M_1$ select a random sample of size $S$ from TRAIN_D and train $M_1$. 
  2. feed whole TRAIN_D to $M_1$ and preserve/mark missclassification samples as error dataset as M_1_MISS_D
  3. Build Model $M_2$ and select a random sample of size $S$ from both M_1_MISS_D and TRAIN_D such that 90% of sample comes from M_1_MISS_D and train $M_2$
  4. continue step 2 and 3 to build $M_n$ models.
  5. feed TEST_D to $M_n$ models and select the result for each sample in TEST_D where the result comes maximum number of times.
   
  * ### Types
    > Different types of boosting occurs in terms of the technique used to select the rows for next $M_{i+1}$ model
    1. Adaptive Boost (Ada Boost) with
        * deicsion tree
      > Misclassified rows are marked by setting higher weight for them. each model is a decision stump.
    2. Gradient Boost with
        * decision tree
      > each model is decision tree with no depth restriction but must set before for all $M_i$. misclassified rows are marked by difference between (actual-predicted)
    3. Extreme Gradient Boost (XG Boost) with
        * decision tree




## Random Forest
Use Bagging where all $M_i$ are decision tree and each random sample with replacement of size $S$ for each Model containts random subset of features column.

## Ada Boost
Use Boosting where $$

## why does boosting with decision tree?
Boosting is again just easier to discuss when talking about decision stumps (1 depth decision trees) as they are simple high bias algorithms which boosting aims at and is good at improving. 