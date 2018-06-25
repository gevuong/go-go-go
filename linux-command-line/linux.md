## Linux from Zero to Expert

### Why is Linux important to learn?
- Linux is used on 80% of all supercomputers
- 


### Part 1 Basics
`date`
`cal`

# Download the working files
`wget http://bit.ly/cli-files`

# Unzip file
`unzip cli-files`
`ls`
`clear`
`ls -alg` - provides more details of each file, and group it belongs in

# Notice that we don’t want any spaces in
# directory names
`cd cli-files-0.0`


### Part 2: Working with Directories
`mkdir temp`
`cd temp`

# Go back to last working directory
`cd -`

# Create hierarchy of directories
# (gives error)
`mkdir d1/d2/d3`
`mkdir -p d1/d2/d3`

# List directory hierarchy and what files each directory contains, recursively
`ls -R d1`


### Part 3: Basics of Files
`pwd` - tells where you are in current directory from root
`touch hello.txt`

# Output contents of file
`cat hello.txt` - cat concatenate contents from files and outputs it on console.

# Word count
`wc dummy-file.txt` - prints newline, word, and byte counts for each file

# Getting help
`whatis wc`
`man wc` - short for manual, which opens a huge file with details of `wc`
`ls lesson-02/*.csv` - returns which files are `.csv` in `lesson-02` dir


### Part 4 Speeding Up

# Use Up/Down arrow keys to cycle through
# previous commands
`history` - list history of commands and numbers each of them in order of execution
`!540 # Enter number after the symbol`
`Ctrl + R` # Search through history - if you don't remember what number the command was in history

// have not tried these commands before 
!! # Repeat last command
Ctrl + L # Clear
Ctrl + U # Cut everything before the cursor
Ctrl + K # Cut everything after the cursor
Ctrl + Y # Paste stuff back in
Ctrl + A # Go to start of line
Ctrl + E # Go to end of line


### Part 5 Inter-Process Communication
`head iris.csv` # show top 10 lines

# Piping and Chaining Commands
- a pipe is used as input/output from one command to another. You can chain them together to your liking. The pipe symbol crates a medium between `cat` and `wc` for example. Whatever cat produces, is fed to wc. It's considered decoupled coding, because cat has not idea what wc wants, and vice versa.
`cat iris.csv | wc`
`grep "setosa" iris.csv` - searches for a specific word and output those lines that contain the word. `grep` can regex as well
`cat iris.csv | grep "setosa"`
`cat iris.csv | grep "setosa" | wc` - find lines containing "setosa" and count those lines
`cat iris.csv | grep "set" | grep "3.5" | wc`
`ls | grep csv` - finds all the files that have `csv` in them


### Part 6 Redirection
# Output to console
`echo "Something"` - basically echos what is written onto the console
`open .` - opens current directory in Finder

# Redirect to file
`echo "Something" > temp` - after running this, the output of echo will be stored in temp file
`cat iris.csv | grep "setosa" > setosat.csv`

# Move/rename files
`mv setosat.csv setosa.csv` - if you're moving a file in same directory, you're basically renaming the file from `setosat.csv` to `setosa.csv`


## Part 7 Remove, Copy
`mkdir backup`
`cp iris.csv iris-backup.csv`
`mv iris-backup.csv backup` - moves `iris-backup.csv` to `backup` dir
`rm setosa.csv`
`cp backup backup-2` # Directory omitted , cp can only work with files by default, not directories
`cp -r backup backup-2` - will create `backup-2`, which contains contents in original `backup`. The `-r` flag means recursive
rm backup # Directory omitted, won't work, what you can do is remove every file in `backup`, then `rmdir backup`
`rm -r backup` # Recursively deleted, **DANGEROUS** because by removing directory, it's gone, it does not go to trash bin. Remember, "with great power, comes great responsibility".


## Part 13 VIM Commands (Lesson 5)
- There are other command line editors other than vi, like emacs and nano, but we'll be using `vi`

`vi utils.py`
# [In Command Mode]
`:q` # Quit
`:wq` # Write file and quit - **huge time saver**
`:q!` # Quit without saving changes
`i` # Go to insert mode
INSERT # Go to insert mode

Why is it important to learn these things?
**It's important to learn these commands because there will be instances when you have no choice but to use `vi`. For instance when logging int a server which may only have a command line**

# Show line numbers
`:set number`
`:<insert line number> `to move to that line 

# Go to end of file
`Shift + G` - will bring you to end of file
`:0` # Go to line number 0
`0` # Go to beginning of line
:100 # Go to line 100
`/num` # Search for ‘num’
`/protein` # Search for ‘protein’
`n` # After search, find next
`f` + `<insert character>` # this will jump to that character in current line
fH # Go forward to character ’H’
FH # Go backwards to character ’H’
`dw` # Delete word
`dd` # Delete (actually cuts) line under cursor, and you can `p` paste it on another line
`d10d` # Delete 10 lines
`x` # Delete character under cursor
`yy` # Copy line
`p` # Paste
. # Repeat last operation
`u` # Undo last operation
`%` # Go to matching bracket
`vi"` # Select inside double quotes - `v` is for VISUAL selection as selecting things, `i` means inside, and `"` means inside double quotes

# [In Insert Mode]
ESC # Go to command mode