INterview Question  :EX = 1:
Description
The Shout-Out-Loud sequence is a sequence of digit strings defined by the recursive formula:

    • ShoutOutLoud(1) = “1”

    • ShoutOutLoud(n) is the way you would "say" the digit string from ShoutOutLoud(n-1), which is then converted into a different digit string.

To determine how you "say" a digit string, split it into the minimal number of groups so that each group is a contiguous section all of the same character. Then for each group, say the number of characters, then say the character. To convert the saying into a digit string, replace the counts with a number and concatenate every saying.

For example, the saying and conversion for digit string “3322251” will be:

Two 3’s, three 2’s, one 5, and one 1 -> 2 3 + 3 2 + 1 5 + 1 1 -> “23321511”

Given an integer n, return the nth term of the ShoutOutLoud sequence as a string.

Input format –

    • The first line of the input contains the integer n.

Output format –

    • A string representing nth term of the ShoutOutLoud sequence.

 

Constraints
1<=n<=30



Examples
Input:
4

Output:
1211

Explanation:
ShoutOutLoud(1) = "1"

ShoutOutLoud (2) = say "1" = one 1 = "11"

ShoutOutLoud (3) = say "11" = two 1's = "21"

ShoutOutLoud (4) = say "21" = one 2 + one 1 = "12" + "11" = "1211"