# OS_HW_3
Part 1
“Don't communicate by sharing memory; share memory by communicating.” This phrase refers to using channels to communicate data between goroutines instead of ussing memory. This both frees up memory for usage in each goroutine, and simplifies memory syncronization bewteen goroutines by making sure memory can be dedication to that goroutine. A disadvantage is this tends to slow down performance, and the need to syncronize between goroutine to ensure any one does not attempt to run without the proper data from another goroutine.

Sources
https://paulwizviz.github.io/go/2024/06/15/go-channel.html
https://go.dev/doc/codewalk/sharemem/
https://blog.carlana.net/post/share-memory-by-communicating/
These sources are all credible because they all come from websites that are dedicated to talking about coding, including the actual Go website that talks about this topic.

Part 2
Instructions
1. Place all files in IDE
2. Navigate to folder where files are saved in terminal
3. Type: go run main.go list.go benchmark.go

AI usage- chat-gpt was used for learning the exact way certain commands in c were written in go, and also for ideas on how to implement a benchmark tool.
