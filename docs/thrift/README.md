Tutorial
========

1) First things first, you'll need to install the Thrift compiler and the
   language libraries. Do that using the instructions in the top level
   README.md file.

2) Read tutorial.thrift to learn about the syntax of a Thrift file

3) Compile the code for the language of your choice:
```shell
# Go to root path
# Edit file ./thrift/sample.thrift
$ thrift -r -out thrift --gen go ./thrift/sample.thrift
# $ thrift -r --gen cpp sample.thrift
```
4) Take a look at the generated code.

5) Look in the language directories for sample client/server code.

6) That's about it for now. This tutorial is intentionally brief. It should be
   just enough to get you started and ready to build your own project.


Docs
========
http://github.com/apache/thrift/tutorial
