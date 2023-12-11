%{
#include <ctype.h>
#include "tabela_simbolos.h"

extern SymbolTable symTable; // Declare the symbol table

int line_num = 1;
int no_of_words = 0;

void yyerror(char* message);
void handleToken(SymbolTable* table, char* token_name);
%}

space          [ ]*
caractere      [^"]
digit          [0-9]
letter   	   [A-Za-z_]
identifier 	   {letter}({letter}|{digit}|_)*
literal        \"{caractere}\"
constant       {digit}+
shiftop        <<|>>|<<<|>>>
mulop          \*|\/|mod|and
addop          \+|\-|or
relop          =|>|>=|<|<=|<>
assignop       :=
negop          not
modop          \%
semicolon      \;
comma          \,
open_p         \(
close_p        \)


%%
\n          	   ++line_num;

"{"[^}\n]*"}"     /* come comentários de linha única */
[ \t\n]+          /* come espaços em branco (white spaces) */;

quit         	   return 0;



program           handleToken(&symTable, "<program>");
declare           handleToken(&symTable, "<declare>");
begin             handleToken(&symTable, "<begin>");
integer           handleToken(&symTable, "<integer>");
decimal           handleToken(&symTable, "<decimal>");
if                handleToken(&symTable, "<if>");
then              handleToken(&symTable, "<then>");
else              handleToken(&symTable, "<else>");
end               handleToken(&symTable, "<end>");
while             handleToken(&symTable, "<while>");
for               handleToken(&symTable, "<for>");
to                handleToken(&symTable, "<to>");
read              handleToken(&symTable, "<read>");
write             handleToken(&symTable, "<write>");
{open_p}          handleToken(&symTable, "<open_p>");
{close_p}         handleToken(&symTable, "<close_p>");
{semicolon}       handleToken(&symTable, "<semicolon>");
{comma}           handleToken(&symTable, "<comma>");
{modop}           handleToken(&symTable, "<modop>");
{negop}           handleToken(&symTable, "<negop>");
{assignop}        handleToken(&symTable, "<assignop>");
{literal}         handleToken(&symTable, "<literal>");
{relop}	         handleToken(&symTable, "<relop>");
{addop}	         handleToken(&symTable, "<addop>");
{mulop}	         handleToken(&symTable, "<mulop>");
{shiftop}	      handleToken(&symTable, "<shiftop>");
{constant}	      handleToken(&symTable, "<constant>");
{digit}{letter}*  yyerror("invalid input");
{identifier}	   handleToken(&symTable, "<identifier>");

.                 yyerror("invalid input");
%%
 
int yywrap(){}

void yyerror(char* message)
{
   fprintf(stderr, "Error: %s in line %d: \"%s\"\n", message, line_num, yytext);
}

void handleToken(SymbolTable* table, char* token_name)
{
   // printf("SCANNER: Adding \"%s\" to symbom table\n", token_name);
   insertSymbol(&symTable, token_name);
}


extern FILE * yyin;
SymbolTable symTable;

void main(int argc, char** argv)
{
   initSymbolTable(&symTable, 10);

    if (argc > 1)
    {
      FILE *file;
      file = fopen(argv[1], "r");      
      if (!file)
      {
         fprintf(stderr, "Could not open %s\n", argv[1]);
         exit(1);
      }
      yyin = file;
    }

    yylex();

    printSymbolTable(&symTable);

    cleanupSymbolTable(&symTable);
}