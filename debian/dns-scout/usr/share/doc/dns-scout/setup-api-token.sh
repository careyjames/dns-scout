#!/bin/bash

# Detect the shell and choose the appropriate configuration file
shell_config_file=""
if [ "$SHELL" == "/bin/bash" ] || [ "$SHELL" == "/usr/bin/bash" ]; then
  shell_config_file="$HOME/.bashrc"
elif [ "$SHELL" == "/bin/zsh" ] || [ "$SHELL" == "/usr/bin/zsh" ]; then
  shell_config_file="$HOME/.zshrc"
else
  echo "Unsupported shell. Exiting."
  exit 1
fi

# Check if the API token is already set in the shell configuration file
if grep -q "export IPINFO_API_TOKEN=" $shell_config_file; then
  echo "API token already set in $shell_config_file. Exiting."
  exit 0
fi

# Prompt the user for the API token
read -p "Please enter your IPInfo API token: " api_token

# Check if the API token was entered
if [ -z "$api_token" ]; then
  echo "API token is required. Exiting."
  exit 1
fi

# Add the API token to the user's shell configuration file
echo "export IPINFO_API_TOKEN=$api_token" >> $shell_config_file

# Reload the shell configuration file to make the new environment variable available immediately
source $shell_config_file

# Print a success message
echo "API token successfully saved."
