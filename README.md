# Thread Worker - James McCabe - 16/05/2022

Open source multi-thread worker for GoLang, supports a maxium workers per thread.
This project relys on github.com/sheerun/queue to run.

go get github.com/sheerun/queue

Currently seutp to read from an input file and queue each line to have a task ran on it.
You can modify this to read from Stdin and to run any task you would like.

You can modify the task at the '// Run Task' comment.
