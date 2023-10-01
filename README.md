# DNS-Scout  

DNS Scout pulls and displays DNS records in a color-coded console output that is **easy to see** and **copy/paste**.   
NS, DMARC, Registrar, MX, and SPF settings for super easy DNS reconnaissance and troubleshooting.
  
<img src="apple-dns.png" alt="Apple DNS records" width="800">  

Features:  
Retrieves Registrar information.  
Fetches Name Servers (NS).  
Shows MX Records (Mail Exchange).  
Displays TXT Records, useful for checking domain verification, SPF settings, etc.  
Retrieves DMARC Records to understand a domain's email authentication settings.  

The essential information needed to resolve client email issues:  
-Registrar  
-The 2NS servers  
-MX records  
-SPF and dmarc  
Provides the exact data with no extra fluff.  
Info is easy to see and copy/paste.  

Python:  
Python 3.x  
Python Libraries:  
colorama: Required for color-coded console output.  
To install the necessary Python library, run: ```pip install colorama  

Usage:  
Navigate to the directory where dns-scout.py is located and run: python3 dns-scout.py  
Follow the on-screen prompts to input a domain and retrieve its information.  
