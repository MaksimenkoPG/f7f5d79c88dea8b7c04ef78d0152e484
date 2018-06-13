require_relative 'libs/route_info/server'

begin
  server = RouteInfo::Server.new
  server.start
rescue Interrupt => _
  server.stop
end
