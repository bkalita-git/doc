#include<stdio.h>

void main(){

/*
Data_Type 								Memory(bytes)		Range 				Format_Specifier 
________________________________________________________________________________________________________________________________________________
char/signed char 								1 	-128 	       to 127 				%c 
unsigned char 									1 	 0 	       to 255 				%c 
short/signed short/signed short int/short int 					2 	-32,768        to 32,767 			%hd 
unsigned short/unsigned short int						2 	 0 	       to 65,535 			%hu 
int/signed int 									4 	-2,147,483,648 to 2,147,483,647 		%d 
unsigned int									4 	 0 	       to 4,294,967,295 		%u 
long/long int/signed long/signed long int 					4 	-2,147,483,648 to 2,147,483,647 		%ld 
unsigned long/unsigned long int 						4 	 0 	       to 4,294,967,295 		%lu 
long long/long long int/signed long long/signed long long int 			8 	-(2^63)        to (2^63)-1 			%lld 
unsigned long long/unsigned long long int 					8 	 0 	       to 18,446,744,073,709,551,615 	%llu 
float 										4 	 1.2E-38       to 3.4E+38			%f 
double										8 	 2.3E-308      to 1.7E+308			%lf 
long double 									16 	 3.4E-4932     to 1.1E+4932			%Lf 
*/


/*
Format_specifier	Description					Supported data types
____________________________________________________________________________________________
%c			Character					char/unsigned char
%d			Signed Integer					short/unsigned short/int/long
%e or %E		Scientific notation of float values		float/double
%f			Floating point					float
%g or %G		Similar as %e or %E				float/double
%hi			Signed Integer(Short)				short
%hu			Unsigned Integer(Short)				unsigned short
%i			Signed Integer					short/unsigned short/int/long
%l or %ld or %li	Signed Integer					long
%lf			Floating point					double
%Lf			Floating point					long double
%lu			Unsigned integer				unsigned int/unsigned long
%lli, %lld		Signed Integer					long long
%llu			Unsigned Integer				unsigned long long
%o			Octal representation of Integer			short/unsigned short/int/unsigned int/long
%p			Address of pointer to void void *		void *
%s			String						char *
%u			Unsigned Integer				unsigned int/unsigned long
%x or %X		Hexadecimal representation of Unsigned Integer	short/unsigned short/int/unsigned int/long
%n			Prints nothing	
%%			Prints % character	
*/


/*
udp header
  0      7 8     15 16    23 24    31
 +--------+--------+--------+--------+
 |          source address           |
 +--------+--------+--------+--------+
 |        destination address        |
 +--------+--------+--------+--------+
 |  zero  |protocol|   UDP length    |
 +--------+--------+--------+--------+
*/


/*
tcp header
0                   1                   2                   3
0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|          Source Port          |       Destination Port        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                        Sequence Number                        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                    Acknowledgment Number                      |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|  Data |           |U|A|P|R|S|F|                               |
| Offset| Reserved  |R|C|S|S|Y|I|            Window             |
|       |           |G|K|H|T|N|N|                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|           Checksum            |         Urgent Pointer        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                    Options                    |    Padding    |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             data                              |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
*/

/*
ip header https://tools.ietf.org/html/rfc791
0                   1                   2                   3
0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|Version|  IHL  |Type of Service|          Total Length         |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|         Identification        |Flags|      Fragment Offset    |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|  Time to Live |    Protocol   |         Header Checksum       |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                       Source Address                          |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                    Destination Address                        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                    Options                    |    Padding    |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
*/
printf("%X\n",10);
}
