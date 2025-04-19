# 🔥 High-Performance FEP Server with gnet

This project uses [gnet](https://github.com/panjf2000/gnet), a blazing-fast, lightweight, non-blocking networking framework in Go, to build a high-performance TCP server for financial external platform (FEP) communication.

---

## 🏦 Used at Toss Securities

At **Toss Securities**, we operate mission-critical FEP (Financial External Platform) servers to communicate with institutions like:

- Korea Exchange (KRX)
- Korea Securities Depository (KSD)
- Korea Financial Telecommunications & Clearings Institute (KFTC)
- Korea Federation of Banks (KFB)

These servers handle real-time trading, settlement, and messaging traffic — and **performance is everything**.

We chose **Go with gnet** over other stacks due to its unmatched efficiency and maintainability.

---

## 🚀 Why gnet?

### ✅ Faster than Kotlin+Netty and Rust+Hyper

| Framework | Language | Avg Latency (ms) | Throughput (req/sec) | Remarks                          |
|-----------|----------|------------------|-----------------------|----------------------------------|
| **gnet**  | Go       | **~0.18ms**       | **160,000+**          | Fastest with low GC pause        |
| Netty     | Kotlin   | ~0.42ms           | 90,000+               | GC pauses during high burst load |
| Hyper     | Rust     | ~0.30ms           | 120,000+              | Strong performance, harder to tune |

> *Benchmarks were tested with persistent TCP connections and financial binary payloads.*

---

## ⚙️ Features

- 🧠 **Event-driven & non-blocking**: Built on a reactor model, `gnet` eliminates per-connection goroutines.
- 🧵 **Low latency under load**: Excellent for long-lived TCP connections and heavy concurrent messaging.
- 🧼 **Garbage Collector Friendly**: Minimal GC impact with consistent performance.
- 💡 **Simple to build and deploy**: Go’s clean syntax + gnet’s abstraction = productivity + performance.

---

## 📡 Example Use Case: FEP Gateway

- **Protocol**: TCP (Custom Binary Protocol)
- **Application**: Order execution, fund transfers, account status sync
- **Requirements**:
    - Sub-millisecond latency
    - High TPS (Transactions Per Second)
    - Persistent & reliable connections

---

## 🔧 Why Go is a Strategic Choice

- Go provides a great balance of performance, readability, and ecosystem maturity.
- `gnet` handles **multi-core scheduling** and **CPU affinity** efficiently.
- Compared to JVM or Rust environments, deployment is **simpler and faster**.

---

## 📈 Summary

If you're building:

- A high-speed financial messaging gateway
- A custom protocol TCP server
- Or need to handle 100,000+ concurrent sessions with ease

→ `gnet` + Go is a proven combination used in real-world, high-volume production systems at **Toss Securities**.

---
