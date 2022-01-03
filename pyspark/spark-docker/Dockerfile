FROM ubuntu:bionic

#setup open-jdk
RUN apt-get update && apt-get install -y openjdk-8-jdk wget
ENV JAVA_HOME /usr/lib/jvm/java-1.8.0-openjdk-amd64/jre

#setup spark
RUN wget -P /tmp/ http://192.168.29.121/spark-3.2.0-bin-hadoop2.7.tgz
ENV SPARK_HOME /opt/spark
RUN tar xvf /tmp/spark-3.2.0-bin-hadoop2.7.tgz -C /tmp && mv /tmp/spark-3.2.0-bin-hadoop2.7 ${SPARK_HOME}
RUN rm /tmp/spark-3.2.0-bin-hadoop2.7.tgz


EXPOSE 7077 4040 8080 22

#run post setup
COPY /script_files/bootstrap.sh /
ENV PATH="${SPARK_HOME}/bin:${SPARK_HOME}/sbin:${PATH}"
ENTRYPOINT ["/bin/bash", "bootstrap.sh"]

