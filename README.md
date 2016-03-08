# go-express

Web frame created by Golang. It is the same structure as express in node.js. It is for easy to setup Web by Golang. It is a MVC web frame.   

# Why should you want this web frame?

- Jquery and Bootstrap are integrated into the go-express. But you can remove them from the go-express if you donot want to use them.
- Go-express is a light web frame for you to create a site. It have the easy-use and simple and powerful routes(controller).

# Usage:

## Dependencies:
- Golanghttp://golang.org/
- jQueryhttp://jquery.com/
- Twitter Bootstraphttp://twitter.github.com/bootstrap/
- google-code-prettifyhttps://github.com/google/code-prettify

## Install and run:

    mkdir ~/go_codes/src
    cd ~/go_codes/src
    git clone https://github.com/DoBetter-pan/go-express.git
    cd go-express
    source setenv.sh
    go build
    ./go-express -host=127.0.0.1 -port=9898 
Then you can visit http://127.0.0.1:9898/express/index in brower.

Note:
You must setup your golang environment first and will change your GOPATH when executing "source setenv.sh".
I tested it in Ubuntu. If you are using other OS, it is the same way as in Ubuntu. Please try.

## Develop:

I wrote a inline example: express. If you want to write a new controller, you can refer to these files.

    controller/expressController.go
    views/express/index.html
    views/express/about.html 
If you want to change the main template, you can change the following files.

    controller/mainController.go
    views/main/main.html
    views/main/header.html
    views/main/content.html
    views/main/footer.html 
    
# What should we do next?

Model is just a directory, not any file in it. It is not completely now. Go-express will support sql and no-sql database as data source.

# Support and contact

If you have any question and advice, you can cantact me in QQ Group: 536069420.
