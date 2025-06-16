/* 412: FizzBuzz
Given an integer n, return a string array answer (1-indexed) where:
    answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
    answer[i] == "Fizz" if i is divisible by 3.
    answer[i] == "Buzz" if i is divisible by 5.
    answer[i] == i (as a string) if none of the above conditions are true.

Solution by: https://www.github.com/waltertaya
*/

Console.Write("Enter the value of n: ");
int n = Convert.ToInt32(Console.ReadLine());

Console.WriteLine("\n");

// output
string[] output = new string[Convert.ToInt32(n)];

// iteration from 1 through to n
for (int i = 1; i <= n; i++)
{
    // if i is divisble by 3 & 5
    if ((i % 3 == 0) && (i % 5 == 0))
    {
        // Console.WriteLine($"{i} is divisible by 3 and 5");
        output[(i-1)] = "FizzBuzz";
    }
    else if ((i % 3 == 0) && (i % 5 != 0)) // if i divisible by 3 only
    {
        // Console.WriteLine($"{i} is divisible by 3");
        output[(i-1)] = "Fizz";
    }
    else if ((i % 3 != 0) && (i % 5 == 0)) // if i divisible by 5 only
    {
        // Console.WriteLine($"{i} is divisible by 5");
        output[(i-1)] = "Buzz";
    }
    else // if none of the above are true return i as string
    {
        // Console.WriteLine($"{i} is NOT divisible by 3 and 5");
        output[(i - 1)] = Convert.ToString(i);
    }
}

Console.WriteLine($"Input: {n}");
// display the output
Console.Write("Output:");
foreach (string i in output)
{
    Console.Write($" {i}");
}

Console.WriteLine("\n");
