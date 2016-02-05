const http = require('http')
const PORT = 8080

function handleRequest(request, response) {
  console.log('Got request')
  response.writeHead(200, { 'Content-Type': 'text/plain' })
  response.end('Hi, ' + process.env.SAY_HELLO);
}

const server = http.createServer(handleRequest)

server.listen(PORT, function() {
  console.log("Server listening on: http://localhost:%s", PORT);
})
