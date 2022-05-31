# LINUX PACKAGE INSTALL

1.  BUILD FROM SOURCE
2.  APPIMAGE //DOUBLE CLICK
3.  PACKAGE MANAGER
4.  CONAINER INSTALL

# System Calls

> Access to the Unix kernel is only available through these system calls
> https://man7.org/linux/man-pages/man2/intro.2.html

> strace to trace syscalls

> libc function “printf” is just a wrapper around syscall “write”

> always use glibc instead of calling sys from assembly

# POSIX function

    https://en.wikipedia.org/wiki/C_POSIX_library
    <sys/socket.h>
    <netdb.h>
    ..etc

# Sockets

> endpoint for communication with process

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

   > opensource cross platform for creating GUI. In one system different version of GTK can be installed.

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

   > if you want to install theme then search for gtk+3 theme. download it and place it in .themes/

2. KDE/GNOME/XFCE
   > desktop environments
   ```
   xfce4-about
   #version 4.16 uses GTKv3
   ```
3. xfwm4

   > The Xfce 4 Window Manager

   ```
   xfwm4-tweaks-settings

   ```

# Virtual Address...

`cpu`->`MMU(Virtual Address)`->`Physical_Address`

- SIZE OF VIRTUAL ADDRESS and Physical Address of a processor (Not relates with Process Virtual Address)
  ```
  cat /proc/cpuinfo
  ```
- The Page Size which was defined at the compiled time of the kernel

  ```
  #outputs in bytes, normally 4KB
  getconf PAGE_SIZE
  ```

- **For each process there exist its own Page Table**. To see the mapping use below. It will show the Virtual addresses ie pages are mapped to what thing. Each range is divisible by the PAGE_SIZE, since the smalled range is equal to PAGE_SIZE.
  ```
  pmap -Xp $PID
  ```
  So, as we can see, stack, heap everything inside the page. SO, even the process allocates a memory for an integer variable and it lives in stack but ultimately it lives in page. and we all know that page can be swapped from RAM to HDD.
  It's all done by `mmap` system call.
- mmap system call maps virtual address space of range PAGE_SIZE to physical address space which can be device or RAM. Now since virtual addresses is only work for process, so mmap must be call from a process.
- SO, Virtual Address space (say 64-bit) is divided into 3 parts Below. Now some of the virtual address space can denote Physical RAM and some can denote Physical HW Devices. So, it does not mean that Userspace means RAM only, for a process Userspace address can denote a HW device too.
  - Kernel Virtual Address Space (umalloc)
  - Kernel Logical Address Space (kmalloc)
    - Here contiguous virtual address space is mapped with contiguous physical RAM Addresses. Good for Direct Memory Transfer.
  - User Space (mmap, brk, sbrk)
- RAM is divided into page frames, say our PAGE_SIZE is 4KB. Now, let's say a file of size 12KB is loaded into RAM. That means, The RAM contains 3 PAGE_FRAMEs for that file. And the process who loads that file into RAM has a PAGE_TABLE and in that table the mapping of virtual addresses for the 3 PAGE_FRAMES are there which was done by `mmap` system call. if we say that each virtual address represents one Byte then there will be 12288 Virtual addresses againts the 2 PAGE_FRAMES. this info can be seen as described above by `pmap` command. Now this scenario is for a whole file in RAM. Now, sometimes or most of the times at first, even there are virtual addresses for a file in the PAGE_TABLE but the file is not in RAM. Now when you `read` system call in that virtual address then page fault occurs and then loads 1 PAGE_FRAME againts the file which contains the required bytes which is done by mmap underlying. (REMEMBER 1 PAGE_FRAME, so must be a portion of that file). Now, these all are for a file. what about `var x int`? Well this is also stored in RAM inside a PAGE_FRAME and mapped through Process Virtual Address, this type of PAGE_FRAME may call STACK or HEAP. So, Everything is page when comes to RAM and Process.

## File descriptor

a process has - Own page Table - Own file descriptor table
Open File table which is global. Now if a process using a file, For that there is a openfile table. so, Fd is linked with openfile.
V-Node table which is linked with the opentable. which contains actual description of the file where it lives.

## Linux command to check openfile

ls -l /proc/273860/fd
remember, always run the compiled program.
pstree -p | grep 275939

## Page

getconf PAGESIZE #result in bytes
sudo ls -l /proc/372569/map_files/

# REFERENCES

1. https://wiki.archlinux.org/title/users_and_groups
1. https://wiki.archlinux.org/title/ArchWiki:Access_levels_and_roles
