#include <stdio.h>

int main() {
    /* Strings and pointers */
    char *ptr = "hello!";
    printf("string from pointer: %s\n", ptr);
    printf("string memory address: %p\n", &ptr);
    printf("first char from pointer dereference: %c\n", *ptr);
    printf("second char from pointer operation dereference: %c\n", *(ptr + 1));
    printf("no parenthesis in pointer operation dereference: %c\n", *ptr + 1);
    printf("\texplanation:\n");
    printf("\tletter 'h' (ASCII 104) + 1 = 'i' (ASCII 105)\n");
    printf("first char in ASCII: %d (using \%%d)\n", *ptr);

    printf("\n-----------\n\n");

    /* Strings and arrays */
    char arr[] = "bye!";
    printf("string from an array: %s\n", arr);
    printf("first char from an array: %c\n", arr[0]);
    printf("second char from an array: %c\n", arr[1]);
    printf("third char from inverse array access: %c\n", 2[arr]);
    printf("\texplanation:\n");
    printf("\t2[arr] and arr[2] yield the same result, as it translates to *(2 + arr) and *(arr + 2)\n");
    // Reference: https://en.cppreference.com/w/c/language/operator_member_access.html

    return 0;
}
