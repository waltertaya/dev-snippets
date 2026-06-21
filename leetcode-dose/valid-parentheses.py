def isValid(s):
    """
    :type s: str
    :rtype: bool
    """
    #base case
    # if len(s) == 1 or set(s) <= set(')]}'):
    #     return False
    if len(s) == 1 or set(s).issubset(set(')]}')) or s[0] in ')]}':
        return False

    stack = []

    for p in s:
        if p in '([{':
            stack.append(p)
        elif p == ')' and len(stack) > 0 and stack[-1] == '(':
            stack.pop()
        elif p == ']' and len(stack) > 0 and stack[-1] == '[':
            stack.pop()
        elif p == '}' and len(stack) > 0 and stack[-1] == '{':
            stack.pop()
        else:
            return False
    
    return len(stack) == 0


if __name__ == '__main__':
    print(isValid("()"))
    print(isValid("()[]{}"))
    print(isValid("(]"))
    print(isValid("([])"))
    print(isValid("([)]"))
    print(isValid("]"))
    print(isValid("}])"))
    print(isValid("){"))
    print(isValid(")(){}"))
    print(isValid("(])"))
