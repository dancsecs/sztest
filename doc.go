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
Package sztest provides a self-contained test helper library built entirely
on the Go standard library. It is designed to make tests cleaner, more
readable, and more reliable, while offering features that go beyond the
default testing framework.

Core features include:

  - Uniform assertions across all built-in types, with consistent reporting.
  - Automatic diffs on failure, rendered with ANSI colors for clarity.
    Diff behavior is configurable, including character- and line-window sizes.
  - Flow control with FailFast, allowing tests to stop on the first error
    or continue gathering results.
  - String helpers (Str, Strf) for concise assertions on string values.
  - Support for slice comparisons and interval checks (bounded and unbounded).
  - Error and panic assertions for verifying expected failures.
  - Output capture of stdout, stderr, and package logs, with diffs against
    expected results.
  - Temporary resource and environment variable helpers to isolate tests.
  - I/O interface shims (io.Reader, io.Writer, io.Seeker, io.Closer) for
    simulating success and failure modes in code under test.
  - Clock utilities to capture and format test timestamps in multiple layouts.
  - Full integration with testing.T through a minimal internal interface,
    enabling sztest to be tested itself with complete coverage.

The library emphasizes a minimal usage pattern:

	chk := sztest.CaptureNothing(t)
	defer chk.Release()
	chk.Str(got, wnt)

By keeping the API uniform and predictable, sztest helps reduce boilerplate
and highlight only what matters in a test: the behavior being verified.

# Dedication

This project is dedicated to Reem. Your brilliance, courage, and quiet
strength continue to inspire me. Every line is written in gratitude for the
light and hope you brought into my life.

---

NOTE: Documentation reviewed and polished with the assistance of ChatGPT from
OpenAI.
*/
package sztest
