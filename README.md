# Name:Anran Li 
# Student ID:C00292764

# 1.barrier
- **Function**：Basic barrier synchronization, all goroutines must complete "Part A" before collectively proceeding to "Part B".
- **Implementation**:
  - Uses 'sync.Mutex' to protect the counter 'count', ensuring thread-safe counting in concurrent scenarios.
  - The last goroutine to finish "Part A" closes the channel 'barriers', which unblocks all goroutines to execute "Part B".

# 2.Rendezvous 
- **Function**：Rendezvous synchronization, after goroutines arrive at "Part A" with random delays, they all proceed to "Part B" collectively.
- **Implementation**:
  - 'sync.Mutex' guards the arrival counter 'arrived'. The last arriving goroutine sends a signal to the channel 'barrier', and other goroutines block on channel reception to achieve "collective execution after all arrive".

# 3.barrier2 
- **Function**：Reusable barrier, supports multiple rounds of barrier synchronization with optimized performance.
- **Implementation**:
  - Replaces 'Mutex' with atomic operations to reduce overhead and improve concurrent counting efficiency.
  - When the last goroutine arrives, it sends a channel signal and resets the atomic counter, enabling the barrier to be reused in multi-round scenarios
  - Leverages an unbuffered channel 'theChan' to implement synchronous blocking between goroutines.
 
# 4. dinPhil
- **Function**： Implements a concurrent solution to the Dining Philosophers problem. Five philosophers simultaneously execute the cycle of thinking, taking forks, eating, and putting down forks, while safely sharing five forks without deadlock.Employs a uniform acquisition order strategy: all philosophers always pick up the smaller-numbered fork first, then the larger-numbered fork. This unified behavior completely eliminates circular wait, ensuring deadlock freedom.
- **Implementation**:
  - Each fork is represented by a channel with buffer size 1.
  - When acquiring forks, the left and right fork numbers are compared to ensure the smaller-numbered fork is taken first.
  - Forks are released in the reverse order, maintaining operational symmetry.
# 5. life_game/gol
- **Function**：Parallelizes the next generation computation of Conway's Game of Life so that multiple goroutines can update the board at the same time.
- **Implementation**:
  - The 300×300 grid is split into 3 row segments, and each segment is handled by one goroutine.
  - Each goroutine reads neighbors from the current grid and writes the results only to its own rows in the buffer, so there is no write conflict.
  - A sync.WaitGroup is used as a barrier: the main goroutine waits until all worker goroutines finish this generation, then swaps buffer and grid to complete the update.

