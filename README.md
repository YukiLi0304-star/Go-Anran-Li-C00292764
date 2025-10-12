# Name:Anran Li 
# Student ID:C00292764

# 1.barrier 
- **Function**：Basic barrier synchronization, all goroutines must complete "Part A" before collectively proceeding to "Part B".
-**Implementation**:
  - Uses 'sync.Mutex' to protect the counter 'count', ensuring thread-safe counting in concurrent scenarios.
  - The last goroutine to finish "Part A" closes the channel 'barriers', which unblocks all goroutines to execute "Part B".

# 2.Rendezvous 
- **Function**：Rendezvous synchronization, after goroutines arrive at "Part A" with random delays, they all proceed to "Part B" collectively.
-**Implementation**:
  - 'sync.Mutex' guards the arrival counter 'arrived'. The last arriving goroutine sends a signal to the channel 'barrier', and other goroutines block on channel reception to achieve "collective execution after all arrive".

 # 3. barrier2 
- **Function**：Reusable barrier, supports multiple rounds of barrier synchronization with optimized performance.
-**Implementation**:
  - Replaces 'Mutex' with atomic operations to reduce overhead and improve concurrent counting efficiency.
  - When the last goroutine arrives, it sends a channel signal and resets the atomic counter, enabling the barrier to be reused in multi-round scenarios
  - Leverages an unbuffered channel 'theChan' to implement synchronous blocking between goroutines.
