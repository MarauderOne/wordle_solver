# Wordle Solver
A Go powered web app for solving Wordle puzzles.

## Description
The purpose of this app is to provide a webpage where users (who are stuck on today's Wordle problem) can enter their current guesses and receive a list of all the possible five-letter words that could potentially be the solution, after taking the characters and colors into account.

The user does by using the soft-keyboard on the page itself (or physical keyboard on desktop PCs) to enter the characters. They can then add the colors returned by Wordle by clicking on the box in each row.

The characters, colors and position of each box will then be computed in order to find all possible solutions to the puzzle (given the words that the current guesses have eliminated).

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
    D{"Is character\nalphabetic?"}
    E([Return Error])
    H[Ignore slices where either\nthe character or color is blank]
    I["Determine position in row\n(First to Fifth)"]
    J["Set regex pattern for Green box\n(Word contains this character any this position)"]
    K["Set regex pattern for Yellow box\n(Word contains this character in any position except this one)"]
    L{"Does this\ncharacter exist\nanywhere else\nin the row?"}
    M["Set regex pattern for Grey box\n(Word does not contain this character in this position)"]
    N["Set regex pattern for Grey box\n(Word does not contain this character any position)"]
    O{"What is the\ncolor of the box\nin this iteration\nof the loop?"}
    P[Revise answer list using\nregex pattern for Green boxes]
    Q[Revise answer list using\nregex pattern for Yellow boxes]
    R[Revise answer list using\nregex pattern for Grey boxes]
    S([Error out for unrecognised colors])
    T{"Has the count\nof answers reached\none or fewer?"}
    U["Break the loop early\n(No point iterating further at this point)"]
    V{"Have we reached\nthe end of the loop?"}
    W{"Is the count of\nanswers still 5237?"}
    X([Tell the user to keep entering guess data])
    Y([Return the results to the user])

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
    K --> L
    L -->|Yes| M
    L -->|No| N
    M --> O
    N --> O
    O -->|Green| P
    O -->|Yellow| Q
    O -->|Grey| R
    O -->|Unrecognised color| S
    P --> T
    Q --> T
    R --> T
    T -->|Yes| U
    T -->|No| V
    U --> W
    V -->|No| G
    V -->|Yes| W
    W -->|Yes| X
    W -->|No| Y
```

## Prerequisites for Local Execution
- Linux
- Go 1.22

## External links
- [NYT Wordle](https://www.nytimes.com/games/wordle/index.html)
