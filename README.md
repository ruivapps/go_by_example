# go\_by\_example

## server.py

small python http server to test code. do not use this to test http\_concurrent\_get.go as the server do not support concurrent 

this server.py is used to test http\_get.go and http\_post.go

~~~bash
pip install -r requirements.txt
python server.py
~~~

## http_get.go

send HTTP GET request. Print responsed content back to stdout

~~~bash
go run http_get.go http://127.0.0.1:8080/fakeapi
~~~

compile and run the binary

~~~bash
go build http_get.go
./http_get http://127.0.0.1:8080/fakeapi
~~~

## http_post.go

send HTTP POST request. Print response content back to stdout

~~~bash
go run http_post.go http://127.0.0.1:8080/fakeapi
~~~

## http_concurrent_get.go

send concurrent HTTP GET request
print out sequence number, len of result and time spent 

~~~bash
go run http_concurrent_get.go http://google.com http://google.com http://google.com http://google.com http://google.com
~~~

## list_files.go

list files in current directory orderedby size

~~~bash
go run list_files.go
~~~

## ssh.go

connect to remote server by ssh and print remote stdout back to local stdout

~~~bash
go run ssh.go 10.1.232.21 df -h
~~~
