# Format Specifiers Quick Reference

| Syntax | Function                                                                           | Example                                                               |
|--------|------------------------------------------------------------------------------------|-----------------------------------------------------------------------|
| %%     | Prints a literal percent sign                                                      | fmt.Printf("%%") will print the string %                              |
| %d     | Prints a decimal integer                                                           | fmt.Printf("%d", 10) will print the string 10                         |
| %o     | Prints an octal integer                                                            | fmt.Printf("%o", 10) will print the string 12                         |
| %x     | Prints a hexadecimal integer                                                       | fmt.Printf("%x", 10) will print the string a                          |
| %f     | Prints a floating-point number                                                     | fmt.Printf("%f", 10.5) will print the string 10.5                     |
| %e     | Prints a floating-point number in scientific notation                              | fmt.Printf("%e", 10.5) will print the string 1.05e+01                 |
| %s     | Prints a string                                                                    | fmt.Printf("%s", "Hello, world!") will print the string Hello, world! |
| %v     | Prints the value in a default format                                               | fmt.Printf("%v", 10.5) will print the string 10.5                     |
| %+v    | Prints the value in a default format, including the plus sign for positive numbers | fmt.Printf("%+v", 10.5) will print the string +10.5                   |
| %#v    | Prints the value in a Go syntax representation                                     | fmt.Printf("%#v", 10.5) will print the string 10.5                    |
