## Interaction 0: GET/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos the pre-requisites the api root responds to a GET (i.e. the server is up and accessible, CORS headers are set up)
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 842
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:31 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200: text/html;charset=utf-8):

```
[{"uid":"cd133afa-6fc5-4b9c-9a41-28c6ad2620cf","title":"blah","order":523,"completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/cd133afa-6fc5-4b9c-9a41-28c6ad2620cf","url":"https://http4k-todo-backend.herokuapp.com/todos/cd133afa-6fc5-4b9c-9a41-28c6ad2620cf"},{"uid":"ebce16d3-ad6d-4d30-ac50-68b844808ae9","title":"blah","order":95,"completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/ebce16d3-ad6d-4d30-ac50-68b844808ae9","url":"https://http4k-todo-backend.herokuapp.com/todos/ebce16d3-ad6d-4d30-ac50-68b844808ae9"},{"uid":"e3594c19-60fa-435e-851d-7025bec510b1","title":"blah","order":95,"completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/e3594c19-60fa-435e-851d-7025bec510b1","url":"https://http4k-todo-backend.herokuapp.com/todos/e3594c19-60fa-435e-851d-7025bec510b1"}]
```

## Interaction 1: POST/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 18
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos the pre-requisites the api root responds to a POST with the todo which was posted to it
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```
{"title":"a todo"}
```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 270
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:32 GMT
Location: https://http4k-todo-backend.herokuapp.com/todos/bea9bc5a-53d3-4206-ba79-0743ce8053cd
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201: text/html;charset=utf-8):

```
{"title":"a todo","uid":"bea9bc5a-53d3-4206-ba79-0743ce8053cd","completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/bea9bc5a-53d3-4206-ba79-0743ce8053cd","url":"https://http4k-todo-backend.herokuapp.com/todos/bea9bc5a-53d3-4206-ba79-0743ce8053cd"}
```

## Interaction 2: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos the pre-requisites the api root responds successfully to a DELETE
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 0
Date: Sat, 19 Sep 2020 16:29:33 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204: ):

```

```

## Interaction 3: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos the pre-requisites after a DELETE the api root responds to a GET with a JSON representation of an empty array
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 0
Date: Sat, 19 Sep 2020 16:29:34 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204: ):

```

```

## Interaction 4: GET/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos the pre-requisites after a DELETE the api root responds to a GET with a JSON representation of an empty array
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 2
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:34 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200: text/html;charset=utf-8):

```
[]
```

## Interaction 5: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at http://http4k-todo-backend.herokuapp.com/todos the pre-requisites after a DELETE the api root responds to a GET with a JSON representation of an empty array
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 0
Date: Sat, 19 Sep 2020 16:29:35 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204: ):

```

```

## Interaction 6: POST/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 24
Content-Type: application/json
Fulltitle: Todo-Backend API residing at http://http4k-todo-backend.herokuapp.com/todos storing new todos by posting to the root url adds a new todo to the list of todos at the root url
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```
{"title":"walk the dog"}
```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 276
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:35 GMT
Location: https://http4k-todo-backend.herokuapp.com/todos/66f99e53-0e3d-4a42-b94c-c5b6a5485ca0
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201: text/html;charset=utf-8):

```
{"title":"walk the dog","uid":"66f99e53-0e3d-4a42-b94c-c5b6a5485ca0","completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/66f99e53-0e3d-4a42-b94c-c5b6a5485ca0","url":"https://http4k-todo-backend.herokuapp.com/todos/66f99e53-0e3d-4a42-b94c-c5b6a5485ca0"}
```

## Interaction 7: GET/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos storing new todos by posting to the root url adds a new todo to the list of todos at the root url
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 288
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:35 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200: text/html;charset=utf-8):

```
[{"uid":"66f99e53-0e3d-4a42-b94c-c5b6a5485ca0","title":"walk the dog","order":0,"completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/66f99e53-0e3d-4a42-b94c-c5b6a5485ca0","url":"https://http4k-todo-backend.herokuapp.com/todos/66f99e53-0e3d-4a42-b94c-c5b6a5485ca0"}]
```

## Interaction 8: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos storing new todos by posting to the root url adds a new todo to the list of todos at the root url
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 0
Date: Sat, 19 Sep 2020 16:29:37 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204: ):

```

```

## Interaction 9: POST/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 16
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos storing new todos by posting to the root url sets up a new todo as initially not completed
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```
{"title":"blah"}
```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 268
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:37 GMT
Location: https://http4k-todo-backend.herokuapp.com/todos/b50f4c92-8895-4171-86fe-196c0823bc2f
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201: text/html;charset=utf-8):

```
{"title":"blah","uid":"b50f4c92-8895-4171-86fe-196c0823bc2f","completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/b50f4c92-8895-4171-86fe-196c0823bc2f","url":"https://http4k-todo-backend.herokuapp.com/todos/b50f4c92-8895-4171-86fe-196c0823bc2f"}
```

## Interaction 10: GET/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos storing new todos by posting to the root url sets up a new todo as initially not completed
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 280
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:37 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200: text/html;charset=utf-8):

```
[{"uid":"b50f4c92-8895-4171-86fe-196c0823bc2f","title":"blah","order":0,"completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/b50f4c92-8895-4171-86fe-196c0823bc2f","url":"https://http4k-todo-backend.herokuapp.com/todos/b50f4c92-8895-4171-86fe-196c0823bc2f"}]
```

## Interaction 11: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos storing new todos by posting to the root url sets up a new todo as initially not completed
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 0
Date: Sat, 19 Sep 2020 16:29:38 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204: ):

```

```

## Interaction 12: POST/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 16
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos storing new todos by posting to the root url each new todo has a url
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```
{"title":"blah"}
```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 268
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:38 GMT
Location: https://http4k-todo-backend.herokuapp.com/todos/3dfa0701-b948-4e04-b8b2-4ce298313de1
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201: text/html;charset=utf-8):

```
{"title":"blah","uid":"3dfa0701-b948-4e04-b8b2-4ce298313de1","completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/3dfa0701-b948-4e04-b8b2-4ce298313de1","url":"https://http4k-todo-backend.herokuapp.com/todos/3dfa0701-b948-4e04-b8b2-4ce298313de1"}
```

## Interaction 13: GET/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos storing new todos by posting to the root url each new todo has a url
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 280
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:38 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200: text/html;charset=utf-8):

```
[{"uid":"3dfa0701-b948-4e04-b8b2-4ce298313de1","title":"blah","order":0,"completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/3dfa0701-b948-4e04-b8b2-4ce298313de1","url":"https://http4k-todo-backend.herokuapp.com/todos/3dfa0701-b948-4e04-b8b2-4ce298313de1"}]
```

## Interaction 14: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos storing new todos by posting to the root url each new todo has a url
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 0
Date: Sat, 19 Sep 2020 16:29:39 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204: ):

```

```

## Interaction 15: POST/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 19
Content-Type: application/json
Fulltitle: Todo-Backend API residing at http://http4k-todo-backend.herokuapp.com/todos storing new todos by posting to the root url each new todo has a url, which returns a todo
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```
{"title":"my todo"}
```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 271
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:39 GMT
Location: https://http4k-todo-backend.herokuapp.com/todos/d4869de4-8479-4c04-8ad7-13d2620fc127
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201: text/html;charset=utf-8):

```
{"title":"my todo","uid":"d4869de4-8479-4c04-8ad7-13d2620fc127","completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/d4869de4-8479-4c04-8ad7-13d2620fc127","url":"https://http4k-todo-backend.herokuapp.com/todos/d4869de4-8479-4c04-8ad7-13d2620fc127"}
```

## Interaction 16: GET/todos/d4869de4-8479-4c04-8ad7-13d2620fc127

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos storing new todos by posting to the root url each new todo has a url, which returns a todo
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 281
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:39 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200: text/html;charset=utf-8):

```
{"uid":"d4869de4-8479-4c04-8ad7-13d2620fc127","title":"my todo","order":0,"completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/d4869de4-8479-4c04-8ad7-13d2620fc127","url":"https://http4k-todo-backend.herokuapp.com/todos/d4869de4-8479-4c04-8ad7-13d2620fc127"}
```

## Interaction 17: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos storing new todos by posting to the root url each new todo has a url, which returns a todo
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 0
Date: Sat, 19 Sep 2020 16:29:40 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204: ):

```

```

## Interaction 18: POST/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 26
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos working with an existing todo can navigate from a list of todos to an individual todo via urls
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```
{"title":"todo the first"}
```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 278
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:40 GMT
Location: https://http4k-todo-backend.herokuapp.com/todos/2a10e28c-89cc-4691-afa7-5ae4b3764a9c
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201: text/html;charset=utf-8):

```
{"title":"todo the first","uid":"2a10e28c-89cc-4691-afa7-5ae4b3764a9c","completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/2a10e28c-89cc-4691-afa7-5ae4b3764a9c","url":"https://http4k-todo-backend.herokuapp.com/todos/2a10e28c-89cc-4691-afa7-5ae4b3764a9c"}
```

## Interaction 19: POST/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 27
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos working with an existing todo can navigate from a list of todos to an individual todo via urls
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```
{"title":"todo the second"}
```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 279
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:40 GMT
Location: https://http4k-todo-backend.herokuapp.com/todos/ef550549-173b-45a3-81ea-37c55f7e18b3
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201: text/html;charset=utf-8):

```
{"title":"todo the second","uid":"ef550549-173b-45a3-81ea-37c55f7e18b3","completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/ef550549-173b-45a3-81ea-37c55f7e18b3","url":"https://http4k-todo-backend.herokuapp.com/todos/ef550549-173b-45a3-81ea-37c55f7e18b3"}
```

## Interaction 20: GET/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos working with an existing todo can navigate from a list of todos to an individual todo via urls
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 580
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:41 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200: text/html;charset=utf-8):

```
[{"uid":"2a10e28c-89cc-4691-afa7-5ae4b3764a9c","title":"todo the first","order":0,"completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/2a10e28c-89cc-4691-afa7-5ae4b3764a9c","url":"https://http4k-todo-backend.herokuapp.com/todos/2a10e28c-89cc-4691-afa7-5ae4b3764a9c"},{"uid":"ef550549-173b-45a3-81ea-37c55f7e18b3","title":"todo the second","order":0,"completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/ef550549-173b-45a3-81ea-37c55f7e18b3","url":"https://http4k-todo-backend.herokuapp.com/todos/ef550549-173b-45a3-81ea-37c55f7e18b3"}]
```

## Interaction 21: GET/todos/2a10e28c-89cc-4691-afa7-5ae4b3764a9c

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos working with an existing todo can navigate from a list of todos to an individual todo via urls
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 288
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:41 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200: text/html;charset=utf-8):

```
{"uid":"2a10e28c-89cc-4691-afa7-5ae4b3764a9c","title":"todo the first","order":0,"completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/2a10e28c-89cc-4691-afa7-5ae4b3764a9c","url":"https://http4k-todo-backend.herokuapp.com/todos/2a10e28c-89cc-4691-afa7-5ae4b3764a9c"}
```

## Interaction 22: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos working with an existing todo can navigate from a list of todos to an individual todo via urls
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 0
Date: Sat, 19 Sep 2020 16:29:41 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204: ):

```

```

## Interaction 23: POST/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 25
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos working with an existing todo can change the todo's title by PATCHing to the todo's url
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```
{"title":"initial title"}
```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 277
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:42 GMT
Location: https://http4k-todo-backend.herokuapp.com/todos/2445817e-4c7f-46ce-bc84-6bbb6fb4a680
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201: text/html;charset=utf-8):

```
{"title":"initial title","uid":"2445817e-4c7f-46ce-bc84-6bbb6fb4a680","completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/2445817e-4c7f-46ce-bc84-6bbb6fb4a680","url":"https://http4k-todo-backend.herokuapp.com/todos/2445817e-4c7f-46ce-bc84-6bbb6fb4a680"}
```

## Interaction 24: PATCH/todos/2445817e-4c7f-46ce-bc84-6bbb6fb4a680

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 25
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos working with an existing todo can change the todo's title by PATCHing to the todo's url
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```
{"title":"bathe the cat"}
```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 287
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:42 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200: text/html;charset=utf-8):

```
{"uid":"2445817e-4c7f-46ce-bc84-6bbb6fb4a680","title":"bathe the cat","order":0,"completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/2445817e-4c7f-46ce-bc84-6bbb6fb4a680","url":"https://http4k-todo-backend.herokuapp.com/todos/2445817e-4c7f-46ce-bc84-6bbb6fb4a680"}
```

## Interaction 25: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos working with an existing todo can change the todo's title by PATCHing to the todo's url
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 0
Date: Sat, 19 Sep 2020 16:29:42 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204: ):

```

```

## Interaction 26: POST/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 16
Content-Type: application/json
Fulltitle: Todo-Backend API residing at http://http4k-todo-backend.herokuapp.com/todos working with an existing todo can change the todo's completedness by PATCHing to the todo's url
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```
{"title":"blah"}
```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 268
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:43 GMT
Location: https://http4k-todo-backend.herokuapp.com/todos/de562f6a-3902-4f17-9b11-d4fe596a2fbf
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201: text/html;charset=utf-8):

```
{"title":"blah","uid":"de562f6a-3902-4f17-9b11-d4fe596a2fbf","completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/de562f6a-3902-4f17-9b11-d4fe596a2fbf","url":"https://http4k-todo-backend.herokuapp.com/todos/de562f6a-3902-4f17-9b11-d4fe596a2fbf"}
```

## Interaction 27: PATCH/todos/de562f6a-3902-4f17-9b11-d4fe596a2fbf

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 18
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos working with an existing todo can change the todo's completedness by PATCHing to the todo's url
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```
{"completed":true}
```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 277
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:43 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200: text/html;charset=utf-8):

```
{"uid":"de562f6a-3902-4f17-9b11-d4fe596a2fbf","title":"blah","order":0,"completed":true,"href":"https://http4k-todo-backend.herokuapp.com/todos/de562f6a-3902-4f17-9b11-d4fe596a2fbf","url":"https://http4k-todo-backend.herokuapp.com/todos/de562f6a-3902-4f17-9b11-d4fe596a2fbf"}
```

## Interaction 28: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos working with an existing todo can change the todo's completedness by PATCHing to the todo's url
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 0
Date: Sat, 19 Sep 2020 16:29:44 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204: ):

```

```

## Interaction 29: POST/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 16
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos working with an existing todo changes to a todo are persisted and show up when re-fetching the todo
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```
{"title":"blah"}
```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 268
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:44 GMT
Location: https://http4k-todo-backend.herokuapp.com/todos/ce27ee58-9cd5-415a-b320-9b2489d90e7e
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201: text/html;charset=utf-8):

```
{"title":"blah","uid":"ce27ee58-9cd5-415a-b320-9b2489d90e7e","completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/ce27ee58-9cd5-415a-b320-9b2489d90e7e","url":"https://http4k-todo-backend.herokuapp.com/todos/ce27ee58-9cd5-415a-b320-9b2489d90e7e"}
```

## Interaction 30: PATCH/todos/ce27ee58-9cd5-415a-b320-9b2489d90e7e

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 42
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos working with an existing todo changes to a todo are persisted and show up when re-fetching the todo
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```
{"title":"changed title","completed":true}
```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 286
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:44 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200: text/html;charset=utf-8):

```
{"uid":"ce27ee58-9cd5-415a-b320-9b2489d90e7e","title":"changed title","order":0,"completed":true,"href":"https://http4k-todo-backend.herokuapp.com/todos/ce27ee58-9cd5-415a-b320-9b2489d90e7e","url":"https://http4k-todo-backend.herokuapp.com/todos/ce27ee58-9cd5-415a-b320-9b2489d90e7e"}
```

## Interaction 31: GET/todos/ce27ee58-9cd5-415a-b320-9b2489d90e7e

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos working with an existing todo changes to a todo are persisted and show up when re-fetching the todo
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 286
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:45 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200: text/html;charset=utf-8):

```
{"uid":"ce27ee58-9cd5-415a-b320-9b2489d90e7e","title":"changed title","order":0,"completed":true,"href":"https://http4k-todo-backend.herokuapp.com/todos/ce27ee58-9cd5-415a-b320-9b2489d90e7e","url":"https://http4k-todo-backend.herokuapp.com/todos/ce27ee58-9cd5-415a-b320-9b2489d90e7e"}
```

## Interaction 32: GET/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos working with an existing todo changes to a todo are persisted and show up when re-fetching the todo
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 288
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:45 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200: text/html;charset=utf-8):

```
[{"uid":"ce27ee58-9cd5-415a-b320-9b2489d90e7e","title":"changed title","order":0,"completed":true,"href":"https://http4k-todo-backend.herokuapp.com/todos/ce27ee58-9cd5-415a-b320-9b2489d90e7e","url":"https://http4k-todo-backend.herokuapp.com/todos/ce27ee58-9cd5-415a-b320-9b2489d90e7e"}]
```

## Interaction 33: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos working with an existing todo changes to a todo are persisted and show up when re-fetching the todo
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 0
Date: Sat, 19 Sep 2020 16:29:45 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204: ):

```

```

## Interaction 34: POST/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 16
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos working with an existing todo can delete a todo making a DELETE request to the todo's url
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```
{"title":"blah"}
```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 268
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:46 GMT
Location: https://http4k-todo-backend.herokuapp.com/todos/bf78bd31-1a2f-40b6-9d29-e8ee3b262e86
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201: text/html;charset=utf-8):

```
{"title":"blah","uid":"bf78bd31-1a2f-40b6-9d29-e8ee3b262e86","completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/bf78bd31-1a2f-40b6-9d29-e8ee3b262e86","url":"https://http4k-todo-backend.herokuapp.com/todos/bf78bd31-1a2f-40b6-9d29-e8ee3b262e86"}
```

## Interaction 35: DELETE/todos/bf78bd31-1a2f-40b6-9d29-e8ee3b262e86

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos working with an existing todo can delete a todo making a DELETE request to the todo's url
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 0
Date: Sat, 19 Sep 2020 16:29:46 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204: ):

```

```

## Interaction 36: GET/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos working with an existing todo can delete a todo making a DELETE request to the todo's url
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 2
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:46 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200: text/html;charset=utf-8):

```
[]
```

## Interaction 37: POST/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 28
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos tracking todo order can create a todo with an order field
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```
{"title":"blah","order":523}
```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 280
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:47 GMT
Location: https://http4k-todo-backend.herokuapp.com/todos/c43e5d1d-d35b-478f-bb4d-be4479da559e
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201: text/html;charset=utf-8):

```
{"title":"blah","order":523,"uid":"c43e5d1d-d35b-478f-bb4d-be4479da559e","completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/c43e5d1d-d35b-478f-bb4d-be4479da559e","url":"https://http4k-todo-backend.herokuapp.com/todos/c43e5d1d-d35b-478f-bb4d-be4479da559e"}
```

## Interaction 38: POST/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 27
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos tracking todo order can PATCH a todo to change its order
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```
{"order":10,"title":"blah"}
```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 279
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:47 GMT
Location: https://http4k-todo-backend.herokuapp.com/todos/022f9a49-09a5-4bff-821a-2cdcceedbfd5
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201: text/html;charset=utf-8):

```
{"order":10,"title":"blah","uid":"022f9a49-09a5-4bff-821a-2cdcceedbfd5","completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/022f9a49-09a5-4bff-821a-2cdcceedbfd5","url":"https://http4k-todo-backend.herokuapp.com/todos/022f9a49-09a5-4bff-821a-2cdcceedbfd5"}
```

## Interaction 39: PATCH/todos/022f9a49-09a5-4bff-821a-2cdcceedbfd5

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 12
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos tracking todo order can PATCH a todo to change its order
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```
{"order":95}
```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 279
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:47 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200: text/html;charset=utf-8):

```
{"uid":"022f9a49-09a5-4bff-821a-2cdcceedbfd5","title":"blah","order":95,"completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/022f9a49-09a5-4bff-821a-2cdcceedbfd5","url":"https://http4k-todo-backend.herokuapp.com/todos/022f9a49-09a5-4bff-821a-2cdcceedbfd5"}
```

## Interaction 40: POST/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 27
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos tracking todo order remembers changes to a todo's order
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```
{"order":10,"title":"blah"}
```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 279
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:48 GMT
Location: https://http4k-todo-backend.herokuapp.com/todos/08b810e3-19b3-4322-af56-f4d8a458bc18
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201: text/html;charset=utf-8):

```
{"order":10,"title":"blah","uid":"08b810e3-19b3-4322-af56-f4d8a458bc18","completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/08b810e3-19b3-4322-af56-f4d8a458bc18","url":"https://http4k-todo-backend.herokuapp.com/todos/08b810e3-19b3-4322-af56-f4d8a458bc18"}
```

## Interaction 41: PATCH/todos/08b810e3-19b3-4322-af56-f4d8a458bc18

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 12
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://http4k-todo-backend.herokuapp.com/todos tracking todo order remembers changes to a todo's order
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```
{"order":95}
```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 279
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:48 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200: text/html;charset=utf-8):

```
{"uid":"08b810e3-19b3-4322-af56-f4d8a458bc18","title":"blah","order":95,"completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/08b810e3-19b3-4322-af56-f4d8a458bc18","url":"https://http4k-todo-backend.herokuapp.com/todos/08b810e3-19b3-4322-af56-f4d8a458bc18"}
```

## Interaction 42: GET/todos/08b810e3-19b3-4322-af56-f4d8a458bc18

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at http://http4k-todo-backend.herokuapp.com/todos tracking todo order remembers changes to a todo's order
Origin: https://servirtium.github.io
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 279
Content-Type: text/html;charset=utf-8
Date: Sat, 19 Sep 2020 16:29:48 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200: text/html;charset=utf-8):

```
{"uid":"08b810e3-19b3-4322-af56-f4d8a458bc18","title":"blah","order":95,"completed":false,"href":"https://http4k-todo-backend.herokuapp.com/todos/08b810e3-19b3-4322-af56-f4d8a458bc18","url":"https://http4k-todo-backend.herokuapp.com/todos/08b810e3-19b3-4322-af56-f4d8a458bc18"}
```
