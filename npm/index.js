#!/usr/bin/env node

const { spawnSync } = require('child_process');
const path = require('path');
const os = require('os');

// Detect OS and Architecture
const platform = os.platform();
const arch = os.arch();

let binaryName = '';
let platformName = '';
if (platform === 'darwin') {
  platformName = 'darwin';
} else if (platform === 'linux') {
  platformName = 'linux';
} else if (platform === 'win32') {
  platformName = 'windows';
} else {
  console.error(`❌ [envsafe] Unsupported platform: ${platform}`);
  process.exit(1);
}

let archName = '';
if (arch === 'x64') {
  archName = 'amd64';
} else if (arch === 'arm64') {
  archName = 'arm64';
} else {
  console.error(`❌ [envsafe] Unsupported architecture: ${arch}`);
  process.exit(1);
}

// Map standard name of our builds
binaryName = `envsafe-${platformName}-${archName}`;
if (platform === 'win32') {
  binaryName += '.exe';
}

// Find path to binary inside this npm package bin/ directory
const binaryPath = path.join(__dirname, 'bin', binaryName);

// Spawn the process and forward all arguments
const result = spawnSync(binaryPath, process.argv.slice(2), {
  stdio: 'inherit',
  shell: false
});

process.exit(result.status);
