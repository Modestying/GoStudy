#include <stdio.h>
#include "demo.h"
#include <stdlib.h>
#include <string.h>

void processInt(int num) {
    printf("processInt: %d\n", num);
}

void processIntPtr(int *num) {
    printf("processIntPointer: %d\n", *num);
}

void processUnsignedInt(unsigned int num) {
    printf("processUnsignedInt: %u\n", num);
}

void processUnsignedIntPtr(unsigned int *num) {
    printf("processUnsignedIntPtr: %u\n", *num);
}

void processChar(char str) {
    printf("processChar: %c\n", str);
}

void processCharPtr(char *str) {
    printf("processCharPtr: %s\n", str);
}

void processConstCharPtr(const char* str){
    printf("processConstCharPtr:%s\n",str);
}

void processUnsignedChar(unsigned char str) {
    printf("processUnsignedChar: %c\n", str);
}

void processUnsignedCharPtr(unsigned char *str) {
    printf("processUnsginedCharPtr: %s\n", str);
}
void processConstUnsignedCharPtr(const unsigned char *data){
    printf("processConstUnsignedCharPtr: %s\n", data);
}

void processStruct( Student stu) {
    printf("processStruct: %d, %s\n", stu.Age, stu.Name);
}

void processStructPtr(Student *stu,int length) {
    printf("processStructPtr:\n");
    for (int i = 0; i < length; i++) {
        stu[i].Age = i+10;
        strcpy(stu[i].Name,"你好啊");
    }
    // for (int i=0;i<length;i++){
    //     printf("name:%s\n",stu[i].Name);
    // }
}

void processVoidPtr(void * data){
    int *intPtr=(int*)data;
    printf("processVoidPtr:%d\n",*intPtr);
    *intPtr=100;
}

void processVoidPtrPtr(void **ptr){
    printf("processVoidPtrPtr\n");
    if (*ptr!=NULL){
        printf("ptr val:%d\n",*(int*)(*ptr));
        // 重新赋值，这里主动释放内存
        free(*ptr);
        *ptr=NULL;
    }
    // value 不能释放
    int *value = (int *)malloc(sizeof(int));
    *value = 42;
    *ptr = value;
}

