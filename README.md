# DNS-Scout  

DNS Scout for Linux/MacOS pulls and displays DNS records in a color-coded console output that is **easy to see** and **copy/paste**.   
Registrar, NS, MX, SPF, DMARC and PTR for easy DNS reconnaissance and troubleshooting.  
 
<img src="apple-dns.png" alt="Apple DNS records" width="800">  

## Features:  
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

### Installation Guide for DNS Scout  
Clone the Repository:  
If you have Git installed, run:  ```git clone https://github.com/YOUR_USERNAME/dns-scout.git```  
Prerequisites:
Python: Python 3.x  
Python Libraries: colorama : Required for color-coded console output.  
To install the necessary Python library, run: 
```pip install colorama``` 

Usage:  
Navigate to the directory where ```dns-scout.py``` is located.  
```cd ~/DNS-Scout/```
Then run: ```python3 dns-scout.py```  
Follow the on-screen prompts to input a domain and retrieve its information.   
OR  
Add this line: ```alias dns='python3 ~/DNS-Scout/dns-scout.py``` in your ```.zshrc``` file for a cool shortcut.  
Then, you can simply type "dns" in terminal to launch.  
<img src="dns-scout.gif" alt="Just type dns" width="800">   

<img src="google-dns.png" alt="Google DNS records" width="800">  

<img src="cia-cisa-dns.png" alt="CIA & CISA DNS records" width="800">  
