# DNS-Scout  

DNS Scout is designed to fetch and display DNS-related records of a domain in a clear, color-coded console output.  
From standard NS records to DMARC settings, it provides a quick overview for DNS reconnaissance and troubleshooting.  
  
![Apple DNS records](apple-dns.png)  
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
To install the necessary Python library, run: pip install colorama  

Usage:  
Navigate to the directory where dns-scout.py is located and run: python3 dns-scout.py  
Follow the on-screen prompts to input a domain and retrieve its information.  
