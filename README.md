# Data structures and Algorithms

## What is it:

- This repository consists of select questions and their solutions picked from leetcode. These questions are assigned ids as in leetcode.

## How to add a new question:

- Create a new directory with the question `id` inside `src`.
- Inside this directory:
  - Create a `README.md` with the link to the question.
  - Create a package with the question id.
  - Add solutions in files named `solution1.go`, `solution2.go`.
  - Add a test file eding with `_test.go` to test the written solutions.

## How to run tests:

- In the home directory, run `go test ./... -v`
