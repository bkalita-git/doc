Pyspark:
	ultimately python code will need a connection with JVM. and that's what it does when getcontext or session is written. Now, we must specify the JVM address(URI) to pyspark so it can communicate. In local mode if jvm is not running then it spins up 1 jvm.
