#include<stdio.h>
#include<sys/io.h>
void main(){
printf("%ld",virt_to_phys(0x456845));
}
