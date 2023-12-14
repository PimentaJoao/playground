# Flex and Bison

Playing around with lexers and parsers and creating a toy compiler.

## Install (setup)

Installing flex:

`sudo apt-get install flex`

Installing bison:

`sudo apt-get install bison`

## Running

Assuming that your lexer and parser definitions are properly written in `lex.l`  and `bison.y` files, respectively.

```shell
# generate parser C code
bison -dy bison.y

# generate lexer C code
# (with tokens based on the parser's definitions)
lex lex.l

# compile generated code
gcc lex.yy.c y.tab.c

# executing the toy compiler
./a.out
```

After this, the toy compiler will start receiving input for it's analysis.

### Processing source code from a file

To use the toy compiler to analyze the source code straight from a file, `<` is used for streaming it's data into our executable.

```shell
./a.out < ./src/my-source-code.txt 
          # example path and file name
```

