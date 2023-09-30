# DNS-Scout

DNS Scout is a utility tool designed to fetch and display DNS-related records of a domain in a clear, color-coded console output. From standard NS records to DMARC settings, it provides a quick overview for DNS reconnaissance and troubleshooting.

Features:
Retrieves Registrar information.
Fetches Name Servers (NS).
Shows MX Records (Mail Exchange).
Displays TXT Records, useful for checking domain verification, SPF settings, etc.
Retrieves DMARC Records to understand a domain's email authentication settings.
dns-scout.py - A powerful and easy-to-use domain DNS lookup utility for the terminal.

The essential information needed to resolve client email issues: 
-Registrar
-The 2NS servers
-MX records
-SPF and dmarc
Provides the exact data with no extra fluff.
Info is easy to see and copy/paste.

Requirements & Dependencies
System Utilities:
whois: Used to retrieve domain registration data.
nslookup: Utilized for DNS record lookups.
These utilities are commonly available on macOS and Linux. If you're on another system, ensure they're installed and available in your PATH.

Python:
Python 3.x
Python Libraries:
colorama: Required for color-coded console output. Install using pip:
Usage
Open Terminal.
sudo apt install whois
sudo apt install dnsutils
Navigate to the directory containing the dns-scout.py script using the cd command.
Run the script using Python 3: python3 whoisns.py
When prompted, you can enter the domain to fetch the WHOIS and DNS information.
The script will display the Registrar, Name Servers, and the relevant DNS records (MX, ALL TXT (SPF), and DMARC if available) for the provided domain.


System Requirements
Python 3.x
whois command-line utility
nslookup command-line utility
The script was written and tested using MacOS Ventura, Python 3.8.5, but it should work with other Python 3 versions. Please ensure you have the whois and nslookup utilities installed on your system.

On macOS, you can check whether these utilities are installed by opening Terminal and running which whois and which nslookup. If these commands return a path, the utilities are installed. If not, you will need to install them. These utilities come pre-installed on most macOS versions.

