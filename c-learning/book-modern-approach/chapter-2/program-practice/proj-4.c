#include <stdio.h>

int main(void){
    float loan, intrate, monpay;

    printf("Enter amount of loan: ");
    scanf("%f", &loan);
    printf("Enter interest rate: ");
    scanf("%f", &intrate);
    printf("Enter monthly payment: ");
    scanf("%f", &monpay);

    loan = loan - (loan * intrate/100) - monpay;
    printf("Balance remaining after first payment: %f\n", loan);
    loan = loan - (loan * intrate/100) - monpay;
    printf("Balance remaining after first payment: %f\n", loan);
    loan = loan - (loan * intrate/100) - monpay;
    printf("Balance remaining after first payment: %f\n", loan);

    return 0;
}
