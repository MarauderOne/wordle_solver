# Wordle Solver
A Go powered web app for solving Wordle puzzles.

## Description
The purpose of this app is to provide a webpage where users (who are stuck on today's Wordle problem) can enter their current guesses and receive a list of all the possible five-letter words that could potentially be the solution, after taking the characters and their colors into account.

The user does this by using the soft-keyboard on the page itself (or physical keyboard on desktop PCs) to enter the characters. They can then add the colors returned by Wordle by clicking on the box in each row.

The characters, their colors and position of each box will then be computed in order to find all possible solutions to the puzzle (given the words that the current guesses have eliminated).

## Isn't That Cheating Though?
Oh, for _you_ it is, yes, absolutely. I mean, you're using a third-party tool to help you solve the Wordle. How can you sleep at night?

For me on the other hand, not so much. The way I see it - Wordle is a test of your mental ability to correctly deduce the answer word based on the clues provided by the characters and colors in the squares. I've simply employed my mental abilities a bit differently, in that I have built an application to help suggest possible answers. :grin:

In reality, I don't use this tool very often. Its primary purpose was to act as a project that allowed me to build up my coding skills in Go, HTML, JavaScript, etc.

## Code Logic
Below is a diagram of the logical steps made by the Wordle Solver Go code in order to return a list of possible answers:
```mermaid
flowchart TD
%%{ init: { 'flowchart': { 'curve': 'stepBefore' } } }%%

    A[/JSON array of user guesses/]
    B[Parsed into Go Struct]
    C[Characters converted to uppercase]
    F[Create initial answer list of all 5237 possible answers]
    G[For each slice in the struct]
    D{Is character alphabetic?}
    E([Return Error])
    H[Ignore slices where either the character or color is blank]
    I[Single-letter logic function]
    J[Determine position in row]
    K{Is the box green?}
    L[Set the regex pattern characters in this postion in green boxes]
    M{Is the box yellow?}
    N[Set the regex pattern characters in this postion in yellow boxes]
    O[Set the regex pattern characters in this postion in grey boxes]
    P[Return regex pattern]
    Q[Multi-letter logic function]
    R[Determine position in row]
    S{Is the box green?}
    T[Determine the presence of concurrent characters and the colors of their corresponding boxes, set regex patterns accordingly]
    U{Is the box yellow?}
    V[Determine the presence of concurrent characters and the colors of their corresponding boxes, set regex patterns accordingly]
    W[Determine the presence of concurrent characters and the colors of their corresponding boxes, set regex patterns accordingly]
    X[Return regex pattern]
    Y{Is the current box at this iteration of the loop a valid color?}
    Z([Error out for unrecognised colors])
    AA[Revise the list of potential answers using regex pattern from the single-letter logic function]
    AB{Were we able to set a regex pattern for multi-letter logic?}
    AC[Revise the list of potential answers using regex pattern from the multi-letter logic function]
    AD{Has the count of answers reached one or fewer?}
    AE(["Break the loop early (No point iterating further at this point)"])
    AF{Have we reached the end of the loop?}
    AG{Is the count of answers still 5237?}
    AH([Tell the user to keep entering guess data])
    AI([Return the results to the user])

    A -->|POST to /guesses| B
    B --> C
    C --> F
    F --> G
    G --> D
    D -->|No| E
    D -->|Yes| H
    H --> I
    I --> J
    J --> K
    K -->|Yes| L
    L --> P
    K -->|No| M
    M -->|Yes| N
    N --> P
    M -->|No| O
    O --> P
    P --> Q
    Q --> R
    R --> S
    S -->|Yes| T
    T --> X
    S -->|No| U
    U -->|Yes| V
    V --> X
    U -->|No| W
    W --> X
    X --> Y
    Y -->|Yes| AA
    Y -->|No| Z
    AA --> AB
    AB -->|Yes| AC
    AB -->|No| AD
    AC --> AD
    AD -->|Yes| AE
    AD -->|No| AF
    AF -->|Yes| AG
    AF -->|No| G
    AG -->|Yes| AH
    AG -->|No| AI

```

## Prerequisites for Local Execution
- Linux
- Go 1.22

## External links
- [NYT Wordle](https://www.nytimes.com/games/wordle/index.html)
