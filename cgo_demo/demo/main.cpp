#include "demo.h"
#include <stdio.h>
#include<stdlib.h>
#include<string.h>
void testData(){
    int a = 0;
    
    processInt(a);
    // 不会修改a的值
    printf("processInt result = %d\n\n", a);

    processIntPtr(&a);
    // 修改指针，会修改a的值
    printf("processIntPtr result= %d\n\n", a);

    unsigned int b = 2;

    processUnsignedInt(b);
    printf("processUnsignedInt result= %d\n\n", b);

    processUnsignedIntPtr(&b);
    printf("processUnsignedIntPtr result= %d\n\n", b);


    char c = 'a';
    processChar(c);
    printf("processChar result= %c\n\n", c);

    char *d = "hello";
    processCharPtr(d);
    printf("processCharPtr result= %s\n\n", d);

    processConstCharPtr(d);
    printf("processConstCharPtr result= %s\n\n", d);

    unsigned char f = 'a';
    processUnsignedChar(f);
    printf("processUnsignedChar result= %c\n\n", f);

    processUnsignedCharPtr(&f);
    printf("processUnsignedCharPtr result= %c\n\n", f);

    unsigned char *g = &f;
    processConstUnsignedCharPtr(g);
}

Student stu[10];

void printStudent(Student stu){
    printf("name:%s, age:%d no:%d \n", stu.Name, stu.Age,*((int *)stu.No));
}
void TestStruct(){
    Student s;
    s.Age = 10;
    strcpy(s.Name, "hello");
    void *no=malloc(sizeof(int));
    *((int *)no)=1;
    s.No=no;
    processStruct(s);

    int number=5;
    processStructPtr(stu,&number);
    printf("real number:%d\n",number);
    for (int i = 0; i < number; i++)
    {   
        printStudent(stu[i]);
    }
}

void TestVoidPtr(){
    int a = 10;
    void *p = &a;
    processVoidPtr(p);
    processVoidPtrPtr(&p);
}
int main(){
    TestStruct();
    return 0;
}
