<!---
   Golang test helper library: sztest.
   Copyright (C) 2023, 2024 Leslie Dancsecs

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
-->

# Package sztest Configuration

- [Example: Default Markup](#example-default-markup)
- [Example: Ascii Markup](#example-ascii-markup)
- [Example: Unicode Markup](#example-unicode-markup)
- [Appendix E: Builtin Ansi Terminal Markup](#appendix-e-builtin-ansi-terminal-markup)

Defaults can be overridden by defining various environment variables. Below
each environment variable will be listed with their default values which are
all strings.

## General

```bash
SZTEST_FAIL_FAST="True"
```

> Sets how check errors are reported to the``` t *testingT ```object.  If true
then errors are reported using the``` t.Fatal ```method causing the current
test to stop at that point.  If set to false then the method``` t.Error ```
which logs the error but permits the test to continue possibly registering
further errors.

```bash
SZTEST_BUFFER_SIZE="10000"
```

> Sets how large internal logging buffers are created.  NOTE:  They will grow
to accommodate more information.

### Temporary Files

```bash
SZTEST_PERM_DIR="0700"
```

> Uses the provided octal permissions when creating temporary directories.

```bash
SZTEST_PERM_FILE="0600"
```

> Uses the provided octal permissions when creating temporary files.

```bash
SZTEST_PERM_EXE="0700"
```

> Uses the provided octal permissions when creating executable temporary files.

```bash
SZTEST_TMP_DIR="" # Uses go's os.TempDir() method as the default
```

> Set the temporary directory.  The default is set to ```os.TempDir()```.  The
directory returned by this function is different from the system temporary
directory when running a test.  This directory is deleted when the test
exits.
>> If a custom directory is set (even if set to the default temporary
directory) then temporary files and directories will only be deleted if the
test is successful.  They are retained (until the test is run again) for
review and debugging on failure.  If the function ```chk.KeepTmpFiles()``` is
invoked from the test then the temporary files and directories will not be
deleted even if the test is successful.

## Difference Windows

```bash
SZTEST_DIFF_CHARS="3"
```

> When comparing two items the minimum number of matching consecutive
characters required to consider that the two sections are equivalent.

```bash
SZTEST_DIFF_SLICE="1"
```

> When comparing two slices the minimum number of matching consecutive lines
required to consider that the two sections are equivalent.

## Difference Markup

Various areas of the output are highlighted (framed) with strings either
defaulted or overridden.

There are builtin values that represent colors
applicable to an ANSI terminal.  Multiple builtins are permitted per entry
separated by the string ```"_and_"``` with one foreground, one background and
multiple styles permitted.  For a complete list of builtins see
[Appendix E](#appendix-e-builtin-ansi-terminal-markup).

```bash
SZTEST_MARK_INS_ON="GREEN_and_REVERSE" # "\x1b[32m\x1b[7m"
SZTEST_MARK_INS_OFF="DEFAULT"          # "\x1b[0m"
```

> Areas found in the got but missing in the want are prefixed/suffixed with
these strings.

```bash
SZTEST_MARK_DEL_ON="RED_and_REVERSE"   # "\x1b[31m\x1b[7m"
SZTEST_MARK_DEL_OFF="DEFAULT"          # "\x1b[0m"
```

> Areas found in the want but missing in the got are prefixed/suffixed with
these strings.

```bash
SZTEST_MARK_CHG_ON="BLUE_and_REVERSE"  # "\x1b[34m\x1b[7m"
SZTEST_MARK_CHG_OFF="DEFAULT"          # "\x1b[0m"
SZTEST_MARK_SEP_ON="BK-YELLOW"         # "\x1b[43m"
SZTEST_MARK_SEP_OFF="DEFAULT"          # "\x1b[0m"
```

> Areas that are different between the got and the want.  The separator markup
highlights the slash '/' used to separate the old value from the new.

```bash
SZTEST_MARK_WNT_ON="CYAN"                           # "\x1b[36m"
SZTEST_MARK_WNT_OFF="DEFAULT"                       # "\x1b[0m"
SZTEST_MARK_GOT_ON="MAGENTA"                        # "\x1b[35m"
SZTEST_MARK_GOT_OFF="DEFAULT"                       # "\x1b[0m"
SZTEST_MARK_MSG_ON="BOLD_and_ITALIC_and_UNDERLINE"  # "\x1b[1m\x1b[3m\x1b[4m"
SZTEST_MARK_MSG_OFF="DEFAULT"                       # "\x1b[0m"
```

> Various portions of the error output are highlighted with provided markup
making the results easier to interpret quickly.

[Contents](README.md#contents)

## Example: Default Markup

```bash
SZTEST_FAIL_FAST="True"
SZTEST_BUFFER_SIZE="10000"

SZTEST_PERM_DIR="0700"
SZTEST_PERM_FILE="0600"
SZTEST_PERM_EXE="0700"
SZTEST_TMP_DIR="/tmp" # Uses go's os.TempDir() method as the default

SZTEST_DIFF_CHARS="3"
SZTEST_DIFF_SLICE="1"

SZTEST_MARK_WNT_ON="CYAN"                           # "\x1b[36m"
SZTEST_MARK_WNT_OFF="DEFAULT"                       # "\x1b[0m"
SZTEST_MARK_GOT_ON="MAGENTA"                        # "\x1b[35m"
SZTEST_MARK_GOT_OFF="DEFAULT"                       # "\x1b[0m"
SZTEST_MARK_MSG_ON="BOLD_and_ITALIC_and_UNDERLINE"  # "\x1b[1m\x1b[3m\x1b[4m"
SZTEST_MARK_MSG_OFF="DEFAULT"                       # "\x1b[0m"

SZTEST_MARK_INS_ON="GREEN_and_REVERSE" # "\x1b[32m\x1b[7m"
SZTEST_MARK_INS_OFF="DEFAULT"          # "\x1b[0m"
SZTEST_MARK_DEL_ON="RED_and_REVERSE"   # "\x1b[31m\x1b[7m"
SZTEST_MARK_DEL_OFF="DEFAULT"          # "\x1b[0m"
SZTEST_MARK_CHG_ON="BLUE_and_REVERSE"  # "\x1b[34m\x1b[7m"
SZTEST_MARK_CHG_OFF="DEFAULT"          # "\x1b[0m"
SZTEST_MARK_SEP_ON="BK-YELLOW"         # "\x1b[43m"
SZTEST_MARK_SEP_OFF="DEFAULT"          # "\x1b[0m"
```

[Contents](README.md#contents)

## Example: Ascii Markup

```bash
SZTEST_FAIL_FAST="True"
SZTEST_BUFFER_SIZE="10000"

SZTEST_PERM_DIR="0700"
SZTEST_PERM_FILE="0600"
SZTEST_PERM_EXE="0700"
SZTEST_TMP_DIR="/tmp" # Uses go's os.TempDir() method as the default

SZTEST_DIFF_CHARS="3"
SZTEST_DIFF_SLICE="1"

SZTEST_MARK_WNT_ON=""
SZTEST_MARK_WNT_OFF=""
SZTEST_MARK_GOT_ON=""
SZTEST_MARK_GOT_OFF=""
SZTEST_MARK_MSG_ON=""
SZTEST_MARK_MSG_OFF=""

SZTEST_MARK_INS_ON="<+"
SZTEST_MARK_INS_OFF="+>"
SZTEST_MARK_DEL_ON="<-"
SZTEST_MARK_DEL_OFF="->"
SZTEST_MARK_CHG_ON="<<"
SZTEST_MARK_CHG_OFF=">>"
SZTEST_MARK_SEP_ON=""
SZTEST_MARK_SEP_OFF=""
```

[Contents](README.md#contents)

## Example: Unicode Markup

```bash
SZTEST_FAIL_FAST="True"
SZTEST_BUFFER_SIZE="10000"

SZTEST_PERM_DIR="0700"
SZTEST_PERM_FILE="0600"
SZTEST_PERM_EXE="0700"
SZTEST_TMP_DIR="/tmp" # Uses go's os.TempDir() method as the default

SZTEST_DIFF_CHARS="3"
SZTEST_DIFF_SLICE="1"

SZTEST_MARK_WNT_ON=""
SZTEST_MARK_WNT_OFF=""
SZTEST_MARK_GOT_ON=""
SZTEST_MARK_GOT_OFF=""
SZTEST_MARK_MSG_ON=""
SZTEST_MARK_MSG_OFF=""

SZTEST_MARK_INS_ON="⨭"
SZTEST_MARK_INS_OFF="⨮"
SZTEST_MARK_DEL_ON="⨴"
SZTEST_MARK_DEL_OFF="⨵"
SZTEST_MARK_CHG_ON="«"
SZTEST_MARK_CHG_OFF="»"
SZTEST_MARK_SEP_ON=""
SZTEST_MARK_SEP_OFF=""
```

[Contents](README.md#contents)

## Appendix E: Builtin Ansi Terminal Markup

Markup overrides may use the following strings to represent standard Ansi
Terminal escape codes representing various colors and styles that may be
represented. Only one background color and one foreground color may be defined
per variable while multiple styles may be specified.

NOTE: If the style ```DEFAULT``` is chosen than no other styles, colors or
custom strings may be specified.

#### Foreground Colors

```go
    // Regular Colors.
    BLACK
    RED
    GREEN
    YELLOW
    BLUE
    MAGENTA
    CYAN
    WHITE

    // Bright Colors.
    HI-BLACK
    HI-RED
    HI-GREEN
    HI-YELLOW
    HI-BLUE
    HI-MAGENTA
    HI-CYAN
    HI-WHITE
```

#### Background Colors

```go
    // Background Regular Colors.
    BK-BLACK
    BK-HI-BLACK
    BK-RED
    BK-HI-RED
    BK-GREEN
    BK-YELLOW
    BK-BLUE
    BK-MAGENTA
    BK-CYAN
    BK-WHITE

    // Background Bright Colors.
    BK-HI-BLACK
    BK-HI-RED
    BK-HI-GREEN
    BK-HI-YELLOW
    BK-HI-BLUE
    BK-HI-MAGENTA
    BK-HI-CYAN
    BK-HI-WHITE
```

#### Styles

```go
    DEFAULT
    BOLD
    ITALIC
    UNDERLINE
    REVERSE
    STRIKEOUT
```

[Contents](README.md#contents)
