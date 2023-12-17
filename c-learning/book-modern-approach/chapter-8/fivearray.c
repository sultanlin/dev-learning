// read 5x5 array and print row sums and column sums
// enter row x:
// repeat x5
// row sum:
// column sum:
#include <stdio.h>

#define ARRAYSIZE 5

int main(void) {
    int fbf[ARRAYSIZE][ARRAYSIZE];

    for (int i = 0; i<ARRAYSIZE ; i++) {
        printf("enter row %d (ex. 14 23 25 234 1): ", i);
        scanf("%d %d %d %d %d", &fbf[i][0], &fbf[i][1], &fbf[i][2], &fbf[i][3], &fbf[i][4]);
    }

    int rowsum, columnsum;
    printf("row sum: ");
    for (int i=0; i<ARRAYSIZE; i++) {
        rowsum=0;
        for (int j=0; j<ARRAYSIZE; j++) {
            rowsum += fbf[i][j];
        }
        printf("%d ", rowsum);
    }
    
    printf("\ncolumn sum: ");
    for (int i=0; i<ARRAYSIZE; i++) {
        columnsum=0;
        for (int j=0; j<ARRAYSIZE; j++) {
            columnsum += fbf[j][i];
        }
        printf("%d ", columnsum);
    }
    printf("\n");

    return 0;
}
