Dir[File.join(__dir__, 'providers', '*.rb')].each { |file| require file }

module RouteInfo
  class Provider
    attr_reader :provider

    def initialize(route_provider: 'google')
      @provider = load_provider(route_provider)
    end

    def get_route_info(route_data)
      provider.get_route_info(route_data)
    end

    private

    def load_provider(route_provider)
      Object.const_get("RouteInfo::Providers::#{provider_class(route_provider)}").new
    end

    def provider_class(route_provider)
      route_provider.split('_').collect!{ |w| w.capitalize }.join
    end
  end
end
