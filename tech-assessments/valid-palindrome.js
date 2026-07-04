/**
 * @param {string} s
 * @return {boolean}
 */
var isPalindrome = function(s) {
    let left = 0, right = s.length - 1

    while (left <= right) {
        if (!isAlnum(s[left])) {
            left += 1
            continue
        }

        if (!isAlnum(s[right])) {
            right -= 1
            continue
        }

        if (s[left].toLocaleLowerCase() !== s[right].toLocaleLowerCase()) {
            return false
        }

        left += 1
        right -= 1
    }

    return true
};

function isAlnum(char) {
    const code = char.charCodeAt(0);
    return (code >= 48 && code <= 57) || (code >= 65 && code <= 90) || (code >= 97 && code <= 122);
}

console.log(isPalindrome("A man, a plan, a canal: Panama"))
console.log(isPalindrome("race a car"))
console.log(isPalindrome(" "))
