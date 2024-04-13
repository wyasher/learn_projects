package com.ash.cache

interface GenericCache<K, V> {
    //    当前缓存的数量
    val size: Int

    // set
    operator fun set(key: K, value: V)

    // get
    operator fun get(key: K): V?

    // 移除单个
    fun remove(key: K): V?

    // clear所有
    fun clear()
}