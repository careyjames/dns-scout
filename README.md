# DNS-Scout  

DNS Scout for Linux/MacOS pulls and displays DNS records in a color-coded console output that is **easy to see** and **copy/paste**.   
Registrar, NS, MX, SPF, DMARC, and PTR for easy DNS reconnaissance and troubleshooting.  
 
<img src="apple-dns.png" alt="Apple DNS records" width="800">  

## Features:  
**Curated Output for Clarity:**  
 DNS Scout stands out by filtering out non-essential information, presenting users with a cleaner,  
 more focused view of the DNS data, and optimizing for clarity and relevance.  
**Enhanced Interactive CLI Interface:**  
 Unlike traditional utilities, DNS Scout leverages prompt_toolkit to offer an advanced command-line interface  
 that's **easy to see and copy/paste** and uses arrow keys up and down to cycle through the history of queried domains.  
**Streamlined WHOIS Lookup:**
 DNS Scout efficiently parses domain registration data, presenting the user with concise registrar details and name servers, eliminating the clutter typically seen in  raw WHOIS outputs.  
**Clear TXT Record Display:**   
 Unlike basic utilities, DNS Scout lists TXT records in an easily digestible format,  
 making tasks like SPF data or domain verification review more straightforward.  
**Retrieves Registrar information**   
**NS Name Servers**  
**MX Records (Mail Exchange)**  
**Displays TXT Records, useful for checking domain verification, SPF settings, etc.**  
**Retrieves DMARC Records**    
**Provides the exact data with no extra fluff**      

### Installation Guide for DNS Scout  
Prerequisites:
Python: Python 3.x  

#### 1. Clone the Repository:    
If you have Git installed, run:  ```git clone https://github.com/careyjames/dns-scout.git```  
If you don't have Git, download the repository as a zip file and extract it in the home folder.  

##### 2. Navigate to the Directory:  
```cd ~/DNS-Scout```  

###### 3. Install Required Python Libraries:
Run the following commands to install the necessary Python libraries:  
```pip install prompt_toolkit colorama```  

 -```colorama``` to print colored output to the terminal.  
 -```prompt_toolkit``` to enhance the input prompt,  
 allowing for functionalities like copy/paste and keeping a history of entered domains.  

***Usage:***  
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
