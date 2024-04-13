# frozen_string_literal: true
require 'zlib'
module Bitcask
  module Serializer
    # Follwing are the header values stored
    # |epoc|keysz|valuesz|key_type|value_type|
    # | 4B |  4B |  4B   |  2B   |    2B    |
    # A total of 16 Bytes
    # L< : unsiged 32 bit int with little endian byte order
    # S< " unsiged 12 bit int with little endian byte order
    # Endian order does not matter, its only used to keep consitent byte ordering to ensure that db file,
    # can be seemlessly interchanged in little/big endian machines
    HEADER_FORMAT = 'L<L<L<S<S<'
    HEADER_SIZE = 16
    CRC32_FORMAT = 'L<'
    CRC32_SIZE = 4
    # 定义支持的数据类型常量 freeze 冻结不允许被修改
    DATA_TYPE = {
      Integer: 1,
      Float: 2,
      String: 3
    }.freeze

    DATA_TYPE_LOOK_UP = {
      DATA_TYPE[:Integer] => :Integer,
      DATA_TYPE[:Float] => :Float,
      DATA_TYPE[:String] => :String
    }.freeze

    DATA_TYPE_DIRECTIVE = {
      DATA_TYPE[:Integer] => 'q<',
      DATA_TYPE[:Float] => 'E'
    }.freeze
    # 序列化
    # @param epoc [Integer] 时间戳
    # @param key [Object] key
    # @param value [Object] value
    def serialize(epoc:, key:, value:)
      key_type = type(key)
      value_type = type(value)

      key_bytes = pack(key, key_type)
      value_bytes = pack(value, value_type)
      # 头部
      header = serialize_header(epoc: epoc, key_type: key_type, key_size: key_bytes.size, value_type: value_type,
                                value_size: value_bytes.size)
      # 数据
      data = key_bytes + value_bytes

      [crc32_header_offset + data.size, crc32(header + data) + header + data]

    end

    # 反序列化
    # @param data  序列化后的数据
    def deserialize(data)
      return 0, '', '' unless crc32_valid?(deserialize_crc32(data[..crc32_offset - 1]), data[crc32_offset..])

      # 获取header信息
      epoc, key_size, _, key_type, value_type = deserialize_header(data[crc32_offset..crc32_header_offset - 1])
      key_bytes = data[crc32_header_offset..crc32_header_offset + key_size - 1]
      value_bytes = data[crc32_header_offset + key_size..]
      [epoc, unpack(key_bytes, key_type), unpack(value_bytes, value_type)]
    end

    # crc32 校验大小 偏移
    def crc32_offset
      CRC32_SIZE
    end

    # header 偏移 大小
    def header_offset
      HEADER_SIZE
    end

    # crc32 + header 偏移 大小
    def crc32_header_offset
      crc32_offset + header_offset
    end

    # 生成crc32 并且序列号
    def crc32(data_bytes)
      [Zlib.crc32(data_bytes, 0)].pack(CRC32_FORMAT)
    end

    def deserialize_crc32(crc)
      crc.unpack1(CRC32_FORMAT)
    end

    # crc32 校验
    # @param digest [String] 校验值
    # @param data_bytes [String] 数据
    def crc32_valid?(digest, data_bytes)
      digest == Zlib.crc32(data_bytes, 0)
    end

    # 序列化头部
    # @param epoc [Integer] 时间戳
    # @param key_type [Symbol] key 类型
    # @param key_size [Integer] key 大小
    # @param value_type [Symbol] value 类型
    # @param value_size [Integer] value 大小
    # @return [String] 序列化后的头部
    def serialize_header(epoc:, key_type:, key_size:, value_type:, value_size:)
      [epoc, key_size, value_size, DATA_TYPE[key_type], DATA_TYPE[value_type]].pack(HEADER_FORMAT)
    end

    # 反序列化头部
    # @param header_data [String]  头部数据
    def deserialize_header(header_data)
      header = header_data.unpack(HEADER_FORMAT)
      [header[0], header[1], header[2], DATA_TYPE_LOOK_UP[header[3]], DATA_TYPE_LOOK_UP[header[4]]]
    end

    #  序列化
    # @param attribute [Object] 序列化对象
    # @param attribute_type [Symbol] 序列化对象类型
    # @return [String] 序列化后的对象
    # @raise [StandardError] 无效的数据类型
    def pack(attribute, attribute_type)
      case attribute_type
      when :Integer, :Float
        [attribute].pack(directive(attribute_type))
      when :String
        attribute.encode('utf-8')
      else
        raise StandardError, "Invalid data type #{attribute_type} for pack"
      end
    end

    # 反序列化
    # @param attribute 序列化后的对象
    # @param attribute_type [Symbol] 序列化对象类型
    # @return [Object] 反序列化后的对象
    # @raise [StandardError] 无效的数据类型
    def unpack(attribute, attribute_type)
      case attribute_type
      when :Integer, :Float
        attribute.unpack1(directive(attribute_type))
      when :String
        attribute
      else
        raise StandardError, "Invalid data type #{attribute_type} for unpack"
      end
    end

    private

    #  获取数据类型指令
    # @param attribute_type [Symbol] 数据类型
    def directive(attribute_type)
      DATA_TYPE_DIRECTIVE[DATA_TYPE[attribute_type]]
    end

    #   获取数据类型
    # @param attribute [Object] 数据对象
    def type(attribute)
      attribute.class.to_s.to_sym
    end
  end
end
