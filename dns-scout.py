# DNS Scout - Version: v5.5

import dns.resolver
from prompt_toolkit import PromptSession, HTML
import whois

MARGIN = '    '

def get_registrar(domain):
    try:
        w = whois.whois(domain)
        if w.registrar:
            return w.registrar
        else:
            return "Unknown"
    except:
        return "Unknown"

def get_info(domain):
    try:
        ns = dns.resolver.resolve(domain, 'NS')
        name_servers = [item.to_text() for item in ns]
    except:
        name_servers = None

    try:
        mx = dns.resolver.resolve(domain, 'MX')
        mx_records = [item.exchange.to_text() for item in mx]
    except:
        mx_records = None

    try:
        txt = dns.resolver.resolve(domain, 'TXT')
        txt_records = [item.strings[0].decode("utf-8") for item in txt]
    except:
        txt_records = []

    try:
        dmarc = dns.resolver.resolve(f'_dmarc.{domain}', 'TXT')
        dmarc_record = dmarc[0].strings[0].decode("utf-8")
        if "v=DMARC1;" not in dmarc_record:
            dmarc_record = None
    except:
        dmarc_record = None

    try:
        ptr = dns.resolver.resolve(domain, 'PTR')
        ptr_records = [item.to_text() for item in ptr]
    except:
        ptr_records = None

    return name_servers, mx_records, txt_records, dmarc_record, ptr_records

def print_dns_records(registrar, name_servers, mx_records, txt_records, dmarc_record, ptr_records):
    print(f"{MARGIN}\033[94m{'-' * 76}\033[0m")
    if registrar == 'Unknown':
        print(f"{MARGIN}\033[95mRegistrar:\033[0m \033[93mUnknown or Classified\033[0m")
    else:
        print(f"{MARGIN}\033[95mRegistrar:\033[0m {registrar}")


    # Name Servers
    print(f"\033[95m{MARGIN}Name Servers:\033[0m")
    if name_servers:
        for server in name_servers:
            print(f"{MARGIN}   {server}")
    else:
        print(f"{MARGIN}   None")

    # MX Records
    print(f"\033[95m{MARGIN}MX Records:\033[0m")
    if mx_records:
        for record in mx_records:
            print(f"{MARGIN}   {record}")
    else:
        print(f"{MARGIN}\033[91m   No MX record, No email!\033[0m")

    # DMARC Records
    print(f"\033[95m{MARGIN}DMARC Records:\033[0m")
    if dmarc_record:
        if "p=quarantine;" in dmarc_record:
            dmarc_record = dmarc_record.replace("p=quarantine;", "\033[93mp=quarantine;\033[0m")
        if "sp=quarantine;" in dmarc_record:
            dmarc_record = dmarc_record.replace("sp=quarantine;", "\033[93msp=quarantine;\033[0m")
        if "p=none;" in dmarc_record:
            dmarc_record = dmarc_record.replace("p=none;", "\033[91mp=none;\033[0m")
        
        for i in range(1, 100):  # Checking pct values from 1 to 99
            if f"pct={i};" in dmarc_record:
                dmarc_record = dmarc_record.replace(f"pct={i};", f"\033[91mpct={i};\033[0m")

        print(f"{MARGIN}   {dmarc_record}")
    else:
        print(f"{MARGIN}\033[91m   None\033[0m")

    # TXT Records
    print(f"\033[95m{MARGIN}TXT Records:\033[0m")
    for record in txt_records:
        if "v=spf1" in record or "v=spf2.0/pra" in record or "~all" in record or "-all" in record:
            if "vspf" in record:
                record = record.replace("vspf", "\033[91mvspf\033[0m")
            if "~all" in record:
                print(f"{MARGIN}\033[32m{record.split('~all')[0]}\033[92m~all\033[0m")
            elif "-all" in record:
                print(f"{MARGIN}\033[32m{record.split('-all')[0]}\033[93m-all\033[0m")
            else:
                print(f"{MARGIN}\033[32m{record}\033[0m")
        else:
            print(f"{MARGIN}\033[32m{record}\033[0m")

    # PTR Records
    print(f"\033[95m{MARGIN}PTR Records:\033[0m")
    if ptr_records:
        for record in ptr_records:
            print(f"{MARGIN}   {record}")
    else:
        print(f"{MARGIN}\033[93m   None\033[0m")

    print(f"{MARGIN}\033[94m{'-' * 76}\033[0m")

session = PromptSession()

while True:
    domain = session.prompt(HTML(f'{MARGIN}<ansibrightblue>Enter domain (or \'exit\' to quit): </ansibrightblue>'))

    if domain.lower() == 'exit':
        break

    registrar = get_registrar(domain) 
    name_servers, mx_records, txt_records, dmarc_record, ptr_records = get_info(domain)
    print_dns_records(registrar, name_servers, mx_records, txt_records, dmarc_record, ptr_records)
