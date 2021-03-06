
## A minimal  example

```
import 'package:flutter/material.dart';
void main()=>runApp(Text("ok",textDirection: TextDirection.ltr));
```

in the above example runApp is a function which takes a widget and make it root.
the runApp will not work if we don't import either flutter/material.dart or flutter/cupertino.dart
again, if you  did create a project using flutter command, then you  will see an error  in App_directory/test/widget_test.dart saying MyApp is not defined, it's due to not making MyApp the root, so it's expecting runApp(MyApp()) 

then What are widget? They are class extending from StatelessWidget, StatefullWidget. widget have a build method.
so, material.dart or cupertino.dart must have implement some basic widget that's  why Text widget  is asking for impoting a library.
Now, material library is use in android  and cupertino is used in ios basically.
whenever we use material.dart it forces us to use components from material library and it's a good idea so we don't mixup material with other cupertino s components.
so, forcing material icons to use we can set 
```
name: my_app
   flutter:
  	uses-material-design: true
```

in pubspec.yaml file. but by default it is automatically set when project is created using flutter command

Now, we can force our root widget to be MaterialApp,
```
void main()=>runApp(
				MaterialApp(
					home:Text("ok",textDirection: TextDirection.ltr)
				)
			);
```
but we can remove the textDirection field from the Text widget
```
void main()=>runApp(MaterialApp(MaterialApp(home:Text("ok"))));
```
but why? because many widget can take advantage from this root widget. so, textDirection is automatically set using the help of the MaterialApp widget.  For a Material app, you can use a Scaffold widget; it provides a default banner, background color, and has API for adding drawers, snack bars, and bottom sheets.like -- 
```
MaterialApp(...home:Scaffold(...)..
```

So, widgets are class extended from StatelessWidget  or  StatefullWidget
a pseudocode
```
class Widget1 extends StatelessWidget{
	@override
	Widget build(BuildContext context)=>Text("This is a new Widget");
}
```
								vs
```
class Widget2 extends StatefullWidget{
	@override
	_Widget2state createState()=> _Widget2State();
}
class _Widget2State extends State<Widget2>{
	@override
	Widget build(BuildContext  context)=>Text("This is a new Widget");
}
```		
from the above example,  why StatefullWidget and State are separate object? they have different life cycle.  when adding the StatefullWidget to a tree, then first createState will be called and it creates an instance of the widget.

Key property in widget:
	A Key is an identifier for Widgets, Elements and SemanticsNodes.A new widget will only be used to update an existing element if its key is the same as the key of the current widget associated with the element

for debugging widget layout use  debugPaintSizeEnabled set to true, so  boundary will get displayed



what are material components? are widget are the only components?
for example GestureDetector widget does not have a visual representation but it detects gesture.
Use a Container when you want to add padding, margins, borders, or background color, to name some of its capabilities. Container adds properties to its child.
consider padding as the widget gets bigger and margins as widget gets further.
the area assigned to an widget. the first visible  root widget takes all the place of the screen. But when resizing, the area get shrinked but inside widget may not, example, the text widget inside container with background color. but what if we don't want our area to get shrinked or grow. **Introducing constraints.
a container widget without height and width wants always to be full screen or the same size as its child except its parent set a rule/constraints to it.** 
***for example, if container is the root widget then screen is the parent and it will force the container to be full size, so width height of container wont work.***
what if the container is inside a center widget and center is the root, then center will be full screen but center widget allows the container to be any size it wants, so width height will work.

```
//determine 
1. is the parent ignored its height or width of the child? if so that means parent does not know how to align it. so, use the child inside a Align(). but still not working? so many widgets, layout, visible, align etc..
2. scaffold is ignoring button but not container height
```

don't use like this
```
Container(color:Colors.red, decoration:BoxDecoration()), 
```
since color is inside decoration this will produce error, instead use this
```
Container(decoration:BoxDecoration(color:Colors.red))
```
lly,
don't use borderRadius and shape inside BoxDecoration

## Different Type of Widget
>it's all about property of elements. Different types of widget sets different properties.
* Visible Widget
* Layout Widget
	* Single Child
		* Align
		* Container
		* Center
		* Expanded
		* SizedBox
		* Transform
	* Multi Child
		* Column
		* ListView
		* Row

```
Layouts  widget:
	standard: Container, GridView, ListView, Stack
	Material: Card,ListTile
Interactive prefabricated widgets: They all extends the StatefullWidget
	standard: Form,FormField
	Material: Checkbox, DropdownButton, TextButton, FloatingActionButton, IcoButton, Radio, ElevatedButton, Slider, Switch, TextField


	class Foo extends StatelessWidget {
	  @override
	  Widget build(BuildContext context) => Text('foo');
	}

	Center(child: Foo())

	//Could equally well be defined and used like this, without defining a new widget class:

	Center(
	  child: Builder(
		builder: (BuildContext context) => Text('foo');
	  ),
	)
```

A BuildContext is nothing else but a reference to the location of a Widget within the tree structure of all the Widgets which are built.
BuildContext instance is the Parent Widget and it is passed/given/called to the child build method by the framework(yes, we don't call it  explicitly) and the  widget can read that instance to findout where it will be attached. 

eg:
```
Scaffold(
Scaffold.of(context).showSnackBar(a_defined_snackBar) 
//since the context is the parent of Scaffold it will not find any so error
)
```

solved by
```
Scaffold(
Builder(
builder:(context) => Scaffold.of(context).showSnackBar(a_defined_snackBar)  
//now context is the scaffold and it will work i.e. will find the Scaffold
)
)
```

WHO OWNS A CONTEXT or TO WHOME A CONTEXT IS GIVEN
eg:
```
 runApp(
    MaterialApp(
      home: Text(Theme.of(context).accentColor.value.toString()), //here's no context so error
    )
  );
```
eg:
```
void main() {
  // debugRepaintRainbowEnabled = true;
  runApp(
    MaterialApp(
      home: MyApp()
    )
  );
}

class MyApp extends StatelessWidget{
  @override 
  Widget build(BuildContext context){
    return Text(Theme.of(context).accentColor.value.toString()); //here's no context error since MaterialApp is the context
  }
} 
```
or simply using a Builder Widget 
```
runApp(
    MaterialApp(
      home: Builder(builder: (context)=>Text(Theme.of(context).accentColor.value.toString()))
    )
);
```

remember: StatefullWidget have many many options, 
Key and controller and context
key to make some actions on various widget
controller to make some actions on single widget
context to popup widget or access data from inherited widget

the : or initializer list is used to do some task with the incoming data in the constructor and after that do the task of constructor.
```
class Car(){
	final int tires;
	Car():tires=3;	//works since tires=3 will execute first before creating the constructor
}

class Car(){
	final int tires;
	Car(){
		tires=3;   //error since tire is final
	}
}
```

assert keyword
dart --enable-asserts example.dart //flutter does in automatically in development mode, in production it is ignored
so,  assert(evenNo % 2 == 0); will give assert error if it is false. and pause the code

ValueNotifier<T> + ValueListenableBuilder<T> + InheritedWidget:
	pass the ValueNotifier instance in the InheritedWidget Constructor and then use ValueListenableBuilder to build the unstable widget
	
Animation Widget:
	Drawbased : something like flying bird, svg images
	Codebased : eg animate a container rotation	
		implicit: using setState  you can make a container bigger but if you used AnimatedContainer with a duration and curve property then the container size will be animated
			Builtin
			TweenAnimationBuilder: creates a custom implicit animation
		explicit: requires animation controller and if animation goes forever=eg while loop  ,discontinuous= circle from small to large, sm to lrg, ... ,multiple widget are animated together in a coordinated fashion
			Builtin
			Customexplicit
				AnimatedWidger :
				AnimatedBuilder: 
			CustomPainer(for every explicit for performance)

CUSTOMPAINT


SUMMARY:  
	#Using key you can access values of stateful widget. just assign the key to the Widget key: your_key for eg, 
		1. accessing the user input using key
		   final GlobalKey<FormFieldState<string>> emailKey = new GlobalKey<FormFieldState<string>>();
		   var email = emailKey.currentState.value
		   
		2. Validating multiple TextFormField , a validator function must be in every TextFormField
		   final GlobalKey<FormState> formKey = new GlobalKey<FormState>();
		   formKey.currentState.validate()
   
   #using controller you can access widgets properties and add listener to it, just assign the controller to the widget, controller:your_controller. using controller you cam make something happen without using setstate. ge: play a video with the controller.play()
   		1. accessing the user input using controller
   		   final TextEditingController _controller = new TextEditingController();
   		   _controller.text
	    2. adding a listener to the controller
	       _controller.addListener((){_controller.text.toLowerCase()});
   
	#using ValueNotifier and ValueListenableBuilder to setState automatically on value changed
	#using ChangeNotifier to add Listner
	#using InheritedWidget to passdown a object and accessing using context 
	#using FutuerBuilder setState automatically on value changed
	#using StreamBuilder setState automatically on value changed
	#using Builder() widget to create a new stateful widget
	#using context to show dialogs,  so widget needs a context to showup
	
	WHAT TO DO:
		if it is not related to future then use ValueNotifier
		
REFERENCES
1. https://flutter.dev/docs/development [explanation]  
2. https://api.flutter.dev/index.html   [code]
3. https://flutter.dev/docs/development/ui/layout/constraints
4. https://api.flutter.dev/flutter/widgets/widgets-library.html
5. https://flutter.dev/docs/development/ui/widgets
6. https://material.io/components?platform=flutter
7. https://flutter.dev/docs/resources/architectural-overview
8. https://liefery-it-legacy.github.io/blog/2019/02/18/flutter-for-newbies-why-you-should-care-about-the-build-context.html



