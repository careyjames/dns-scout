# whoisns Mac Help Nashville, Inc.
a Python3 script for MacOS Terminal that gets right to the mission criticals of a domain!

WhoisNS.py - Python Script for DNS & WHOIS Information

Description:
This Python script retrieves and displays WHOIS data (Registrar and Name Servers) and DNS records (MX and TXT records, including DMARC records if they exist) for a given domain.
The most important things I need to know to help resolve client issues are: 
*Registrar
*The 2NS servers
*MX records
*SPF and dmarc

System Requirements
Python 3.x
whois command-line utility
nslookup command-line utility
The script was written and tested using Python 3.8.5, but it should work with other Python 3 versions. Please ensure you have the whois and nslookup utilities installed on your system.

On macOS, you can check whether these utilities are installed by opening Terminal and running which whois and which nslookup. If these commands return a path, the utilities are installed. If not, you will need to install them. These utilities come pre-installed on most macOS versions.

Usage
Open Terminal.
Navigate to the directory containing the whoisns.py script using the cd command.
Run the script using Python 3: python3 whoisns.py
When prompted, enter the domain for which you want to fetch the WHOIS and DNS information.
The script will then display the Registrar, Name Servers, and the relevant DNS records (MX, TXT (SPF), and DMARC if available) for the provided domain.

Example:
$ python3 whoisns.py
Enter domain: example.com

The script will output the domain's WHOIS and DNS records in the following format:

Registrar: GoDaddy.com, LLC
Name Servers: ns1.example.com
Name Servers: ns2.example.com
MX Records: 10 mail.example.com.
TXT Records: "v=spf1 include:_spf.google.com ~all"
DMARC Record: "v=DMARC1; p=none"
