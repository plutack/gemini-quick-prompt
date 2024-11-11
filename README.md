#  Gemini-quick-prompt
A CLI tool designed to quickly provide insights into programming functions, frameworks, and more. Ideal for developers seeking instant guidance without searching through documentation. Manpage styled output to the terminal is favoured for easy information consumption

## How to install
1. ensure go is properly installed. See [Go official docs](https://go.dev/doc/install) for proper installation of operating system of choice
2. run the command below.
    ```sh
    go install github.com/plutack/gemini-quick-prompt@latest
    ```
    <!-- OR -->
<!-- 1. Alternatively, download the pre-built binary from the releases page or run the command below, place it in a directory of your choice, and add that directory to your PATH. --> 
<!---
    ```sh
    curl -fSL -o gemini-quick-prompt https://github.com/plutack/gemini-quick-prompt/releases/latest/download/gemini-quick-prompt
    ```
-->
##  How to use
1. get a [gemini API key](https://ai.google.dev/gemini-api/docs/api-key) here
2. set gemini_api_key as an enviroment variable. This will differ depending on the operating system you use
3. command are expected in the format;
    ```sh
    gemini-quick-prompt "js array includes"
    ```
4. user may also alias gemini-quick-prompt to a shorter alias if they choose.

https://github.com/user-attachments/assets/a5b0f81c-112a-4857-bdf6-fd688a48de93


