%{
#include <stdio.h>

void yyerror(char *message);
int yylex(void);
%}

%token PROGRAM DECLARE BGN INTEGER DECIMAL IF THEN ELSE END DO WHILE FOR TO READ WRITE;
%token NOT SEMICOLON COLON COMMA OPEN_P CLOSE_P ASSIGN EQ GT GTE LT LTE NEQ CONDITIONAL;
%token MUL_OP DIV_OP PLUS_OP MINUS_OP MOD_OP AND_OP OR_OP LS_OP RS_OP LSF_OP RSF_OP LITERAL;
%token CONSTANT IDENTIFIER;

%left PLUS_OP MINUS_OP
%left MUL_OP DIV_OP MOD_OP
%left AND_OP
%left OR_OP
%nonassoc LT GT LTE GTE EQ NEQ
%left NOT
%nonassoc IF
%nonassoc ELSE
%nonassoc COLON
%nonassoc CONDITIONAL
%left OPEN_P CLOSE_P

%start program

%%

program: PROGRAM IDENTIFIER body;

body: DECLARE decl_list BGN stmt_list END;

decl_list: decl SEMICOLON decl_list_star;

decl_list_star: decl SEMICOLON decl_list_star 
             |
             ;

decl: type ident_list;

ident_list: IDENTIFIER ident_list_star;

ident_list_star: COMMA IDENTIFIER ident_list_star 
              |
              ;

type: INTEGER
    | DECIMAL
    ;

stmt_list: stmt SEMICOLON stmt_list_star;

stmt_list_star: stmt SEMICOLON stmt_list_star 
             |
             ;

stmt: assign_stmt
    | if_stmt
    | do_while_stmt
    | read_stmt
    | write_stmt
    ;

assign_stmt: IDENTIFIER ASSIGN simple_expr;

if_stmt: IF condition THEN stmt_list END 
       | IF condition THEN stmt_list ELSE stmt_list END
       ;

do_while_stmt: DO stmt_list stmt_suffix;

stmt_suffix: WHILE condition;

read_stmt: READ OPEN_P IDENTIFIER CLOSE_P;

write_stmt: WRITE OPEN_P writable CLOSE_P;

writable: simple_expr 
        | LITERAL
        ;

condition: expression;

expression: simple_expr
          | simple_expr EQ simple_expr
          | simple_expr GTE simple_expr
          | simple_expr LTE simple_expr
          | simple_expr LT simple_expr
          | simple_expr GT simple_expr
          | simple_expr NEQ simple_expr
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

int main(int argc, char** argv)
{
    yyparse();
    return 0;
}
