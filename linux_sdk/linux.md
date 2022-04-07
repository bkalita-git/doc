# LINUX PACKAGE INSTALL
 1. BUILD FROM SOURCE
 2. APPIMAGE //DOUBLE CLICK
 3. PACKAGE MANAGER
 4. CONAINER INSTALL
# System Calls
> Access to the Unix kernel is only available through these system calls
https://man7.org/linux/man-pages/man2/intro.2.html

>strace to trace syscalls

>libc function “printf” is just a wrapper around syscall “write”

> always use glibc instead of calling sys from assembly

# POSIX function
	https://en.wikipedia.org/wiki/C_POSIX_library
	<sys/socket.h>
	<netdb.h> 
	..etc

# Sockets
>endpoint for communication with process  

transport layer
1. send and receive
2. bind and listen  
linux installation:

# CRON JOB
install crontab for arch linux

# display:
	$ ls -l folder  									    #list folder/file owner
	$ cat /etc/passwd 										#all local users in format account:password:UID:GID:GECOS:directory:shell
	$ cat /etc/group  										#all groups
	$ groups user_name 										#all groups of user_name
# change:
	$ sudo chown user_name_of_new_owner file/folder_loc		#change owner of the file/folder
# add:
	$ gpasswd -a user_name group_name 						# add to group
	$ gpasswd -d user_name group_name 						# delete from group
# create
	$ adduser -m user_name									#create new user
	$ passwd user_name										#create password for user_name user
	$ groupadd group_name 									#create new group
# Access levels and roles
for example i can't execute a database command from a normal user.  
also '$ passwd postgres' does not let me because I am a 'bipul' user and it does not have the permission. but I can do the same thing from root user.


# TOOLS
1. GTK
	>opensource cross platform for creating GUI. In one system different version of GTK can be installed.
	```
	configuration of a gtk
	/usr/share/gtk-2.0/gtkrc
	usr/share/gtk-3.0/settings.ini

	installed themes are
	~/.themes/Sweet-master/gtk-3.0
	~/.local/share/themes/gtk-3.0
	/usr/share/themes/

	css used by selected themes are at
	resource:///org/adapta-project/gtk-3.24-nokto-eta/
	```
	>if you want to install theme then search for gtk+3 theme. download it and place it in .themes/
2. KDE/GNOME/XFCE
   >desktop environments
   ```
   xfce4-about
   #version 4.16 uses GTKv3
   ```
3. xfwm4
   > The Xfce 4 Window Manager
   ```
   xfwm4-tweaks-settings
   ```

# REFERENCES
1. https://wiki.archlinux.org/title/users_and_groups
1. https://wiki.archlinux.org/title/ArchWiki:Access_levels_and_roles
