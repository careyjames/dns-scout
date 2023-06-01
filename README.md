# whoisns by Carey James Balboa - Mac Help Nashville, Inc.
a Python3 script for MacOS Terminal that gets right to the mission criticals of a domain!

WhoisNS.py - Python Script for DNS & WHOIS Information

Description:
Retrieves and displays WHOIS data (Registrar and Name Servers) and DNS records (MX and TXT SPF records, including DMARC records if they exist) for a given domain.

The most important things I need to know to help resolve client email issues are: 
-Registrar
-The 2NS servers
-MX records
-SPF and dmarc
This script provides the exact data without all the extra fluff.

Usage
Open Terminal.
Navigate to the directory containing the whoisns.py script using the cd command.
Run the script using Python 3: python3 whoisns.py
When prompted, enter the domain for which you want to fetch the WHOIS and DNS information.
The script will then display the Registrar, Name Servers, and the relevant DNS records (MX, TXT (SPF), and DMARC if available) for the provided domain.

Example:
```
$ python3 whoisns.py
Enter domain: example.com
```
The script will output the domain's WHOIS and DNS records in the following format:
```
python3 whoisns.py                                                                                                            ─╯
Enter domain: google.com
Registrar: MarkMonitor, Inc.
Name Servers: ns3.google.com
Name Servers: ns1.google.com
Name Servers: ns2.google.com
Name Servers: ns4.google.com
MX Records: 10 smtp.google.com.
TXT Records: "onetrust-domain-verification=de01ed21f2fa4d8781cbc3ffb89cf4ef"
TXT Records: "globalsign-smime-dv=CDYX+XFHUw2wml6/Gb8+59BsH31KzUr6c1l2BPvqKX8="
TXT Records: "webexdomainverification.8YX6G=6e6922db-e3e6-4a36-904e-a805c28087fa"
TXT Records: "v=spf1 include:_spf.google.com ~all"
TXT Records: "apple-domain-verification=30afIBcvSuDV2PLX"
TXT Records: "google-site-verification=TV9-DBe4R80X4v0M4U_bd_J9cpOJM0nikft0jAgjmsQ"
TXT Records: "atlassian-domain-verification=5YjTmWmjI92ewqkx2oXmBaD60Td9zWon9r6eakvHX6B77zzkFQto8PQ9QsKnbf4I"
TXT Records: "docusign=1b0a6754-49b1-4db5-8540-d2c12664b289"
TXT Records: "MS=E4A68B9AB2BB9670BCE15412F62916164C0B20BB"
TXT Records: "google-site-verification=wD8N7i1JTNTkezJ49swvWW48f8_9xveREV4oB-0Hf5o"
TXT Records: "facebook-domain-verification=22rm551cu4k0ab0bxsw536tlds4h95"
TXT Records: "docusign=05958488-4752-4ef2-95eb-aa7ba8a3bd0e"
TXT Records: "v=DMARC1; p=reject; rua=mailto:mailauth-reports@google.com"
```
System Requirements
Python 3.x
whois command-line utility
nslookup command-line utility
The script was written and tested using Python 3.8.5, but it should work with other Python 3 versions. Please ensure you have the whois and nslookup utilities installed on your system.

On macOS, you can check whether these utilities are installed by opening Terminal and running which whois and which nslookup. If these commands return a path, the utilities are installed. If not, you will need to install them. These utilities come pre-installed on most macOS versions.

