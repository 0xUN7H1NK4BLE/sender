# sender
Sender is a command-line tool written in Go for uploading files and sending messages via Discord webhooks.

# Installation
To install Sender, you can use the go install command:

```bash
go install github.com/0xUN7H1NK4BLE/sender@latest
```
# Usage

Sender requires a Discord webhook URL to send messages and/or upload files. Here are the available options:

-w: Discord webhook URL (required)
-f: Path to the file to upload (optional)
-m: Message to send with the file (optional)

# Examples
Upload a file and send a message:
```bash
sender -w <webhook_URL> -f <file_path> -m "Message text"
```

Send a message without uploading a file:
```bash
sender -w <webhook_URL> -m "Just a message"
```

Send a file only without message
```bash
sender -w <webhook_URL> -f <file_path>
```

# Note
* If only a message is provided (-m flag), Sender sends only the message to the specified Discord webhook.
* If only a file is provided (-f flag), Sender uploads the file to the specified Discord webhook.
* If both a file and a message are provided, Sender uploads the file and sends the message together.