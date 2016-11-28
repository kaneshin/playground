require "http/server"

server = HTTP::Server.new(8080) do |context|
  context.response.content_type = "text/plain"
  context.response.print "Hello world! The time is #{Time.now}"
end

puts "Listening on http://0.0.0.0:8080"
server.listen
