## Overview
We will create a REST API that allows us to CREATE, READ, UPDATE and DELETE the articles on our website. When we talk about CRUD APIs we are referring to an API that can handle all of these tasks: Creating, Reading, Updating and Deleting. If you are writing any form of web application, then you are most likely interfacing with 1 or more REST APIs in order to populate the dynamic parts of your application and to perform tasks such as updating or deleting data within a database.

In this tutorial, you are going to be building a fully-fledged REST API that exposes GET, POST, DELETE and PUT endpoints that will subsequently allow you to perform the full range of CRUD operations. 

In order to keep this simple and focus on the basic concepts, we won’t be interacting with any backend database technologies to store the articles that we’ll be playing with. However, we will be writing this REST API in such a way that it will be easy to update the functions we will be defining so that they make subsequent calls to a database to perform any necessary CRUD operations.

### Installation
1. **Make a new project directory**
   
   ```
   $ mkdir articles-api
   ```

2. **Clone the repo**

   ```
   $ cd articles-api
   $ git clone git@github.com:ianlimle/articles-api-in-go.git 
   ```

3. **Run the API**

   ```
   $ go get github.com/gorilla/mux
   $ go run main.go
   ```

4. **Call HTTP Endpoints**
   
   Open a new terminal, change the current working directory to 'articles-api'

   GET all articles:
   ```
   $ curl localhost:10000/all                          
   ```
   
   POST new article with {id}:
   ```
   $ curl -X POST -H 'Content-Type: application/json' -d '{"Id":"{id}","New Title":"New Title","desc":"New description","content":"New content"}' localhost:10000/article      
   ```
   
   PUT (update) article with {id}:
   ```
   $ curl -X PUT -H 'Content-Type: application/json' -d '{"Id":"{id}","Title":"Edited Title","desc":"Edited description","content":"Edited content"}' localhost:10000/article/{id}
   ```
   
   GET single article with {id}:
   ```
   $ curl localhost:10000/article/{id}      
   ```
   
   DELETE article with {id}:
   ```
   $ curl -X DELETE localhost:10000/article/{id}      
   ```

   


