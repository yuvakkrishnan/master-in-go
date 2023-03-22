Now, imagine that printNumbers and generateNumbers each takes three seconds to run. When running synchronously, or one after the other like the last example, your program would take six seconds to run. 
First, printNumbers would run for three seconds, and then generateNumbers would run for three seconds. In your program, however, these two functions are independent of the other because they don’t rely on data from the other to run. 
You can take advantage of this to speed up this hypothetical program by running the functions concurrently using goroutines. When both functions are running concurrently the program could, in theory, run in half the time. 
If both the printNumbers and the generateNumbers functions take three seconds to run and both start at exactly the same time, the program could finish in three seconds. 
(The actual speed could vary due to outside factors, though, such as how many cores the computer has or how many other programs are running on the computer at the same time.)

Running a function concurrently as a goroutine is similar to running a function synchronously. 
To run a function as a goroutine (as opposed to a standard synchronous function), you only need to add the go keyword before the function call.

However, for the program to run the goroutines concurrently, you’ll need to make one additional change. You’ll need to add a way for your program to wait until both goroutines have finished running. 
If you don’t wait for your goroutines to finish and your main function completes, the goroutines could potentially never run, or only part of them could run and not complete running.

To wait for the functions to finish, you’ll use a WaitGroup from Go’s sync package. 
The sync package contains “synchronization primitives”, such as WaitGroup, that are designed to synchronize various parts of a program. 
In your case, the synchronization keeps track of when both functions have finished running so you can exit the program.

The WaitGroup primitive works by counting how many things it needs to wait for using the Add, Done, and Wait functions. 
The Add function increases the count by the number provided to the function, and Done decreases the count by one. 
The Wait function can then be used to wait until the count reaches zero, meaning that Done has been called enough times to offset the calls to Add. Once the count reaches zero, the Wait function will return and the program will continue running.

Next, update the code in your main.go file to run both of your functions as goroutines using the go keyword, and add a sync.WaitGroup to the program: