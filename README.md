Game-a-holic is an End to End "Gamification for work" portal. Intention is to use gaming features that makes the player engaging and trying to apply it to work. This is an alpha version with very basic scoring system. Features to enhance it further is under development.

Mission is to provide an end to end platform for teams to include this as part of this work.

Current version is aiming, software development team members to gain scores, levels (badges) and avatars based on their daily activity like Pull Requests, Ticket closing etc to be part of the game. Application works based on webhooks for most of the integrations and so doesnt require anything from developers. Just enable webhooks for supported integrations and application will take care of the rest.

This can be made generic for any kind of teams with right triggers but needs to be discussed as per usecases.

Following Integrations are currently supported

1. Bitbucket Server
   1. Pull Request Merged
   2. Pull Request Approved
2. Jira
   1. Generic for any activity

Please details on integrations below

#### Table of Contents

1. [Build](#build)
   1. [Prerequisites](#prerequisites)
   2. [Build And Run Command](#build-and-run-commands)
2. [Navigating the repository](#navigating-the-repository)
4. [Extending gameaholic with Modules](#extending-gameaholic-with-modules)
5. [Documentation](#documentation)
   1. [Integrations](#integrations)
   2. [Developer guides](#developer-guides)
   3. [Wiki](#wiki)
   4. [Website](#website)
7. [Issues](#issues)
8. [Community](#community)
9. [Support](#support)
10. [License](#license)


## Build

### Prerequisites

#### Golang
Currently gameaholic is tested to work with golang version.
Refer golang official [documentation](https://golang.org/doc/install) for installation.

#### Buffalo

As gameaholic was developed in 72 hrs during a hackathon we used rapid application development toolchain called Buffalo.

To install buffalo refer buffalo official [documentation](https://gobuffalo.io/en/docs/getting-started/installation)

>[!WARNING]
>For windows 10, sometimes GCC needs to be installed separately refer [here](https://blog.gobuffalo.io/install-buffalo-on-windows-10-e08b3aa304a3)


#### Git

Install the version control tool [git](https://git-scm.com/) and clone this repository with

```bash
git clone https://github.com/bribebybytes/gameaholic.git
```

#### mysql

gameaholic ships with a docker-compose file, which can download supported mysql version and run it. But you can also use your mysql database. gameaholic is tested with mysql 8.0.19.

Configure database.yml or setup following environment variables to override defaults

1. **`DATABASE_HOST`** - defaults to `localhost`
2. **`DATABASE_PORT`**- defaults to `3306`
3. **`DATABASE_USER`** - defaults to `root`
4. **`DATABASE_PASSWORD`** - defaults to `root`

{{box op="start" cssClass="boxed noteBox"}}
> Please use defaults only for development setup
{{box op="end"}}

{{box op="start" cssClass="boxed noteBox"}}
> Setup your own database backup restore solution. We suggest taking mysql dumps and restoring periodically.
{{box op="end"}}

##### Default Setup
Bootstrap setup can be invoked using buffalo pop. Ensure mysql is running and database.yml or environment variables are setup correctly. 

Either this can be setup manually or using buffalo pop

###### Manual setup
Use mysqldump or run contents of files under Migration directory with file name like *.create_bootstrap.mysql.up.sql



```bash
# Replace <<password>> with your root password
docker exec -it gameaholic_mysql sh -c  "mysql -p<<password>> -e 'create database gameaholic;'"
docker cp .\migrations\*_create_bootstraps.mysql.up.sql  gameaholic_mysql:\
docker exec -it gameaholic_mysql sh -c "mysql -p<<password>> gameaholic < *_create_bootstraps.mysql.up.sql" 
```

###### Using Buffalo Pop (Soda)
1. Ensure mysqldump is available 
2. Ensure database.yml or database environment variables are configured correctly 



{{box op="start" cssClass="boxed noteBox"}}
> Even if using docker-compose up for mysql. We need to have [mysqldump](https://dev.mysql.com/downloads/installer/) available locally to run migration steps if you want to use buffalo pop [soda wrapper].
{{box op="end"}}

### Build and Run Commands

After you have taken care of the [Prerequisites](#prerequisites)

Execute the following

For quick development run

```bash
cd gameaholic
buffalo dev
```
For generating go binary

```bash
cd gameaholic
buffalo build
cd bin
./game-a-holic #for windows run game-a-holic.exe generated
```
Please create github issues if you face issues during build.



## Navigating the repository

As project uses buffalo as framework, we follow same [folder structure](https://gobuffalo.io/en/docs/getting-started/directory-structure) for now. Unless we get a better recommendations.


## Extending gameaholic with Modules

TODO

## Documentation

### Integrations

TODO


### Contributing

If you want to contribute please [github pull requests](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/about-pull-requests)


### Code Reviews

Get started with re-:eyes: pull requests!

You might not have the time to develop yourself but enough experience with reviewing code, your help on code reviews will be much
appreciated!

### Translation

TODO

## Feature Requests and Issues

Pls use github issues to raise any issues or feature requests

## Community

[![BribeByByte Slack](https://img.shields.io/badge/bribebybytes-Slack-%26blue)](https://join.slack.com/t/bribebybytes/shared_invite/enQtOTU0NTAzMjI2MjczLWVjYzZlYzFhZWYxNDUxZDcwNmQxM2FlY2M3OThiODZiNDRiYmFlN2FjN2FjYWZiYzU5NjkxMzE1ZDk0M2RjMTk)
