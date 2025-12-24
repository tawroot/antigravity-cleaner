#!/bin/bash

# Session Transfer Tool for Linux/Mac
# Export or Import sessions between systems

ACTION=${1:-export}
SESSION_DIR="$HOME/.antigravity-cleaner/sessions"
EXPORT_DIR="$HOME/Desktop/antigravity-sessions-backup"

# Colors
CYAN='\033[0;36m'
YELLOW='\033[1;33m'
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "\n${CYAN}======================================================${NC}"
echo -e "       SESSION TRANSFER TOOL (Linux/Mac)"
echo -e "${CYAN}======================================================${NC}"

if [ "$ACTION" == "export" ]; then
    echo -e "\n[EXPORT MODE] Backing up sessions to Desktop..."
    
    if [ ! -d "$SESSION_DIR" ]; then
        echo -e "${RED}[ERROR] No sessions found at: $SESSION_DIR${NC}"
        exit 1
    fi
    
    # Create export folder
    rm -rf "$EXPORT_DIR"
    mkdir -p "$EXPORT_DIR"
    
    # Copy all session files and key
    cp -r "$SESSION_DIR"/* "$EXPORT_DIR/"
    
    echo -e "${GREEN}[OK] Sessions exported to: $EXPORT_DIR${NC}"
    echo -e "\nFiles exported:"
    ls "$EXPORT_DIR" | while read line; do echo -e "  - $line"; done
    
    echo -e "\n${YELLOW}[!] IMPORTANT: Copy the entire folder to the target system!${NC}"
    echo -e "    Then run: ./transfer-session.sh import"
    
else
    echo -e "\n[IMPORT MODE] Restoring sessions from Desktop..."
    
    if [ ! -d "$EXPORT_DIR" ]; then
        echo -e "${RED}[ERROR] No backup folder found at: $EXPORT_DIR${NC}"
        echo -e "${YELLOW}Please copy the 'antigravity-sessions-backup' folder to your Desktop first.${NC}"
        exit 1
    fi
    
    # Create session directory if not exists
    mkdir -p "$SESSION_DIR"
    
    # Copy files
    cp -r "$EXPORT_DIR"/* "$SESSION_DIR/"
    
    echo -e "${GREEN}[OK] Sessions imported successfully!${NC}"
    echo -e "\nImported files:"
    ls "$SESSION_DIR" | while read line; do echo -e "  - $line"; done
    
    echo -e "\nYou can now use Session Manager to restore these sessions."
fi

echo ""
