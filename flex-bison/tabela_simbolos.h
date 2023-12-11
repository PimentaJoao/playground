#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
    char* name;
} Token;

typedef struct {
    Token* tokens;
    size_t size;
    size_t capacity;
} SymbolTable;

// Initialize the symbol table
void initSymbolTable(SymbolTable* table, size_t initialCapacity) {
    table->tokens = malloc(initialCapacity * sizeof(Token));
    table->size = 0;
    table->capacity = initialCapacity;
}

// Insert a new entry into the symbol table
void insertSymbol(SymbolTable* table, const char* name) {
    
    // Resize the symbol table if necessary
    if (table->size >= table->capacity) {
        table->capacity *= 2;
        table->tokens = realloc(table->tokens, table->capacity * sizeof(Token));
    }

    Token newToken;
    newToken.name = strdup(name);

    table->tokens[table->size++] = newToken;
}

// Cleanup the symbol table
void cleanupSymbolTable(SymbolTable* table) {
    for (size_t i = 0; i < table->size; ++i) {
        free(table->tokens[i].name);
        // Free other attributes as needed
    }
    free(table->tokens);
}

// Print all entry names in the symbol table
void printSymbolTable(const SymbolTable* table) {
    printf("Symbol Table Entries:\n");
    for (size_t i = 0; i < table->size; ++i) {
        printf("%s\n", table->tokens[i].name);
    }
}