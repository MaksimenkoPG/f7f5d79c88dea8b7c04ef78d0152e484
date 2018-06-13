require 'bunny'
require 'json'
require_relative 'provider'
require_relative 'config'

module RouteInfo
  class Server
    ATTEMPTS_LIMIT = 60
    attr_reader :channel, :exchange, :queue, :connection, :queue_name

    def initialize
      attempt = 0

      @connection = Bunny.new host: 'rabbitmq',
                              user: RouteInfo::Config::RABBITMQ[:user],
                              password: RouteInfo::Config::RABBITMQ[:password]
      @connection.start
      @channel = connection.create_channel
      @queue_name = 'rpc_queue'
    rescue Bunny::TCPConnectionFailedForAllHosts => e
      attempt += 1

      if attempt == ATTEMPTS_LIMIT
        fail e
      else
        sleep 1
        retry
      end
    end

    def start
      @queue = channel.queue(queue_name)
      @exchange = channel.default_exchange
      subscribe_to_queue
    end

    def stop
      channel.close
      connection.close
    end

    private

    def subscribe_to_queue
      queue.subscribe(block: true) do |_delivery_info, properties, payload|
        provider = RouteInfo::Provider.new
        route_info = provider.get_route_info JSON.parse(payload)

        exchange.publish(
          route_info.to_json,
          routing_key: properties.reply_to,
          correlation_id: properties.correlation_id
        )
      end
    end
  end
end
