const { execSync } = require('child_process');

if (process.platform === 'win32') {
  execSync('.\\\\dist\\\\apps\\\\go-compose.exe', { stdio: 'inherit' });
} else {
  execSync('chmod +x ./dist/apps/go-compose && ./dist/apps/go-compose', {
    stdio: 'inherit',
  });
}
