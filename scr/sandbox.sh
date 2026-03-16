#!/usr/bin/env bash

BASE="$HOME/.z0m4sandbox"
SCRIPTS="$BASE/scripts"
LOGS="$BASE/logs"

mkdir -p "$SCRIPTS" "$LOGS"

TIME=$(date +"%Y%m%d_%H%M%S")
SCRIPT="$SCRIPTS/run_$TIME.sh"
LOG="$LOGS/run_$TIME.log"

echo "paste script — CTRL+D to run"
echo ""

cat > "$SCRIPT"

# agregar shebang si no existe
head -n1 "$SCRIPT" | grep -q "^#!" || sed -i '1i#!/usr/bin/env bash' "$SCRIPT"

# scanner básico
danger=(
"rm -rf /"
"rm -rf /*"
"mkfs"
"dd if="
":(){ :|:& };:"
)

for d in "${danger[@]}"; do
    if grep -q "$d" "$SCRIPT"; then
        echo "⚠ dangerous command detected"
        echo "✗ sandbox blocked"
        exit 1
    fi
done

# verificar sintaxis
bash -n "$SCRIPT" >> "$LOG" 2>&1
if [ $? -ne 0 ]; then
    echo "✗ sandbox error — see log"
    exit 1
fi

# ejecutar
bash "$SCRIPT" >> "$LOG" 2>&1

if [ $? -eq 0 ]; then
    echo "✓ sandbox ok"
else
    echo "✗ sandbox error — see log"
fi
