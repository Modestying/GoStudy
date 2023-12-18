#ifndef DEMO_H
#define DEMO_H

typedef struct __Student
{
    int Age;
    char Name[20];
    void * No;
}Student;


void processInt(int num);
void processIntPtr(int *num);


void processUnsignedInt(unsigned int num);
void processUnsignedIntPtr(unsigned int *num);

void processChar(char str);
void processCharPtr(char *str);
void processConstCharPtr(const char* str);


void processUnsignedChar(unsigned char data);
void processUnsignedCharPtr(unsigned char *data);
void processConstUnsignedCharPtr(const unsigned char *data);


void processStruct( Student stu);
void processStructPtr( Student *stu,int *length);

void processVoidPtr(void *);
void processVoidPtrPtr(void **);

#endif
