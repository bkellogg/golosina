# Getting Started

Here I'll try to explain all the elements that will need to be configured to start working with the framework.

## Installation

The most important part of the installation of the framework is to have all the `GOPATH` and `GOROOT` defined. Usually I don't think too much on what I have to do here, I just follow this awesome guide.

> https://ahmadawais.com/install-go-lang-on-macos-with-homebrew/

After you're ready with that, move to your project folder and simpli clone the repo. As simple as running:

```bash
cd /path/to/your/project
git clone git@github.com:eaperezc/golosina.git
```

With that you'll have the starting project ready to be modified.

## Configuration

Since I wanted to mimic closely what Laravel is doing you can find a `.env.example` file in the project. Create your own as a copy of it and update the values inside to match you're desired configuration.

```bash
cp .env.example .env
```

## Directory Structure

The project folders are a little bit different from the one we are used to but not that much. This was designed like this because the language requires it to be as such, but don't worry it's still a very common structure:

* controllers
    * A controller per resource (usually) or any independent controllers you will need in your app will live here. The naming convension we defined for these files is `EntityController.go` (e.g UserController.go)
* models
    * All application models and db logic. Those models are basically gorm models and they are named with upper camelCasing like `User.go` or `Post.go`
* routes
    * api.go: has the definition of our API routes
    * web.go: has the definition of our web pages routes
* views
    * The definition of our view code will live here. This code will contain all the parameters for our views and also the structs that we will send as data to the templates
        * templates
            * Here we will store all of our view templates
* framework
    * Internal framework package with wrappers to other libraries to make your life easy. Most of the cases you don't need to worry what happens here but maybe you'll want to customize the internal structure to fit your needs
* docs
    * The folder that contains this documentation. Built with [Docsify](https://docsify.js.org/#/quickstart)
* server.go
    * This is the `main` file for the project. We will run `go run server.go` to start our app

!> You'll not see the `.env` file listed but make sure you create it and have all the configurations up to date
