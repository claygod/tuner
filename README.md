# Tuner
The loader in Go ini files with simultaneous support for environment variables and command line.

The tuner is designed for easy configuration of the application.
The parameters it takes from (in decreasing order of priority):
- Command Line
- environment
- Configuration File

# Usage

Example

```Go
	tun, err := tuner.New(confile)
	if err != nil {
		panic(err)
	}
	path := tun.Section("Main").Get("path").(string)
	length := tun.Section("Main").Get("length").(float64)
	mas := tun.Section("Main").Get("mas").([]interface{})
	first := mas[0].(int)
 ```
 
 The configuration file is a usual ini-file, which supports strings, integers, floats and lists.
 
# Comments
Comments are denoted by the symbol `#`. All that follows this symbol in the line is a comment.

# Srtings
Strings are enclosed in single or double quotes. Embedded quotation marks must be distinguished from the exterior.

# Integers
Variable, consisting only of numbers is treated as a whole int. An integer can not be more than 2147483648

# Floats
Chance of numbers containing the point (not a comma!) Treated as `float64`

# Lists
Lists may contain both types of strings and numbers. Nested lists are not supported. Lists are enclosed in single or braces. list of members separated by commas.

# The principle of the formation of the key
In the environment variables, and command-line switch variable must be specified as a pair of `Section_key` symbol `_` connected. Connecting the character can be changed in the configuration file. The lines should be the same way as in the file, enclosed in quotation marks, and the list in parentheses.

### Example
For the section of Main set value for the path variable must be so: `Main_path`. Command line at the same time could look in Windows as follows:
`Tuner.exe -Main_path =" config.ini "-Second_mas = (6,7," abc "," efj ", 3.14)`

### Getting the value
When a value is first necessary to use the method `Section` and then the method of` Get`. The resulting value should lead to the desired type. Lists return a list of interfaces in the list and you will also need to give value to the desired type.
`Path: = tun.Section (" Main ") Get (" path ") (string)`..

# Changing the value
If necessary, you can change the value of the configuration parameter, the contents of the configuration file is not changed:

```Go
	tnr, err := New("config.ini")
	y := tnr.Section("Main").Set("path", "now.ini")
	if x := tnr.Section("Main").Get("path"); y != nil || x != "now.ini" {
		fmt.Print("Error in key `path`: ", x, "! ", err)
	}
  ```
  


