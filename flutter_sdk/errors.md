1. post request is not hitting in django post method due to CORS problem, can be checked in browser. Usually the flow for developing app using flutter web:  

    Development stage: by turning off the CORS on google chrome #46904.  
      
    Release stage: we need to comply with the CORS using some mechanism.
    For example by putting the flutter frontend and backend on the same domain and using https.  

    SO, using middleware in django we can add         
    response["Access-Control-Allow-Origin"] = "*"  
    response["Access-Control-Allow-Headers"] = "*"          

    so, flutter http request will also send cors-Headers like origin, and then django should also send it in response. what is preflight request.

2. flutter run -d chrome --web-port 9999 --verbose
3. redirection  
   when redirect by server the response status code is 302 and the target url is sent in location header by the server and when it is received the client whill make the request to the location header value. we can say, redirection is a special type of response from server which has location header and a status code 302 and then client again send a request to that location.

4. google openid connect flow with flutter and django  
   solution is the redirect link should hit flutter page not server side. and then the flutter page can fetch the `code` param and then if code is not null then send a post request and check the response.