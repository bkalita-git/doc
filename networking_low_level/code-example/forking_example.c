#include <stdio.h>
#include <unistd.h>

void main(){

	/*pid_t par_pid   = getpid();
	pid_t child_pid = fork(); //here two child_pid variable willbe created for both
	//Upon successful completion, fork() returns 0 to the child process and returns the process ID of the child process to the parent process.
	//printf("fork returned %lu and now my pid = %lu\n",(unsigned long)child_pid, (unsigned long)getpid() );
	
	if(getpid()!=par_pid){
		printf("I AM CHILD\n");
	}else{
		printf("I AM MAIN\n");
	}*/
	//char MSG[100];
	//fgets(MSG, sizeof(MSG), stdin);
	//printf("%s\n",MSG);
	
	//4 threads including the main as thread1
	/*pid_t thread1_id = getpid();
	fork();
	if(getpid()==thread1_id){
		printf("I AM THREAD1\n");
	}else{
		pid_t thread2_id = getpid();
		fork();
		if(getpid()==thread2_id){
			printf("I AM THREAD2\n");
		}else{
			pid_t thread3_id = getpid();
			fork();
			if(getpid()==thread3_id){
				printf("I AM THREAD3\n");
			}else{
				printf("I AM THREAD4\n");
			}
		}
	}*/
	
	
	//or
	/*
	pid_t thread1_id = getpid();
	fork();
	if(getpid()==thread1_id){
		fork();
		if(getpid()==thread1_id){
			fork();
			if(getpid()==thread1_id){
				printf("I AM THREAD1\n");
				
			}else{
				printf("I AM THREAD2\n");
			}
			
		}else{
			printf("I AM THREAD3\n");
		}
		
	}else{
		printf("I AM THREAD4\n");
	}*/
	
	
	//or 3threads
	pid_t thread1_id = getpid();
	fork();
	if(getpid()==thread1_id){
		fork();
		if(getpid()==thread1_id){
			printf("I AM THREAD1\n");
			
		}else{
			printf("I AM THREAD3\n");
		}
		
	}else{
		printf("I AM THREAD4\n");
	}
}
