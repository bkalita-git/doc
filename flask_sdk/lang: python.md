#THE RUNTIME
	So Python is not defined by its language reference only. It would be also wrong to say that Python is defined by its reference implementation, CPython, since there are some implementation details that are not a part of the language. The garbage collector that relies on a reference counting is one example. Since there is no single source of truth, we may say that Python is defined partly by the Python Language Reference and partly by its main implementation, CPython.

	cpython_interpreter->AST(abstract syntax tree)->bytecode->vm

	Bytecode is a series of instructions. Each instruction consists of two bytes: one for an opcode and one for an argument. Consider an example:
	def g(x):
		return x + 3

	$ python -m dis example1.py
	...
	2           0 LOAD_FAST            0 (x)
		        2 LOAD_CONST           1 (3)
		        4 BINARY_ADD
		        6 RETURN_VALUE

	cpython(is a python interpreter written in c) compiles python language to bytecode and run in the cpython virtual machine.
	
	code Object
	function Object
	frame Object
	thread State
	interpreter State
	runtime State
	
	STACK FRAME: [interpreter init process]
		the interpreter stroes FRAME OBJECT in STACK FRAME while in runtime.
		STACK FRAME  is just a 	STACK
		FRAME OBJECT is just an OBJECT
		
		for example: 
			foo() is a function call, when it iterpretes by the interpreter it creates a FRAME OBJECT for the function foo() using the information in 'def foo()' and push it to the STACK FRAME
			so, FRAME OBJECT for foo() can contain:
			1. local and global variables of the function as dictionary 
			[but, locals() and globals() function already hold them, so this FRAME OBJECT can use them instead of storing local, global variables directly inside thr FRAME OBJECT]
			2. FRAME Attribute
				2.1. CODE OBJECT
			AND MANY MORE.

#ALMOST EVERYTHING IN PYTHON IS AN OBJECT
	it's like everything we can do using metaclass.

	A type can produce another_type
	-------------------------------
	metaclass --------- metaclass
	metaclass --------- class
	class     --------- instance
	instance  --------- 

	class:
		implementation:
			either using class keyword		[manually]
			either using metaclass 		    [dynamically]
	function:
		object of a class having __call__ method
		implementation:
			either using metaclasss 	[dynamically]
			either using class keyword  [manually]
			either using def keyword	[manually]
	iterator:
		a class having __next__ method

	generator:
		object of a class having __next__ and __iter__ method
		implementation:
			either using metaclass
			either using class keyword
			either using def and yield keyword
			either using parenthesis lietrals ()
			
	certain instance/object:
		implementation:
			using class_name()
			using literals such as, "string", [2,3,4],{"key":"word"}...
			
			
	<class 'type'>
	<class 'function'>
	<class 'NoneType'>
	<class 'int'>
	<class 'float'>
	<class 'generator'>
	<class 'tuple'>
	<class 'list'>
	<class 'str'>
	<class 'range'>
	
	we can achieve same things, i.e. object by different techniques
	
#META CLASS 'type'
	The new version of python and old version of python uses this creation of a class differently.
	'type' class can be used to create class dynamically. 
	any class in Python 3, is an instance of the 'type' metaclass
	type is also an instance of the type metaclass, so it is an instance of itself.
	Calling type() in this manner creates a new instance of the 'type' metaclass. In other words, it dynamically creates a new class.
	eg. MyShinyClass = type('MyShinyClass', (), {}) # returns a class object
	
	>>> class Foo(object):
	...       bar = True
			or
	>>> Foo = type('Foo', (), {'bar':True})

	#Classes are objects too.
	#This object (the class) is itself capable of creating objects (the instances), and this is why it's a class.
	#type can take the description of a class as parameters, and return a class.
	class Foo:
		pass
	x=Foo()
	#when python encounters Foo():
		#As soon as you use the keyword class, Python executes it and creates an object. creates in memory an object with the name Foo
		1. Is there a __metaclass__ attribute in Foo, 
			->If yes, create in-memory a class object (I said a class object, stay with me here), with the name Foo by using what is in __metaclass__
			->If no, look for __metaclass__ in the class where it's inherited from. (in 'type')
			->If no parent class, look for a __metaclass__ at the MODULE level,basically old-style classes
		1. The __call__() method of Foo’s parent class is called. which is inside type metaclass
		2. the __call__() method invokes __new__() and __init__()
	
	x.__class__.__class__ is <type 'type'>

#UNDERSTANDING Names and Object, NO VARIABLES
	literals have addresses too.
	>>>id(323)
	139998696512656
	>>>id(323)
	139998696512656

	x = 323
	>>> id(323)
	139998696512816
	>>> id(x)
	139998696348688


	>>> x = x-1
	>>> id(x)
	139998696512816
	the address of x is changed. it's not a variable, it's a name.


	x = 323 as follows:

		Create a PyObject
		Set the typecode to integer for the PyObject
		Set the value to 323 for the PyObject
		Create a name called x
		Point x to the new PyObject
		Increase the refcount of the PyObject by 1

	so, Type, Value and ReferenceCount.
	**a name is only references to the address of the object.
	**x points to the memory address of 323 object
	**id(x) will display the address of 323


	Varaiables are passed by reference in Python, Bydefault.
	>>>from os import sys
	>>>sys.getrefcount(x)
	2
	the output is 2 because it creates another reference to name x.

	y = 9
	y1 = y
	y2 = y

	the id of y,y1,y2 will be same as the address of 9. called **interning.
	>>>y1 = y.copy() will make a different copy of 9
	>>> y==y1 tels if content are same, y is y1 tells if they are referencing to same object


	>>>sys.getsizeof(x) in bytes

#GENERATORS USING def keyword
	#below, each block will be triggered as every next(gen)
	#yield keyword pauses execution
	#yield from behave like replaces instruction and return from a yield from another function takes the value 
	def generator():
		return "this will throw StopIteration and below code will not be proceed but we can retrieve this string using exception" #BLOCK0
		print("this execute till the next yield") #BLOCK1
		print("this execute till the next yield") #BLOCK1
		yield 1 #returns 1	#BLOCK1
		yield 2 #returns 2  #BLOCK2
		print("this will be executed")  #BLOCK3
		return "this will throw StopIteration" #but we can retrieve this string using exception #BLOCK3
	gen = generator() #returns object
	
	next(gen)
		will execute till the next yield
		will execute till the next yield
		1
	
	next(gen)
		2
		
	next(gen)
		exception since there is no yield.
		exception StopIteration as exc:
			exc.value
					
	--------------------------------------------------------------------------------
	
	def generator():
		val = yield 1 #BLOCK0
		print(val) #BLOCK1
		print(val) #BLOCK1
		yield 2    #BLOCK1
	
	BLOCK1 expecting a value from the BLOCK0 in print function but next(gen) will pop the BLOCK0, so we can manually provide it using gen.send(BLOCK0 value) and BLOCk1 will be executed.
	
	Generator expression:
	(x for x in range(10))

#GIL: GLOBAL INTERPRETER LOCK (Memory Management Feature)

#LAMBDA
	x = lambda a : a + 10
	print(x(5)) 


#UNDERSTANDING PACKAGE AND MODULE
	Modules are scripts, like the normal python file which is a main module. we can import function, names from a module. like, 
	>>>import module_name
	>>>from module import func as function
	if __name__ == "__main__":
		import sys
		fib(int(sys.argv[1]))
	Inside Package there are modules, and __init__.py, __init__.py for importing modules, as >>>from .ModuleA import something, see the dot(.)
	
	when you import a package, only variables/functions/classes in the __init__.py file of that package are directly visible, not sub-packages or modules. 

#BUILTIN PACKAGES/STANDARD LIBRARY IN PYTHON
	https://docs.python.org/3/library/

#TYPE ANNOTATION
	def get(name:str)->str: #return type is expected to be a string (->str)
		return "string"
		
#DECORATOR
	def decorator(f):
		def new_fun():
			return f()+" hi"
		return new_fun
	
	@decorator #similar to f = decorator(f); f() #decorator is a function which takes a function and return a modified function
	def f():
		string = "I want to change"
		return string

#ENCODING DECODING BASE64
	string-->ascii_byte-->base64_byte--->base64_string (encode)
	
	base64_string->base64_byte->ascii_byte->string	   (decode)
	import base64
	string = "hi"
	byte_in_ascii = string.encode("ascii")
	byte_in_base64 = base64.b64encode(byte_in_ascii)
	string_in_base64 = byte_in_base64.decode("ascii")
	
	byte_in_base64 = string_in_base64.encode("ascii")
	byte_in_ascii = base64.b64decode(byte_in_base64)
	string = byte_in_ascii.decode("ascii")
	
	
#*args vs **kwargs:
	kwargs['key_name']
	args[0]

#3 dots (...) or ellipises
	... is a token having the singleton value Ellipsis, similar to how None is a singleton value
	as a placeholder for not yet written, can be used in variable length of arguments.

#<function signature> specifies the input arguments to the function

#Optional[int] from typing module is, Union[int,None], ie. either int or None,  isinstance(x, T) where T is our custom TypeVar

#with statement
	 It allows you to ensure that a resource is "cleaned up" when the code that uses it finishes running, even if exceptions are thrown. It provides 'syntactic sugar' for try/finally blocks.
  with open('output.txt', 'w') as f: #automatically close the file after writing. 
 	f.write('Hi there!')
	#If an exception occurs before the end of the block, it will close the file before the exception is caught by an outer exception handler.

## Abstract class in python
> normally in python if one function in base class is not reimplemented in derived class then while instantiating derived class object there would be no errors. What if we want to make that base class **abstarct**? Use **abc** module. 
```
from abc import ABCMeta, abstractmethod
class Base(metaclass=ABCMeta):
	@abstractmethod
	def fun(self):
		...

Class Derived(Base):
	#reimeplement fun here otherwise error.
``` 
## Calling as Classname.method() also as classObject.method()
  * >using @classmethod
	```
	class Base:
		@classmethod
		def fun(cls):
			#it's not self
			print(cls.name)

	Base.fun()
	```
  * >using @staticmethod
  	```
	class Base:
		@staticsmethod
		def fun():
			#no param here
			print("ok")

	Base.fun()
	```

## Mixin
> what if you want to inherit a base class and then add functionality which are written in another class? then those fuctionalities should be written in a class called Mixin. here below is an example
```
class Base():
	def fun1(self):
		...

class Feature1Mixin():
	def featureFun1(self):
		...

Class A(Base, FeatureMixin):
	...

#A().featureFun1() can be used.
```

## getter setter using @property
```
class User:
	def __init__(self):
		self.__name = ''
	
	@property
	#it's become getter automatically
	def name(self):
		return self.__name
	
	@name.setter
	def name(self,val):
		self.__name=val
	
	@name.deleter
	def name(self):
		del self.__name
u = User()
u.name='bipul'
print(u.name)
del p.name
print(p.name) #it will trow an error saying no attribute
```

## regex
```
import re
str= 'There is a tech\nfsdf dfsf\nfdsfsdf. \n\n Ioooop\nLooop\n\n'
pattern = re.compile('\w+\n\w+', re.S)
re.sub(pattern, lambda match: match.group().replace('\n',''),str)

output >>> ''There is a techfsdf dfsffdsfsdf. \n\n IoooopLooop\n\n''
```
## The art of doing async tasks
An example
```
import asyncio

async def hello(i):
    print(f"hello {i} started")
    await asyncio.sleep(4)
    print(f"hello {i} done")

async def main(): #task3
    task1 = asyncio.create_task(hello(1)) #task 1
    task2 = asyncio.create_task(hello(2)) #task 2
    await task1
    await task2

asyncio.run(main())  # task4

>>> hello 1 started
hello 2 started
hello 1 done
hello 2 done
```
What can be run asynchronously? for example if we have two tasks and we want them to run asyncronously then we will need a 3rd task called an  event loop/event controller and we will need 4th task to start the task3. in the above example task1 and task2 are two tasks we wanted to run asyncronously and main() is the 3rd task i.e. event loop. and we are feeding the event loop in task4.  
Now what if we want to run two async tasks where one task returning a response to a client? then what will be our event loop and at what code block we will feed the event loop? there is no way in WSGI. Let's say we made the view as async still it will not be possible to run views and task2 concurrently since there's no code to build the event loop.

## Changes
* >In python3.3+ you don't need to pass the class name explicitly, you can just do:
	```
	#instead of super(MdiChild, self).__init__()
	super().__init__()
	```
* >in python 3+ class Base(object) and class Base() and class Base are same.



## REFERENCES
 1. https://stackoverflow.com/questions/49005651/how-does-asyncio-actually-work
 2. https://tenthousandmeters.com/blog/python-behind-the-scenes-1-how-the-cpython-vm-works/
 3. https://www.javatpoint.com/pointer-in-python
