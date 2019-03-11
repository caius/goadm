#!/usr/bin/env ruby
#
#  wrapper.rb
#  goadm
#
#  Basic wrapper intended for vmadm/imgadm given a node,
#  and some commands to run
#
#  Created by Caius Durling on 2019-03-06.
#  Copyright 2019 SwedishCampground. All rights reserved.
#

require "net/ssh"
require "optparse"
require "shellwords"

host, user, port = ARGV.first(3)
user ||= "root"
port ||= 22
port = Integer(port)

remote_command = ARGV[4]

Net::SSH.start(host, user, port: port, non_interactive: true) do |ssh|
  ssh.exec!(remote_command) do |channel, stream,
    data|
    case stream
    when :stdout
      $stdout.puts data
    when :stderr
      $stderr.puts data
    end
  end
end
