module RouteInfo
  module Config
    GOOGLE_API_KEY = ENV.fetch('GOOGLE_API_KEY') { 'GOOGLE_API_KEY' }
    RABBITMQ = {
      user: ENV.fetch('RABBITMQ_USER') { 'user' },
      password: ENV.fetch('RABBITMQ_PASS') { 'password' }
    }.freeze
  end
end
