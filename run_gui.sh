#!/bin/bash
# ========================================
# Antigravity Cleaner - GUI Launcher
# ========================================

echo ""
echo "========================================"
echo "  Antigravity Cleaner v2.1 - GUI"
echo "========================================"
echo ""

# Check if Python is installed
if ! command -v python3 &> /dev/null; then
    echo "Error: Python 3 is not installed!"
    echo "Please install Python 3.8 or higher."
    exit 1
fi

# Check Python version
PYTHON_VERSION=$(python3 --version 2>&1 | awk '{print $2}')
echo "Python version: $PYTHON_VERSION"

echo "Starting GUI..."
echo ""

# Get script directory
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Run the GUI
cd "$SCRIPT_DIR"
python3 src/gui_cleaner.py

if [ $? -ne 0 ]; then
    echo ""
    echo "Error: Failed to start GUI"
    read -p "Press Enter to exit..."
fi

exit 0
