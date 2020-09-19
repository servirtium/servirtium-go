## Interaction 0: GET/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos the pre-requisites the api root responds to a GET (i.e. the server is up and accessible, CORS headers are set up)
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
Date: Sat, 19 Sep 2020 14:56:12 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
[{"uid":"a610da12-82f6-4fc5-8b01-d3274efb403b","title":"blah","order":523,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/a610da12-82f6-4fc5-8b01-d3274efb403b","url":"https://todo-backend-sinatra.herokuapp.com/todos/a610da12-82f6-4fc5-8b01-d3274efb403b"},{"uid":"71a0325d-02f1-413f-924d-a4d1dac4ebd9","title":"blah","order":10,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/71a0325d-02f1-413f-924d-a4d1dac4ebd9","url":"https://todo-backend-sinatra.herokuapp.com/todos/71a0325d-02f1-413f-924d-a4d1dac4ebd9"},{"uid":"3a0549a7-663f-4940-bdf0-421f18ba771f","title":"blah","order":10,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/3a0549a7-663f-4940-bdf0-421f18ba771f","url":"https://todo-backend-sinatra.herokuapp.com/todos/3a0549a7-663f-4940-bdf0-421f18ba771f"}]
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
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos the pre-requisites the api root responds to a POST with the todo which was posted to it
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
Date: Sat, 19 Sep 2020 14:56:12 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/1022dfde-da59-4eda-a43e-b0250a80854e
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"a todo","uid":"1022dfde-da59-4eda-a43e-b0250a80854e","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/1022dfde-da59-4eda-a43e-b0250a80854e","url":"https://todo-backend-sinatra.herokuapp.com/todos/1022dfde-da59-4eda-a43e-b0250a80854e"}
```

## Interaction 2: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos the pre-requisites the api root responds successfully to a DELETE
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
Date: Sat, 19 Sep 2020 14:56:13 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204 No Content:):

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
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos the pre-requisites after a DELETE the api root responds to a GET with a JSON representation of an empty array
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
Date: Sat, 19 Sep 2020 14:56:14 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204 No Content:):

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
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos the pre-requisites after a DELETE the api root responds to a GET with a JSON representation of an empty array
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
Date: Sat, 19 Sep 2020 14:56:14 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

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
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos the pre-requisites after a DELETE the api root responds to a GET with a JSON representation of an empty array
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
Date: Sat, 19 Sep 2020 14:56:14 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204 No Content:):

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
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos storing new todos by posting to the root url adds a new todo to the list of todos at the root url
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
Date: Sat, 19 Sep 2020 14:56:15 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/f1791699-98cb-4feb-a5d5-803a75327a2c
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"walk the dog","uid":"f1791699-98cb-4feb-a5d5-803a75327a2c","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/f1791699-98cb-4feb-a5d5-803a75327a2c","url":"https://todo-backend-sinatra.herokuapp.com/todos/f1791699-98cb-4feb-a5d5-803a75327a2c"}
```

## Interaction 7: GET/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos storing new todos by posting to the root url adds a new todo to the list of todos at the root url
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
Date: Sat, 19 Sep 2020 14:56:16 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
[{"uid":"f1791699-98cb-4feb-a5d5-803a75327a2c","title":"walk the dog","order":0,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/f1791699-98cb-4feb-a5d5-803a75327a2c","url":"https://todo-backend-sinatra.herokuapp.com/todos/f1791699-98cb-4feb-a5d5-803a75327a2c"}]
```

## Interaction 8: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos storing new todos by posting to the root url adds a new todo to the list of todos at the root url
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
Date: Sat, 19 Sep 2020 14:56:16 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204 No Content:):

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
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos storing new todos by posting to the root url sets up a new todo as initially not completed
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
Date: Sat, 19 Sep 2020 14:56:16 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/e1048fc0-fe31-4593-b11f-6b15fb92f9cc
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"blah","uid":"e1048fc0-fe31-4593-b11f-6b15fb92f9cc","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/e1048fc0-fe31-4593-b11f-6b15fb92f9cc","url":"https://todo-backend-sinatra.herokuapp.com/todos/e1048fc0-fe31-4593-b11f-6b15fb92f9cc"}
```

## Interaction 10: GET/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at http://todo-backend-sinatra.herokuapp.com/todos storing new todos by posting to the root url sets up a new todo as initially not completed
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
Date: Sat, 19 Sep 2020 14:56:17 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
[{"uid":"e1048fc0-fe31-4593-b11f-6b15fb92f9cc","title":"blah","order":0,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/e1048fc0-fe31-4593-b11f-6b15fb92f9cc","url":"https://todo-backend-sinatra.herokuapp.com/todos/e1048fc0-fe31-4593-b11f-6b15fb92f9cc"}]
```

## Interaction 11: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos storing new todos by posting to the root url sets up a new todo as initially not completed
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
Date: Sat, 19 Sep 2020 14:56:17 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204 No Content:):

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
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos storing new todos by posting to the root url each new todo has a url
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
Date: Sat, 19 Sep 2020 14:56:17 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/bcc93ad7-bbca-4829-90a9-62eeb33ab9ea
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"blah","uid":"bcc93ad7-bbca-4829-90a9-62eeb33ab9ea","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/bcc93ad7-bbca-4829-90a9-62eeb33ab9ea","url":"https://todo-backend-sinatra.herokuapp.com/todos/bcc93ad7-bbca-4829-90a9-62eeb33ab9ea"}
```

## Interaction 13: GET/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos storing new todos by posting to the root url each new todo has a url
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
Date: Sat, 19 Sep 2020 14:56:18 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
[{"uid":"bcc93ad7-bbca-4829-90a9-62eeb33ab9ea","title":"blah","order":0,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/bcc93ad7-bbca-4829-90a9-62eeb33ab9ea","url":"https://todo-backend-sinatra.herokuapp.com/todos/bcc93ad7-bbca-4829-90a9-62eeb33ab9ea"}]
```

## Interaction 14: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos storing new todos by posting to the root url each new todo has a url
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
Date: Sat, 19 Sep 2020 14:56:18 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204 No Content:):

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
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos storing new todos by posting to the root url each new todo has a url, which returns a todo
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
Date: Sat, 19 Sep 2020 14:56:18 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/401d613d-c6f2-4e02-a824-bd09e751bcff
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"my todo","uid":"401d613d-c6f2-4e02-a824-bd09e751bcff","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/401d613d-c6f2-4e02-a824-bd09e751bcff","url":"https://todo-backend-sinatra.herokuapp.com/todos/401d613d-c6f2-4e02-a824-bd09e751bcff"}
```

## Interaction 16: GET/todos/401d613d-c6f2-4e02-a824-bd09e751bcff

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos storing new todos by posting to the root url each new todo has a url, which returns a todo
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
Date: Sat, 19 Sep 2020 14:56:18 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
{"uid":"401d613d-c6f2-4e02-a824-bd09e751bcff","title":"my todo","order":0,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/401d613d-c6f2-4e02-a824-bd09e751bcff","url":"https://todo-backend-sinatra.herokuapp.com/todos/401d613d-c6f2-4e02-a824-bd09e751bcff"}
```

## Interaction 17: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos storing new todos by posting to the root url each new todo has a url, which returns a todo
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
Date: Sat, 19 Sep 2020 14:56:19 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204 No Content:):

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
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos working with an existing todo can navigate from a list of todos to an individual todo via urls
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
Date: Sat, 19 Sep 2020 14:56:19 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/caf17296-9394-4931-b6ff-5ca2c29c9863
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"todo the first","uid":"caf17296-9394-4931-b6ff-5ca2c29c9863","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/caf17296-9394-4931-b6ff-5ca2c29c9863","url":"https://todo-backend-sinatra.herokuapp.com/todos/caf17296-9394-4931-b6ff-5ca2c29c9863"}
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
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos working with an existing todo can navigate from a list of todos to an individual todo via urls
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
Date: Sat, 19 Sep 2020 14:56:19 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/f1346a7a-9820-4fe4-a651-7c9a488566a6
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"todo the second","uid":"f1346a7a-9820-4fe4-a651-7c9a488566a6","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/f1346a7a-9820-4fe4-a651-7c9a488566a6","url":"https://todo-backend-sinatra.herokuapp.com/todos/f1346a7a-9820-4fe4-a651-7c9a488566a6"}
```

## Interaction 20: GET/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at http://todo-backend-sinatra.herokuapp.com/todos working with an existing todo can navigate from a list of todos to an individual todo via urls
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
Date: Sat, 19 Sep 2020 14:56:20 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
[{"uid":"caf17296-9394-4931-b6ff-5ca2c29c9863","title":"todo the first","order":0,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/caf17296-9394-4931-b6ff-5ca2c29c9863","url":"https://todo-backend-sinatra.herokuapp.com/todos/caf17296-9394-4931-b6ff-5ca2c29c9863"},{"uid":"f1346a7a-9820-4fe4-a651-7c9a488566a6","title":"todo the second","order":0,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/f1346a7a-9820-4fe4-a651-7c9a488566a6","url":"https://todo-backend-sinatra.herokuapp.com/todos/f1346a7a-9820-4fe4-a651-7c9a488566a6"}]
```

## Interaction 21: GET/todos/caf17296-9394-4931-b6ff-5ca2c29c9863

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos working with an existing todo can navigate from a list of todos to an individual todo via urls
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
Date: Sat, 19 Sep 2020 14:56:20 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
{"uid":"caf17296-9394-4931-b6ff-5ca2c29c9863","title":"todo the first","order":0,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/caf17296-9394-4931-b6ff-5ca2c29c9863","url":"https://todo-backend-sinatra.herokuapp.com/todos/caf17296-9394-4931-b6ff-5ca2c29c9863"}
```

## Interaction 22: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos working with an existing todo can navigate from a list of todos to an individual todo via urls
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
Date: Sat, 19 Sep 2020 14:56:20 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204 No Content:):

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
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos working with an existing todo can change the todo's title by PATCHing to the todo's url
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
Date: Sat, 19 Sep 2020 14:56:21 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/2273a78d-99cb-4cef-8588-9501c3936f35
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"initial title","uid":"2273a78d-99cb-4cef-8588-9501c3936f35","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/2273a78d-99cb-4cef-8588-9501c3936f35","url":"https://todo-backend-sinatra.herokuapp.com/todos/2273a78d-99cb-4cef-8588-9501c3936f35"}
```

## Interaction 24: PATCH/todos/2273a78d-99cb-4cef-8588-9501c3936f35

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 25
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos working with an existing todo can change the todo's title by PATCHing to the todo's url
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
Date: Sat, 19 Sep 2020 14:56:21 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
{"uid":"2273a78d-99cb-4cef-8588-9501c3936f35","title":"bathe the cat","order":0,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/2273a78d-99cb-4cef-8588-9501c3936f35","url":"https://todo-backend-sinatra.herokuapp.com/todos/2273a78d-99cb-4cef-8588-9501c3936f35"}
```

## Interaction 25: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos working with an existing todo can change the todo's title by PATCHing to the todo's url
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
Date: Sat, 19 Sep 2020 14:56:22 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204 No Content:):

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
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos working with an existing todo can change the todo's completedness by PATCHing to the todo's url
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
Date: Sat, 19 Sep 2020 14:56:22 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/ef1a3258-d4a3-4aad-b646-c1dbc4ccfda6
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"blah","uid":"ef1a3258-d4a3-4aad-b646-c1dbc4ccfda6","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/ef1a3258-d4a3-4aad-b646-c1dbc4ccfda6","url":"https://todo-backend-sinatra.herokuapp.com/todos/ef1a3258-d4a3-4aad-b646-c1dbc4ccfda6"}
```

## Interaction 27: PATCH/todos/ef1a3258-d4a3-4aad-b646-c1dbc4ccfda6

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 18
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos working with an existing todo can change the todo's completedness by PATCHing to the todo's url
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
Date: Sat, 19 Sep 2020 14:56:22 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
{"uid":"ef1a3258-d4a3-4aad-b646-c1dbc4ccfda6","title":"blah","order":0,"completed":true,"href":"https://todo-backend-sinatra.herokuapp.com/todos/ef1a3258-d4a3-4aad-b646-c1dbc4ccfda6","url":"https://todo-backend-sinatra.herokuapp.com/todos/ef1a3258-d4a3-4aad-b646-c1dbc4ccfda6"}
```

## Interaction 28: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at http://todo-backend-sinatra.herokuapp.com/todos working with an existing todo can change the todo's completedness by PATCHing to the todo's url
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
Date: Sat, 19 Sep 2020 14:56:23 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204 No Content:):

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
Fulltitle: Todo-Backend API residing at http://todo-backend-sinatra.herokuapp.com/todos working with an existing todo changes to a todo are persisted and show up when re-fetching the todo
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
Date: Sat, 19 Sep 2020 14:56:23 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/b138ac2f-a4ed-4fcc-a721-325630101ff3
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"blah","uid":"b138ac2f-a4ed-4fcc-a721-325630101ff3","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/b138ac2f-a4ed-4fcc-a721-325630101ff3","url":"https://todo-backend-sinatra.herokuapp.com/todos/b138ac2f-a4ed-4fcc-a721-325630101ff3"}
```

## Interaction 30: PATCH/todos/b138ac2f-a4ed-4fcc-a721-325630101ff3

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 42
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos working with an existing todo changes to a todo are persisted and show up when re-fetching the todo
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
Date: Sat, 19 Sep 2020 14:56:23 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
{"uid":"b138ac2f-a4ed-4fcc-a721-325630101ff3","title":"changed title","order":0,"completed":true,"href":"https://todo-backend-sinatra.herokuapp.com/todos/b138ac2f-a4ed-4fcc-a721-325630101ff3","url":"https://todo-backend-sinatra.herokuapp.com/todos/b138ac2f-a4ed-4fcc-a721-325630101ff3"}
```

## Interaction 31: GET/todos/b138ac2f-a4ed-4fcc-a721-325630101ff3

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos working with an existing todo changes to a todo are persisted and show up when re-fetching the todo
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
Date: Sat, 19 Sep 2020 14:56:23 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
{"uid":"b138ac2f-a4ed-4fcc-a721-325630101ff3","title":"changed title","order":0,"completed":true,"href":"https://todo-backend-sinatra.herokuapp.com/todos/b138ac2f-a4ed-4fcc-a721-325630101ff3","url":"https://todo-backend-sinatra.herokuapp.com/todos/b138ac2f-a4ed-4fcc-a721-325630101ff3"}
```

## Interaction 32: GET/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos working with an existing todo changes to a todo are persisted and show up when re-fetching the todo
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
Date: Sat, 19 Sep 2020 14:56:23 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
[{"uid":"b138ac2f-a4ed-4fcc-a721-325630101ff3","title":"changed title","order":0,"completed":true,"href":"https://todo-backend-sinatra.herokuapp.com/todos/b138ac2f-a4ed-4fcc-a721-325630101ff3","url":"https://todo-backend-sinatra.herokuapp.com/todos/b138ac2f-a4ed-4fcc-a721-325630101ff3"}]
```

## Interaction 33: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos working with an existing todo changes to a todo are persisted and show up when re-fetching the todo
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
Date: Sat, 19 Sep 2020 14:56:23 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204 No Content:):

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
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos working with an existing todo can delete a todo making a DELETE request to the todo's url
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
Date: Sat, 19 Sep 2020 14:56:24 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/891aedb9-e374-497e-8e55-b65fba0ac5b0
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"blah","uid":"891aedb9-e374-497e-8e55-b65fba0ac5b0","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/891aedb9-e374-497e-8e55-b65fba0ac5b0","url":"https://todo-backend-sinatra.herokuapp.com/todos/891aedb9-e374-497e-8e55-b65fba0ac5b0"}
```

## Interaction 35: DELETE/todos/891aedb9-e374-497e-8e55-b65fba0ac5b0

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos working with an existing todo can delete a todo making a DELETE request to the todo's url
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
Date: Sat, 19 Sep 2020 14:56:25 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204 No Content:):

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
Fulltitle: Todo-Backend API residing at http://todo-backend-sinatra.herokuapp.com/todos working with an existing todo can delete a todo making a DELETE request to the todo's url
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
Date: Sat, 19 Sep 2020 14:56:25 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

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
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos tracking todo order can create a todo with an order field
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
Date: Sat, 19 Sep 2020 14:56:25 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/da19dcfe-5fdc-420b-a089-a972481a7441
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"blah","order":523,"uid":"da19dcfe-5fdc-420b-a089-a972481a7441","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/da19dcfe-5fdc-420b-a089-a972481a7441","url":"https://todo-backend-sinatra.herokuapp.com/todos/da19dcfe-5fdc-420b-a089-a972481a7441"}
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
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos tracking todo order can PATCH a todo to change its order
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
Date: Sat, 19 Sep 2020 14:56:26 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/9ce0866a-d10f-4d10-b7a4-957b3fabb355
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"order":10,"title":"blah","uid":"9ce0866a-d10f-4d10-b7a4-957b3fabb355","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/9ce0866a-d10f-4d10-b7a4-957b3fabb355","url":"https://todo-backend-sinatra.herokuapp.com/todos/9ce0866a-d10f-4d10-b7a4-957b3fabb355"}
```

## Interaction 39: PATCH/todos/9ce0866a-d10f-4d10-b7a4-957b3fabb355

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 12
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos tracking todo order can PATCH a todo to change its order
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
Date: Sat, 19 Sep 2020 14:56:26 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
{"uid":"9ce0866a-d10f-4d10-b7a4-957b3fabb355","title":"blah","order":95,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/9ce0866a-d10f-4d10-b7a4-957b3fabb355","url":"https://todo-backend-sinatra.herokuapp.com/todos/9ce0866a-d10f-4d10-b7a4-957b3fabb355"}
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
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos tracking todo order remembers changes to a todo's order
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
Date: Sat, 19 Sep 2020 14:56:26 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/7f1bb201-0736-42ea-be1c-a17afb1f55e5
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"order":10,"title":"blah","uid":"7f1bb201-0736-42ea-be1c-a17afb1f55e5","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/7f1bb201-0736-42ea-be1c-a17afb1f55e5","url":"https://todo-backend-sinatra.herokuapp.com/todos/7f1bb201-0736-42ea-be1c-a17afb1f55e5"}
```

## Interaction 41: PATCH/todos/7f1bb201-0736-42ea-be1c-a17afb1f55e5

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 12
Content-Type: application/json
Fulltitle: Todo-Backend API residing at http://todo-backend-sinatra.herokuapp.com/todos tracking todo order remembers changes to a todo's order
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
Date: Sat, 19 Sep 2020 14:56:27 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
{"uid":"7f1bb201-0736-42ea-be1c-a17afb1f55e5","title":"blah","order":95,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/7f1bb201-0736-42ea-be1c-a17afb1f55e5","url":"https://todo-backend-sinatra.herokuapp.com/todos/7f1bb201-0736-42ea-be1c-a17afb1f55e5"}
```

## Interaction 42: GET/todos/7f1bb201-0736-42ea-be1c-a17afb1f55e5

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Fulltitle: Todo-Backend API residing at https://todo-backend-sinatra.herokuapp.com/todos tracking todo order remembers changes to a todo's order
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
Date: Sat, 19 Sep 2020 14:56:27 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
{"uid":"7f1bb201-0736-42ea-be1c-a17afb1f55e5","title":"blah","order":95,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/7f1bb201-0736-42ea-be1c-a17afb1f55e5","url":"https://todo-backend-sinatra.herokuapp.com/todos/7f1bb201-0736-42ea-be1c-a17afb1f55e5"}
```
