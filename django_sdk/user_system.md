## User
The user table is_staff column value to True will allow to use admin system. Models should be registered in admin for access/write.
```
from django.contrib import admin
from .models import User

#registering table with admin
admin.site.register(User)

#custom admin representations
class CustomAdmin(admin.ModelAdmin):
    ...
admin.site.register(CustomAdmin)
``` 
TABLES RELATED WITH User system are - 

> **auth_group**, **auth_group_permissions**, **auth_permission**, **auth_user**, **auth_user_groups**, **auth_user_user_permissions** are related with **django.contrib.auth** app


> **django_admin_log** related with **django.contrib.admin**

> **django_content_type** related with **django.contrib.contenttypes**

> **django_sessions** related with **django.contrib.sessions**

> **authtoken_token** related with **rest_framework.authtoken**

> django_migrations

> sqlite_sequence