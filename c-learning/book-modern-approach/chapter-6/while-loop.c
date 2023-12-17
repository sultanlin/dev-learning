#include <stdio.h>

int main(void){

    int i = 0;

    do{
        i++;
        printf("%d", i);
    }while (i < 100);

    for(int i = 10; i > 0; i--){
        printf("hello %d", i);
    }

    return 0;
}
