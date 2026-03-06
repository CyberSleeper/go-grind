The Problem: The High-Throughput Order Sequencer

Your task is to build the core ID generator for our matching engine. Every single order that hits our servers must immediately request a SequenceID from your component before it goes to the execution queue.
The Requirements

    Strict Monotonicity: Every generated ID must be strictly greater than the previous one, even if 10,000 goroutines call it at the exact same millisecond.

    Time-Awareness: The ID should ideally embed the current UNIX millisecond timestamp so that we can globally sort orders by time across different microservices.

    Collision Resolution: If multiple requests happen in the same millisecond, you must use an internal logical sequence counter to separate them.

    High Throughput: This is on the critical path for every trade. It needs to be incredibly fast and thread-safe.

Your Objective

Implement the NextID() method for the OrderSequencer struct.

To keep it simple for this 45-minute window, you don't need to build a full 64-bit bitwise Snowflake ID. Instead, return a composite int64 (or string if you prefer) that guarantees strict FCFS ordering under massive concurrent load.