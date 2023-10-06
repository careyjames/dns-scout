# DNS-Scout  

DNS Scout for Linux/MacOS pulls and displays DNS records in a color-coded console output that is **easy to see** and **copy/paste**.   
Registrar, NS, MX, SPF, DMARC, and PTR for easy DNS reconnaissance and troubleshooting.  
 
<img src="example-domains.png" alt="Apple DNS records" width="800">  

## Features:   

**Curated Output for Clarity:**  
 DNS Scout stands out by filtering out non-essential information, presenting users with a cleaner,  
 more focused view of the DNS data, and optimizing for clarity and relevance.  
 
**Enhanced Interactive CLI Interface:**  
 DNS Scout leverages ```prompt_toolkit``` to offer an advanced command-line interface  
 that's **easy to see and copy/paste**     
 **Session-based Memory Cycling**  
DNS Scout's interactive interface has a memory cycle feature, controlled by the up and down arrow keys.  
It helps navigate recent domain lookups for the session quickly.  
This feature is useful when conducting multiple lookups, and you need to refer to a previous entry.   
The memory is session-based, which means it clears once you exit DNS Scout, maintaining your privacy and a clean slate for your next session.  
 
**Streamlined WHOIS Lookup:**
 DNS Scout efficiently parses domain registration data, presenting the user with concise registrar details and name servers, eliminating the clutter typically seen in  raw WHOIS outputs.  
 
**Clear TXT Record Display:**   
 DNS Scout lists TXT records in an easily digestible format,  
 making tasks like SPF data or domain verification review more straightforward.  
 
**Registrar**   
**NS Name Servers**  
**MX Records**  
**Displays TXT Records**, useful for checking domain verification, SPF settings, etc.  
**DMARC Records**
**PTR**  
**Exact DNS data, no scrolling**      

### Installation Guide for DNS Scout  
Prerequisites: Python 3.x  

1. Clone the Repository:    
If you have Git installed, run:  ```git clone https://github.com/careyjames/dns-scout.git```  
If you don't have Git, download the repository as a zip file and extract it in your home folder.  

2. Navigate to the Directory:  
```cd ~/DNS-Scout```  

3. Install Required Python Libraries:  
 Run the following commands to install the necessary Python libraries:  
 ```pip install dnspython whois prompt-toolkit colorama```  

  -```colorama``` to print colored output to the terminal.  
  -```prompt_toolkit``` to enhance the input prompt,  
  allowing for functionalities like copy/paste and keeping a history of entered domains.  

#### **Usage:**  
Navigate to the directory where ```dns-scout.py``` is located.  
```cd ~/DNS-Scout/```
Then run: ```python3 dns-scout.py```  
Follow the on-screen prompts to input a domain and retrieve its information.   
OR  
Add this line: ```alias dns='python3 ~/DNS-Scout/dns-scout.py``` in your ```.zshrc``` file for a cool shortcut.  
Then, you can simply type "dns" in terminal to launch.  

 
