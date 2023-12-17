#include <stdio.h>
#define F_TO_C (5.0f/9.0f)
#define FREEZING_PT 32.0f

int main (void){
    float fahrenheit;

    printf("Enter the fahrenheit value: ");
    scanf("%f", &fahrenheit);
    float celsius = F_TO_C * (fahrenheit - FREEZING_PT);
    printf("The converted celsius is: %.2f \n", celsius);

    return 0;
}
