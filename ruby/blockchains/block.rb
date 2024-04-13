# frozen_string_literal: true

require 'digest'
require 'pp'

# Block class
class Block
  attr_reader :data
  attr_reader :prev
  attr_reader :difficulty
  attr_reader :time
  attr_reader :nonce

  def hash
    Digest::SHA256.hexdigest("#{nonce}#{time}#{difficulty}#{prev}#{data}")
  end

  def initialize(data, prev, difficulty: '00000')
    @data = data
    @difficulty = difficulty
    @prev = prev
    @nonce, @time = compute_hash_with_proof_of_work difficulty
  end

  def compute_hash_with_proof_of_work(difficulty = '00')
    nonce = 0
    time = Time.now.to_i
    loop do
      hash = Digest::SHA256.hexdigest("#{nonce}#{time}#{difficulty}#{prev}#{data}")
      return [nonce, time] if hash.start_with?(difficulty)
      nonce += 1
    end
  end
end

b0 = Block.new 'Hello, Cryptos!', '0000000000000000000000000000000000000000000000000000000000000000'
pp b0
b1 = Block.new 'Hello, Cryptos! - Hello, Cryptos!', b0.hash
pp b1
b2 = Block.new 'Your Name Here', b1.hash
pp b2

puts b1.time >= b0.time
puts b2.time >= b1.time

puts b0.hash.start_with?(b0.difficulty)
puts b1.hash.start_with?(b1.difficulty)
puts b2.hash.start_with?(b2.difficulty)