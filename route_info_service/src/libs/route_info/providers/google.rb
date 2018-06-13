require 'net/http'
require_relative '../config'

module RouteInfo
  module Providers
    class Google
      API_URL = 'https://maps.googleapis.com/maps/api/distancematrix/json'.freeze
      API_UNITS = 'metric'.freeze
      APU_MODE = 'driving'.freeze

      def get_route_info(route_data)
        uri = URI API_URL
        params = default_params.merge origins: get_origins(route_data), destinations: get_destinations(route_data)
        uri.query = URI.encode_www_form(params) + "&key=#{RouteInfo::Config::GOOGLE_API_KEY}"

        response = Net::HTTP.get_response(uri)
        parse_response response.body
      end

      private

      def default_params
        { units: API_UNITS, mode: APU_MODE }
      end

      def get_origins(route_data)
        [route_data['origin']['latitude'], route_data['origin']['longitude']].join(',')
      end

      def get_destinations(route_data)
        [route_data['destination']['latitude'], route_data['destination']['longitude']].join(',')
      end

      def parse_response(response_body)
        data = JSON.parse response_body

        if data['status'] == 'OK'
          {
            distance: data['rows'].first['elements'].first['distance']['value'],
            duration: data['rows'].first['elements'].first['duration']['value'],
            status: 'OK'
          }
        else
          { error_message: data['error_message'], status: data['status'] }
        end
      end
    end
  end
end
