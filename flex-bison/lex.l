%{
#include <ctype.h>
#include "y.tab.h"

int line_num = 1;

void yyerror(char* message);
// void handleToken(SymbolTable* table, char* token_name);
%}

caractere  [^"]
digit      [0-9]
letter     [A-Za-z]
identifier {letter}({letter}|{digit})*
literal    \"{caractere}\"
constant   {digit}+

%%
\n           {++line_num;}
[ \t]+       /* Ignora espaços em branco (white spaces) */
"%"[^\n]*    /* Ignora comentários em linha única (iniciados por %) */
\"[^\"\n]*\" /* Ignora cadeias dentro de aspas duplas */

"quit"       {return 0;}

program      {return PROGRAM;}
declare      {return DECLARE;}
begin        {return BGN;}
integer      {return INTEGER;}
decimal      {return DECIMAL;}
if           {return IF;}
then         {return THEN;}
else         {return ELSE;}
end          {return END;}
do           {return DO;}
while        {return WHILE;}
for          {return FOR;}
to           {return TO;}
read         {return READ;}
write        {return WRITE;}
not          {return NOT;}
"\;"         {return SEMICOLON;}
"\:"         {return COLON;}
"\,"         {return COMMA;}
"\("         {return OPEN_P;}
"\)"         {return CLOSE_P;}
":="         {return ASSIGN;}
"="          {return EQ;}
">"          {return GT;}
">="         {return GTE;}
"<"          {return LT;}
"<="         {return LTE;}
"<>"         {return NEQ;}
"?"          {return CONDITIONAL;}
"\*"         {return MUL_OP;}
"\/"         {return DIV_OP;}
"\+"         {return PLUS_OP;}
"\-"         {return MINUS_OP;}
mod          {return MOD_OP;}
and          {return AND_OP;}
or           {return OR_OP;}
"<<"         {return LS_OP;}
">>"         {return RS_OP;}
"<<<"        {return LSF_OP;}
">>>"        {return RSF_OP;}
{literal}    {return LITERAL;}
{constant}	 {return CONSTANT;}
{identifier} {return IDENTIFIER;}

.            yyerror("unknown input");
%%
 
int yywrap()
{
   return 1;
}

void yyerror(char* message)
{
   fprintf(stderr, "Error: %s in line %d: \"%s\"\n", message, line_num, yytext);
}

// void handleToken(SymbolTable* table, char* token_name)
// {
//    // printf("SCANNER: Adding \"%s\" to symbom table\n", token_name);
//    insertSymbol(&symTable, token_name);
// }


// extern FILE * yyin;
// SymbolTable symTable;
