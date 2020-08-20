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
Content-Length: 842
Content-Type: text/html;charset=utf-8
Date: Thu, 20 Aug 2020 17:30:14 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
[{"uid":"47886b47-fe17-4a99-a097-3b7bbe4a1d60","title":"blah","order":523,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/47886b47-fe17-4a99-a097-3b7bbe4a1d60","url":"https://todo-backend-sinatra.herokuapp.com/todos/47886b47-fe17-4a99-a097-3b7bbe4a1d60"},{"uid":"809a59a3-a9aa-44bf-b5b3-f2778a010bbd","title":"blah","order":95,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/809a59a3-a9aa-44bf-b5b3-f2778a010bbd","url":"https://todo-backend-sinatra.herokuapp.com/todos/809a59a3-a9aa-44bf-b5b3-f2778a010bbd"},{"uid":"7a3597c0-13f1-4997-8da4-dbf0e8511daa","title":"blah","order":95,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/7a3597c0-13f1-4997-8da4-dbf0e8511daa","url":"https://todo-backend-sinatra.herokuapp.com/todos/7a3597c0-13f1-4997-8da4-dbf0e8511daa"}]
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
Date: Thu, 20 Aug 2020 17:30:14 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/24a1a77d-4b3c-4b59-adc9-0887181a4dbe
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"a todo","uid":"24a1a77d-4b3c-4b59-adc9-0887181a4dbe","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/24a1a77d-4b3c-4b59-adc9-0887181a4dbe","url":"https://todo-backend-sinatra.herokuapp.com/todos/24a1a77d-4b3c-4b59-adc9-0887181a4dbe"}
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
Date: Thu, 20 Aug 2020 17:30:14 GMT
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
Date: Thu, 20 Aug 2020 17:30:15 GMT
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
Date: Thu, 20 Aug 2020 17:30:15 GMT
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
Date: Thu, 20 Aug 2020 17:30:15 GMT
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
Date: Thu, 20 Aug 2020 17:30:15 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/5a2cf1b7-5bb3-4e8e-9b8f-761854cef57f
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"walk the dog","uid":"5a2cf1b7-5bb3-4e8e-9b8f-761854cef57f","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/5a2cf1b7-5bb3-4e8e-9b8f-761854cef57f","url":"https://todo-backend-sinatra.herokuapp.com/todos/5a2cf1b7-5bb3-4e8e-9b8f-761854cef57f"}
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
Date: Thu, 20 Aug 2020 17:30:16 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
[{"uid":"5a2cf1b7-5bb3-4e8e-9b8f-761854cef57f","title":"walk the dog","order":0,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/5a2cf1b7-5bb3-4e8e-9b8f-761854cef57f","url":"https://todo-backend-sinatra.herokuapp.com/todos/5a2cf1b7-5bb3-4e8e-9b8f-761854cef57f"}]
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
Date: Thu, 20 Aug 2020 17:30:16 GMT
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
Date: Thu, 20 Aug 2020 17:30:16 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/f792a371-501b-4160-8705-fff93f5dd3ba
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"blah","uid":"f792a371-501b-4160-8705-fff93f5dd3ba","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/f792a371-501b-4160-8705-fff93f5dd3ba","url":"https://todo-backend-sinatra.herokuapp.com/todos/f792a371-501b-4160-8705-fff93f5dd3ba"}
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
Date: Thu, 20 Aug 2020 17:30:17 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
[{"uid":"f792a371-501b-4160-8705-fff93f5dd3ba","title":"blah","order":0,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/f792a371-501b-4160-8705-fff93f5dd3ba","url":"https://todo-backend-sinatra.herokuapp.com/todos/f792a371-501b-4160-8705-fff93f5dd3ba"}]
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
Date: Thu, 20 Aug 2020 17:30:17 GMT
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
Date: Thu, 20 Aug 2020 17:30:17 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/04a59516-3486-4e5f-92dc-f9a07de3b346
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"blah","uid":"04a59516-3486-4e5f-92dc-f9a07de3b346","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/04a59516-3486-4e5f-92dc-f9a07de3b346","url":"https://todo-backend-sinatra.herokuapp.com/todos/04a59516-3486-4e5f-92dc-f9a07de3b346"}
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
Date: Thu, 20 Aug 2020 17:30:17 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
[{"uid":"04a59516-3486-4e5f-92dc-f9a07de3b346","title":"blah","order":0,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/04a59516-3486-4e5f-92dc-f9a07de3b346","url":"https://todo-backend-sinatra.herokuapp.com/todos/04a59516-3486-4e5f-92dc-f9a07de3b346"}]
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
Date: Thu, 20 Aug 2020 17:30:18 GMT
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
Date: Thu, 20 Aug 2020 17:30:18 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/b3dce1fb-127b-4b53-ba83-41bd5905408f
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"my todo","uid":"b3dce1fb-127b-4b53-ba83-41bd5905408f","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/b3dce1fb-127b-4b53-ba83-41bd5905408f","url":"https://todo-backend-sinatra.herokuapp.com/todos/b3dce1fb-127b-4b53-ba83-41bd5905408f"}
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
Date: Thu, 20 Aug 2020 17:30:20 GMT
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
Date: Thu, 20 Aug 2020 17:30:20 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/b9ec18d0-f23b-4843-ab6a-b470f64c0d49
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"todo the first","uid":"b9ec18d0-f23b-4843-ab6a-b470f64c0d49","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/b9ec18d0-f23b-4843-ab6a-b470f64c0d49","url":"https://todo-backend-sinatra.herokuapp.com/todos/b9ec18d0-f23b-4843-ab6a-b470f64c0d49"}
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
Date: Thu, 20 Aug 2020 17:30:20 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/1e926ef2-3764-4218-b223-a84bb2945c34
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"todo the second","uid":"1e926ef2-3764-4218-b223-a84bb2945c34","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/1e926ef2-3764-4218-b223-a84bb2945c34","url":"https://todo-backend-sinatra.herokuapp.com/todos/1e926ef2-3764-4218-b223-a84bb2945c34"}
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
Date: Thu, 20 Aug 2020 17:30:20 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
[{"uid":"b9ec18d0-f23b-4843-ab6a-b470f64c0d49","title":"todo the first","order":0,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/b9ec18d0-f23b-4843-ab6a-b470f64c0d49","url":"https://todo-backend-sinatra.herokuapp.com/todos/b9ec18d0-f23b-4843-ab6a-b470f64c0d49"},{"uid":"1e926ef2-3764-4218-b223-a84bb2945c34","title":"todo the second","order":0,"completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/1e926ef2-3764-4218-b223-a84bb2945c34","url":"https://todo-backend-sinatra.herokuapp.com/todos/1e926ef2-3764-4218-b223-a84bb2945c34"}]
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
Date: Thu, 20 Aug 2020 17:30:22 GMT
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
Date: Thu, 20 Aug 2020 17:30:22 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/b1a9e6ec-6986-49b1-92eb-175167aaa57e
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"initial title","uid":"b1a9e6ec-6986-49b1-92eb-175167aaa57e","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/b1a9e6ec-6986-49b1-92eb-175167aaa57e","url":"https://todo-backend-sinatra.herokuapp.com/todos/b1a9e6ec-6986-49b1-92eb-175167aaa57e"}
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
Date: Thu, 20 Aug 2020 17:30:24 GMT
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
Date: Thu, 20 Aug 2020 17:30:24 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/7f44c660-95b2-4ab4-8dc4-4d09ba80c9c9
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"blah","uid":"7f44c660-95b2-4ab4-8dc4-4d09ba80c9c9","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/7f44c660-95b2-4ab4-8dc4-4d09ba80c9c9","url":"https://todo-backend-sinatra.herokuapp.com/todos/7f44c660-95b2-4ab4-8dc4-4d09ba80c9c9"}
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
Date: Thu, 20 Aug 2020 17:30:25 GMT
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
Date: Thu, 20 Aug 2020 17:30:26 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/5d02b088-c25d-41e5-b943-9342cca21a3c
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"blah","uid":"5d02b088-c25d-41e5-b943-9342cca21a3c","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/5d02b088-c25d-41e5-b943-9342cca21a3c","url":"https://todo-backend-sinatra.herokuapp.com/todos/5d02b088-c25d-41e5-b943-9342cca21a3c"}
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
Date: Thu, 20 Aug 2020 17:30:27 GMT
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (200 OK:text/html;charset=utf-8):

```
[{"uid":"5d02b088-c25d-41e5-b943-9342cca21a3c","title":"changed title","order":0,"completed":true,"href":"https://todo-backend-sinatra.herokuapp.com/todos/5d02b088-c25d-41e5-b943-9342cca21a3c","url":"https://todo-backend-sinatra.herokuapp.com/todos/5d02b088-c25d-41e5-b943-9342cca21a3c"}]
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
Date: Thu, 20 Aug 2020 17:30:27 GMT
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
Date: Thu, 20 Aug 2020 17:30:28 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/e385b432-c8ea-4ac1-bf30-0a56fadbc0fa
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"blah","uid":"e385b432-c8ea-4ac1-bf30-0a56fadbc0fa","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/e385b432-c8ea-4ac1-bf30-0a56fadbc0fa","url":"https://todo-backend-sinatra.herokuapp.com/todos/e385b432-c8ea-4ac1-bf30-0a56fadbc0fa"}
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
Date: Thu, 20 Aug 2020 17:30:29 GMT
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
Date: Thu, 20 Aug 2020 17:30:29 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/d8bcea51-4400-4ad3-99c1-1f110753cc17
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"title":"blah","order":523,"uid":"d8bcea51-4400-4ad3-99c1-1f110753cc17","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/d8bcea51-4400-4ad3-99c1-1f110753cc17","url":"https://todo-backend-sinatra.herokuapp.com/todos/d8bcea51-4400-4ad3-99c1-1f110753cc17"}
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
Date: Thu, 20 Aug 2020 17:30:30 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/e94238ec-6be2-480b-b6c1-8b7b90550da2
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"order":10,"title":"blah","uid":"e94238ec-6be2-480b-b6c1-8b7b90550da2","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/e94238ec-6be2-480b-b6c1-8b7b90550da2","url":"https://todo-backend-sinatra.herokuapp.com/todos/e94238ec-6be2-480b-b6c1-8b7b90550da2"}
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
Date: Thu, 20 Aug 2020 17:30:31 GMT
Location: https://todo-backend-sinatra.herokuapp.com/todos/b339670e-5348-4317-b2d7-5a4f909993a1
Server: thin 1.6.2 codename Doc Brown
Via: 1.1 vegur
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

```

### Response body recorded for playback (201 Created:text/html;charset=utf-8):

```
{"order":10,"title":"blah","uid":"b339670e-5348-4317-b2d7-5a4f909993a1","completed":false,"href":"https://todo-backend-sinatra.herokuapp.com/todos/b339670e-5348-4317-b2d7-5a4f909993a1","url":"https://todo-backend-sinatra.herokuapp.com/todos/b339670e-5348-4317-b2d7-5a4f909993a1"}
```
