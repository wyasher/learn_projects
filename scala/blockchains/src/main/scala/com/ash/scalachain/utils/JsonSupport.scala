package com.ash.scalachain.utils

import com.ash.scalachain.blockchain.{Chain, ChainLink, EmptyChain, Transaction}
import spray.json.*

object JsonSupport extends DefaultJsonProtocol {
  implicit object ListTransactionJsonFormat extends RootJsonFormat[List[Transaction]] {
    override def read(json: JsValue): List[Transaction] = json match
      case JsArray(elements) => elements.map {
        case JsObject(fields) =>
          val sender = fields("sender").asInstanceOf[JsString].value
          val recipient = fields("recipient").asInstanceOf[JsString].value
          val value = fields("value").asInstanceOf[JsNumber].value.toLong
          Transaction(sender, recipient, value)
        case _ => throw new IllegalArgumentException("错误的数据")
      }.toList
      case _ => throw new IllegalArgumentException("错误的数据")

    override def write(list: List[Transaction]): JsArray = JsArray(
      list.map(TransactionJsonFormat.write).toVector
    )
  }

  implicit object ListStringJsonFormat extends RootJsonFormat[List[String]] {
    override def read(value: JsValue): List[String] = value match
      case JsArray(elements) => elements.map {
        case JsString(str) => str
        case _ => throw new IllegalArgumentException("Invalid JSON format for List[String]")
      }.toList
      case _ => throw new IllegalArgumentException("Invalid JSON format for List[String]")

    override def write(list: List[String]): JsArray = JsArray(list.map(JsString(_)).toVector)

  }


  implicit object TransactionJsonFormat extends RootJsonFormat[Transaction] {
    override def write(t: Transaction): JsObject = JsObject(
      "sender" -> JsString(t.sender),
      "recipient" -> JsString(t.recipient),
      "value" -> JsNumber(t.value),
    )

    override def read(value: JsValue): Transaction =
      value.asJsObject().getFields("sender", "recipient", "value") match
        case Seq(JsString(sender), JsString(recipient), JsNumber(amount))
        => Transaction(sender, recipient, amount.toLong)
        case _ => throw DeserializationException("Transaction expected")
  }

  implicit object ChainLinkJsonFormat extends RootJsonFormat[ChainLink] {
    override def read(json: JsValue): ChainLink = json.asJsObject.getFields("index", "proof", "values", "previousHash", "timestamp", "tail") match {
      case Seq(JsNumber(index), JsNumber(proof), values, JsString(previousHash), JsNumber(timestamp), tail) =>
        ChainLink(index.toInt, proof.toLong, values.convertTo[List[Transaction]], previousHash, timestamp.toLong, tail.convertTo(ChainJsonFormat))
      case _ => throw DeserializationException("Cannot deserialize: Chainlink expected")
    }

    override def write(obj: ChainLink): JsValue = JsObject(
      "index" -> JsNumber(obj.index),
      "proof" -> JsNumber(obj.proof),
      "values" -> JsArray(obj.values.map(_.toJson).toVector),
      "previousHash" -> JsString(obj.previousHash),
      "timestamp" -> JsNumber(obj.timestamp),
      "tail" -> ChainJsonFormat.write(obj.tail)
    )
  }

  implicit object ChainJsonFormat extends RootJsonFormat[Chain] {
    override def read(json: JsValue): Chain = {
      json.asJsObject.getFields("previousHash") match {
        case Seq(_) => json.convertTo[ChainLink]
        case Seq() => EmptyChain
      }
    }

    override def write(obj: Chain): JsValue = obj match {
      case link: ChainLink => link.toJson
      case EmptyChain => JsObject(
        "index" -> JsNumber(EmptyChain.index),
        "hash" -> JsString(EmptyChain.hash),
        "values" -> JsArray(),
        "proof" -> JsNumber(EmptyChain.proof),
        "timeStamp" -> JsNumber(EmptyChain.timestamp)
      )
    }
  }
}