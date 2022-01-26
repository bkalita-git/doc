## database
## data warehouse
>Data Warehouse is data that you have in other databases but organized specifically for the questions you want to ask.

## Datamart
subject-oriented relational database

multidimensional schema

Typical datafields include numerical order, timevalue and reference to one or more objects.
Three main type of schemas are
* Star Schema
  * fact table
    >here Customer ID, Time ID and Product ID are foreign key to dimension table.

    |Time ID|Product ID|Customer ID|Unit Sold|
    |-------|----------|-----------|---------|
    |4      |17        |2          |1        |
    |8      |21        |3          |2        |
    |-------|----------|-----------|---------|

  * Metric dimension table
    >More than one star schema can use a metric dimension table. 
  * dimension table
    > there is no dependency between dimension tables, so star schema needs fewer joins.

    Table Name: Customer ID
    |Customer ID|Name       |Gender|Income|Education|Region|
    |-----------|-----------|------|------|---------|------|
    |2          |Rajan      |M     |2     |3        |4     |
    |3          |prithoo    |M     |3     |4        |2     |
    |-----------|-----------|------|------|---------|------|

    >Likewise there will be Time ID and Product ID dimensions table.
    
* Snowflake Schema
  >Not to confuse with Snowflake company. A schema is known as a snowflake if one or more dimension tables do not connect directly to the fact table but must join through other dimension tables. Like in star schema we can see that fact table has foreign keys to dimension table. Consider Snowflake schema as it has one fact table and dimension tables like star schema but each dimension table can have another dimension table and so on.

* vault Schema

## data lake
## data cube
## data swamp