package com.ash.scalachain.blockchain

import spray.json.*
import com.ash.scalachain.crypto.Crypto

import java.security.InvalidParameterException
import scala.annotation.targetName
import com.ash.scalachain.utils.JsonSupport.ChainLinkJsonFormat

// 链
sealed trait Chain {
  val index: Int
  val hash: String
  val values: List[Transaction]
  val proof: Long
  val timestamp: Long

  def ::(link: Chain): Chain = link match {
    case l: ChainLink => ChainLink(l.index, l.proof, l.values, l.previousHash, l.timestamp, this)
    case _ => throw new InvalidParameterException("Cannot add invalid link to chain")
  }
}

object Chain {
  def apply[T](b: Chain*): Chain = {
    if (b.isEmpty) EmptyChain
    else {
      val link = b.head.asInstanceOf[ChainLink]
      ChainLink(link.index, link.proof, link.values, link.previousHash, link.timestamp, apply(b.tail: _*))
    }
  }
}

case class ChainLink(index: Int, proof: Long, values: List[Transaction], previousHash: String="", timestamp: Long = System.currentTimeMillis(), tail: Chain = EmptyChain) extends Chain {
  override val hash: String = Crypto.sha256Hash(this.toJson.toString)
}

//创世区块
case object EmptyChain extends Chain {
  override val index: Int = 0
  override val hash: String = "1"
  override val values: List[Transaction] = Nil
  override val proof: Long = 100L
  override val timestamp: Long = System.currentTimeMillis()
}