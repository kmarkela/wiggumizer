
        __       __  __                                              __
        |  \  _  |  \|  \                                            |  \
        | $$ / \ | $$ \$$  ______    ______   __    __  ______ ____   \$$ ________   ______     ______      
        | $$/  $\| $$|  \ /      \  /      \ |  \  |  \|      \    \ |  \|        \ /      \  /      \       
        | $$  $$$\ $$| $$|  $$$$$$\|  $$$$$$\| $$  | $$| $$$$$$\$$$$\| $$ \$$$$$$$$|  $$$$$$\|  $$$$$$\  
        | $$ $$\$$\$$| $$| $$  | $$| $$  | $$| $$  | $$| $$ | $$ | $$| $$  /    $$ | $$    $$| $$   \$$   
        | $$$$  \$$$$| $$| $$__| $$| $$__| $$| $$__/ $$| $$ | $$ | $$| $$ /  $$$$_ | $$$$$$$$| $$
        | $$$    \$$$| $$ \$$    $$ \$$    $$ \$$    $$| $$ | $$ | $$| $$|  $$    \ \$$     \| $$
         \$$      \$$ \$$ _\$$$$$$$ _\$$$$$$$  \$$$$$$  \$$  \$$  \$$ \$$ \$$$$$$$$  \$$$$$$$ \$$
                         |  \__| $$|  \__| $$   Web Traffic 4nalizer
                          \$$    $$ \$$    $$
                           \$$$$$$   \$$$$$$


# Wiggumizer: Web Traffic analyzer

Wiggumizer is a tool designed for security researchers, penetration testers, and ethical hackers.  It is equped with list of checks to identify potential security vulnerabilities in a Web Application. It is not a vulnarability scanner and it does not execute any active scanning/enumeration/testing. It is rather anaylises history exported from Web Proxy (only Burp Suite is currently supported) to identify potential security vulnerabilities, enabling users to focus their investigative efforts efficiently.

## Disclaimer: Ethical Use Only

Wiggumize is intended to be used exclusively for ethical and legitimate purposes, such as security assessments, penetration testing, and vulnerability research. Any use of Wiggumize for malicious, unauthorized, or unethical activities is strictly prohibited.

By using Wiggumize, you acknowledge and agree to adhere to all applicable laws and regulations governing your activities. You are solely responsible for obtaining proper authorization before conducting security assessments on any systems, networks, or applications. The developers and maintainers of Wiggumize disclaim any liability for any misuse, damage, or legal consequences resulting from the misuse of this tool.


## Installation

### Building from src:

```bash
git clone https://github.com/kmarkela/Wiggumizeng.git
cd Wiggumizeng
go build
```

## Usage

```shell
Web Traffic 4nalizer

Usage:
  Wiggumize [flags]
  Wiggumize [command]

Available Commands:
  help        Help about any command
  scan        scan analysis web history and run list of checks on Req\Res body and headers
  search      powerfull search in browse history

Flags:
  -h, --help          help for Wiggumize
  -V, --version       print version
  -w, --workers int   amount of workers (default 5)

Use "Wiggumize [command] --help" for more information about a command.
```

## Scan

In scan mode, Wiggumizer analysis web history and run list of checks on Req\Res body and headers.  The result of checks saved in `md` format (default file: `report.md`).

### List of checks 

- **LFI Checker**: This module is searching for filenames in request parameters. Could be an indication of possible LFI
- **Redirect Checker**: This module is searching for Redirects
- **Secret Checker**: This module lokking for sensitive information, such as API keys
- **SSRF Checker**: This module is searching for URL in request parameters.
- **Subdomain Checker**: This module is searching for 404 messages form hosting platformas
- **XML Checker**:  This module is searching for XML in request parameters


## Search

It allows for powerfull search in browse history. 

### Avaliable search fields: 
- `Method`
- `ReqHeader`
- `ReqContentType`
- `ReqBody`
- `ResHeader`
- `ResContentType`
- `ResBody`

### Avaliable search operators: 
- `&` - AND
- `!` - NOT


### Avaliable config flags: 
- `-i`  - Case insensitive search
- `-br` - brief output (only list uniq endpoints)
- `-h`  - only headers in output
- `-f`  - full output

### Search Example:

Search for requests that satisfy the following criteria:
- Request method is POST
- Request body contains the term "admin"
- Response content type is not HTML
- Response body contains the term "success"

> Make search case insensitive and output only list uniq endpoints.  

```shell
ReqMethod POST & ReqBody *admin* & ! ResContentType HTML & ResBody success -br -i
```