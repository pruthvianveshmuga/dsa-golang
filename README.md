# Data structures and Algorithms

## What is it:

- This repository consists of select questions and their solutions picked from leetcode. These questions are assigned ids as in leetcode.

## How to add a new question:

- Create a new directory with the question `id`.
- Inside this directory:
  - Create a `README.md` with the description and the link to the question.
  - Create a unique package with the question id.
  - Add multiple solutions in different files named `solution1.go`, `solution2.go`, in that package.
  - Add a test file eding with `_test.go` and write tests for all your solutions.

## How to run tests:

- In the home directory, run `go test ./... -v`
