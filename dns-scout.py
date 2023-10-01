# DNS Scout - Version: v5.0

import subprocess
import sys
from prompt_toolkit import prompt
from prompt_toolkit.history import InMemoryHistory
from prompt_toolkit.clipboard.in_memory import InMemoryClipboard
from prompt_toolkit.key_binding import KeyBindings
from prompt_toolkit.keys import Keys
from prompt_toolkit.formatted_text import FormattedText

# Ensure Python dependencies
def ensure_python_dependencies():
    try:
        from prompt_toolkit import prompt
        import colorama
    except ImportError:
        print("Some Python modules are missing. Attempting to install necessary modules...")
        subprocess.run([sys.executable, "-m", "pip", "install", "prompt_toolkit", "colorama"])

# Check Python dependencies
ensure_python_dependencies()

from colorama import Fore, init
init(autoreset=True)  # Automatically reset terminal color to default after print statements

# Define the pasting keybindings
kb = KeyBindings()

@kb.add(Keys.BracketedPaste)
def _(event):
    event.current_buffer.insert_text(event.data, fire_event=False)

# Setup command history
history = InMemoryHistory()

def get_info(start_line, lines):
    registrar = None
    name_servers = []
    i = start_line
    for i in range(start_line, len(lines)):
        line = lines[i]
        if "Registrar:" in line and registrar is None:
            registrar = line.split("Registrar:")[1].strip()
            print(Fore.GREEN + "Registrar:", registrar)
        if "Name Server:" in line:
            name_servers.append(line.split("Name Server:")[1].strip())
        if line.strip() == "" and name_servers:
            break

    if name_servers:
        print(Fore.YELLOW + "Name Servers:", ", ".join(name_servers))
        return i, True
    return i, False

def print_dns_records(mx_records, txt_records, dmarc_record, ptr_records):
    # MX Records
    print(Fore.MAGENTA + "\nMX Records:")
    for line in mx_records:
        if "mail exchanger" in line:
            print("  ", line.split("=")[1].strip())

    # TXT Records
    print(Fore.MAGENTA + "\nTXT Records:")
    for line in txt_records:
        if "text =" in line:
            txt = line.split("text =")[1].strip()
            print("  ", txt)

    # DMARC Records
    print(Fore.MAGENTA + "\nDMARC Records:")
    for line in dmarc_record:
        if "text =" in line:
            txt = line.split("text =")[1].strip()
            print("  ", txt)

    # PTR Records
    print(Fore.MAGENTA + "\nPTR Records:")
    ptr_found = False
    for line in ptr_records:
        if "name =" in line:
            ptr = line.split("name =")[1].strip()
            print("  ", ptr)
            ptr_found = True
    if not ptr_found:
        print("  None found")

# Main loop
while True:
    print('-' * 80)  # Line separator
    domain = prompt(
        FormattedText([('cyan', 'Enter domain (or \'exit\' to quit): ')]),
        history=history,
        key_bindings=kb,
        clipboard=InMemoryClipboard()
    ).strip()

    if domain.lower() in ['exit', 'q', 'quit']:
        break

    result = subprocess.run(["whois", domain], capture_output=True, text=True)
    lines = result.stdout.splitlines()
    start_line = 0
    registrar_printed = False

    while not registrar_printed:
        start_line, registrar_printed = get_info(start_line, lines)
        start_line += 1
        if start_line >= len(lines):
            break

    try:
        mx_records = subprocess.check_output(["nslookup", "-type=mx", domain]).decode("utf-8").splitlines()
    except subprocess.CalledProcessError:
        mx_records = []

    try:
        txt_records = subprocess.check_output(["nslookup", "-type=txt", domain]).decode("utf-8").splitlines()
    except subprocess.CalledProcessError:
        txt_records = []

    try:
        dmarc_record = subprocess.check_output(["nslookup", "-type=txt", "_dmarc."+domain]).decode("utf-8").splitlines()
    except subprocess.CalledProcessError:
        dmarc_record = []

    try:
        ptr_records = subprocess.check_output(["nslookup", "-q=ptr", domain]).decode("utf-8").splitlines()
    except subprocess.CalledProcessError:
        ptr_records = []

    print_dns_records(mx_records, txt_records, dmarc_record, ptr_records)
