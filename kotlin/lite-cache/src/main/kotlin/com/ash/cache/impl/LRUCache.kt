package com.ash.cache.impl

import com.ash.cache.Constants.DEFAULT_SIZE
import com.ash.cache.Constants.PRESENT
import com.ash.cache.GenericCache

class LRUCache<K, V>(
    private val delegate: GenericCache<K, V>,
    private val minimalSize: Int = DEFAULT_SIZE
) : GenericCache<K, V> by delegate {
    // 按照访问顺序
    private val keyMap = object : LinkedHashMap<K, Boolean>(minimalSize, .75f, true) {
        override fun removeEldestEntry(eldest: MutableMap.MutableEntry<K, Boolean>): Boolean {
            val tooManyCachedItems = size > minimalSize
            if (tooManyCachedItems) eldestKeyToRemove = eldest.key
            return tooManyCachedItems
        }
    }

    private var eldestKeyToRemove: K? = null

    override fun get(key: K): V? {
        keyMap[key]
        return delegate[key]
    }

    override fun set(key: K, value: V) {
        delegate[key] = value
        cycleKeyMap(key)
    }

    override fun clear() {
        keyMap.clear()
        delegate.clear()
    }

    private fun cycleKeyMap(key: K) {
        keyMap[key] = PRESENT
        eldestKeyToRemove?.let { delegate.remove(it) }
        eldestKeyToRemove = null
    }

}