require "digest"
module Block
    extend self
    def difficulty
        5
    end
    def create(index,timestamp,data,prev_hash)
        block = {
            index:index,
            timestamp:timestamp,
            data:data,
            prev_hash:prev_hash,
            difficulty:self.difficulty,
            nonce:""
        }

        block.merge({hash: self.calculate_hash(block)})
    end
    # 计算hash
    def calculate_hash(block)
        plain_text = "
            #{block[:index]}
            #{block[:timestamp]}
            #{block[:data]}
            #{block[:prev_hash]}
            #{block[:nonce]}
            "
        hash = Digest::SHA256.digest(plain_text)
        hash.to_slice.hexstring
    end
    # 生成一个block
    def generate(last_block,data)
        new_block = self.create(
            last_block[:index] + 1,
            Time.utc.to_s,
            data,
            last_block[:hash]
        )
        i = 0
        loop do
            hex = i.to_s(16)
            # 随机数
            new_block = new_block.merge({nonce:hex})    
            hash = self.calculate_hash(new_block)
            if !self.is_hash_valid?(hash,new_block[:difficulty])
                puts "Mining:try another nonce ... "
                puts hash
                # 新增随机数
                i += 1
                next
            else
                puts "\nMining complete! Nonce for this block is #{new_block[:nonce]}."
                # 设置当前区块hash
                new_block = new_block.merge({hash:hash})
                break
            end
        end
        new_block
    end
    # 验证hash
    def is_hash_valid?(hash,difficulty)
        # 难度为几个0 开始
        prefix = "0" * difficulty
        hash.starts_with?(prefix)
    end

    def is_valid?(new_block, old_block)
        if old_block[:index] + 1 != new_block[:index]
          return false
        elsif old_block[:hash] != new_block[:prev_hash]
          return false
        elsif self.calculate_hash(new_block) != new_block[:hash]
          return false
        end

        true
    end

end