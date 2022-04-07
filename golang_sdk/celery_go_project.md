1. check rabbitmq is running
    ```
    sudo systemctl status rabbitmq
    ``` 
2. configure rabbitmq for celery
   ```
    # add user 'bipul' with password 'password'
    $ rabbitmqctl add_user bipul password

    # add virtual host 'bipul_vhost'
    $ rabbitmqctl add_vhost bipul_vhost

    # add user tag 'bipul_tag' for user 'bipul'
    $ rabbitmqctl set_user_tags bipul bipul_tag
    
    # set permission for user 'bipul' on virtual host 'bipul_vhost'
    $ rabbitmqctl set_permissions -p bipul_vhost bipul ".*" ".*" ".*"
   ```
3. Install Celery
   >celery is a python package so can be installed with pip inside a venv
    ```
    python3.7 -m venv venv
    . venv/bin/activate
    pip install celery
    ```
4. BROKER_URL: "amqp://user:password@remote.server.com:port//vhost"
5. celery -A tasks worker --loglevel=INFO
6. calling the tasks
   ```
   from tasks import add
   add.delay(4, 4)
   ```
celery_client->rabbitmq->celery_worker


7. Testing celery using rabbitMQ client
```
pip install pika
```
run the code below and it will show something in celery daemon.
```
import pika

credentials = pika.PlainCredentials('bipul', 'password')
parameters = pika.ConnectionParameters('localhost',5672,'bipul_vhost',credentials)

connection = pika.BlockingConnection(parameters)
channel = connection.channel()
#channel.queue_declare(queue='hello')
channel.basic_publish(exchange='',routing_key='celery',body='Hello World!')
print(" [x] Sent 'Hello World!'")
connection.close()
```


## submitting to rabbitMQ from python using celery client and compute in golang celery worker.
so, the celery daemon will be in golang and client will be in python. it turns out that celery worker in go is new and not matured. celery needs broker URI and backend URI. the broker is the message broker for example rabbitMQ. and the broker will be the place where celery worker put its result. Backend can be SQLAlchemy/Django ORM, MongoDB, Memcached, Redis, RPC (RabbitMQ/AMQP).
```
$ go mod init techvariable.com/celery_go_service
#it will add requirements in go.mod file
$ go get -u gorm.io/gorm
$ go get -u gorm.io/driver/sqlite
$ go get -u github.com/gocelery/gocelery@latest

$ pip install flower
$ celery -A test_celery flower
```

## error in django
if the file structure is this -
```
MainProjectDir
   MainApp
   AnotherApp
      - __init__.py
      - celeryapp.py
      - celeryappworker.py
```
then run celery worker using this
```
#set your working dir to 'MainApp'
$ celery -A AnotherApp.celeryworker worker --loglevel=INFO
```
if celery project is in other directory outside django then do this in django's client celery file
also, specify the task name in django with this 
```
@app.task(name='celery_livy_service.celerytracksession.track_session')
```

## Conclusion
when we run celery daemon, there we provide the message broker url, i.e. rabbitMQ because the celery daemon makes a queue in that rabbitMQ and whenever something comes up in that rabbutMQ celery worker will get run automatically. Now who can push to the rabbitMQ? both pika client and celery client can push to that rabbitMQ from python. the difference is that in celery client we don't need to explicitly provide the connection info to rabbitMQ.

## REFERENCES
1. https://pkg.go.dev/github.com/gocelery/gocelery