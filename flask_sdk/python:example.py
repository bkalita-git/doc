'''
main_event_loop(coroutine_obj_1):
	list_tasks = [(coroutine_obj_1,None)]
	while list_tasks:
		queue, list_tasks = list_tasks, []
		for task, data in queue:
			data = task.send(data)
			if returns data:
				print data
			else:
				if data an object:
					list_tasks.append(data) #data is here coroutine_obj_2
				list_tasks.append(task)		#task is parent task, here, coroutine_obj_1
			
coroutine_obj_1:
	#coroutine_obj_2 is child coroutine
	yield coroutine_obj_2 [pause and return coroutine_obj_2 and in main_event_loop: list_tasks.append(coroutine_obj_2)]
	print
	
coroutine_obj_2:
	yield 1 [pause and return 1]
	print
'''
def level0_coroutine_obj1():
	yield 2
	print("i am level 0")

def level1_coroutine_obj1():
	val = yield level0_coroutine_obj1
	print(" i am level 1",val)



def main_event_loop(level1_coroutine_obj1):
	list_tasks = [(level1_coroutine_obj1,None)]
	while list_tasks:
		#print("list_tasks =  ",list_tasks)
		queue, list_tasks = list_tasks, []
		#print("queue = ",queue)
		for task,data in queue:
			print("insider queue= ",queue)
			print("task = ", task)
			#print("data = ", data)
			data = task.send(data)
			#print("result data = ", data)
			if type(data) is int:
				#add to parent
				#list_tasks.append((,data))
				pass
			else:
				list_tasks.append((data(),None))
				
x = level1_coroutine_obj1()
main_event_loop(x)



