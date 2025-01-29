#!/bin/bash

# Function to rename directories and files recursively
rename_files() {
  # Traverse directories and files recursively
  find "$1" -depth | while read -r file; do
    # Get the new name by replacing spaces and underscores with dashes
    new_name=$(echo "$file" | tr ' _' '--')

    # Rename the file or directory if the name is different
    if [ "$file" != "$new_name" ]; then
      echo "Renaming '$file' to '$new_name'"
      mv "$file" "$new_name"
    fi
  done
}

# Main script execution
BASE_DIR="dsa-golang"  # Set the base directory
echo "Starting renaming process for files and directories..."

rename_files "$BASE_DIR"

echo "Renaming process completed!"
