# Go-PLSQL Native Function

This is a template application that creates a **PostgreSQL** native function in the GO language. This application can be used as a statring point to build more complex PostgreSQL native functions.

## Build and Install

To build and install the function the following has to be done:  

* edit `install.sh` and change the `GO_DB` parameter to point to the correct PostgreSQL database that will be used to install the function. 

* run the `compile.sh` script to create **libGoPgFunc.so** library file.

```
./compile.sh
```

* run the `install.sh` to create a SQL  function in the PostgreSQL database that will call the native function in the created library file.

```
./install.sh
CREATE FUNCTION
```

## Test

By this point, the sample function `get_arg_text` already has been installed and can now be run.

```
go=# \df get_arg_text
                                 List of functions
 Schema |     Name     | Result data type |       Argument data types       | Type 
--------+--------------+------------------+---------------------------------+------
 public | get_arg_text | text             | text, integer, double precision | func
(1 row)

```

and to actually use it:

```
go=# select get_arg_text('マリオ・バサ',100,3.141618);
                 get_arg_text                  
-----------------------------------------------
 Hello マリオ・バサ, int: 101, float: 4.141618
(1 row)

```
