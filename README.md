# GPT cli

GPT cli is a simple command-line interface (CLI) tool that interacts with OpenAI's GPT models (like GPT-3) using their API. This CLI tool allows you to send a prompt to the API and receive generated text as output in your terminal.

## Prerequisites

1. Go programming language installed on your system: https://golang.org/dl/
2. An OpenAI API key: https://beta.openai.com/signup

## Setup

1. Clone this repository to your local machine:

```
git clone https://github.com/greatwebguy/gpt-cli.git
```

2. Navigate to the `gpt-cli` directory:

```
cd gpt-cli
```

3. Open the `askgpt.go` file and replace `your_api_key_here` with your actual OpenAI API key.

4. Compile the program:

```
go build askgpt.go
```

This will generate an executable binary named `askgpt` (or `askgpt.exe` on Windows).

## Usage

To use the CLI tool, run the following command:

```
./askgpt "<your_prompt_here>"
```

Replace `<your_prompt_here>` with the text prompt you'd like to send to the GPT API. The tool will send the prompt to the API and return the generated text in the terminal.

### Example

```
./askgpt "create a curl command to download a file from https://downloads.com/img.png"
```

## Customization

This CLI tool is a basic implementation and can be further improved. Some potential enhancements include:

1. Add error handling and custom error messages.
2. Implement a proper command-line parsing library.
3. Allow users to customize API parameters (e.g., temperature, model, and token limit) via command-line flags.
4. Add support for additional GPT models as they become available.
