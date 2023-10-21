# General Notes

## Write Tests First

- Write a test
- Write basic code to make the compiler pass
- Run the test, see that if it fails and check the errors
- Write code to make the test pass
- Refactor

By sticking to this routine helps us write well designed software.<br/>
Seeing the tests fail also lets you see what the error messages look like, which will be very important while refactoring the code later down the line. <br/>
By not writing test you are committing to manually checking your code by running your software.

## Type, Struct, Interface

As opposed to languages like python and javascript, arrays in Go are **homogeneous** (meaning they can contain data of same type).<br/>
To create an array with multiple types we can use ```interface{}``` or ```struct```<br/>
In Go interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile.
