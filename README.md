# RegexProject 

Graph Theory Project 2017
Regular Expression Engine

This repository contains an regular expression application written in the programming language Go that defines and implements a non-deterministic automoton to test strings against regular expressions. 
The author is Micheal Curley.

The objective of this years assignment to build a regular expression engine which posed many challenges. I spent a large portion of my time on converting regular expressions
to non-deterministic finite automata. I'm fairly satisfied with my attempt but there is definite room for improvement.

Refrences used: 
Learnonline course page (videos) - https://learnonline.gmit.ie/course/view.php?id=4161
GoLang - https://golang.org/pkg/regexp


Running the code

To run the code in this repository, the files must first be compiled. The Go compiler must first be installed on your machine. 
Once that is installed, the code can be compiled and run by following these steps. We assume you are using the command line.

Clone this repository using Git.
> git clone https://github.com/curley147/RegexProject.git
Change into the folder.
> cd regexengine
Compile the first file with the following command.
> go build regexengine.go
Run the executable produced.
> regexengine.exe

Special characters included in my regular expression engine are '*', '+', '?', '|' and '.' 
When entering the desired regular expression no two non-special characters should be together. Use 
the concatenating operater '.' to match one character after another(i.e enter a.b*.c+ instead of ab*c+)
