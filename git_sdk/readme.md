1. $HOME/.ssh/config
   #hastetough
   Host github.com-hastetough
   HostName github.com
   User git
   AddKeysToAgent yes
   IdentityFile ~/.ssh/id_rsa_hastetough

   #bkalita-git
   Host github.com-bkalita-git
   HostName github.com
   User git
   AddKeysToAgent yes
   IdentityFile ~/.ssh/id_rsa_bkalita-git

2. ls $HOME/.ssh/
   /home/bipul/.ssh/id_rsa_bkalita-git ->private key
   /home/bipul/.ssh/id_rsa_bkalita-git.pub ->in github

   /home/bipul/.ssh/id_rsa_hastetough ->private key
   /home/bipul/.ssh/id_rsa_hastetough.pub ->in github

3. undo commit before push
   git reset HEAD~1
4. clone other's git repo which have accessed to the user "bkalita-techvariable-git"

```
git clone git@github.com-bkalita-techvariable-git:sivgos-tv/CDP-microservices.git

[bipul@archlinux CDP-microservices]$ git branch -a
* main
  remotes/origin/HEAD -> origin/main
  remotes/origin/celery-livy-service
  remotes/origin/main


[bipul@archlinux CDP-microservices]$ git checkout -b celery-livy-service
Switched to a new branch 'celery-livy-service'
[bipul@archlinux CDP-microservices]$ git branch -a
* celery-livy-service
  main
  remotes/origin/HEAD -> origin/main
  remotes/origin/celery-livy-service
  remotes/origin/main


# the * represents current branch, and push command will always do with current branch by default.

[bipul@archlinux CDP-microservices]$ git push
fatal: The current branch celery-livy-service has no upstream branch.
To push the current branch and set the remote as upstream, use

    git push --set-upstream origin celery-livy-service

once done then next can do git push
```

5. connect a directory to a git repo

```
$ cd directory
$ git init .
$ git config --global user.name "Bipul Kalita"
$ git config --global user.email "bipul@techvariable.com"
$ git remote add origin sshhost:username/reponame
$ git remote -v #tocheck existing remote url
$ git pull origin branch_name
$ git branch -a
$ git checkout -b branch_name
$ git add -A
$ git commit --allow-empty-message -m ""
$ git push --set-upstream origin branch_name

$ git submodule add techvariable:techvariable/go-fileupload <destination_folder>

```

6. to change a repo url so it can push to this

```
git remote set-url origin bkalita:username/reponame
```
