# frozen_string_literal: true

require 'ostruct'

module Bitcask
  class DiskStore
    include Serializer

    def initialize(db_file = 'bitcask.db')
      @db_fh = File.open(db_file, 'a+b')
      @write_pos = 0
      @key_dir = {}

      init_key_dir
    end

    def [](key)
      get(key)
    end

    def []=(key, value)
      put(key, value)
    end

    def get(key)
      key_struct = @key_dir[key]
      return '' if key_struct.nil?
      @db_fh.seek(key_struct[:write_pos])
      epoc, key, value = deserialize(@db_fh.read(key_struct[:log_size]))
      value
    end

    def put(key, value)
      log_size, data = serialize(epoc: Time.now.to_i, key: key, value: value)
      @key_dir[key] = key_struct(@write_pos, log_size, key)
      persist(data)
      incr_write_pos(log_size)
      nil
    end

    def keys
      @key_dir.keys
    end

    def size
      @db_fh.size
    end

    def flush
      @db_fh.flush
    end

    def close
      flush
      @db_fh.close
    end

    private

    def key_struct(write_pos, log_size, key)
      { write_pos:, log_size:, key: }
    end

    # 持久化保存数据
    def persist(data)
      @db_fh.write(data)
      @db_fh.flush
    end

    def incr_write_pos(pos)
      @write_pos += pos
    end

    def init_key_dir
      while (crc_and_header_bytes = @db_fh.read(crc32_header_offset))
        header_bytes = crc_and_header_bytes[crc32_offset..]
        epoc, key_size, value_size, key_type, value_type = deserialize_header(header_bytes)
        key_bytes = @db_fh.read(key_size)
        value_bytes = @db_fh.read(value_size)
        key = unpack(key_bytes, key_type)
        value = unpack(value_bytes, value_type)

        crc = crc_and_header_bytes[..crc32_offset - 1]
        raise StandardError, 'file corrupted' unless crc32_valid?(deserialize_crc32(crc), header_bytes + key_bytes + value_bytes)
        log_size = crc32_header_offset + key_size + value_size
        @key_dir[key] = key_struct(@write_pos, log_size, key)
        incr_write_pos(log_size)
      end
    end

  end
end
