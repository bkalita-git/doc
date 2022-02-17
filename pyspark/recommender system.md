## Techniques
 1. collaborative filtering
   > Collaborative filtering approaches build a model from a user's past behavior (items previously purchased or selected and/or numerical ratings given to those items) as well as similar decisions made by other users.
   * memory-based
     * user-based algorithm
   * model-based
     * Kernel-Mapping Recommender
   * Data Collection
     * explicit data collection include the following: 
       * Asking a user to rate an item on a sliding scale.
       * Asking a user to search.
     * implicit data collection include the following: 
       * Observing the items that a user views in an online store.
       * Analyzing item/user viewing times
   * Faced Problem
     * cold start
       * For a new user or item, there isn't enough data to make accurate recommendations.
     * scalability
       * There are millions of users and products in many of the environments in which these systems make recommendations. Thus, a large amount of computation power is often necessary to calculate recommendations.

     * sparsity
       * The number of items sold on major e-commerce sites is extremely large. The most active users will only have rated a small subset of the overall database. Thus, even the most popular items have very few ratings.
 2. content-based filtering 
   > Content-based filtering approaches utilize a series of discrete, pre-tagged characteristics of an item in order to recommend additional items with similar properties. Content-based filtering methods are based on a description of the item and a profile of the user's preferences.

 3. Knowledge based system
 4. Session-based recommender systems
 5. Reinforcement learning for recommender systems
 6. Multi-criteria recommender systems
 7. Risk-aware recommender systems
 8. Mobile recommender systems
 9. Hybrid recommender systems