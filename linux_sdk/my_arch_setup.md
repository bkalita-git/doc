## running services
```
sudo systemctl list-units --state running
sudo systemctl list-unit-files --state enabled 
```
## Disabled services
```
sudo systemctl stop bluetooth bumblebeed httpd mariadb postgresql cups docker
sudo systemctl disable bluetooth bumblebeed httpd mariadb postgresql cups docker
```

## optional dependency and dependency
every package have optional dependency and mandatory dependency. when installing a package using ```sudo pacman -S package_name``` it installs all the mandatory dependency and later it shows all the optional dependencies. To view the optional and madatory dependency do following. go to pacman https://wiki.archlinux.org/title/Pacman/Rosetta#Querying_package_dependencies for performing more operations.
```
sudo pacman -Ss search_package_name
sudo pacman -Q #to list all installed package
pacman -Qdtq | sudo pacman -Rs - # to list all packages which have no dependencies
sudo pacman -Sii package_name #information of a package
```
## disk usage
```
sudo du -sh /directory/*
```
## wifi
sometimes broadcom-wl or wpa_supplicant update leads to no wifi. the current working version of wpa_supplicant and broadcom-wl is
```
sudo pacman -U wpa_supplicant-2\:2.9-8-x86_64.pkg.tar.zst
sudo pacman -U broadcom-wl-6.30.223.271-357-x86_64.pkg.tar.zst
```

glibc:
    linux-api-headers:
    tzdata
    filesystem:
        iana-etc

linux:
    coreutils:
        glibc
        acl:
            attr
            libattr.so
        attr:
            glibc
        gmp
        libcap:
            pam:
                glibc
                libtirpc
                pambase
                audit
                libaudit.so
                libxcrypt
                libxcrypt.so
        openssl:
            glibc
    kmod
        glibc
        zlib
        openssl
        xz
        zstd
    initramfs/mkinitcpio:
        awk
        kmod
        util-linux
        libarchive
        coreutils
        bash
        binutils
        diffutils
        findutils
        grep
        filesysem
        zstd
        systemd

## sym link
```
sudo ln -s firefox-developer-edition /usr/bin/google-chrome-stable
```

## desktop entries
/usr/share/applications/
/usr/local/share/applications/
/home/bipul/.local/share/applications/



## bash terminal color profile
nano /etc/bash.bashrc

PS1='\[\033[01;32m\][\u@\h \W]\[\033[00m\]\$ '

if \e[5m is on then \e[25m is off for that by adding the 2
\e[34m for foreground text color and 4 is the value, and \e[37 where 7 is the white value
we can setup different color schemes for different command output.
use trap command
for example bash to see all color it supports
```
for C in {0..255}; do
    tput setab $C
    echo -n "$C "
done
```
finally this is powerful
```
PS1='\e[23m\e[37m\e[1m\e[36m\e[7m\u@\h \e[5m¦\e[25m\e[27m \e[37m\e[1m\n\e[3m\e[31m' #at first unset whatever set at the end of PS1
trap 'echo -ne "\e[23m\e[37m"' DEBUG #unset whatever set at the end of PS1
```
## What are terminal and terminal emulator and what are shells

`shell` are command line interpreter. using shell command user communicates with kernel. eg. bash, cshell, kshell, sh.
```
#!/bin/bash
#!/bin/sh
```
`xterm` is the terminal emulator for `X windows system`. A terminal-emulator/pty is the GUI way to write shell command.
`bashrc` file is a bash script and get excuted when `bash` command is run.

Terminal is a real keyboard and communicates with a computer.
Terminal emulator(TTY) emulates terminal: example - xterm,uxterm,xfce-terminal
Pseudo Terminal
## tmux
in ~/.profile
[ -z "$DISPLAY" ] && [ $XDG_VTNR -eq 1 ] && exec startx 
then add into /etc/bash.bashrc  at first
[[ $TERM != "screen" ]] && exec tmux


