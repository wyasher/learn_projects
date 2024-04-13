package com.ash.cache.impl

import com.ash.cache.GenericCache

//永久缓存
class PerpetualCache<K, V> : GenericCache<K, V> {
    private val cache = HashMap<K, V>()

    override val size: Int
        get() = cache.size

    override fun clear() {
        cache.clear()
    }

    override fun remove(key: K): V? = cache.remove(key)

    override fun get(key: K): V? = cache[key]

    override fun set(key: K, value: V) {
        cache[key] = value
    }


}