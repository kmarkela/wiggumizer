# Wiggumize: Burp Suite History Analysis Tool


Wiggumize is a powerful and versatile tool developed in Golang specifically designed for thorough analysis of Burp Suite history files in XML format. Tailored to assist security researchers, penetration testers, and ethical hackers, Wiggumize is equipped with an array of  features to identify potential security vulnerabilities and provide actionable insights, enabling users to focus their investigative efforts efficiently. With its comprehensive suite of checks and flexible search capabilities, Wiggumize empowers security professionals to enhance their web application assessments and contribute to a more secure digital landscape.

## Disclaimer: Ethical Use Only

Wiggumize is intended to be used exclusively for ethical and legitimate purposes, such as security assessments, penetration testing, and vulnerability research. The tool is designed to assist security professionals in identifying potential vulnerabilities within web applications. Any use of Wiggumize for malicious, unauthorized, or unethical activities is strictly prohibited.

By using Wiggumize, you acknowledge and agree to adhere to all applicable laws and regulations governing your activities. You are solely responsible for obtaining proper authorization before conducting security assessments on any systems, networks, or applications. The developers and maintainers of Wiggumize disclaim any liability for any misuse, damage, or legal consequences resulting from the misuse of this tool.

Always prioritize the principles of responsible disclosure and collaboration with relevant stakeholders when identifying and reporting security vulnerabilities. Use Wiggumize as a tool for positive contributions to cybersecurity and the enhancement of web application security.

## Features and Modules

Wiggumize offers an impressive set of modules to address various aspects of web application security:

- **SSRF Detection**: Unveil potential Server-Side Request Forgery (SSRF) vulnerabilities by identifying URLs within request parameters.

- **404 Detection**: Hunt for 404 error messages in responses, a sign of misconfigurations or vulnerabilities in web application hosting.

- **XML Analysis**: Scrutinize request parameters for XML data to uncover XML-related vulnerabilities.

- **Redirect Detection**: Identify redirects that may indicate open redirection vulnerabilities.

- **Secrets Detection**: Focus on uncovering sensitive information, such as API keys, within the Burp Suite history.

- **LFI Indication**: Analyze filenames within request parameters to pinpoint possible Local File Inclusion (LFI) vulnerabilities.

- **Parameter Parsing**: Wiggumize intelligently parses GET and POST (JSON) parameters for a deeper layer of analysis, enhancing the security assessment process.

## Installation

To harness the power of Wiggumize, follow these installation steps:

```bash
git clone https://github.com/kmarkela/Wiggumize.git
cd Wiggumize
go build
```

## Usage

Wiggumize offers a user-friendly command-line interface with the following parameters:

```shell
-f   Path to XML file with Burp history
-o   Path to output file (default: report.md)
-a   Action. 'scan' for history analysis (default), 'search' for pattern search
```

### Scan Mode (Default)

In scan mode, Wiggumize performs an in-depth analysis of the specified Burp history XML file, executing various security checks. The results are meticulously compiled into a structured Markdown report (default file: `report.md`).

### Search Mode

Search mode allows researchers to execute targeted pattern searches within the Burp history dataset. The robust search capabilities accommodate the usage of regular expressions across diverse fields such as request methods, headers, content types, request and response bodies, and more.

#### Search Example:

Search for requests that satisfy the following criteria:
- Request method is POST
- Request body contains the term "admin"
- Response content type is not HTML
- Response body contains the term "success"

```shell
ReqMethod POST & ReqBody *admin* & ! ResContentType HTML & ResBody success
```

## Contributing

Wiggumize thrives on collaboration and community contributions. If you possess ideas for new checks, enhancements, or bug fixes, we encourage you to initiate discussions through issue creation or by submitting pull requests.

---

Wiggumize has been meticulously crafted to empower security researchers and professionals to unearth potential vulnerabilities and elevate the quality of web application assessments utilizing Burp Suite history. With its modular architecture and extensibility, Wiggumize stands as an indispensable tool in the toolkit of every security enthusiast. Explore, contribute, and together let's fortify the web's security posture with Wiggumize!