package com.ash.scalachain.actor

import akka.actor.{Actor, ActorLogging, Props}
import Miner.*
import akka.actor.Status.{Failure, Success}
import com.ash.scalachain.exception.{InvalidProofException, MinerBusyException}
import com.ash.scalachain.proof.ProofOfWork

import scala.concurrent.Future

object Miner {
  sealed trait MinerMessage

  case class Validate(hash: String, proof: Long) extends MinerMessage

  case class Mine(hash: String) extends MinerMessage

  case object Ready extends MinerMessage

  val props: Props = Props(new Miner)
}

class Miner extends Actor with ActorLogging {

  import context._

  def validate: Receive =
    case Validate(hash: String, proof: Long) =>
      log.info(s"Validating proof $proof")
      if (ProofOfWork.validProof(hash, proof)) {
        log.info("proof is valid!")
        sender() ! Success
      } else {
        log.info("proof is not valid")
        sender() ! Failure(new InvalidProofException(hash, proof))
      }

  def ready: Receive = validate orElse {
    case Mine(hash) =>
      log.info(s"Mining hash $hash ...")
      val proof = Future {
        //        开始挖掘区块
        ProofOfWork.proofOfWork(hash)
      }
      sender() ! proof
      become(busy)
    case Ready =>
      log.info("I'm ready to mine!")
      sender() ! Success("ok")
  }

  def busy: Receive = validate orElse {
    case Mine(_) =>
      log.info("i'm already mining")
      sender() ! Failure(new MinerBusyException("Miner is busy"))
    case Ready =>
      log.info("Ready to mine a new block")
      become(ready)
  }

  override def receive: Receive =
    case Ready => become(ready)
}
