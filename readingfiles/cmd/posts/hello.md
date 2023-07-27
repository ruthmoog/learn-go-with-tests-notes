Title: Learn Go With Tests
Description: Learn Go and TDD at the same time
Tags: tdd, go
---
In this chapter we're going to learn how to read some files, get some data out of them, and do something useful.

Pretend you're working with your friend to create some blog software. The idea is an author will write their posts in markdown, with some metadata at the top of the file. On startup, the web server will read a folder to create some Posts, and then a separate NewHandler function will use those Posts as a datasource for the blog's webserver.

We've been asked to create the package that converts a given folder of blog post files into a collection of Posts.