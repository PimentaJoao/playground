/* A Bison parser, made by GNU Bison 3.8.2.  */

/* Bison interface for Yacc-like parsers in C

   Copyright (C) 1984, 1989-1990, 2000-2015, 2018-2021 Free Software Foundation,
   Inc.

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.  */

/* As a special exception, you may create a larger work that contains
   part or all of the Bison parser skeleton and distribute that work
   under terms of your choice, so long as that work isn't itself a
   parser generator using the skeleton or a modified version thereof
   as a parser skeleton.  Alternatively, if you modify or redistribute
   the parser skeleton itself, you may (at your option) remove this
   special exception, which will cause the skeleton and the resulting
   Bison output files to be licensed under the GNU General Public
   License without this special exception.

   This special exception was added by the Free Software Foundation in
   version 2.2 of Bison.  */

/* DO NOT RELY ON FEATURES THAT ARE NOT DOCUMENTED in the manual,
   especially those whose name start with YY_ or yy_.  They are
   private implementation details that can be changed or removed.  */

#ifndef YY_YY_BISON_TAB_H_INCLUDED
# define YY_YY_BISON_TAB_H_INCLUDED
/* Debug traces.  */
#ifndef YYDEBUG
# define YYDEBUG 0
#endif
#if YYDEBUG
extern int yydebug;
#endif

/* Token kinds.  */
#ifndef YYTOKENTYPE
# define YYTOKENTYPE
  enum yytokentype
  {
    YYEMPTY = -2,
    YYEOF = 0,                     /* "end of file"  */
    YYerror = 256,                 /* error  */
    YYUNDEF = 257,                 /* "invalid token"  */
    PROGRAM = 258,                 /* PROGRAM  */
    DECLARE = 259,                 /* DECLARE  */
    BEGIN = 260,                   /* BEGIN  */
    INTEGER = 261,                 /* INTEGER  */
    DECIMAL = 262,                 /* DECIMAL  */
    IF = 263,                      /* IF  */
    THEN = 264,                    /* THEN  */
    ELSE = 265,                    /* ELSE  */
    END = 266,                     /* END  */
    DO = 267,                      /* DO  */
    WHILE = 268,                   /* WHILE  */
    FOR = 269,                     /* FOR  */
    TO = 270,                      /* TO  */
    READ = 271,                    /* READ  */
    WRITE = 272,                   /* WRITE  */
    NOT = 273,                     /* NOT  */
    SEMICOLON = 274,               /* SEMICOLON  */
    COLON = 275,                   /* COLON  */
    COMMA = 276,                   /* COMMA  */
    OPEN_P = 277,                  /* OPEN_P  */
    CLOSE_P = 278,                 /* CLOSE_P  */
    ASSIGN = 279,                  /* ASSIGN  */
    EQ = 280,                      /* EQ  */
    GT = 281,                      /* GT  */
    GTE = 282,                     /* GTE  */
    LT = 283,                      /* LT  */
    LTE = 284,                     /* LTE  */
    NEQ = 285,                     /* NEQ  */
    CONDITIONAL = 286,             /* CONDITIONAL  */
    MUL_OP = 287,                  /* MUL_OP  */
    DIV_OP = 288,                  /* DIV_OP  */
    PLUS_OP = 289,                 /* PLUS_OP  */
    MINUS_OP = 290,                /* MINUS_OP  */
    MOD_OP = 291,                  /* MOD_OP  */
    AND_OP = 292,                  /* AND_OP  */
    OR_OP = 293,                   /* OR_OP  */
    LS_OP = 294,                   /* LS_OP  */
    RS_OP = 295,                   /* RS_OP  */
    LSF_OP = 296,                  /* LSF_OP  */
    RSF_OP = 297,                  /* RSF_OP  */
    LITERAL = 298,                 /* LITERAL  */
    CONSTANT = 299,                /* CONSTANT  */
    IDENTIFIER = 300,              /* IDENTIFIER  */
    MULTI_OP = 301                 /* MULTI_OP  */
  };
  typedef enum yytokentype yytoken_kind_t;
#endif

/* Value type.  */
#if ! defined YYSTYPE && ! defined YYSTYPE_IS_DECLARED
typedef int YYSTYPE;
# define YYSTYPE_IS_TRIVIAL 1
# define YYSTYPE_IS_DECLARED 1
#endif


extern YYSTYPE yylval;


int yyparse (void);


#endif /* !YY_YY_BISON_TAB_H_INCLUDED  */
