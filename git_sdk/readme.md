## Maintaining Multiple Git Account
1. In `~/.ssh/config` file

```
   #hastetough
   Host github.com-hastetough
   HostName github.com
   User git
   AddKeysToAgent yes
   IdentityFile ~/.ssh/id_rsa_hastetough


   #bkalita-git
   Host bkalita
   HostName github.com
   User git
   AddKeysToAgent yes
   IdentityFile ~/.ssh/id_rsa_bkalita-git


   #bkalita-techvariable-git
   Host techvariable
   HostName github.com
   user git
   AddKeysToAgent yes
   IdentityFile ~/.ssh/id_rsa_bkalita-techvariable-git

   #ec2
   Host ec2
   HostName 13.233.39.180
   User ec2-user
   IdentityFile ~/.ssh/id_rsa_ec2
   Compression yes
```
2. Place Private Keyfile in `$HOME/.ssh/`

```
   /home/bipul/.ssh/id_rsa_bkalita-git #private key
   /home/bipul/.ssh/id_rsa_bkalita-git.pub #in github

   /home/bipul/.ssh/id_rsa_hastetough #private key
   /home/bipul/.ssh/id_rsa_hastetough.pub #in github
```

3. Now you'll be able to do `git clone git@{Host}:{Username}/{repository}`

## Git Terminology

### Who is the person going to push?

```
git config --global user.name
```

We can also change the fields globally (not for a particular directory) by below

```
$ git config --global user.name "Your Name Comes Here"
$ git config --global user.email you@yourdomain.example.com
```

And for a particular git repository change `--global` to `--local`

### Where files are going to be pushed?

`git init` inside a directory will make configuration files to communicate with the git server for example github.

### Take a Snapshot and store to the temporary Staging area called "index"

```
git add . #a snapshot containing all files(including recursive files) under the current directory
```

To undo the above step

```
git rm --cached filename
git rm -r --cached .
```

## How it works

if I `add` a file which had "hello" sentence and then I changed the content of the file to "hello world!" then git knows that this file is modified but it only has the snapshot of "hello" since we ran `add` command before `hellow world!` lines modification.

the snapshot is a 1 file which may contain information of many files

```
      #index hash      #index
index ce01362..a042389 100644

```

If we modify a file after `add` command we can see the changes using

```
git diff
```

and if we want to see what's in the staging

```
git diff --cached
```

we can restore to the last snapshot after even modified.

```
git restore filename
```

```
git show hash_of_object
```

###

```
branch object head
head is a reference to commit object
by default there is a Head in every repository called Master
'HEAD' is alias to 'current head'
git log --raw
git log --patch

```

### permanently store contents of "index" in the repository(locally)

```
git commit -m "...."
```

### Status

```
git status
```

3.

4. undo commit before push
   git reset HEAD~1
5. clone other's git repo which have accessed to the user "bkalita-techvariable-git"

```

git clone git@github.com-bkalita-techvariable-git:sivgos-tv/CDP-microservices.git

[bipul@archlinux CDP-microservices]$ git branch -a

- main
  remotes/origin/HEAD -> origin/main
  remotes/origin/celery-livy-service
  remotes/origin/main

[bipul@archlinux CDP-microservices]$ git checkout -b celery-livy-service
Switched to a new branch 'celery-livy-service'
[bipul@archlinux CDP-microservices]$ git branch -a

- celery-livy-service
  main
  remotes/origin/HEAD -> origin/main
  remotes/origin/celery-livy-service
  remotes/origin/main

# the \* represents current branch, and push command will always do with current branch by default.

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

7. Merge one branch to the current branch

```

git merge other_branch

```
## How To use CICD for a repository
* inside the repository create a `.yml` file like this `.github/workflows/build.yml`
* remember, each `.yml` file is a script which will get run on condition written inside that file. for example, an `.yml` can be written so that it gets executed each time when  something is pushed to that corresponding repository.
* Now, what gets executed? first of all content inside that script will get executed by the github runner. it's like a computing environment for those yml files. let's look at a simple such file - 
```
name: ict_health build

on:
  push:
    branches: [ dev_branch ]
    
  workflow_dispatch:


jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Set up Python 3.10.5
        uses: actions/setup-python@v2
        with:
          python-version: 3.10.5
      - name: Run Python commands
        run: |
          pip install --upgrade pip
          
          python3.10 -m venv venv-ictchat
          python3.10 -m venv venv-ictservice
          
          source venv-ictchat/bin/activate
          pip install -r ictchat/requirements.txt
          deactivate
          
          source venv-ictservice/bin/activate
          pip install -r ictservice/requirements.txt
          deactivate
```
Look at the above line `action/checkout@v2` what it does that it copies the repository to the action runner computing environment. 