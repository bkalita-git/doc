1. handler:  
	A variable of any type that implements ServeMux(rw http.ResponseWriter)
	```
	type a_handler_type int
	func (x a_handler_type) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
		d, _ := ioutil.ReadAll(r.Body)
		print(string(d))
	}s
	var a_handler a_handler_type
	```
2. servemux  
	A variable that is initialized by the address of an initialized ServeMux[it is a type and underlying type is a structure]
	```
	var servemux = &http.ServeMux{}
	```
3. handler registrerer  
	the ServeMux type also implements a method called Handle(pattern string,handler Handler), the job is to  register the Handler to the string
	```
	servemux.Handle("/",a_handler)
	```
4. listen to a servemux
	```
	http.ListenAndServe(":9091", servemux)
	```
		
5. gracefully shutdown:
   do somthig before termination of the server

REFERENCES:
	
