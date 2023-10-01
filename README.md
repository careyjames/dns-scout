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
Prerequisites:
Python: Python 3.x  

#### 1. **Clone the Repository:**    
If you have Git installed, run:  ```git clone https://github.com/careyjames/dns-scout.git```  
If you don't have Git, download the repository as a zip file and extract it in the home folder.  

#### Navigate to the Directory:  
```cd ~/DNS-Scout```  

#####  Install Required Python Libraries:
Run the following commands to install the necessary Python libraries:  
```pip install prompt_toolkit colorama```  

 -```colorama``` to print colored output to the terminal.  
 -```prompt_toolkit``` to enhance the input prompt,  
 allowing for functionalities like copy/paste and keeping a history of entered domains.  

###### Usage:  
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
