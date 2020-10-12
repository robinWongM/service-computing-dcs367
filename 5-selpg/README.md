# 设计说明

参照 `selpg.c` 的思路进行实现，分为 `processFlag` 和 `processInput` 两部分。

命令行参数部分不使用 `selpg.c` 中手动处理的方式，而是使用 `pflag` 进行解析，再手动对后置条件进行判断。（故未能完全模拟 `selpg.c` 的行为。

# 单元或集成测试结果

暂无。

# 功能测试结果

参照 IBM 用法介绍进行若干测试：

```
codespace ➜ ~/workspace/service-computing-dcs367/5-selpg (main ✗) $ ./selpg -s1 -e1 input_file 
Hello, World!

./selpg: done
```

```
codespace ➜ ~/workspace/service-computing-dcs367/5-selpg (main ✗) $ ./selpg -s1 -e1 < input_file 
Hello, World!

./selpg: done
```

```
codespace ➜ ~/workspace/service-computing-dcs367/5-selpg (main ✗) $ man bash | ./selpg -s10 -e20
              programmable completion facilities (see Programmable Completion below).
       COMP_TYPE
              Set  to an integer value corresponding to the type of completion attempted that caused a completion func‐
              tion to be called: TAB, for normal completion, ?, for listing completions after successive tabs,  !,  for
              listing alternatives on partial word completion, @, to list completions if the word is not unmodified, or
              %, for menu completion.  This variable is available only in shell functions and external commands invoked
              by the programmable completion facilities (see Programmable Completion below).
       COMP_WORDBREAKS
              The  set  of  characters that the readline library treats as word separators when performing word comple‐
              tion.  If COMP_WORDBREAKS is unset, it loses its special properties, even if it is subsequently reset.
       COMP_WORDS
              An array variable (see Arrays below) consisting of the individual words in the current command line.  The
              line  is  split  into  words  as readline would split it, using COMP_WORDBREAKS as described above.  This
              variable is available only in shell functions invoked by the programmable completion facilities (see Pro‐
              grammable Completion below).
       COPROC An array variable (see Arrays below) created to hold the file descriptors for output from and input to an
              unnamed coprocess (see Coprocesses above).
       DIRSTACK
              An array variable (see Arrays below) containing the current contents of the directory stack.  Directories
              appear  in  the  stack in the order they are displayed by the dirs builtin.  Assigning to members of this
              array variable may be used to modify directories already in the stack, but the pushd  and  popd  builtins
              must  be  used  to  add  and remove directories.  Assignment to this variable will not change the current
              directory.  If DIRSTACK is unset, it loses its special properties, even if it is subsequently reset.
       EUID   Expands to the effective user ID of the current user, initialized at shell  startup.   This  variable  is
              readonly.
       FUNCNAME
              An array variable containing the names of all shell functions currently in the execution call stack.  The
              element with index 0 is the name of any currently-executing shell function.  The bottom-most element (the
              one  with  the  highest  index) is "main".  This variable exists only when a shell function is executing.
              Assignments to FUNCNAME have no effect.  If FUNCNAME is unset, it loses its special properties,  even  if
              it is subsequently reset.

              This  variable  can be used with BASH_LINENO and BASH_SOURCE.  Each element of FUNCNAME has corresponding
              elements in BASH_LINENO and BASH_SOURCE to describe the call stack.  For  instance,  ${FUNCNAME[$i]}  was
              called from the file ${BASH_SOURCE[$i+1]} at line number ${BASH_LINENO[$i]}.  The caller builtin displays
              the current call stack using this information.
       GROUPS An array variable containing the list of groups of which the current user is a  member.   Assignments  to
              GROUPS  have  no effect.  If GROUPS is unset, it loses its special properties, even if it is subsequently
              reset.
       HISTCMD
              The history number, or index in the history list, of the current command.  If HISTCMD is unset, it  loses
              its special properties, even if it is subsequently reset.
       HOSTNAME
              Automatically set to the name of the current host.
       HOSTTYPE
              Automatically  set  to  a  string that uniquely describes the type of machine on which bash is executing.
              The default is system-dependent.
       LINENO Each time this parameter is referenced, the shell substitutes a decimal number representing  the  current
              sequential  line number (starting with 1) within a script or function.  When not in a script or function,
              the value substituted is not guaranteed to be meaningful.  If LINENO is unset, it loses its special prop‐
              erties, even if it is subsequently reset.
       MACHTYPE
              Automatically  set  to  a  string that fully describes the system type on which bash is executing, in the
              standard GNU cpu-company-system format.  The default is system-dependent.
       MAPFILE
              An array variable (see Arrays below) created to hold the text read by the mapfile builtin when  no  vari‐
              able name is supplied.
       OLDPWD The previous working directory as set by the cd command.
       OPTARG The  value  of  the last option argument processed by the getopts builtin command (see SHELL BUILTIN COM‐
              MANDS below).
       OPTIND The index of the next argument to be processed by the getopts builtin command (see SHELL BUILTIN COMMANDS
              below).
       OSTYPE Automatically  set  to  a  string  that  describes  the operating system on which bash is executing.  The
              default is system-dependent.
       PIPESTATUS
              An array variable (see Arrays below) containing a list of exit status values from the  processes  in  the
              most-recently-executed foreground pipeline (which may contain only a single command).
       PPID   The process ID of the shell's parent.  This variable is readonly.
       PWD    The current working directory as set by the cd command.
       RANDOM Each  time this parameter is referenced, a random integer between 0 and 32767 is generated.  The sequence
              of random numbers may be initialized by assigning a value to RANDOM.  If RANDOM is unset,  it  loses  its
              special properties, even if it is subsequently reset.
       READLINE_LINE
              The contents of the readline line buffer, for use with "bind -x" (see SHELL BUILTIN COMMANDS below).
       READLINE_POINT
              The  position  of  the  insertion  point  in  the readline line buffer, for use with "bind -x" (see SHELL
              BUILTIN COMMANDS below).
       REPLY  Set to the line of input read by the read builtin command when no arguments are supplied.
       SECONDS
              Each time this parameter is referenced, the number of seconds since shell invocation is returned.   If  a
              value  is  assigned  to  SECONDS,  the value returned upon subsequent references is the number of seconds
              since the assignment plus the value assigned.  If SECONDS is unset, it loses its special properties, even
              if it is subsequently reset.
       SHELLOPTS
              A  colon-separated  list  of enabled shell options.  Each word in the list is a valid argument for the -o
              option to the set builtin command (see SHELL BUILTIN COMMANDS below).  The options appearing in SHELLOPTS
              are  those  reported  as  on by set -o.  If this variable is in the environment when bash starts up, each
              shell option in the list will be enabled before reading any startup files.  This variable is read-only.
       SHLVL  Incremented by one each time an instance of bash is started.
       UID    Expands to the user ID of the current user, initialized at shell startup.  This variable is readonly.

       The following variables are used by the shell.  In some cases, bash assigns a default value to a variable; these
       cases are noted below.

       BASH_COMPAT
              The value is used to set the shell's compatibility level.  See the description of the shopt builtin below
              under SHELL BUILTIN COMMANDS for a description of the various compatibility  levels  and  their  effects.
              The  value may be a decimal number (e.g., 4.2) or an integer (e.g., 42) corresponding to the desired com‐
              patibility level.  If BASH_COMPAT is unset or set to the empty string, the compatibility level is set  to
              the  default for the current version.  If BASH_COMPAT is set to a value that is not one of the valid com‐
              patibility levels, the shell prints an error message and sets the compatibility level to the default  for
              the  current version.  The valid compatibility levels correspond to the compatibility options accepted by
              the shopt builtin described below (for example, compat42 means that 4.2 and 42 are  valid  values).   The
              current version is also a valid value.
       BASH_ENV
              If  this  parameter  is set when bash is executing a shell script, its value is interpreted as a filename
              containing commands to initialize the shell, as in ~/.bashrc.  The value  of  BASH_ENV  is  subjected  to
              parameter  expansion,  command substitution, and arithmetic expansion before being interpreted as a file‐
              name.  PATH is not used to search for the resultant filename.
       BASH_XTRACEFD
              If set to an integer corresponding to a valid file descriptor, bash will write the trace output generated
              when  set  -x  is  enabled  to that file descriptor.  The file descriptor is closed when BASH_XTRACEFD is
              unset or assigned a new value.  Unsetting BASH_XTRACEFD or assigning it the empty string causes the trace
              output  to  be sent to the standard error.  Note that setting BASH_XTRACEFD to 2 (the standard error file
              descriptor) and then unsetting it will result in the standard error being closed.
       CDPATH The search path for the cd command.  This is a colon-separated list of directories  in  which  the  shell
              looks for destination directories specified by the cd command.  A sample value is ".:~:/usr".
       CHILD_MAX
              Set  the  number of exited child status values for the shell to remember.  Bash will not allow this value
              to be decreased below a POSIX-mandated minimum, and there is a maximum value (currently 8192)  that  this
              may not exceed.  The minimum value is system-dependent.
       COLUMNS
              Used by the select compound command to determine the terminal width when printing selection lists.  Auto‐
              matically set if the checkwinsize option is enabled or in an interactive shell upon  receipt  of  a  SIG‐
              WINCH.
       COMPREPLY
              An array variable from which bash reads the possible completions generated by a shell function invoked by
              the programmable completion facility (see Programmable Completion below).  Each  array  element  contains
              one possible completion.
       EMACS  If  bash finds this variable in the environment when the shell starts with value "t", it assumes that the
              shell is running in an Emacs shell buffer and disables line editing.
       ENV    Similar to BASH_ENV; used when the shell is invoked in POSIX mode.
       EXECIGNORE
              A colon-separated list of shell patterns (see Pattern Matching) defining the  list  of  filenames  to  be
              ignored  by  command  search  using PATH.  Files whose full pathnames match one of these patterns are not
              considered executable files for the purposes of completion and command execution via PATH  lookup.   This
              does  not  affect the behavior of the [, test, and [[ commands.  Full pathnames in the command hash table
              are not subject to EXECIGNORE.  Use this variable to ignore shared library files that have the executable
              bit  set,  but  are  not  executable files.  The pattern matching honors the setting of the extglob shell
              option.
       FCEDIT The default editor for the fc builtin command.
       FIGNORE
              A colon-separated list of suffixes to ignore when performing filename completion (see READLINE below).  A
              filename  whose  suffix  matches one of the entries in FIGNORE is excluded from the list of matched file‐
              names.  A sample value is ".o:~" (Quoting is needed when assigning a value to this variable,  which  con‐
              tains tildes).
       FUNCNEST
              If set to a numeric value greater than 0, defines a maximum function nesting level.  Function invocations
              that exceed this nesting level will cause the current command to abort.
       GLOBIGNORE
              A colon-separated list of patterns defining the set of filenames to be ignored by pathname expansion.  If
              a  filename matched by a pathname expansion pattern also matches one of the patterns in GLOBIGNORE, it is
              removed from the list of matches.
       HISTCONTROL
              A colon-separated list of values controlling how commands are saved on the history list.  If the list  of
              values  includes ignorespace, lines which begin with a space character are not saved in the history list.
              A value of ignoredups causes lines matching the previous history entry to  not  be  saved.   A  value  of
              ignoreboth  is  shorthand for ignorespace and ignoredups.  A value of erasedups causes all previous lines
              matching the current line to be removed from the history list before that line is saved.  Any  value  not
              in the above list is ignored.  If HISTCONTROL is unset, or does not include a valid value, all lines read
              by the shell parser are saved on the history list, subject to the value of HISTIGNORE.   The  second  and
              subsequent lines of a multi-line compound command are not tested, and are added to the history regardless
              of the value of HISTCONTROL.
       HISTFILE
              The name of the file in which command history is  saved  (see  HISTORY  below).   The  default  value  is
              ~/.bash_history.  If unset, the command history is not saved when a shell exits.
       HISTFILESIZE
              The  maximum  number of lines contained in the history file.  When this variable is assigned a value, the
              history file is truncated, if necessary, to contain no more than that number of  lines  by  removing  the
              oldest entries.  The history file is also truncated to this size after writing it when a shell exits.  If
              the value is 0, the history file is truncated to zero size.  Non-numeric values and numeric  values  less
              than  zero  inhibit  truncation.  The shell sets the default value to the value of HISTSIZE after reading
              any startup files.
       HISTIGNORE
              A colon-separated list of patterns used to decide which command lines should  be  saved  on  the  history
              list.   Each  pattern  is  anchored  at  the  beginning  of the line and must match the complete line (no
              implicit `*' is appended).  Each pattern is tested against the line after the checks specified  by  HIST‐
              CONTROL are applied.  In addition to the normal shell pattern matching characters, `&' matches the previ‐
              ous history line.  `&' may be escaped using a backslash; the backslash is  removed  before  attempting  a
              match.  The second and subsequent lines of a multi-line compound command are not tested, and are added to
              the history regardless of the value of HISTIGNORE.  The pattern matching honors the setting of  the  ext‐
              glob shell option.
       HISTSIZE
              The  number  of commands to remember in the command history (see HISTORY below).  If the value is 0, com‐
              mands are not saved in the history list.  Numeric values less than zero result  in  every  command  being
              saved on the history list (there is no limit).  The shell sets the default value to 500 after reading any
              startup files.
       HISTTIMEFORMAT
              If this variable is set and not null, its value is used as a format string for strftime(3) to  print  the
              time stamp associated with each history entry displayed by the history builtin.  If this variable is set,
              time stamps are written to the history file so they may be preserved across shell  sessions.   This  uses
              the history comment character to distinguish timestamps from other history lines.
       HOME   The  home  directory  of the current user; the default argument for the cd builtin command.  The value of
              this variable is also used when performing tilde expansion.
       HOSTFILE
              Contains the name of a file in the same format as /etc/hosts that should be read when the shell needs  to
              complete  a  hostname.   The list of possible hostname completions may be changed while the shell is run‐
              ning; the next time hostname completion is attempted after the value is changed, bash adds  the  contents
              of  the new file to the existing list.  If HOSTFILE is set, but has no value, or does not name a readable
              file, bash attempts to read /etc/hosts to obtain the list of possible hostname completions.   When  HOST‐
              FILE is unset, the hostname list is cleared.
       IFS    The  Internal  Field  Separator  that  is used for word splitting after expansion and to split lines into
              words with the read builtin command.  The default value is ``<space><tab><newline>''.
       IGNOREEOF
              Controls the action of an interactive shell on receipt of an EOF character as the sole  input.   If  set,
              the  value  is the number of consecutive EOF characters which must be typed as the first characters on an
              input line before bash exits.  If the variable exists but does not have a numeric value, or has no value,
              the default value is 10.  If it does not exist, EOF signifies the end of input to the shell.
       INPUTRC
              The filename for the readline startup file, overriding the default of ~/.inputrc (see READLINE below).
       LANG   Used to determine the locale category for any category not specifically selected with a variable starting
              with LC_.
       LC_ALL This variable overrides the value of LANG and any other LC_ variable specifying a locale category.
       LC_COLLATE
              This variable determines the collation order used when sorting the results  of  pathname  expansion,  and
              determines  the  behavior of range expressions, equivalence classes, and collating sequences within path‐
              name expansion and pattern matching.
       LC_CTYPE
              This variable determines the interpretation of characters and the behavior of  character  classes  within
              pathname expansion and pattern matching.
       LC_MESSAGES
              This variable determines the locale used to translate double-quoted strings preceded by a $.
       LC_NUMERIC
              This variable determines the locale category used for number formatting.
       LC_TIME
              This variable determines the locale category used for data and time formatting.
       LINES  Used  by  the select compound command to determine the column length for printing selection lists.  Auto‐
              matically set if the checkwinsize option is enabled or in an interactive shell upon  receipt  of  a  SIG‐
              WINCH.
       MAIL   If  this  parameter is set to a file or directory name and the MAILPATH variable is not set, bash informs
              the user of the arrival of mail in the specified file or Maildir-format directory.
       MAILCHECK
              Specifies how often (in seconds) bash checks for mail.  The default is 60 seconds.  When it  is  time  to
              check  for  mail,  the shell does so before displaying the primary prompt.  If this variable is unset, or
              set to a value that is not a number greater than or equal to zero, the shell disables mail checking.
       MAILPATH
              A colon-separated list of filenames to be checked for mail.  The message to be printed when mail  arrives
              in  a particular file may be specified by separating the filename from the message with a `?'.  When used
              in the text of the message, $_ expands to the name of the current mailfile.  Example:
              MAILPATH='/var/mail/bfox?"You have mail":~/shell-mail?"$_ has mail!"'
              Bash can be configured to supply a default value for this variable (there is no value  by  default),  but
              the location of the user mail files that it uses is system dependent (e.g., /var/mail/$USER).
       OPTERR If  set  to the value 1, bash displays error messages generated by the getopts builtin command (see SHELL
              BUILTIN COMMANDS below).  OPTERR is initialized to 1 each time the shell is invoked or a shell script  is
              executed.
       PATH   The  search  path for commands.  It is a colon-separated list of directories in which the shell looks for
              commands (see COMMAND EXECUTION below).  A zero-length (null) directory name in the value of  PATH  indi‐
              cates  the  current directory.  A null directory name may appear as two adjacent colons, or as an initial
              or trailing colon.  The default path is system-dependent, and is set by the  administrator  who  installs
              bash.  A common value is ``/usr/local/bin:/usr/local/sbin:/usr/bin:/usr/sbin:/bin:/sbin''.
       POSIXLY_CORRECT
              If  this  variable is in the environment when bash starts, the shell enters posix mode before reading the
              startup files, as if the --posix invocation option had been supplied.  If it is set while  the  shell  is
              running, bash enables posix mode, as if the command set -o posix had been executed.
       PROMPT_COMMAND
              If set, the value is executed as a command prior to issuing each primary prompt.
       PROMPT_DIRTRIM
              If set to a number greater than zero, the value is used as the number of trailing directory components to
              retain when expanding the \w and \W prompt string escapes (see PROMPTING below).  Characters removed  are
              replaced with an ellipsis.
       PS0    The  value  of this parameter is expanded (see PROMPTING below) and displayed by interactive shells after
              reading a command and before the command is executed.
       PS1    The value of this parameter is expanded (see PROMPTING below) and used as the primary prompt string.  The
              default value is ``\s-\v\$ ''.
       PS2    The value of this parameter is expanded as with PS1 and used as the secondary prompt string.  The default
              is ``> ''.
       PS3    The value of this parameter is used as the prompt for the select command (see SHELL GRAMMAR above).
       PS4    The value of this parameter is expanded as with PS1 and the value is printed  before  each  command  bash
              displays  during  an execution trace.  The first character of PS4 is replicated multiple times, as neces‐
              sary, to indicate multiple levels of indirection.  The default is ``+ ''.
       SHELL  The full pathname to the shell is kept in this environment variable.  If it is not  set  when  the  shell
              starts, bash assigns to it the full pathname of the current user's login shell.
       TIMEFORMAT
              The  value  of  this parameter is used as a format string specifying how the timing information for pipe‐
              lines prefixed with the time reserved word should be displayed.  The %  character  introduces  an  escape
              sequence  that is expanded to a time value or other information.  The escape sequences and their meanings
              are as follows; the braces denote optional portions.
              %%        A literal %.
              %[p][l]R  The elapsed time in seconds.
              %[p][l]U  The number of CPU seconds spent in user mode.
              %[p][l]S  The number of CPU seconds spent in system mode.
              %P        The CPU percentage, computed as (%U + %S) / %R.

              The optional p is a digit specifying the precision, the number  of  fractional  digits  after  a  decimal
              point.   A  value  of 0 causes no decimal point or fraction to be output.  At most three places after the
              decimal point may be specified; values of p greater than 3 are changed to 3.  If p is not specified,  the
              value 3 is used.

              The  optional  l  specifies  a  longer  format, including minutes, of the form MMmSS.FFs.  The value of p
              determines whether or not the fraction is included.

              If this variable is not set, bash acts as if it had the value $'\nreal\t%3lR\nuser\t%3lU\nsys\t%3lS'.  If
              the  value  is  null,  no  timing  information is displayed.  A trailing newline is added when the format
              string is displayed.
       TMOUT  If set to a value greater than zero, TMOUT is treated as the default timeout for the read  builtin.   The
              select command terminates if input does not arrive after TMOUT seconds when input is coming from a termi‐
              nal.  In an interactive shell, the value is interpreted as the number of seconds to wait for  a  line  of
              input  after  issuing  the primary prompt.  Bash terminates after waiting for that number of seconds if a
              complete line of input does not arrive.
       TMPDIR If set, bash uses its value as the name of a directory in which bash  creates  temporary  files  for  the
              shell's use.
       auto_resume
              This  variable  controls how the shell interacts with the user and job control.  If this variable is set,
              single word simple commands without redirections are treated as candidates for resumption of an  existing
              stopped  job.   There  is  no  ambiguity allowed; if there is more than one job beginning with the string
              typed, the job most recently accessed is selected.  The name of a stopped job, in this  context,  is  the
              command  line  used to start it.  If set to the value exact, the string supplied must match the name of a
              stopped job exactly; if set to substring, the string supplied needs to match a substring of the name of a
              stopped  job.   The  substring  value provides functionality analogous to the %?  job identifier (see JOB
              CONTROL below).  If set to any other value, the supplied string must be a prefix of a stopped job's name;
              this provides functionality analogous to the %string job identifier.
       histchars
              The  two  or  three  characters  which  control history expansion and tokenization (see HISTORY EXPANSION
              below).  The first character is the history expansion character, the character which signals the start of
              a  history  expansion,  normally `!'.  The second character is the quick substitution character, which is
              used as shorthand for re-running the previous command entered, substituting one string for another in the
              command.   The  default  is  `^'.  The optional third character is the character which indicates that the
              remainder of the line is a comment when found as the first character of a word, normally `#'.   The  his‐
              tory comment character causes history substitution to be skipped for the remaining words on the line.  It
              does not necessarily cause the shell parser to treat the rest of the line as a comment.

   Arrays
       Bash provides one-dimensional indexed and associative array variables.  Any variable may be used as  an  indexed
       array; the declare builtin will explicitly declare an array.  There is no maximum limit on the size of an array,
       nor any requirement that members be indexed or assigned contiguously.  Indexed arrays are referenced using inte‐
       gers  (including  arithmetic  expressions) and are zero-based; associative arrays are referenced using arbitrary
       strings.  Unless otherwise noted, indexed array indices must be non-negative integers.

       An indexed array is created automatically if any variable is assigned to using the syntax name[subscript]=value.
       The  subscript  is treated as an arithmetic expression that must evaluate to a number.  To explicitly declare an
       indexed array, use declare -a name (see SHELL BUILTIN COMMANDS  below).   declare  -a  name[subscript]  is  also
       accepted; the subscript is ignored.

       Associative arrays are created using declare -A name.

       Attributes  may  be  specified  for  an  array variable using the declare and readonly builtins.  Each attribute
       applies to all members of an array.

       Arrays are assigned to using compound assignments of the form name=(value1 ... valuen), where each value  is  of
       the  form  [subscript]=string.  Indexed array assignments do not require anything but string.  When assigning to
       indexed arrays, if the optional brackets and subscript are supplied, that index is assigned  to;  otherwise  the
       index of the element assigned is the last index assigned to by the statement plus one.  Indexing starts at zero.

       When assigning to an associative array, the subscript is required.

       This  syntax  is  also  accepted by the declare builtin.  Individual array elements may be assigned to using the
       name[subscript]=value syntax introduced above.  When assigning to an indexed array, if name is subscripted by  a
       negative  number, that number is interpreted as relative to one greater than the maximum index of name, so nega‐
       tive indices count back from the end of the array, and an index of -1 references the last element.

       Any element of an array may be referenced using ${name[subscript]}.  The braces are required to avoid  conflicts
       with  pathname  expansion.   If  subscript is @ or *, the word expands to all members of name.  These subscripts
       differ only when the word appears within double quotes.  If the word is double-quoted, ${name[*]} expands  to  a
       single  word  with  the value of each array member separated by the first character of the IFS special variable,
       and ${name[@]} expands each element of name to a separate word.  When there are  no  array  members,  ${name[@]}
       expands  to  nothing.  If the double-quoted expansion occurs within a word, the expansion of the first parameter
       is joined with the beginning part of the original word, and the expansion of the last parameter is  joined  with
       the  last  part of the original word.  This is analogous to the expansion of the special parameters * and @ (see
       Special Parameters above).  ${#name[subscript]} expands to the length of ${name[subscript]}.  If subscript is  *
       or @, the expansion is the number of elements in the array.  If the subscript used to reference an element of an
       indexed array evaluates to a number less than zero, it is interpreted as relative to one greater than the  maxi‐
       mum  index of the array, so negative indices count back from the end of the array, and an index of -1 references
       the last element.

       Referencing an array variable without a subscript is equivalent to referencing the array with a subscript of  0.
       Any reference to a variable using a valid subscript is legal, and bash will create an array if necessary.

       An array variable is considered set if a subscript has been assigned a value.  The null string is a valid value.

       It  is  possible  to  obtain  the keys (indices) of an array as well as the values.  ${!name[@]} and ${!name[*]}
       expand to the indices assigned in array variable name.  The treatment when in double quotes is  similar  to  the
       expansion of the special parameters @ and * within double quotes.

       The  unset  builtin  is  used to destroy arrays.  unset name[subscript] destroys the array element at index sub‐
       script.  Negative subscripts to indexed arrays are interpreted as described above.  Care must be taken to  avoid
       unwanted  side  effects  caused  by  pathname expansion.  unset name, where name is an array, or unset name[sub‐
       script], where subscript is * or @, removes the entire array.

       The declare, local, and readonly builtins each accept a -a option to specify an indexed array and a -A option to
       specify an associative array.  If both options are supplied, -A takes precedence.  The read builtin accepts a -a
       option to assign a list of words read from the standard input to an array.  The set and declare builtins display
       array values in a way that allows them to be reused as assignments.

EXPANSION
       Expansion  is performed on the command line after it has been split into words.  There are seven kinds of expan‐
       sion performed: brace expansion, tilde expansion, parameter and variable expansion, command substitution, arith‐
       metic expansion, word splitting, and pathname expansion.

       The  order  of  expansions  is:  brace  expansion; tilde expansion, parameter and variable expansion, arithmetic
       expansion, and command substitution (done in a left-to-right fashion); word splitting; and pathname expansion.

       On systems that can support it, there is an additional expansion available: process substitution.  This is  per‐
       formed at the same time as tilde, parameter, variable, and arithmetic expansion and command substitution.

       After these expansions are performed, quote characters present in the original word are removed unless they have
       been quoted themselves (quote removal).

       Only brace expansion, word splitting, and pathname expansion can change the number of words  of  the  expansion;
       other  expansions expand a single word to a single word.  The only exceptions to this are the expansions of "$@"
       and "${name[@]}" as explained above (see PARAMETERS).

   Brace Expansion
       Brace expansion is a mechanism by which arbitrary strings may be generated.  This mechanism is similar to  path‐
       name  expansion,  but the filenames generated need not exist.  Patterns to be brace expanded take the form of an
       optional preamble, followed by either a series of comma-separated strings or a  sequence  expression  between  a
       pair  of  braces,  followed by an optional postscript.  The preamble is prefixed to each string contained within
       the braces, and the postscript is then appended to each resulting string, expanding left to right.

       Brace expansions may be nested.  The results of each expanded string are not sorted; left to right order is pre‐
       served.  For example, a{d,c,b}e expands into `ade ace abe'.

       A sequence expression takes the form {x..y[..incr]}, where x and y are either integers or single characters, and
       incr, an optional increment, is an integer.  When integers are supplied, the expression expands to  each  number
       between x and y, inclusive.  Supplied integers may be prefixed with 0 to force each term to have the same width.
       When either x or y begins with a zero, the shell attempts to force all generated terms to contain the same  num‐
       ber of digits, zero-padding where necessary.  When characters are supplied, the expression expands to each char‐
       acter lexicographically between x and y, inclusive, using the default C locale.  Note that both x and y must  be
       of  the same type.  When the increment is supplied, it is used as the difference between each term.  The default
       increment is 1 or -1 as appropriate.

       Brace expansion is performed before any other expansions, and any characters special  to  other  expansions  are
       preserved  in the result.  It is strictly textual.  Bash does not apply any syntactic interpretation to the con‐
       text of the expansion or the text between the braces.

       A correctly-formed brace expansion must contain unquoted opening and closing braces, and at least  one  unquoted
       comma  or  a valid sequence expression.  Any incorrectly formed brace expansion is left unchanged.  A { or , may
       be quoted with a backslash to prevent its being considered part of a brace expression.  To avoid conflicts  with
       parameter expansion, the string ${ is not considered eligible for brace expansion.

       This  construct  is  typically used as shorthand when the common prefix of the strings to be generated is longer
       than in the above example:

              mkdir /usr/local/src/bash/{old,new,dist,bugs}
       or
              chown root /usr/{ucb/{ex,edit},lib/{ex?.?*,how_ex}}

       Brace expansion introduces a slight incompatibility with historical versions of sh.  sh does not  treat  opening
       or  closing braces specially when they appear as part of a word, and preserves them in the output.  Bash removes
       braces from words as a consequence of brace expansion.  For example, a word entered to sh as  file{1,2}  appears
       identically  in the output.  The same word is output as file1 file2 after expansion by bash.  If strict compati‐
       bility with sh is desired, start bash with the +B option or disable brace expansion with the +B  option  to  the
       set command (see SHELL BUILTIN COMMANDS below).

   Tilde Expansion
       If  a  word  begins  with  an unquoted tilde character (`~'), all of the characters preceding the first unquoted
       slash (or all characters, if there is no unquoted slash) are considered a tilde-prefix.  If none of the  charac‐
       ters  in  the  tilde-prefix  are quoted, the characters in the tilde-prefix following the tilde are treated as a
       possible login name.  If this login name is the null string, the tilde is replaced with the value of  the  shell
       parameter  HOME.   If  HOME is unset, the home directory of the user executing the shell is substituted instead.
       Otherwise, the tilde-prefix is replaced with the home directory associated with the specified login name.

       If the tilde-prefix is a `~+', the value of the shell variable PWD replaces the tilde-prefix.  If the tilde-pre‐
       fix  is a `~-', the value of the shell variable OLDPWD, if it is set, is substituted.  If the characters follow‐
       ing the tilde in the tilde-prefix consist of a number N, optionally prefixed by a `+' or a `-', the tilde-prefix
       is  replaced  with  the  corresponding  element  from  the directory stack, as it would be displayed by the dirs
       builtin invoked with the tilde-prefix as an argument.  If the characters following the tilde in the tilde-prefix
       consist of a number without a leading `+' or `-', `+' is assumed.

       If the login name is invalid, or the tilde expansion fails, the word is unchanged.

       Each  variable  assignment  is checked for unquoted tilde-prefixes immediately following a : or the first =.  In
       these cases, tilde expansion is also performed.  Consequently, one may use filenames with tildes in  assignments
       to PATH, MAILPATH, and CDPATH, and the shell assigns the expanded value.

   Parameter Expansion
       The  `$' character introduces parameter expansion, command substitution, or arithmetic expansion.  The parameter
       name or symbol to be expanded may be enclosed in braces, which are optional but serve to protect the variable to
       be expanded from characters immediately following it which could be interpreted as part of the name.

       When  braces  are used, the matching ending brace is the first `}' not escaped by a backslash or within a quoted
       string, and not within an embedded arithmetic expansion, command substitution, or parameter expansion.

       ${parameter}
              The value of parameter is substituted.  The braces are required when parameter is a positional  parameter
              with  more than one digit, or when parameter is followed by a character which is not to be interpreted as
              part of its name.  The parameter is a shell parameter as described above PARAMETERS) or an  array  refer‐
              ence (Arrays).

       If the first character of parameter is an exclamation point (!), and parameter is not a nameref, it introduces a
       level of variable indirection.  Bash uses the value of the variable formed from the rest  of  parameter  as  the
       name  of  the  variable;  this variable is then expanded and that value is used in the rest of the substitution,
       rather than the value of parameter itself.  This is known as indirect expansion.  If  parameter  is  a  nameref,
       this  expands  to  the  name of the variable referenced by parameter instead of performing the complete indirect
       expansion.  The exceptions to this are the expansions of  ${!prefix*}  and  ${!name[@]}  described  below.   The
       exclamation point must immediately follow the left brace in order to introduce indirection.

       In  each  of the cases below, word is subject to tilde expansion, parameter expansion, command substitution, and
       arithmetic expansion.

       When not performing substring expansion, using the forms documented below (e.g., :-), bash tests for a parameter
       that is unset or null.  Omitting the colon results in a test only for a parameter that is unset.

       ${parameter:-word}
              Use Default Values.  If parameter is unset or null, the expansion of word is substituted.  Otherwise, the
              value of parameter is substituted.
       ${parameter:=word}
              Assign Default Values.  If parameter is unset or null, the expansion of word is  assigned  to  parameter.
              The  value  of  parameter  is  then substituted.  Positional parameters and special parameters may not be
              assigned to in this way.
       ${parameter:?word}
              Display Error if Null or Unset.  If parameter is null or unset, the expansion of word (or  a  message  to
              that effect if word is not present) is written to the standard error and the shell, if it is not interac‐
              tive, exits.  Otherwise, the value of parameter is substituted.
       ${parameter:+word}
              Use Alternate Value.  If parameter is null or unset, nothing is substituted, otherwise the  expansion  of
              word is substituted.
       ${parameter:offset}
       ${parameter:offset:length}
              Substring  Expansion.  Expands to up to length characters of the value of parameter starting at the char‐
              acter specified by offset.  If parameter is @, an indexed array subscripted by @ or *, or an  associative
              array name, the results differ as described below.  If length is omitted, expands to the substring of the
              value of parameter starting at the character specified by offset and extending to the end of  the  value.
              length and offset are arithmetic expressions (see ARITHMETIC EVALUATION below).

              If offset evaluates to a number less than zero, the value is used as an offset in characters from the end
              of the value of parameter.  If length evaluates to a number less than zero, it is interpreted as an  off‐
              set  in  characters  from  the  end of the value of parameter rather than a number of characters, and the
              expansion is the characters between offset and that result.  Note that a negative offset  must  be  sepa‐
              rated from the colon by at least one space to avoid being confused with the :- expansion.

              If  parameter is @, the result is length positional parameters beginning at offset.  A negative offset is
              taken relative to one greater than the greatest positional parameter, so an offset of -1 evaluates to the
              last positional parameter.  It is an expansion error if length evaluates to a number less than zero.

              If  parameter  is  an  indexed  array name subscripted by @ or *, the result is the length members of the
              array beginning with ${parameter[offset]}.  A negative offset is taken relative to one greater  than  the
              maximum index of the specified array.  It is an expansion error if length evaluates to a number less than
              zero.

              Substring expansion applied to an associative array produces undefined results.

              Substring indexing is zero-based unless the positional parameters are used, in which  case  the  indexing
              starts  at  1  by default.  If offset is 0, and the positional parameters are used, $0 is prefixed to the
              list.

       ${!prefix*}
       ${!prefix@}
              Names matching prefix.  Expands to the names of variables whose names begin with prefix, separated by the
              first  character  of  the  IFS  special variable.  When @ is used and the expansion appears within double
              quotes, each variable name expands to a separate word.

       ${!name[@]}
       ${!name[*]}
              List of array keys.  If name is an array variable, expands to the list of array indices  (keys)  assigned
              in  name.   If  name is not an array, expands to 0 if name is set and null otherwise.  When @ is used and
              the expansion appears within double quotes, each key expands to a separate word.

       ${#parameter}
              Parameter length.  The length in characters of the value of parameter is substituted.  If parameter is  *
              or  @,  the value substituted is the number of positional parameters.  If parameter is an array name sub‐
              scripted by * or @, the value substituted is the number of elements in the array.   If  parameter  is  an
              indexed  array  name  subscripted  by  a  negative  number, that number is interpreted as relative to one
              greater than the maximum index of parameter, so negative indices count back from the end  of  the  array,
              and an index of -1 references the last element.

       ${parameter#word}
       ${parameter##word}
              Remove matching prefix pattern.  The word is expanded to produce a pattern just as in pathname expansion.
              If the pattern matches the beginning of the value of parameter, then the result of the expansion  is  the
              expanded  value  of parameter with the shortest matching pattern (the ``#'' case) or the longest matching
              pattern (the ``##'' case) deleted.  If parameter is @ or *, the pattern removal operation is  applied  to
              each  positional  parameter  in  turn, and the expansion is the resultant list.  If parameter is an array
              variable subscripted with @ or *, the pattern removal operation is applied to each member of the array in
              turn, and the expansion is the resultant list.

       ${parameter%word}
       ${parameter%%word}
              Remove matching suffix pattern.  The word is expanded to produce a pattern just as in pathname expansion.
              If the pattern matches a trailing portion of the expanded value of parameter,  then  the  result  of  the
              expansion  is  the expanded value of parameter with the shortest matching pattern (the ``%'' case) or the
              longest matching pattern (the ``%%'' case) deleted.  If parameter is @ or *, the pattern  removal  opera‐
              tion is applied to each positional parameter in turn, and the expansion is the resultant list.  If param‐
              eter is an array variable subscripted with @ or *, the pattern removal operation is applied to each  mem‐
              ber of the array in turn, and the expansion is the resultant list.

       ${parameter/pattern/string}
              Pattern  substitution.   The  pattern  is  expanded  to  produce a pattern just as in pathname expansion.
              Parameter is expanded and the longest match of pattern against its value is  replaced  with  string.   If
              pattern begins with /, all matches of pattern are replaced with string.  Normally only the first match is
              replaced.  If pattern begins with #, it must match at the beginning of the expanded value  of  parameter.
              If  pattern  begins  with  %,  it must match at the end of the expanded value of parameter.  If string is
              null, matches of pattern are deleted and the / following pattern may  be  omitted.   If  the  nocasematch
              shell  option is enabled, the match is performed without regard to the case of alphabetic characters.  If
              parameter is @ or *, the substitution operation is applied to each positional parameter in turn, and  the
              expansion  is the resultant list.  If parameter is an array variable subscripted with @ or *, the substi‐
              tution operation is applied to each member of the array in turn, and the expansion is the resultant list.

       ${parameter^pattern}
       ${parameter^^pattern}
       ${parameter,pattern}
       ${parameter,,pattern}
              Case modification.  This expansion modifies the case of alphabetic characters in parameter.  The  pattern
              is  expanded to produce a pattern just as in pathname expansion.  Each character in the expanded value of
              parameter is tested against pattern, and, if it matches the pattern, its case is converted.  The  pattern
              should  not attempt to match more than one character.  The ^ operator converts lowercase letters matching
              pattern to uppercase; the , operator converts matching uppercase letters to lowercase.   The  ^^  and  ,,
              expansions convert each matched character in the expanded value; the ^ and , expansions match and convert
              only the first character in the expanded value.  If pattern is omitted, it is treated  like  a  ?,  which
              matches  every  character.   If  parameter  is @ or *, the case modification operation is applied to each
              positional parameter in turn, and the expansion is the resultant list.  If parameter is an array variable
              subscripted  with @ or *, the case modification operation is applied to each member of the array in turn,
              and the expansion is the resultant list.

       ${parameter@operator}
              Parameter transformation.  The expansion is either a transformation of the value of parameter or informa‐
              tion about parameter itself, depending on the value of operator.  Each operator is a single letter:

              Q      The  expansion is a string that is the value of parameter quoted in a format that can be reused as
                     input.
              E      The expansion is a string that is the value of parameter with backslash escape sequences  expanded
                     as with the $'...' quoting mechansim.
              P      The  expansion  is a string that is the result of expanding the value of parameter as if it were a
                     prompt string (see PROMPTING below).
              A      The expansion is a string in the form of an assignment statement or declare command that, if eval‐
                     uated, will recreate parameter with its attributes and value.
              a      The expansion is a string consisting of flag values representing parameter's attributes.

              If  parameter is @ or *, the operation is applied to each positional parameter in turn, and the expansion
              is the resultant list.  If parameter is an array variable subscripted with @ or *, the case  modification
              operation is applied to each member of the array in turn, and the expansion is the resultant list.

              The result of the expansion is subject to word splitting and pathname expansion as described below.

   Command Substitution
       Command substitution allows the output of a command to replace the command name.  There are two forms:

              $(command)
       or
              `command`

       Bash  performs  the expansion by executing command in a subshell environment and replacing the command substitu‐
       tion with the standard output of the command, with any trailing newlines deleted.   Embedded  newlines  are  not
       deleted, but they may be removed during word splitting.  The command substitution $(cat file) can be replaced by
       the equivalent but faster $(< file).

       When the old-style backquote form of substitution is used, backslash retains its  literal  meaning  except  when
       followed  by  $,  `, or \.  The first backquote not preceded by a backslash terminates the command substitution.
       When using the $(command) form, all characters between the parentheses make up the  command;  none  are  treated
       specially.

       Command  substitutions  may be nested.  To nest when using the backquoted form, escape the inner backquotes with
       backslashes.

       If the substitution appears within double quotes, word splitting and pathname expansion are not performed on the
       results.

   Arithmetic Expansion
       Arithmetic  expansion allows the evaluation of an arithmetic expression and the substitution of the result.  The
       format for arithmetic expansion is:

              $((expression))

       The old format $[expression] is deprecated and will be removed in upcoming versions of bash.

       The expression is treated as if it were within double quotes, but a double quote inside the parentheses  is  not
       treated specially.  All tokens in the expression undergo parameter and variable expansion, command substitution,
       and quote removal.  The result is treated as the arithmetic expression to be evaluated.   Arithmetic  expansions
       may be nested.

       The  evaluation  is performed according to the rules listed below under ARITHMETIC EVALUATION.  If expression is
       invalid, bash prints a message indicating failure and no substitution occurs.

   Process Substitution
       Process substitution allows a process's input or output to be referred to using a filename.  It takes  the  form
       of  <(list)  or >(list).  The process list is run asynchronously, and its input or output appears as a filename.
       This filename is passed as an argument to the current command as the result of the expansion.   If  the  >(list)
       form  is used, writing to the file will provide input for list.  If the <(list) form is used, the file passed as
       an argument should be read to obtain the output of list.  Process substitution is supported on systems that sup‐
       port named pipes (FIFOs) or the /dev/fd method of naming open files.

       When  available, process substitution is performed simultaneously with parameter and variable expansion, command
       substitution, and arithmetic expansion.

   Word Splitting
       The shell scans the results of parameter expansion, command substitution, and arithmetic expansion that did  not
       occur within double quotes for word splitting.

       The shell treats each character of IFS as a delimiter, and splits the results of the other expansions into words
       using these characters as field terminators.  If IFS is unset, or its value  is  exactly  <space><tab><newline>,
       the default, then sequences of <space>, <tab>, and <newline> at the beginning and end of the results of the pre‐
       vious expansions are ignored, and any sequence of IFS characters not at the beginning or end serves  to  delimit
       words.   If  IFS has a value other than the default, then sequences of the whitespace characters space, tab, and
       newline are ignored at the beginning and end of the word, as long as the whitespace character is in the value of
       IFS (an IFS whitespace character).  Any character in IFS that is not IFS whitespace, along with any adjacent IFS
       whitespace characters, delimits a field.  A sequence of IFS whitespace characters is also treated  as  a  delim‐
       iter.  If the value of IFS is null, no word splitting occurs.

       Explicit null arguments ("" or '') are retained and passed to commands as empty strings.  Unquoted implicit null
       arguments, resulting from the expansion of parameters that have no values, are removed.  If a parameter with  no
       value  is  expanded  within double quotes, a null argument results and is retained and passed to a command as an
       empty string.  When a quoted null argument appears as part of a word whose expansion is non-null, the null argu‐
       ment is removed.  That is, the word -d'' becomes -d after word splitting and null argument removal.

       Note that if no expansion occurs, no splitting is performed.

   Pathname Expansion
       After  word  splitting,  unless the -f option has been set, bash scans each word for the characters *, ?, and [.
       If one of these characters appears, then the word is regarded as a pattern, and replaced with an  alphabetically
       sorted list of filenames matching the pattern (see Pattern Matching below).  If no matching filenames are found,
       and the shell option nullglob is not enabled, the word is left unchanged.  If the nullglob option is set, and no
       matches are found, the word is removed.  If the failglob shell option is set, and no matches are found, an error
       message is printed and the command is not executed.  If the shell option nocaseglob is  enabled,  the  match  is
       performed  without regard to the case of alphabetic characters.  Note that when using range expressions like [a-
       z] (see below), letters of the other case may be included, depending on the setting of LC_COLLATE.  When a  pat‐
       tern  is  used  for  pathname  expansion, the character ``.''  at the start of a name or immediately following a
       slash must be matched explicitly, unless the shell option dotglob is set.  When matching a pathname,  the  slash
       character  must  always  be  matched explicitly.  In other cases, the ``.''  character is not treated specially.
       See the description of shopt below under SHELL BUILTIN COMMANDS for a description of the  nocaseglob,  nullglob,
       failglob, and dotglob shell options.

       The GLOBIGNORE shell variable may be used to restrict the set of filenames matching a pattern.  If GLOBIGNORE is
       set, each matching filename that also matches one of the patterns in GLOBIGNORE is  removed  from  the  list  of
       matches.   If the nocaseglob option is set, the matching against the patterns in GLOBIGNORE is performed without
       regard to case.  The filenames ``.''  and ``..''  are always ignored when GLOBIGNORE is set and not null.   How‐
       ever,  setting  GLOBIGNORE to a non-null value has the effect of enabling the dotglob shell option, so all other
       filenames beginning with a ``.''  will match.  To get the old behavior of ignoring filenames  beginning  with  a
       ``.'', make ``.*''  one of the patterns in GLOBIGNORE.  The dotglob option is disabled when GLOBIGNORE is unset.
       The pattern matching honors the setting of the extglob shell option.

       Pattern Matching

       Any character that appears in a pattern, other than the special  pattern  characters  described  below,  matches
       itself.   The NUL character may not occur in a pattern.  A backslash escapes the following character; the escap‐
       ing backslash is discarded when matching.  The special pattern characters must be  quoted  if  they  are  to  be
       matched literally.

       The special pattern characters have the following meanings:

              *      Matches  any  string, including the null string.  When the globstar shell option is enabled, and *
                     is used in a pathname expansion context, two adjacent *s used as a single pattern will  match  all
                     files  and  zero or more directories and subdirectories.  If followed by a /, two adjacent *s will
                     match only directories and subdirectories.
              ?      Matches any single character.
              [...]  Matches any one of the enclosed characters.  A pair of characters separated by a hyphen denotes  a
                     range expression; any character that falls between those two characters, inclusive, using the cur‐
                     rent locale's collating sequence and character set, is matched.  If the first character  following
                     the  [ is a !  or a ^ then any character not enclosed is matched.  The sorting order of characters
                     in range expressions is determined by the current locale and  the  values  of  the  LC_COLLATE  or
                     LC_ALL  shell  variables,  if set.  To obtain the traditional interpretation of range expressions,
                     where [a-d] is equivalent to [abcd], set value of the LC_ALL shell variable to C,  or  enable  the
                     globasciiranges  shell  option.  A - may be matched by including it as the first or last character
                     in the set.  A ] may be matched by including it as the first character in the set.

                     Within [ and ], character classes can be specified using the syntax [:class:], where class is  one
                     of the following classes defined in the POSIX standard:
                     alnum alpha ascii blank cntrl digit graph lower print punct space upper word xdigit
                     A character class matches any character belonging to that class.  The word character class matches
                     letters, digits, and the character _.

                     Within [ and ], an equivalence class can be specified using the syntax [=c=],  which  matches  all
                     characters with the same collation weight (as defined by the current locale) as the character c.

                     Within [ and ], the syntax [.symbol.] matches the collating symbol symbol.

       If  the extglob shell option is enabled using the shopt builtin, several extended pattern matching operators are
       recognized.  In the following description, a pattern-list is a list of one or more patterns separated  by  a  |.
       Composite patterns may be formed using one or more of the following sub-patterns:

              ?(pattern-list)
                     Matches zero or one occurrence of the given patterns
              *(pattern-list)
                     Matches zero or more occurrences of the given patterns
              +(pattern-list)
                     Matches one or more occurrences of the given patterns
              @(pattern-list)
                     Matches one of the given patterns
              !(pattern-list)
                     Matches anything except one of the given patterns

   Quote Removal
       After  the preceding expansions, all unquoted occurrences of the characters \, ', and " that did not result from
       one of the above expansions are removed.

REDIRECTION
       Before a command is executed, its input and output may be redirected using a special notation interpreted by the
       shell.   Redirection  allows commands' file handles to be duplicated, opened, closed, made to refer to different
       files, and can change the files the command reads from and writes to.  Redirection may also be  used  to  modify
       file  handles  in  the  current shell execution environment.  The following redirection operators may precede or
       appear anywhere within a simple command or may follow a command.  Redirections are processed in the  order  they
       appear, from left to right.

       Each  redirection that may be preceded by a file descriptor number may instead be preceded by a word of the form
       {varname}.  In this case, for each redirection operator except >&- and <&-,  the  shell  will  allocate  a  file
       descriptor  greater  than  or equal to 10 and assign it to varname.  If >&- or <&- is preceded by {varname}, the
       value of varname defines the file descriptor to close.

       In the following descriptions, if the file descriptor number is omitted, and the first character  of  the  redi‐
       rection operator is <, the redirection refers to the standard input (file descriptor 0).  If the first character
       of the redirection operator is >, the redirection refers to the standard output (file descriptor 1).

       The word following the redirection operator in the following descriptions, unless otherwise noted, is  subjected
       to  brace  expansion, tilde expansion, parameter and variable expansion, command substitution, arithmetic expan‐
./selpg: done
```

