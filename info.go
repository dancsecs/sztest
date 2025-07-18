/*
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
*/

/*
Package sztest implements some general go testing helper functions to
provide for cleaner more readable tests as well as automatic diffs
of unexpected results.  In addition to providing general tests it
also provides builtin io interfaces that can be used to simulate io
errors for code tests.  Finally it provides for the capturing of logs and
standard output streams with automatic diffs.
*/
package sztest
