#!/usr/bin/env bash
set -euo pipefail

APP_NAME="Galactic Calculator"
BUNDLE_NAME="${APP_NAME}.app"
BUNDLE_ID="com.example.galacticcalculator"
BUILD_DIR="build"
APP_DIR="${BUILD_DIR}/${BUNDLE_NAME}"
MACOS_DIR="${APP_DIR}/Contents/MacOS"
RESOURCES_DIR="${APP_DIR}/Contents/Resources"

mkdir -p "${MACOS_DIR}" "${RESOURCES_DIR}"

echo "[1/3] Building macOS binary..."
CGO_ENABLED=1 go build -o "${MACOS_DIR}/galactic-calculator" .

cat > "${APP_DIR}/Contents/Info.plist" <<PLIST
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
  <key>CFBundleName</key>
  <string>${APP_NAME}</string>
  <key>CFBundleDisplayName</key>
  <string>${APP_NAME}</string>
  <key>CFBundleIdentifier</key>
  <string>${BUNDLE_ID}</string>
  <key>CFBundleVersion</key>
  <string>1.0.0</string>
  <key>CFBundleShortVersionString</key>
  <string>1.0.0</string>
  <key>CFBundleExecutable</key>
  <string>galactic-calculator</string>
  <key>CFBundlePackageType</key>
  <string>APPL</string>
  <key>LSMinimumSystemVersion</key>
  <string>12.0</string>
</dict>
</plist>
PLIST

echo "[2/3] App bundle created at: ${APP_DIR}"
echo "[3/3] You can launch it with: open \"${APP_DIR}\""
