#!/bin/bash

# Function to show help message
show_help() {
    echo "Name Fixer

Features:
  Batch process filenames by extracting episode numbers and renaming in specified format

Usage:
  nf <series-name> [custom-rules...]
  nf --help or nf -h    Show help information

Parameters:
  <series-name>      Required, used for matching and renaming files
  [custom-rules...]  Optional, multiple custom rules for processing. If a rule contains '=>',
                     the string before it will be replaced by the string after it.
                     If not, the string will be removed.
                     Note: Each rule must be quoted.

Processing Rules:
  1. If no custom rules are defined, extract the last occurring number as episode number
     and generate new filename format: 'series-name-episode-number.extension'
  2. If custom rules are defined, process according to the rules.
  3. If a rule contains '=>', the string before it will be replaced by the string after it.
     If not, the string will be removed.

Examples:
  # Basic usage: Process files containing 'Attack on Titan'
  nf 'Attack on Titan'

  # Advanced usage: Specify strings to remove
  nf 'Attack on Titan' '[SubGroup]' '1080P'

  # Using string replacement feature
  nf 'Attack on Titan' 'Chapter=>Episode ' 'Season=>Season '"
}

# Function to clean filename
clean_filename() {
    local filename="$1"
    local series_name="$2"
    shift 2
    local patterns=($@)
    
    # Remove extension first
    local ext="${filename##*.}"
    local name="${filename%.*}"
    
    # Apply custom patterns if provided
    if [ ${#patterns[@]} -gt 0 ]; then
        for pattern in "${patterns[@]}"; do
            if [[ $pattern == *"=>"* ]]; then
                # Handle replacement pattern
                local old_str="${pattern%%=>*}"
                local new_str="${pattern##*=>}"
                name="${name//$old_str/$new_str}"
            else
                # Handle removal pattern
                name="${name//$pattern/}"
            fi
        done
    else
        # Extract the last number sequence as episode number
        if [[ $name =~ ([0-9]+) ]]; then
            local last_num="${BASH_REMATCH[1]}"
            name="$series_name-$last_num"
        fi
    fi
    
    # Clean up any remaining special characters and spaces
    name="$(echo "$name" | tr -s ' ' | sed 's/^ *//;s/ *$//')"
    echo "$name.$ext"
}

# Check if help is requested
if [[ "$1" == "--help" || "$1" == "-h" || "$1" == "--h" || "$1" == "-help" ]]; then
    show_help
    exit 0
fi

# Check if series name is provided
if [ $# -lt 1 ]; then
    echo "Please provide at least one argument: series name. Use --help for help information."
    exit 1
fi

# Get series name and remove patterns
series_name="$1"
shift
remove_patterns=($@)

# Check series name length
if [ -z "$series_name" ]; then
    echo "Series name length cannot be less than one character!"
    exit 1
fi

# Get working directory
working_dir="$(pwd)"

# Check if series name should be replaced
new_series_name="$series_name"
for pattern in "${remove_patterns[@]}"; do
    if [[ $pattern == *"=>"* ]]; then
        old_str="${pattern%%=>*}"
        new_str="${pattern##*=>}"
        if [ "$old_str" == "$series_name" ]; then
            new_series_name="$new_str"
            break
        fi
    fi
done

# Create target directory
target_dir="$working_dir/$new_series_name"
if [ -e "$target_dir" ]; then
    if [ ! -d "$target_dir" ]; then
        echo "Error: $target_dir already exists and is not a directory."
        exit 1
    fi
else
    mkdir -p "$target_dir"
    if [ $? -ne 0 ]; then
        echo "Failed to create target directory: $target_dir"
        exit 1
    fi
fi

# Find matching files
declare -A source_dirs
matches=()
while IFS= read -r -d '' file; do
    if [[ "${file,,}" == *"${series_name,,}"* ]]; then
        matches+=("$file")
        source_dirs["$(dirname "$file")"]=1
    fi
done < <(find "$working_dir" -type f -print0)

if [ ${#matches[@]} -eq 0 ]; then
    echo "No matching files found."
    exit 0
fi

# Move and clean matching files
for file in "${matches[@]}"; do
    filename="$(basename "$file")"
    clean_name="$(clean_filename "$filename" "$new_series_name" "${remove_patterns[@]}")"
    target_path="$target_dir/$clean_name"
    
    mv "$file" "$target_path"
    if [ $? -eq 0 ]; then
        echo "Moved and cleaned: $file -> $target_path"
    else
        echo "Failed to move file $file"
    fi
done

# Remove empty source directories
for dir in "${!source_dirs[@]}"; do
    while [ "$dir" != "$working_dir" ]; do
        if [ -z "$(ls -A "$dir")" ]; then
            rmdir "$dir"
            dir="$(dirname "$dir")"
        else
            break
        fi
    done
done

echo "Operation completed."