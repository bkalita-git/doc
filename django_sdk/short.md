## ModelManager

Models can have ModelManager, for example User.objects.all() executes a query which is written in "objects" model manager .all() method.
example:

```
from django.db import models

class StudentManager(models.Manager):
    def get_queryset(self):
        return super().get_querysey().filter(desg="student")

class Person(models.Model):
    desg = models.CharField(max_length=100)
    name = models.CharField(max_length=100)

    #here Person.objects.all() returns all rows
    objects = models.Manager()

    #here Person.students.all() returns all students rows only
    students = StudentManager()

```

## Model Helper Function

method inside a model becomes helper function. instance of that model can use those like, save() method is an example.

## Templates folder

if you add a directory named "templates" inside any **app** not inside the root directory, then these will be used as templates path. this is specified in settings.py in **TEMPLATES** where **_'APP_DIRS':True_** and can be specified extra directory in **_'DIRS':[]_**

## settings.PASSWORD_HASHERS

in **django/contrib/auth/models.py** in class UserManager(BaseUserManager) it has a function named def create*user(...) which calls def \_create_user(...) which is calls **django/contrib/auth/models.py** def make_password(password) inside that it calls hasher=get_hasher(hasher='default') and return hasher.encode(password,salt or hasher.salt()). inside def get_hasher() and it calls def get_hashers() and there it's written \*\*\_for hasher_path in settings.PASSWORD_HASHERS:*\*\*

## django form

just like serializer django form is also same.

```
from django.forms import ModelForm
from .models import User
class UserForm(ModelForm):
    class Meta:
        model = User
        fields = '__all__'

from django.shortcuts import render

def simpleView(request):
    form = UserForm()
    context = {'form':form}
    return render(request,'register/signup.html',context)

```

in app/templates/register/signup.html

```
<form action="" method="">
    {{form}}
    <input type='submit'/>
</form>
```

## django auth system architecture

> 3 main components are Model, ModelManager and **init**.py

APP_DIR = **django/contrib/auth/**

APP_DIR/**init**.py contains methods like login(), logout(), authenticate() which interacts with User model.

APP_DIR/models.py contains Model classes and ModelManager classes. each Model contains one or more Helperfuncion and ModelManager() instances. Helperfunction only be called using instance of a Model class and methods of ModelManager can be called without instantiating Model class. Methods of ModelManager interacts with the builtin ModelManager object which then interacts with the DatabaseTable. for example.

```
#Model Manager
class UserManager(models.Manager):
    def create(self):
        #interacts with User Table in DB
        ...
    def all(self):
        #interacts with User Table in DB
        ...
#Model
class User(models.Model):
    ...
    #Helperfunction
    def save(self):
        ...

    #instance of ModelManager
    objects = models.UserManager()

u1=User.objects.create(username="bipul") #uses ModelManager
u1.save() #uses helper function
```

Configurable Settings for auth system in settings.py

- for **init**.py

  1. settings.AUTHENTICATION_BACKENDS
  2. settings.AUTH_USER_MODEL

- for hashers.py

  1. settings.PASSWORD_HASHERS

- for password_validation.py
  1. **settings.AUTH_PASSWORD_VALIDATORS**
     > which is only used in forms.py. and which is used in a custom admin class called 'UserAdmin(admin.ModelAdmin)' which is registered with admin system.
- for tokens.py
  1. settings.SECRET_KEY
  2. settings.PASSWORD_RESET_TIMEOUT
- for views.py
  1. settings.LOGIN_REDIRECT_URL
  2. settings.LOGOUT_REDIRECT_URL
  3. settings.LOGIN_URL

How to interact with the auth system?

- It gives 6 systems to interact.
  1. using **views** which endpoints are in **django.contrib.auth.urls**
  2. using **forms** which are in **django.contrib.auth.forms**
  3. using **admin app** since auth system is registered with admin
  4. directly using **modelmanager** methods which already included in User model. **django.contrib.auth.models.User**
  5. directly using the **django.contrib.auth.models.User** class which is not recomended.
  6. the commands for manage.py **createsuperuser**
  7. Also using **rest_framework** we can interact with it.

## Django request vs rest_framework request

## Default middlewares in django

> middlewares provide configurable settings in settings.py. Some middlewares provided by Django framework (eg: django.middleware...) and some middlewares are provided by other apps(eg: django.contrib.auth.middleware...).

- > django.middleware.security.SecurityMiddleware

  provides following configurable settings

  ```
  SECURE_BROWSER_XSS_FILTER
  SECURE_CONTENT_TYPE_NOSNIFF
  SECURE_CROSS_ORIGIN_OPENER_POLICY
  SECURE_HSTS_INCLUDE_SUBDOMAINS
  SECURE_HSTS_PRELOAD
  SECURE_HSTS_SECONDS
  SECURE_REDIRECT_EXEMPT
  SECURE_REFERRER_POLICY
  SECURE_SSL_HOST
  SECURE_SSL_REDIRECT
  ```

- > django.contrib.sessions.middleware.SessionMiddleware

- > django.middleware.common.CommonMiddleware

  ```
  #list of compiled regular expressions objects
  DISALLOWED_USER_AGENTS
  APPEND_SLASH
  PREPEND_WWW
  ```

- > django.middleware.csrf.CsrfViewMiddleware

- > django.contrib.auth.middleware.AuthenticationMiddleware

- > django.contrib.messages.middleware.MessageMiddleware

- > django.middleware.clickjacking.XFrameOptionsMiddleware

  It is used to set response header with below option. By default this middleware sets DENY for every response.

  ```
  X-Frame-Options : SAMEORIGIN OR DENY
  ```

  If SAMEORIGIN then client can load the resorce in a iframe iff both origin is same. and DENY then this resource will not be rendered in any iframe. this is used to prevent clickjacking hack.

## DJANGO APPS AND MIDDLEWARES AND TABLES CONNECTIVITY

| directory      | app            | tables[fields]                                                                                                                                                                                                                                                                                                                                |
| -------------- | -------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|                | contenttypes   | django_content_type[id,app_label,model]                                                                                                                                                                                                                                                                                                       |
| contrib        | auth           | auth_permission[id,content_type_id,codename,name],<br> auth_group[id,name],<br> auth_group_permissions[id,group_id,permission_id]<br>auth_user[password,last_login,is_staff,id,email,is_super,first_name,last_name,date_joined,is_active],<br> auth_user_user_permissions[id,user_id,permission_id],<br>auth_user_groups[id,user_id,group_id] |
|                | sessions       | django_session[session_key,session_data,expire_date]                                                                                                                                                                                                                                                                                          |
|                | message        |                                                                                                                                                                                                                                                                                                                                               |
|                | admin          | django_admin_log[id,action_time,object_id,object_repr,change_message,content_type_id,user_id,action_flag]                                                                                                                                                                                                                                     |
|                | staticfiles    |                                                                                                                                                                                                                                                                                                                                               |
| rest_framework | auth_token     | authtoken_token[key,created,user_id]                                                                                                                                                                                                                                                                                                          |
|                | rest_framework |                                                                                                                                                                                                                                                                                                                                               |

## restframework authenticate vs permission

authenticate class should contain code that validates the user otherwise raise exceptions.AuthenticationFailed('No such user')
.
permission class should contain code that check if the user is permitted to access something.

## Cross Site Request Forgery protection (CSRFP)

an attack where cookie is used by attacker site to make a request.
site A send request to site B with site B's cookie  
solution 1: for every request the server send a token(anti-forgery token) to client and client's next request should contain this token and then again server will send a new token and this will not be stored in cookie.  
solution 2: every session have 1 token( anti-forgery token). this will also not stored in cookie.  
solution 3: check the referer and accept.
both solution 2 and 3 can be stored in cookie if "double submit cookie" pattern" is used.

Response header:
Cross-Origin-Opener-Policy
Referrer-Policy

## Headers related to CORS

headers are not case sensitive but values are.
forbidden header, meaning that you cannot change it programatically
**Preflight Request must read.**  
So, Preflight request is a http request of method OPTION. it is sent by client to check what the server accept and then the client sends the main request if the response of OPTION match with the requirments of main Request. if requirements are not fullfiled then it will cause CORS error. Remember the server must response the OPTIONS request. in golang you have to manually write it.

- Response Headers
  1. Access-Control-Allow-Origin
  2. Access-Control-Allow-Credentials
  3. Access-Control-Allow-Headers
  4. Access-Control-Allow-Methods
  5. Access-Control-Expose-Headers
  6. Access-Control-Max-Age
  7. Timing-Allow-Origin
- Request Headers
  1. Access-Control-Request-Headers
  2. Access-Control-Request-Method
  3. Origin

REFERENCES:

1. https://www.kite.com/python/docs/django.middleware
2. https://cheatsheetseries.owasp.org/index.html
3. https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#cors
