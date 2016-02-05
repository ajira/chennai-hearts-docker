const http = require('http')
const PORT = 8080
const redis = require('redis')
const client = redis.createClient(process.env.REDIS_URL)

function handleRequest(request, response) {
  client.info(function(err, info) {
    console.log('Got request')
    response.writeHead(200, { 'Content-Type': 'application/json' })
    response.end(info)
  })
}

const server = http.createServer(handleRequest)

server.listen(PORT, function() {
  console.log("Server listening on: http://localhost:%s", PORT)
})
