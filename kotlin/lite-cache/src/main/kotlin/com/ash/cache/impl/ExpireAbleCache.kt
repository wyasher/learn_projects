package com.ash.cache.impl

import com.ash.cache.GenericCache
import java.util.concurrent.TimeUnit

//可过期的缓存
class ExpireAbleCache<K, V>(
    private val delegate: GenericCache<K, V>,
    private val flushInternal: Long = TimeUnit.MINUTES.toMillis(1)
) : GenericCache<K, V> by delegate {
    private var lastFlushTime = System.nanoTime()

    override val size: Int
        get() {
            recycle()
            return delegate.size
        }

    override fun remove(key: K): V? {
        recycle()
        return delegate.remove(key)
    }

    override fun get(key: K): V? {
        recycle()
        return delegate[key]
    }

    /**
     * 回收
     */
    private fun recycle() {
        val shouldRecycle = System.nanoTime() - lastFlushTime >=TimeUnit.MICROSECONDS.toNanos(flushInternal)
        if (shouldRecycle){
            delegate.clear()
            lastFlushTime = System.nanoTime()
        }
    }

}