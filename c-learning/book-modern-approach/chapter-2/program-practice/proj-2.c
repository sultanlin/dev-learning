#include <stdio.h>
#include <math.h>

#define SPHERE_VOLUME (4.0f/3.0f) * M_PI

int main(void){
    int radius = 10;
    int volume = radius * radius * radius * SPHERE_VOLUME;
    printf("%d\n", volume);
}
