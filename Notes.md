# Notes

## File Stricture

In Go, it is generally recommended to use lowercase, hyphen-seperated names for directories, packages, and modules. Using spaces or uppercase letters may cause package discovery and imports. I was getting a warning that the file is not within module ".". This was resolved when I updated the folder structure and named the folders properly.

## Types

As opposed to languages like python and javascript, arrays in Go are **homogeneous** (meaning they can contain data of same type).
To create an array with multiple types we can use ```interface{}``` or ```struct```

## Slice

A slice in Go is a portion of array, Unlike arrays, the length of a slice is dynamic and it can change during the execution of a program.

- A slice does not own any data of it's own
- Any modifications done to the slice will affect the array
- Slices can be created directly without an initial array
