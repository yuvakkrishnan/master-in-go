Before we dive into the code, let’s have a quick look at the theory and why we need mutexes at all.

So, a mutex, or a mutual exclusion is a mechanism that allows us to prevent concurrent processes from entering a critical section of data whilst it’s already being executed by a given process.

Let’s think about an example where we have a bank balance and a system that both deposits and withdraws sums of money from that bank balance. Within a single threaded, synchronous program, this would be incredibly easy. We could effectively guarantee that it would work as intended every time with a small battery of unit tests.

However, if we were to start introducing multiple threads, or multiple goroutines in Go’s case, we may start to see issues within our code.

Imagine we had a customer with a balance of £1,000.
The customer deposits £500 to his account
One goroutine would see this transaction, read the value at £1,000 and proceed to add the £500 to the existing balance.
However, at the same moment, a charge of £700 is applied to his account to pay for his mortgage.
This second process reads the account balance of £1,000 before the first is able to add the additional deposit of £500 and proceeds to subtract £700 from his account.
The customer checks his bank balance the next day and notices he is down to £300 as the second process was unaware of the first deposit and overwrote the value upon completion.
Obviously, if you were the customer, you would be pretty pissed that the bank had somehow “lost” your deposit of £500 and you would switch banks.

This right here, is an example of a race condition and how, if we aren’t careful, our concurrent programs can start to see issues when we don’t carefully guard the critical sections in our code.