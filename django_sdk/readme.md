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
	python manage.py createsuperuser-fail                 Fail silently (no output at all) on HTTP errors
 -h, --help <category>      Get help for commands
 -i, --include              Include protocol response headers in the output
 -o, --output <file>        Write to file instead of stdout
 -O, --remote-name          Write output to a file named as the remote file
 -s, --silent               Silent mode
 -T, --upload-file <file>   Transfer local FILE to destination
 -u, --user <user:password> Server user and password
 -A, --user-agent <name>    Send User-Agent <name> to server
 -v, --verbose              Make the operation more talkative
 -V, --version              Show version number and quit

This is not the full help, this menu is stripped into categories.
Use "--help category" to get an overview of all categories.
For all options use the manual or "--help all".
[bipul@archlinux ~]$ 

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
* django:
	```
	#python manage.py makemigratations #to generate sql from models.
	#python manage.py migrate #to apply those above code
	#results = Users.objects.all() here results is "querysets"
	#django model by default creates "id" field in each class.

    # below "user" becomes "user_id" in real table and UserFiles().user.email can be used. also User().files can be used
	Model UserFiles:
		user = models.ForeignKey(
			to=User, on_delete=models.CASCADE, related_name='files')

	# the ManyToOneRel is different and not to be used.
	```


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
## Extras - OpenID Connect
>OpenID Connect 1.0 is a simple identity layer on top of the OAuth 2.0 protocol. OpenID Connect performs many of the same tasks as OpenID 2.0. in OpenID Connect, OAuth 2.0 capabilities are integrated with the protocol itself. basically Open ID Connect and Open ID 2.0 and OAuth 2.0 is a server-client thing. So, we need the server version of it and then connect using its client which must be registered with the server version. for example, google OAuth2.0 server to communicate with it, it provides HTTP Request, Google Sign in and Google client Libraries for different programming languages to iteract. the service will provide **document discovery** where all necessary endpoints will be included. The main differene between openid connect and OAuth2.0 is Openid connect gives an extra token called id_token along with access_token which tells the information about the user, but access_token does not contain information about user. **if say request scope="openid profile email" then the access_token won't get the profile info of user but id_token will display the information about the user.** But Google way is different. Google have many APIs. and each API has some scopes. scope defines the information we want from that specific API. what if we send a scope = "https://www.googleapis.com/auth/userinfo.email" in request param?. then after login it adds two more scope with it in response, these are "email", "https://www.googleapis.com/auth/userinfo.email" and "openid". but same scope is linked with more than 1 API so? which API was used? the answer is to whom you are sending the request. for now, https://accounts.google.com/o/oauth2/v2/auth is where the scope was sent which is a Google OAuth2 API, v2. so How it added "openid" in response even though we did not define in google oauthconsent screen and also we did not sent it in request. the thing is openid is not an API, by default google OAuth2.0 uses openid by default which will give us a id_token. simply, access_token is linked with a specific API and id_token is the information of user. so why we can't get a id_token without access_token? the access_token always come. it's because we ask for it to oauth2 API which is a service and it provides the access_token to use its service along with id_token. so, basically, google OAuth2.0 API provides open id connect features too. so, using access_token you can use a service but using id_token you can't use a service. because id_token is like a status that the user logged in with gmail. no need to store those in server.
***the confusion is using access_token also an django server app can call userinfo endpoint to obtain more information about the user. But, userinfo endpoint is the part of openid connect system. Also, another thing is that, django server app gets two tokens at the same time. Access_token and id_token, that means we can verify that access_token is not send to the django server app by some other site since django app server also gets id_token so that user must have logged in recently. Also the confusion is before openid connect Google, Facebook used oauth2.0 for authentication also, it's like a hack for oauth2 to use authentication.***
```
#OAuth2.0 web server flow
server.tempCode = client.get_temporary_code_from_OAuth2.0()
server.access_token = server.get_real_acccess_token_from_OAuth2.0(server.tempCode)
result = google.contact_service(server.access_token)
#this access_token is for authorization with google contact service. this token contains the information of scope and key to authorize the google contact service for example.  

#OpenIDConnect web server flow
server.tempCode = client.get_temporary_code_from_OAuth2.0()
(server.access_token, server.id_token) = server.get_real_acccess_token_from_OAuth2.0(server.tempCode)
server.user_profile = decode(server.id_token)
#this id_token is a JWT and can be used to authenticate a client. this token contains the information of the client.
```

* using HTTP Request	
  * Create an anti-forgery state token. which will be `state`
  * get the "`authorization_endpoint`" value from this link. it's called **document discovery**. https://accounts.google.com/.well-known/openid-configuration. for now as of updated by google it's 
	```
	{"authorization_endpoint":"https://accounts.google.com/o/oauth2/v2/auth"}
	```
	which will be used as BASE URI
  * get `client_id` from https://console.developers.google.com/apis/credentials. which is a django app for example. that client must be registered before.
  * `response_type` which can be retrieve from document discovery.
  * `scope` which can be retrieve from document discovery. this value describes if it's simple oauth or open id connect.
  * `redirect_uri` this value must be matched with the redirect_uri defined in while making the client credentials in oauth server.
  * `login_hint` which is a user's email id which will allow the Oauth server to access its data.
  * send HTTPS GET request to `authorization_endpoint` to get `code` from google which is a temorary grant. but not very secure so again get access token in next step.
	```
	https://accounts.google.com/o/oauth2/v2/auth?
	response_type=code&
	client_id=117939782928-pouc8jel9tf0iu31fm2vlb1k84es2h5r.apps.googleusercontent.com&
	scope=openid%20email&
	redirect_uri=http%3A//localhost:8000/uploadfile/&
	state=security_token%3D138r5719ru3e1%26url%3Dhttps%3A%2F%2Foauth2-login-demo.example.com%2FmyHome&
	login_hint=hastetough1@gmail.com

	#it will redirect this below and extract the `code` value which is a token. but the value `state` must be checked with the `state` which was sent so that this request is the same request that was sent by the server.
	http://localhost:8000/uploadfile/?state=security_token%3D138r5719ru3e1%26url%3Dhttps%3A%2F%2Foauth2-login-demo.example.com%2FmyHome&code=4%2F0AX4XfWgT8B3oPMbmhorP1Che4KvPtI0iy-3T07Um9amKxWLgYLq1KSMVLQ&scope=email+openid+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email&authuser=1&prompt=consent
	```
  * After getting the `code` from google exchange get access_token and id_token from  google with POST request to `token_endpoint`
	```
	POST /token HTTP/1.1
	Host: oauth2.googleapis.com
	Content-Type: application/x-www-form-urlencoded

	code=4/P7q7W91a-oMsCeLvIaQm6bTrgtp7& #which was sent by google
	client_id=your-client-id&
	client_secret=your-client-secret&
	redirect_uri=http%3A//localhost:8000/uploadfile/&
	grant_type=authorization_code

	#Response will be 
	access_token #A token that can be sent to a Google API.
	expires_in 	#The remaining lifetime of the access token in seconds.
	id_token  #A JWT that contains identity information about the user that is digitally signed by Google.
	scope #The scopes of access granted by the access_token expressed as a list of space-delimited, case-sensitive strings.
	token_type 	#Identifies the type of token returned. At this time, this field always has the value Bearer.
	refresh_token #(optional) This field is only present if the access_type parameter was set to offline in the authentication request. For details, see Refresh tokens.
	```
  * decode `id_token` which is a JWT using encoding information in `jwks_uri` which is in **discovery document**. so, jwt has 3 parts, first one is header then payload and then signature. both header and payload can be decoded using BASE64URL decoder which does not need any secret key. but the signature string must be validate. to do so, 
	```
	curl --location --request GET 'https://www.googleapis.com/oauth2/v3/certs'
	```
  * decode `id_token` JWT for development purpose.
	```
	curl --location --request GET 'https://oauth2.googleapis.com/tokeninfo?id_token=the_id_token_sent_by_google
	```
  * Obtain user information using the `access_token` by sending GET request to `userinfo_endpoint` in discovery document.
	```
	curl --location --request GET 'https://openidconnect.googleapis.com/v1/userinfo?scope=openid%20profile%20email' \
	--header 'Authorization: Bearer the access_token_sent_by_google'
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
7. https://developers.google.com/identity/protocols/oauth2/openid-connect
8. https://developers.google.com/identity/protocols/oauth2/scopes
9. https://firebase.google.com/docs/auth/admin/verify-id-tokens#python
		
