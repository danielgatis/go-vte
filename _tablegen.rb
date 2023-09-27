#!/usr/bin/env ruby

require "erb"

action_names = [
  :None,
  :Clear,
  :Collect,
  :CsiDispatch,
  :EscDispatch,
  :Execute,
  :Hook,
  :Ignore,
  :OscEnd,
  :OscPut,
  :OscStart,
  :Param,
  :Print,
  :Put,
  :Unhook,
  :BeginUtf8,
]

state_names = [
  :Anywhere,
  :CsiEntry,
  :CsiIgnore,
  :CsiIntermediate,
  :CsiParam,
  :DcsEntry,
  :DcsIgnore,
  :DcsIntermediate,
  :DcsParam,
  :DcsPassthrough,
  :Escape,
  :EscapeIntermediate,
  :Ground,
  :OscString,
  :SosPmApcString,
  :Utf8,
]

states = {}

states[:Anywhere] = {
  0x18 => [:Ground, :Execute],
  0x1a => [:Ground, :Execute],
  0x1b => [:Escape, :None],
}

states[:Ground] = {
  0x00..0x17 => [:Anywhere, :Execute],
  0x19       => [:Anywhere, :Execute],
  0x1c..0x1f => [:Anywhere, :Execute],
  0x20..0x7f => [:Anywhere, :Print],
  0x80..0x8f => [:Anywhere, :Execute],
  0x91..0x9a => [:Anywhere, :Execute],
  0x9c       => [:Anywhere, :Execute],
  0xc2..0xdf => [:Utf8, :BeginUtf8], #Beginning of UTF-8 2 byte sequence
  0xe0..0xef => [:Utf8, :BeginUtf8], #Beginning of UTF-8 3 byte sequence
  0xf0..0xf4 => [:Utf8, :BeginUtf8], #Beginning of UTF-8 4 byte sequence
}

states[:Escape] = {
  0x00..0x17 => [:Anywhere, :Execute],
  0x19       => [:Anywhere, :Execute],
  0x1c..0x1f => [:Anywhere, :Execute],
  0x7f       => [:Anywhere, :Ignore],
  0x20..0x2f => [:EscapeIntermediate, :Collect],
  0x30..0x4f => [:Ground, :EscDispatch],
  0x51..0x57 => [:Ground, :EscDispatch],
  0x59       => [:Ground, :EscDispatch],
  0x5a       => [:Ground, :EscDispatch],
  0x5c       => [:Ground, :EscDispatch],
  0x60..0x7e => [:Ground, :EscDispatch],
  0x5b       => [:CsiEntry, :None],
  0x5d       => [:OscString, :None],
  0x50       => [:DcsEntry, :None],
  0x58       => [:SosPmApcString, :None],
  0x5e       => [:SosPmApcString, :None],
  0x5f       => [:SosPmApcString, :None],
}

states[:EscapeIntermediate] = {
  0x00..0x17 => [:Anywhere, :Execute],
  0x19       => [:Anywhere, :Execute],
  0x1c..0x1f => [:Anywhere, :Execute],
  0x20..0x2f => [:Anywhere, :Collect],
  0x7f       => [:Anywhere, :Ignore],
  0x30..0x7e => [:Ground, :EscDispatch],
}

states[:CsiEntry] = {
  0x00..0x17 => [:Anywhere, :Execute],
  0x19       => [:Anywhere, :Execute],
  0x1c..0x1f => [:Anywhere, :Execute],
  0x7f       => [:Anywhere, :Ignore],
  0x20..0x2f => [:CsiIntermediate, :Collect],
  0x30..0x39 => [:CsiParam, :Param],
  0x3a..0x3b => [:CsiParam, :Param],
  0x3c..0x3f => [:CsiParam, :Collect],
  0x40..0x7e => [:Ground, :CsiDispatch],
}

states[:CsiIgnore] = {
  0x00..0x17 => [:Anywhere, :Execute],
  0x19       => [:Anywhere, :Execute],
  0x1c..0x1f => [:Anywhere, :Execute],
  0x20..0x3f => [:Anywhere, :Ignore],
  0x7f       => [:Anywhere, :Ignore],
  0x40..0x7e => [:Ground, :None],
}

states[:CsiParam] = {
  0x00..0x17 => [:Anywhere, :Execute],
  0x19       => [:Anywhere, :Execute],
  0x1c..0x1f => [:Anywhere, :Execute],
  0x30..0x39 => [:Anywhere, :Param],
  0x3a..0x3b => [:Anywhere, :Param],
  0x7f       => [:Anywhere, :Ignore],
  0x3c..0x3f => [:CsiIgnore, :None],
  0x20..0x2f => [:CsiIntermediate, :Collect],
  0x40..0x7e => [:Ground, :CsiDispatch],
}

states[:CsiIntermediate] = {
  0x00..0x17 => [:Anywhere, :Execute],
  0x19       => [:Anywhere, :Execute],
  0x1c..0x1f => [:Anywhere, :Execute],
  0x20..0x2f => [:Anywhere, :Collect],
  0x7f       => [:Anywhere, :Ignore],
  0x30..0x3f => [:CsiIgnore, :None],
  0x40..0x7e => [:Ground, :CsiDispatch],
}

states[:DcsEntry] = {
  0x00..0x17 => [:Anywhere, :Ignore],
  0x19       => [:Anywhere, :Ignore],
  0x1c..0x1f => [:Anywhere, :Ignore],
  0x7f       => [:Anywhere, :Ignore],
  0x20..0x2f => [:DcsIntermediate, :Collect],
  0x30..0x39 => [:DcsParam, :Param],
  0x3a..0x3b => [:DcsParam, :Param],
  0x3c..0x3f => [:DcsParam, :Collect],
  0x40..0x7e => [:DcsPassthrough, :None],
}

states[:DcsIntermediate] = {
  0x00..0x17 => [:Anywhere, :Ignore],
  0x19       => [:Anywhere, :Ignore],
  0x1c..0x1f => [:Anywhere, :Ignore],
  0x20..0x2f => [:Anywhere, :Collect],
  0x7f       => [:Anywhere, :Ignore],
  0x30..0x3f => [:DcsIgnore, :None],
  0x40..0x7e => [:DcsPassthrough, :None],
}

states[:DcsIgnore] = {
  0x00..0x17 => [:Anywhere, :Ignore],
  0x19       => [:Anywhere, :Ignore],
  0x1c..0x1f => [:Anywhere, :Ignore],
  0x20..0x7f => [:Anywhere, :Ignore],
  0x9c       => [:Ground, :None],
}

states[:DcsParam] = {
  0x00..0x17 => [:Anywhere, :Ignore],
  0x19       => [:Anywhere, :Ignore],
  0x1c..0x1f => [:Anywhere, :Ignore],
  0x30..0x39 => [:Anywhere, :Param],
  0x3a..0x3b => [:Anywhere, :Param],
  0x7f       => [:Anywhere, :Ignore],
  0x3c..0x3f => [:DcsIgnore, :None],
  0x20..0x2f => [:DcsIntermediate, :Collect],
  0x40..0x7e => [:DcsPassthrough, :None],
}

states[:DcsPassthrough] = {
  0x00..0x17 => [:Anywhere, :Put],
  0x19       => [:Anywhere, :Put],
  0x1c..0x1f => [:Anywhere, :Put],
  0x20..0x7e => [:Anywhere, :Put],
  0x7f       => [:Anywhere, :Ignore],
  0x9c       => [:Ground, :None],
}

states[:SosPmApcString] = {
  0x00..0x17 => [:Anywhere, :Ignore],
  0x19       => [:Anywhere, :Ignore],
  0x1c..0x1f => [:Anywhere, :Ignore],
  0x20..0x7f => [:Anywhere, :Ignore],
  0x9c       => [:Ground, :None],
}

states[:OscString] = {
  0x00..0x06 => [:Anywhere, :Ignore],
  0x07       => [:Ground, :None],
  0x08..0x17 => [:Anywhere, :Ignore],
  0x19       => [:Anywhere, :Ignore],
  0x1c..0x1f => [:Anywhere, :Ignore],
  0x20..0xff => [:Anywhere, :OscPut],
}

exit_actions = [
  nil,
  nil,
  nil,
  nil,
  nil,
  nil,
  nil,
  nil,
  nil,
  :Unhook,
  nil,
  nil,
  nil,
  :OscEnd,
  nil,
  nil,
]

table = state_names.map do |state|
  actions = states[state]
  next if actions.nil?

  state_actions = Array.new(256, 0)

  actions.each do |range, act|
    Array(range).each do |r|
      state_actions[r] = act
    end
  end

  state_actions
end

table.compact!

if __FILE__ == $0
  template_path = File.read(File.expand_path("_table.go.erb", File.dirname(__FILE__)))
  out_path = File.expand_path("table.go", File.dirname(__FILE__))

  renderer = ERB.new(template_path, nil, "-")
  output = renderer.result(binding)

  File.write(out_path, output)
  `go fmt #{out_path}`
end
