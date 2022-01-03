#V0.2 ENABLING HISTORY SERVER

from pyspark import SparkContext, SparkConf
from pyspark.sql import SparkSession

#SPARK CONF FOR CONTEXT
conf = SparkConf().setAppName("main_app").setMaster("local[*]")
conf = conf.set("spark.eventLog.enabled",True).set("spark.history.fs.logDirectory","/tmp/spark-events")

#CREATING THE SPARK CONTEXT
sc = SparkContext(conf=conf)

#CREATING MULTIPLE SPARK SESSIONS
#SESSION 1
s1 = SparkSession(sc).builder.appName("u1").getOrCreate() #use .config for builder
#SESSION 2
s2 = s1.newSession()
s2.conf.set("spark.app.name","u2")
#SESSION 3
s3 = s1.newSession()
s3.conf.set("spark.app.name","u3")

#PRINTING THE SPARK CONTEXT OF EACH OF THE SPARK SESSION
print("Session 1 context ID=",id(s1.sparkContext),"\tSession 2 context ID=",id(s2.sparkContext),"\tSession 3 context ID=",id(s3.sparkContext))
'''
Session 1 context ID= 140024041727408 	Session 2 context ID= 140024041727408 	Session 3 context ID= 140024041727408
'''


#PRINTING THE SESSION ID
print("Session 1 ID",s1._jsparkSession.sessionUUID(),"\tSession 2 ID",s2._jsparkSession.sessionUUID(),"\tSession 3 ID",s3._jsparkSession.sessionUUID())
'''
Session 1 ID 4a3d8e06-33d8-4b2a-ba1e-c54525c42cd7 	Session 2 ID 48006897-a216-4609-9ed7-fb0060e5cc86 	Session 3 ID f1027e66-33e4-405c-a663-3eb1f81a58be
'''

# CREATING DATAFRAME IN EACH OF THE SPARK SESSION
df1 = s1.read.format("csv").option("header",True).option("inferSchema",True).option("delimiter","\t").option("path","marketing_campaign.csv").load()
df1 = df1.sample(False,0.01).select(df1[0])
df2 = s2.read.format("csv").option("header",True).option("inferSchema",True).option("delimiter","\t").option("path","marketing_campaign.csv").load()
df2 = df2.sample(False,0.01).select(df2[1])
df3 = s3.read.format("csv").option("header",True).option("inferSchema",True).option("delimiter","\t").option("path","marketing_campaign.csv").load()
df3 = df3.sample(False,0.01).select(df3[2])
'''
>>> df1.show()
+-----+
|   ID|
+-----+
| 6864|
| 9701|
|  624|
| 4785|
| 2607|
| 6131|
| 2258|
| 3277|
| 9826|
| 7034|
| 4597|
| 1970|
| 1307|
|11114|
|10475|
|10133|
| 2525|
| 3865|
|  202|
|10212|
+-----+

>>> df2.show()
+----------+
|Year_Birth|
+----------+
|      1971|
|      1972|
|      1966|
|      1973|
|      1959|
|      1958|
|      1949|
|      1968|
|      1971|
|      1975|
|      1965|
|      1960|
|      1947|
|      1987|
|      1948|
|      1969|
|      1976|
|      1975|
|      1947|
|      1974|
+----------+

>>> df3.show()
+----------+
| Education|
+----------+
|Graduation|
|Graduation|
|       PhD|
|Graduation|
|    Master|
|Graduation|
|       PhD|
|  2n Cycle|
|Graduation|
|    Master|
|Graduation|
|Graduation|
|Graduation|
|Graduation|
|Graduation|
|Graduation|
|     Basic|
|     Basic|
|Graduation|
|       PhD|
+----------+

'''


#PRINTING SESSION ID OF EACH DATAFRAME
print("DF 1 Session ID ",df1._jdf.sparkSession().sessionUUID(),"\tDF 2 Session ID ",df2._jdf.sparkSession().sessionUUID(),"\tDF 3 Session ID ",df3._jdf.sparkSession().sessionUUID())
'''
DF 1 Session ID  4a3d8e06-33d8-4b2a-ba1e-c54525c42cd7   DF 2 Session ID  48006897-a216-4609-9ed7-fb0060e5cc86 	DF 3 Session ID  f1027e66-33e4-405c-a663-3eb1f81a58be
'''

#GET THE DATAFRAME FROM EACH OF THE SESSION AND MERGED INTO A SINGLE ONE.
#The leftmost  DF's Session ID will be the final DF's session ID.
df = df3.join(df2).join(df1)
print(df.limit(20).show())
'''
>>> df.show()
+----------+----------+-----+
| Education|Year_Birth|   ID|
+----------+----------+-----+
|Graduation|      1971| 6864|
|Graduation|      1971| 9701|
|Graduation|      1971|  624|
|Graduation|      1971| 4785|
|Graduation|      1971| 2607|
|Graduation|      1971| 6131|
|Graduation|      1971| 2258|
|Graduation|      1971| 3277|
|Graduation|      1971| 9826|
|Graduation|      1971| 7034|
|Graduation|      1971| 4597|
|Graduation|      1971| 1970|
|Graduation|      1971| 1307|
|Graduation|      1971|11114|
|Graduation|      1971|10475|
|Graduation|      1971|10133|
|Graduation|      1971| 2525|
|Graduation|      1971| 3865|
|Graduation|      1971|  202|
|Graduation|      1971|10212|
+----------+----------+-----+

'''

#PRINT THE SESSION ID OF RESULTANT DATAFRAME
print("FINAL DF Session ID ",df._jdf.sparkSession().sessionUUID())
'''
FINAL DF Session ID  8758af8f-b868-4ad5-bd66-eb07f3f4394f
'''

#PRINT DEFAULT SESSION ID
#It can be accessed through any session, i.e. s1,s2,s3
print("DEFAULT SESSION ID ",s1.getActiveSession()._jsparkSession.sessionUUID())
'''
DEFAULT SESSION ID  8eaf18b0-db9b-4fd2-8e16-722908a54243
'''



