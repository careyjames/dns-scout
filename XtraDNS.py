import subprocess

# Define a function to get the Registrar and Name Server information from WHOIS data
def get_info(start_line, lines):
    for i in range(start_line, len(lines)):
        # If the line starts with 'Name Server:', print the Name Servers
        if lines[i].startswith('Name Server:'):
            print('Name Servers:', lines[i].split(': ')[1])
        # If the line starts with 'Registrar:', print the Registrar
        elif lines[i].startswith('Registrar:'):
            print('Registrar:', lines[i].split(': ')[1])
    # Return the current line number
    return i

# Define a function to print DNS record information
def print_dns_records(records):
    # Loop through each record
    for record in records:
        # If the record is a mail exchanger (MX) record, print it
        if 'mail exchanger' in record:
            print('MX Records:', record.split('= ')[1])
        # If the record is a text (TXT) record, print it
        elif 'text =' in record:
            print('TXT Records:', record.split('= ')[1])
        # If the record is a DMARC record, print it
        elif '_dmarc' in record:
            print('DMARC Record:', record.split('= ')[1])

# Ask the user to input a domain
domain = input("Enter domain: ")

# Run the WHOIS command on the domain and capture the output
result = subprocess.run(["whois", domain], capture_output=True, text=True)

# Split the output into lines
lines = result.stdout.splitlines()
start_line = 0

# Keep getting info until we've gone through all the lines
while True:
    start_line = get_info(start_line, lines) + 1
    if start_line >= len(lines):
        break

# Run the NSLOOKUP command on the domain for MX records and capture the output, then split it into lines
mx_records = subprocess.check_output(["nslookup", "-type=mx", domain]).decode("utf-8").splitlines()

# Run the NSLOOKUP command on the domain for TXT records and capture the output, then split it into lines
txt_records = subprocess.check_output(["nslookup", "-type=txt", domain]).decode("utf-8").splitlines()

# Try to run the NSLOOKUP command on the domain for DMARC records and capture the output, then split it into lines
try:
    dmarc_record = subprocess.check_output(["nslookup", "-type=txt", "_dmarc."+domain]).decode("utf-8").splitlines()
except subprocess.CalledProcessError:
    # If the command fails (which might happen if there is no DMARC record), make dmarc_record an empty list
    dmarc_record = []

# Print the DNS records
print_dns_records(mx_records)
print_dns_records(txt_records)
print_dns_records(dmarc_record)
