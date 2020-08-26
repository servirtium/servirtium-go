## Interaction 0: GET/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 562
Content-Type: text/html;charset=utf-8
Date: Wed, 26 Aug 2020 17:39:24 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
[{"uid":"434c2fb2-639b-40bf-99c7-2be701f3296c","title":"blah","order":523,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/434c2fb2-639b-40bf-99c7-2be701f3296c","url":"https://todo-backend-sinatra.herokuapp.com/todos/434c2fb2-639b-40bf-99c7-2be701f3296c"},{"uid":"1a9d93bc-479b-4078-bc2f-622c412114c8","title":"blah","order":10,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/1a9d93bc-479b-4078-bc2f-622c412114c8","url":"https://todo-backend-sinatra.herokuapp.com/todos/1a9d93bc-479b-4078-bc2f-622c412114c8"}]
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
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

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
Date: Wed, 26 Aug 2020 17:39:24 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/b7d00c22-54ca-4425-98ac-e33672eade0d
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"a todo","uid":"b7d00c22-54ca-4425-98ac-e33672eade0d","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/b7d00c22-54ca-4425-98ac-e33672eade0d","url":"https://todo-backend-sinatra.herokuapp.com/todos/b7d00c22-54ca-4425-98ac-e33672eade0d"}
```

## Interaction 2: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 0
Date: Wed, 26 Aug 2020 17:39:24 GMT
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
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 0
Date: Wed, 26 Aug 2020 17:39:25 GMT
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
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

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
Date: Wed, 26 Aug 2020 17:39:25 GMT
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
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 0
Date: Wed, 26 Aug 2020 17:39:25 GMT
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
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

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
Date: Wed, 26 Aug 2020 17:39:26 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/b77bb73b-c21a-413e-b684-731f86903782
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"walk the dog","uid":"b77bb73b-c21a-413e-b684-731f86903782","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/b77bb73b-c21a-413e-b684-731f86903782","url":"https://todo-backend-sinatra.herokuapp.com/todos/b77bb73b-c21a-413e-b684-731f86903782"}
```

## Interaction 7: GET/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

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
Date: Wed, 26 Aug 2020 17:39:26 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
[{"uid":"b77bb73b-c21a-413e-b684-731f86903782","title":"walk the dog","order":0,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/b77bb73b-c21a-413e-b684-731f86903782","url":"https://todo-backend-sinatra.herokuapp.com/todos/b77bb73b-c21a-413e-b684-731f86903782"}]
```

## Interaction 8: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 0
Date: Wed, 26 Aug 2020 17:39:26 GMT
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
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

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
Date: Wed, 26 Aug 2020 17:39:26 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/e843a487-cff4-4599-9cb1-13cbcc579a9f
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"blah","uid":"e843a487-cff4-4599-9cb1-13cbcc579a9f","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/e843a487-cff4-4599-9cb1-13cbcc579a9f","url":"https://todo-backend-sinatra.herokuapp.com/todos/e843a487-cff4-4599-9cb1-13cbcc579a9f"}
```

## Interaction 10: GET/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

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
Date: Wed, 26 Aug 2020 17:39:27 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
[{"uid":"e843a487-cff4-4599-9cb1-13cbcc579a9f","title":"blah","order":0,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/e843a487-cff4-4599-9cb1-13cbcc579a9f","url":"https://todo-backend-sinatra.herokuapp.com/todos/e843a487-cff4-4599-9cb1-13cbcc579a9f"}]
```

## Interaction 11: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 0
Date: Wed, 26 Aug 2020 17:39:27 GMT
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
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

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
Date: Wed, 26 Aug 2020 17:39:27 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/293e4926-10c3-4924-b4b2-54c728f5fb5a
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"blah","uid":"293e4926-10c3-4924-b4b2-54c728f5fb5a","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/293e4926-10c3-4924-b4b2-54c728f5fb5a","url":"https://todo-backend-sinatra.herokuapp.com/todos/293e4926-10c3-4924-b4b2-54c728f5fb5a"}
```

## Interaction 13: GET/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

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
Date: Wed, 26 Aug 2020 17:39:28 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
[{"uid":"293e4926-10c3-4924-b4b2-54c728f5fb5a","title":"blah","order":0,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/293e4926-10c3-4924-b4b2-54c728f5fb5a","url":"https://todo-backend-sinatra.herokuapp.com/todos/293e4926-10c3-4924-b4b2-54c728f5fb5a"}]
```

## Interaction 14: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 0
Date: Wed, 26 Aug 2020 17:39:28 GMT
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
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

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
Date: Wed, 26 Aug 2020 17:39:28 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/833fb42e-938b-4ba8-8e94-4f273c59b3fa
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"my todo","uid":"833fb42e-938b-4ba8-8e94-4f273c59b3fa","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/833fb42e-938b-4ba8-8e94-4f273c59b3fa","url":"https://todo-backend-sinatra.herokuapp.com/todos/833fb42e-938b-4ba8-8e94-4f273c59b3fa"}
```

## Interaction 16: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 0
Date: Wed, 26 Aug 2020 17:39:30 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204 No Content:):

```

```

## Interaction 17: POST/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 26
Content-Type: application/json
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

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
Date: Wed, 26 Aug 2020 17:39:30 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/8f2a2f48-9f92-4625-855a-f8e72019e5a3
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"todo the first","uid":"8f2a2f48-9f92-4625-855a-f8e72019e5a3","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/8f2a2f48-9f92-4625-855a-f8e72019e5a3","url":"https://todo-backend-sinatra.herokuapp.com/todos/8f2a2f48-9f92-4625-855a-f8e72019e5a3"}
```

## Interaction 18: POST/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 27
Content-Type: application/json
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

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
Date: Wed, 26 Aug 2020 17:39:31 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/0d655306-795b-41d8-951d-546489a2b039
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"todo the second","uid":"0d655306-795b-41d8-951d-546489a2b039","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/0d655306-795b-41d8-951d-546489a2b039","url":"https://todo-backend-sinatra.herokuapp.com/todos/0d655306-795b-41d8-951d-546489a2b039"}
```

## Interaction 19: GET/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

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
Date: Wed, 26 Aug 2020 17:39:31 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
[{"uid":"8f2a2f48-9f92-4625-855a-f8e72019e5a3","title":"todo the first","order":0,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/8f2a2f48-9f92-4625-855a-f8e72019e5a3","url":"https://todo-backend-sinatra.herokuapp.com/todos/8f2a2f48-9f92-4625-855a-f8e72019e5a3"},{"uid":"0d655306-795b-41d8-951d-546489a2b039","title":"todo the second","order":0,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/0d655306-795b-41d8-951d-546489a2b039","url":"https://todo-backend-sinatra.herokuapp.com/todos/0d655306-795b-41d8-951d-546489a2b039"}]
```

## Interaction 20: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 0
Date: Wed, 26 Aug 2020 17:39:32 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204 No Content:):

```

```

## Interaction 21: POST/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 25
Content-Type: application/json
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

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
Date: Wed, 26 Aug 2020 17:39:33 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/7b2f2d44-7563-4be1-96e8-d7a33ce020d3
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"initial title","uid":"7b2f2d44-7563-4be1-96e8-d7a33ce020d3","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/7b2f2d44-7563-4be1-96e8-d7a33ce020d3","url":"https://todo-backend-sinatra.herokuapp.com/todos/7b2f2d44-7563-4be1-96e8-d7a33ce020d3"}
```

## Interaction 22: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 0
Date: Wed, 26 Aug 2020 17:39:33 GMT
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
Content-Length: 16
Content-Type: application/json
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

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
Date: Wed, 26 Aug 2020 17:39:34 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/15e36e90-4fd5-4967-b9f4-5284ca918dce
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"blah","uid":"15e36e90-4fd5-4967-b9f4-5284ca918dce","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/15e36e90-4fd5-4967-b9f4-5284ca918dce","url":"https://todo-backend-sinatra.herokuapp.com/todos/15e36e90-4fd5-4967-b9f4-5284ca918dce"}
```

## Interaction 24: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 0
Date: Wed, 26 Aug 2020 17:39:35 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204 No Content:):

```

```

## Interaction 25: POST/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 16
Content-Type: application/json
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

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
Date: Wed, 26 Aug 2020 17:39:36 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/11265769-20b0-46b0-97cb-4d8ed0502f9c
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"blah","uid":"11265769-20b0-46b0-97cb-4d8ed0502f9c","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/11265769-20b0-46b0-97cb-4d8ed0502f9c","url":"https://todo-backend-sinatra.herokuapp.com/todos/11265769-20b0-46b0-97cb-4d8ed0502f9c"}
```

## Interaction 26: GET/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

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
Date: Wed, 26 Aug 2020 17:39:37 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
[{"uid":"11265769-20b0-46b0-97cb-4d8ed0502f9c","title":"changed title","order":0,"completed":true,"href":"https://todo-backend-sinatra.herokuapp.com/todos/11265769-20b0-46b0-97cb-4d8ed0502f9c","url":"https://todo-backend-sinatra.herokuapp.com/todos/11265769-20b0-46b0-97cb-4d8ed0502f9c"}]
```

## Interaction 27: DELETE/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

```

### Request body recorded for playback ():

```

```


### Response headers recorded for playback:

```
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Length: 0
Date: Wed, 26 Aug 2020 17:39:37 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff

```

### Response body recorded for playback (204 No Content:):

```

```

## Interaction 28: POST/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 16
Content-Type: application/json
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

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
Date: Wed, 26 Aug 2020 17:39:37 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/7bfd1166-e44b-4bec-8903-2cd53dd40243
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"blah","uid":"7bfd1166-e44b-4bec-8903-2cd53dd40243","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/7bfd1166-e44b-4bec-8903-2cd53dd40243","url":"https://todo-backend-sinatra.herokuapp.com/todos/7bfd1166-e44b-4bec-8903-2cd53dd40243"}
```

## Interaction 29: GET/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Type: application/json
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

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
Date: Wed, 26 Aug 2020 17:39:39 GMT
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

## Interaction 30: POST/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 28
Content-Type: application/json
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

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
Date: Wed, 26 Aug 2020 17:39:39 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/840d2abe-b7a3-463d-9079-af871a7e353b
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"blah","order":523,"uid":"840d2abe-b7a3-463d-9079-af871a7e353b","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/840d2abe-b7a3-463d-9079-af871a7e353b","url":"https://todo-backend-sinatra.herokuapp.com/todos/840d2abe-b7a3-463d-9079-af871a7e353b"}
```

## Interaction 31: POST/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 27
Content-Type: application/json
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

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
Date: Wed, 26 Aug 2020 17:39:39 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/db803c1d-0c97-4043-b358-a76e51b78d5e
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"order":10,"title":"blah","uid":"db803c1d-0c97-4043-b358-a76e51b78d5e","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/db803c1d-0c97-4043-b358-a76e51b78d5e","url":"https://todo-backend-sinatra.herokuapp.com/todos/db803c1d-0c97-4043-b358-a76e51b78d5e"}
```

## Interaction 32: POST/todos

### Request headers recorded for playback:

```
Accept: text/plain, */*; q=0.01
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Content-Length: 27
Content-Type: application/json
Origin: https://www.todobackend.com
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36

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
Date: Wed, 26 Aug 2020 17:39:41 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/4b1327e8-03b1-4906-8d5a-461c224071b9
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"order":10,"title":"blah","uid":"4b1327e8-03b1-4906-8d5a-461c224071b9","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/4b1327e8-03b1-4906-8d5a-461c224071b9","url":"https://todo-backend-sinatra.herokuapp.com/todos/4b1327e8-03b1-4906-8d5a-461c224071b9"}
```
