%{
#include <stdio.h>

void yyerror(char *message);
int yylex(void);
%}

%token PROGRAM DECLARE BEGIN INTEGER DECIMAL IF THEN ELSE END DO WHILE FOR TO READ WRITE NOT
%token SEMICOLON COLON COMMA OPEN_P CLOSE_P ASSIGN
%token EQ GT GTE LT LTE NEQ CONDITIONAL
%token MUL_OP DIV_OP PLUS_OP MINUS_OP MOD_OP AND_OP OR_OP LS_OP RS_OP LSF_OP RSF_OP
%token LITERAL CONSTANT IDENTIFIER

%left EQ GTE LTE LT GT NEQ
%left PLUS_OP MINUS_OP OR_OP
%left MULTI_OP DIV_OP MOD_OP AND_OP
%left NOT CONDITIONAL

%%

program: PROGRAM IDENTIFIER body;

body: DECLARE decl_list BEGIN stmt_list END;

decl_list: decl SEMICOLON decl_list_aux;

decl_list_aux: decl SEMICOLON decl 
             |
             ;

decl: type ident_list;

ident_list: IDENTIFIER ident_list_aux;

ident_list_aux: COMMA IDENTIFIER ident_list_aux 
              |
              ;

type: INTEGER
    | DECIMAL
    ;

stmt_list: stmt SEMICOLON stmt_list_aux;

stmt_list_aux: stmt SEMICOLON stmt_list_aux 
             |
             ;

stmt: assign_stmt
    | if_stmt
    | while_stmt
    | read_stmt
    | write_stmt
    ;

assign_stmt: IDENTIFIER ASSIGN simple_expr;

if_stmt: IF condition THEN stmt_list END 
       | IF condition THEN stmt_list ELSE stmt_list END
       ;

while_stmt: DO stmt_list stmt_suffix;

stmt_suffix: WHILE condition;

read_stmt: READ OPEN_P IDENTIFIER CLOSE_P;

write_stmt: WRITE OPEN_P writable CLOSE_P;

writable: simple_expr 
        | LITERAL
        ;

condition: expression;

expression: simple_expr expression_aux;

expression_aux: EQ simple_expr
              | GTE simple_expr
              | LTE simple_expr
              | LT simple_expr
              | GT simple_expr
              | NEQ simple_expr
              ;

simple_expr: term 
           | simple_expr PLUS_OP term
           | simple_expr MINUS_OP term
           | simple_expr OR_OP term
           | OPEN_P simple_expr CLOSE_P CONDITIONAL simple_expr COLON simple_expr
           ;

term: factor_a 
    | term MUL_OP factor_a
    | term DIV_OP factor_a
    | term MOD_OP factor_a
    | term AND_OP factor_a
    ;

factor_a: factor 
        | NOT factor 
        | MINUS_OP factor
        ;

factor: IDENTIFIER
      | CONSTANT
      | OPEN_P expression CLOSE_P
      ;
%%

int main() {
    yyparse();
    return 0;
}