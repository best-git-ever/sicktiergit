#!/bin/bash

# sg (sicktiergit) Demo Script
# Demonstrates the full power of chaos-driven development

echo "╔═══════════════════════════════════════════════════╗"
echo "║      SG (SICKTIERGIT) DEMONSTRATION               ║"
echo "║      Prepare for enlightenment (or confusion)     ║"
echo "╚═══════════════════════════════════════════════════╝"
echo ""

sleep 1

echo "→ Initializing a new repository..."
./bin/sg init demo-project
echo ""

sleep 2

echo "→ Checking repository status..."
./bin/sg status
echo ""

sleep 2

echo "→ Committing changes..."
./bin/sg commit -m "Initial commit of chaos"
echo ""

sleep 2

echo "→ Viewing commit history..."
./bin/sg log
echo ""

sleep 2

echo "→ Listing branches..."
./bin/sg branch
echo ""

sleep 2

echo "→ Pushing to remote (maybe)..."
./bin/sg push --maybe
echo ""

sleep 2

echo "→ Running blame analysis..."
./bin/sg blame --all
echo ""

sleep 2

echo "→ Attempting to undo (spoiler: it won't work)..."
./bin/sg undo
echo ""

sleep 2

echo "→ Seeking emotional support..."
./bin/sg therapy
echo ""

sleep 2

echo "╔═══════════════════════════════════════════════════╗"
echo "║      DEMONSTRATION COMPLETE                       ║"
echo "║      You are now ready to embrace the chaos       ║"
echo "╚═══════════════════════════════════════════════════╝"
echo ""
echo "For more information:"
echo "  • Run: sg --help"
echo "  • Read: USAGE.md"
echo "  • Or just: sg therapy --crisis"
