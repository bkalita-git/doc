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