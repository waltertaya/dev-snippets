from typing import List
'''
A simple code folding engine needs to identify every scope in a source string.
Each scope begins with { and ends with its matching }.
The input may contain ordinary soource text between scopes.
This text does not affect the result.
Braces inside // line comments or /*...*/ block comments must also be ignored.

Every scope receives a hierarchical identifier.
The single outer scope is named 1. Children are numbered from one in source order.
A child identifier is formed by appending its sibling number to its parent identifier with a period.

For example, the second child of scope 1.1 is named 1.1.2.

REQUIREMENTS:
Implement a function named Parse that accepts a source string and returns the identifiers of every scope it contains.
1. Return an empty collection when the input is empty
2. Assign 1 to the outer scope.
3. Number sibling scope from one according to the order in which their opening braces appear.
4. Build each nested identifier by appending a period and the scope's sibling number to its parent identifier.
5. Ignore all nonbrace source text.
6. Ignore braces inside // line comments and /* ... */ block comments.
7. Return the identifiers in any order

USE CASES:
# Use case 1: Parse a source string containing only its outer scope.
print(" ".join(Parse("{ execute(); }"))) # output: 1

Use case 2: Number nested siblings according to source order
print(" ".join(Parse("{ { { } { } } }"))) # Output: 1 1.1 1.1.1 1.1.2

Use case 3: Ignore braces that occur inside comments.
print(" ".join(Parse("{ /* ignored { } */ { } // ignored {\n}"))) # Output 1 1.1
'''

def Parse(input_code: str) -> List[str]:
    stack = []
    identifiers = []
    i = 0
    n = len(input_code)

    in_block_comment = False
    in_line_comment = False

    while i < n:
        if in_block_comment and i + 1 < n:
            if input_code[i] == '*' and input_code[i+1] == '/':
                in_block_comment = False
                i += 2
            else:
                i += 1
            continue

        if in_line_comment:
            if input_code[i] == '\n':
                in_line_comment = False
            i += 1
            continue

        if input_code[i] == '/' and i + 1 < n:
            if input_code[i + 1] == '*':
                in_block_comment = True
                i += 2
                continue

            if input_code[i + 1] == '/':
                in_line_comment = True
                i += 2
                continue

        if input_code[i] == '{':
            if not stack:
                identifier = '1'
            else:
                parent_id, child_number = stack[-1]
                child_number += 1

                stack[-1] = (parent_id, child_number)
                identifier = f'{parent_id}.{child_number}'

            stack.append((identifier, 0))
            identifiers.append(identifier)

        if input_code[i] == '}':
            if stack:
                stack.pop()

        i += 1

    return identifiers


if __name__ == '__main__':
    print(" ".join(Parse("{ execute(); }")))
    print(" ".join(Parse("{ { { } { } } }")))
    print(" ".join(Parse("{ /* ignored { } */ { } // ignored {\n}")))
    print(" ".join(Parse("{ { } { } }")))
    print(" ".join(Parse("{ { } { { } } }")))
