## Installation of django and rest_framework
*
	```
	#install django 
	pip install django
	#install rest_framework
	pip install djangorestframework
	#create a project
	django-admin startproject project_namepython manage.py runserver
	#create an app inside project
	python manage.py startapp app_name
	#create a user for db
	python manage.py createsuperuser
	#apply model changes
	python manage.py makemigrations
	python manage.py migrate
	```
*	*project_name and app_name will be on same directory level and will be inside project_name folder.*


## cURL for sending different types of requests.
* >data can be passed as queryparam or payload. payload can be formatted as form-data, json-data, text-data, etc. queryparam goes in the url not in payload remember that.
	```
	#query data
	curl --location --request POST 'localhost:8000/products/' -d "pid=23424"
	#json data
	curl --location --header "Content-Type: application/json" --request POST 'localhost:8000/products/' -d '{"pid":23424}'
	#form data, which is bydefault
	curl --location --header "Content-Type: application/x-www-form-urlencoded" --request POST 'localhost:8000/products/' --form "pid=23424" --form "key=value"
	#for asking a token with username and password
	curl --location --request POST 'localhost:8000/login/api-token-auth/' --header "Authorization: Basic encoded_data_from_user_password"
	#for asking data with token
	curl --location --request GET 'localhost:8000/' --header "Authorization: Token token_value"
	```

## where's our wsgi app?
* in Flask:
* in FastAPI:
* in Django:
  ```
  #in wsgi.py
  application = get_wsgi_application()
  ```

## where's our config file?
* in Flask:
* in FastAPI:
* in Django:
  >In settings.py in project_name/project_name and consumed in wsgi.py. Also urls are consumed by settings 

## Middleware
  >In our WSGI_SERVER(Apache(request/response))<->WSGI_app(request/response)<->request/Response <-> Middleware <-> Views. SO request hits middleware before views.
* flask
* FastAPI
* Django
  > in settings.py  MIDDLEWARE=[], order of middleware is important.Create a directory called middleware inside app_name as a package and add in settings.py
  * **Built-in**
  * **Custom**

## Route/Controller options?
	django:
		everything inside urlpatterns as list
			#can use a function for view, they can be from DJANGORESTFRAMEWORK or BUILTIN
			#class-based view, can use a class for view: they can be from DJANGORESTFRAMEWORK or BUILTIN
			#can add another urls file with include(filepath)
## Accesing and validating request Data and Response as well
> The request object contains requested *methods, headers, and payload.* the payload may comes in different forms, like- formData, QueryParam/KeyValue, Json data etc. Most of the time we first convert the payload to our needs and then we process the data in views. Even though we can access data with **request.data** it's better to use a **serializer**.
* in Flask

* in FastAPI

* in Django:
  * **Built in**
	1. ```
		from django.cores import serializers
	   ```
  
  * **rest_framework**
    1. ```
	    from rest_framework import serializers

		#extends serializer.Serializer

		#extends serializer.ModelSerializer

		#extends serializer.ListSerializer

		#extends serializer.BaseSerializer
  	   ```
## different types of views
* in Flask:
  > use **flask_restful** package
* in FastAPI:
* in Django:
  > a view can be example of ListUser, CreateUser etc. and a viewset is set of views. Thre are Built in and **rest_framework** views. views can be reusable/subclass. 

  * **Built in**:
	* ***class-based views***:
		```
		from django.views import View
		```
	* ***function-based views***:
		```
		#in the function check request.method to each method and return http response accordingly
		```
  * **rest_framework**:
	* ***class-based views***:
    	1. ```
		   from rest_framework.views

		   #extends views.APIView
		   #this APIView class extends from django's View class
		   ```
		2. ```
		   #this class extends APIView
		   #there are prebuilt views and to use them use generics
		   from rest_framework import generics

		   #extends generics.ListCreateAPIView
           ```
  		3. ```
		   from rest_framework import viewsets

		   #extends viewsets.GenericViewSet

		   #extends viewsets.ViewSet

		   #extends viewsets.ModelViewSet

		   #extends viewsets.ReadOnlyModelViewSet
   		   ```

		4. ```
      	   from rest_framework import mixins
			
		   #extends mixins.CreateModelMixin

		   #extends mixins.ListModelMixin

		   #extends mixins.RetrieveModelMixin
		   ```
	* ***function-based views***:
		1. ```
           from rest_framework.decorators import api_view

		   #use @api_view
		   ```

## the request object and response object:  
	django:
		option 1. add 'request' param in every view function for receiving request
		option 2. for making a new request to others. from django.http import HttpRequest
		option 2. return httpresponse from django.http import HttpResponse

	
## how to do dependency injection?
	django:

## database_orm?
	django:
		#python manage.py makemigratations #to make the tables in db from models.
		#python manage.py migrate
		#results = Users.objects.all() here results is "querysets"
		#django model by default creates "id" field in each class. 


## Extras - Authentication then Authorization
> As we know rest architecture should not involve any transaction where the server has to store a state of client, for example *Session*. Authentication and Authorization can be done manually in views but why implementing if they are already available for us.
* In Flask
  
* In FastAPI
  
* In Django
  * **Built in**
  
  * **rest_framework**
	> client sends credentials in headers and payload and based on that **permission** and **throttling** policies will determine whether to execute main login inside view. **rest_framework** provides 3 types of **authentication** namely **BasicAuthentication**, **SessionAuthentication**, **TokenAuthentication**, each of these will try to validate the request by checking existing DB along with them we can set **permission** and **throttling**.
     1. ```
		#globally in settings.py
		REST_FRAMEWORK = {
    		'DEFAULT_AUTHENTICATION_CLASSES': [
        			'rest_framework.authentication.BasicAuthentication',
        			'rest_framework.authentication.SessionAuthentication',
    		]
		}
        ```
	2. ```
		#as per view separately inside APIView
		from rest_framework.authentication import SessionAuthentication, BasicAuthentication
		from rest_framework.permissions import IsAuthenticated
		
		#can be accessed through user&password or TokenAuthentication or SessionAuthentication
		authentication_classes = [SessionAuthentication, BasicAuthentication, TokenAuthentication]
    	permission_classes = [IsAuthenticated]
	   ```
	3. ```
		#token authentication, in installed apps in settings.py
		#also migrate db
		'rest_framework.authtoken'
		from rest_framework.authtoken.models import Token
		token = Token.objects.create(user='username')
	   ```

	4. ```
		#use built in view to retrieve token but must send user name and password in body
		from rest_framework.authtoken import views
		path('get_token/',views.obtain_auth_token)
	   ```
	5. ```
		#for jwt
		pip install djangorestframework-simplejwt
       ```
## Extras - Template as Rensponse
* In Flask
  
* In FastAPI
  
* In Django
  * **Built in**
  * **rest_framework**
	1. 
## Extras - Already Provided Apps/models/templates
* In Django
  * >User model. There are 5 ways to extends this. 1. using a proxy model 2. using one-to-one link 3. extending AbstractBaseUser 4. extending AbstractUser
  	```
	from django.contrib.auth.models import User
	#this User is extended from AbstractBaseUser
	#fields are username,first_name,last_name,email,password,groups,user_permissions,is_staff,is_active,is_superuser,last_login,date_joined.
	#Also this model has attributes which are helper function and use it rather than creating directly with objects.create..
    ```
  * >django has many apps which can then included in INSTALLED_APPS. if you want to change the default behaviour of prebuilt apps then use settings.py to store that changes. For example to change the default User table in django.contrib.auth use "AUTH_USER_MODEL = 'login.User'" in settings.py, where User is a model but we write app_label.model_name.
  so, if INSTALLED_APPS=[django.contrib.auth] then all the models/tables inside that app. ie. app.models.ModelName will be created when doing makemigrations and migrates.
  By default for django project these two models must be needed. There are 
  django.contrib.auth.models.Permission and django.contrib.contenttypes.models.ContentType so in DB, the two actual tables will be auth_permission and django_content_type.

  * **App django.contrib.auth**
	>The user system which can be customized. 
	1. ```
		# Custom User table in app_name.models
		class User(AbstractUser):
			...
		#in settings.py
		INSTALLED_APPS=[
			'django.contrib.contenttypes',
			'django.contrib.auth', 
			'app_name'
		]
		AUTH_USER_MODEL = 'app_name.User'
	   ```
	2. ```
		#
   	   ```
  * **App django.contrib.admin** 
	>used to interact with *django.contrib.auth* with GUI


## Extras - browse sqlite file
* In Django
  > Install DB browser
  ```
  #in settings.py
  ```
## Extras - Testing
```
from rest_framework.test import APIRequestFactory
#Using the standard RequestFactory API to create a form POST request
factory = APIRequestFactory()
request = factory.post('/notes/', {'title': 'new idea'})
```
REFERENCES:
1. https://www.django-rest-framework.org/
2. https://docs.djangoproject.com/en/4.0/topics/
3. https://django-rest-framework-simplejwt.readthedocs.io/en/latest/index.html
4. https://www.cdrf.co/
5. https://simpleisbetterthancomplex.com/tutorial/2016/07/22/how-to-extend-django-user-model.html
6. https://docs.djangoproject.com/en/4.0/topics/auth/
		
