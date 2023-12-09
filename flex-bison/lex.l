%{
int line_num = 0;
int no_of_words = 0;

void yyerror(char *message);
%}

caractere	   [^"]
digit          [0-9]
letter   	   [A-Za-z_]
identifier 	   ({letter}({letter}|{digit})*)
constant       {digit}+
shiftop        (<<|>>|<<<|>>>)
mulop          (\*|\/|mod|and)

%%
end         	return 0;

if				   printf("<IF>\n");
{mulop}	      printf("<MULOP>\n");
{shiftop}	   printf("<SHIFTOP>\n");
{constant}	   printf("<CONST>\n");
{identifier}	printf("<IDENTIFIER>\n");

\n          	++line_num;
.              yyerror("invalid!");
%%
 
/*** User code section***/
int yywrap(){}

void yyerror(char *message)
{
   fprintf(stderr, "Error: \"%s\" in line %d.\n", message, line_num);
   exit(1);
}

int main() {
	yylex();

	return 0;
}
