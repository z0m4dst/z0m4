#!/data/data/com.termux/files/usr/bin/bash

echo "→ cleaning proot-distro..."

# salir de cualquier sesión
exit 2>/dev/null

# listar instaladas
echo "→ removing installed distros..."
for d in $(proot-distro list-installed 2>/dev/null); do
    echo "  removing $d"
    proot-distro remove "$d"
done

# limpiar cache
echo "→ cleaning dlcache..."
rm -rf $PREFIX/var/lib/proot-distro/dlcache/*

# limpiar data local
echo "→ cleaning local proot data..."
rm -rf ~/.local/share/proot-distro

# limpiar zf config
echo "→ cleaning zf config..."
rm -rf ~/.zf

# limpiar go build cache
echo "→ cleaning go cache..."
go clean

echo "✓ clean environment ready"
