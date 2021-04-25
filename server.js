const http = require('http');

let count = 0;

const server = http.createServer((req, res) => {
  count++;
  res.setHeader('Content-Type', 'text/plain');
  res.write(count + '\n');
  res.end();
});

console.log('Node.js listening on port 3000');
server.listen(3000);
