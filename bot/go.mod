module github.com/mateuszpoland/chatbot/bot

go 1.18

require github.com/gorilla/mux v1.8.0

require github.com/mateuszpoland/chatbot/models v0.0.0-20220814130858-93e0a1d19e43 // indirect

replace github.com/mateuszpoland/chatbot/models => ../models
